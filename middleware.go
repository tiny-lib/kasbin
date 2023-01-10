/*
 * MIT License
 *
 * Copyright (c) 2023  czyt
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package kasbin

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/czyt/kasbin/internal/builtinmodel"
	"github.com/go-kratos/kratos/v2/middleware"
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
			allowed, err = o.enforcer.Enforce(ctxCreator.CreateEnforcerContext()()...)
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
