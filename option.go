/*
 * Copyright (c) 2023  czyt
 * All rights reserved.
 *
 *  Use of this source code is governed by a MIT License that can be
 * found in the LICENSE file.
 */

package kasbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

type options struct {
	model                  model.Model
	policy                 persist.Adapter
	enforcer               *casbin.SyncedEnforcer
	watcher                persist.Watcher
	enforcerContextCreator EnforcerContextCreator
	useBuiltinModel        bool
}

type Option interface {
	apply(options *options)
}

type ModelOpt struct {
	m model.Model
}

func (m ModelOpt) apply(opt *options) {
	opt.model = m.m
}

func WithModel(m model.Model) Option {
	return ModelOpt{m: m}
}

type PolicyOpt struct {
	p persist.Adapter
}

func (p PolicyOpt) apply(opt *options) {
	opt.policy = p.p
}

func WithPolicy(p persist.Adapter) Option {
	return PolicyOpt{p: p}
}

type EnforcerOpt struct {
	e *casbin.SyncedEnforcer
}

func (e EnforcerOpt) apply(opt *options) {
	opt.enforcer = e.e
}

func WithEnforcer(e *casbin.SyncedEnforcer) Option {
	return EnforcerOpt{e: e}
}

type EnforcerCtxCreatorOpt struct {
	c EnforcerContextCreator
}

func (e EnforcerCtxCreatorOpt) apply(opt *options) {
	opt.enforcerContextCreator = e.c
}

func WithEnforcerContextCreator(c EnforcerContextCreator) Option {
	return EnforcerCtxCreatorOpt{c: c}
}

type UseBuiltinFlagOpt struct {
	useBuiltinModel bool
}

func (u UseBuiltinFlagOpt) apply(opt *options) {
	opt.useBuiltinModel = u.useBuiltinModel
}

func UseBuiltinRBACIfModelUnset(flag bool) Option {
	return UseBuiltinFlagOpt{useBuiltinModel: flag}
}

type WatcherOpt struct {
	watcher persist.Watcher
}

func (w WatcherOpt) apply(opt *options) {
	opt.watcher = w.watcher
}

func WithWatcher(watcher persist.Watcher) Option {
	return WatcherOpt{watcher: watcher}
}
