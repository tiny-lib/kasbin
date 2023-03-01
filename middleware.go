/*
 * Copyright (c) 2023  czyt
 * All rights reserved.
 *
 *  Use of this source code is governed by a MIT License that can be
 * found in the LICENSE file.
 */

package kasbin

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/czyt/kasbin/internal/builtinmodel"
	"github.com/go-kratos/kratos/v2/middleware"
	"time"
)

type contextKey string

const (
	enforcerContextCreatorKey contextKey = "enforcerContextCreator"
)

func Server(opts ...Option) middleware.Middleware {
	o := &options{
		enforcerContextCreator: nil,
	}
	for _, opt := range opts {
		opt.apply(o)
	}

	if o.model == nil && o.useBuiltinModel {
		o.model, _ = builtinmodel.LoadDefaultRBACModel()
	}
	o.enforcer, _ = casbin.NewSyncedEnforcer(o.model, o.policy)
	// add watcher to its enforcer
	if o.watcher != nil && o.enforcer != nil {
		o.watcher.SetUpdateCallback(func(s string) {
			o.enforcer.LoadPolicy()
		})
		o.enforcer.SetWatcher(o.watcher)
	}
	// add policy autoload
	if o.autoLoadPolicy && o.enforcer != nil {
		if !o.enforcer.IsAutoLoadingRunning() && o.autoLoadPolicyInterval > time.Duration(0) {
			o.enforcer.StartAutoLoadPolicy(o.autoLoadPolicyInterval)
		}
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var (
				allowed bool
				err     error
			)
			if o.model == nil {
				return nil, ErrModelMissing
			}
			if o.enforcer == nil {
				return nil, ErrEnforcerMissing
			}

			if o.enforcerContextCreator == nil {
				return nil, ErrEnforcerContextCreatorMissing
			}

			ctxCreator := o.enforcerContextCreator
			if err := ctxCreator.ParseContext(ctx); err != nil {
				return nil, ErrParseContextFailed
			}
			ctx = context.WithValue(ctx, enforcerContextCreatorKey, ctxCreator)
			allowed, err = o.enforcer.Enforce(ctxCreator.GetEnforcerContext()...)
			if err != nil {
				return nil, err
			}
			if !allowed {
				return nil, ErrUnauthorized
			}
			return handler(ctx, req)
		}
	}
}

func Client(opts ...Option) middleware.Middleware {
	o := &options{
		enforcerContextCreator: nil,
	}
	for _, opt := range opts {
		opt.apply(o)
	}

	if o.model == nil && o.useBuiltinModel {
		o.model, _ = builtinmodel.LoadDefaultRBACModel()
	}
	o.enforcer, _ = casbin.NewSyncedEnforcer(o.model, o.policy)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			return handler(ctx, req)
		}
	}
}

func EnforceContextCreatorFromContext(ctx context.Context) (EnforcerContextCreator, bool) {
	creator, ok := ctx.Value(enforcerContextCreatorKey).(EnforcerContextCreator)
	return creator, ok
}
