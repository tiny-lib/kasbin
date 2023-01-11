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
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

type options struct {
	model                  model.Model
	policy                 persist.Adapter
	enforcer               *casbin.SyncedEnforcer
	enforcerContextCreator EnforcerContextCreator
	useBuiltinModel        bool
}

type Option interface {
	apply(options *options)
}

type casModel struct {
	m model.Model
}

func (c casModel) apply(opt *options) {
	opt.model = c.m
}

func WithModel(m model.Model) Option {
	return casModel{m: m}
}

type casPolicy struct {
	p persist.Adapter
}

func (c casPolicy) apply(opt *options) {
	opt.policy = c.p
}

func WithPolicy(p persist.Adapter) Option {
	return casPolicy{p: p}
}

type casEnforcer struct {
	e *casbin.SyncedEnforcer
}

func (c casEnforcer) apply(opt *options) {
	opt.enforcer = c.e
}

func WithEnforcer(e *casbin.SyncedEnforcer) Option {
	return casEnforcer{e: e}
}

type casEnforcerCtxCreator struct {
	c EnforcerContextCreator
}

func (c casEnforcerCtxCreator) apply(opt *options) {
	opt.enforcerContextCreator = c.c
}

func WithEnforcerContextCreator(c EnforcerContextCreator) Option {
	return casEnforcerCtxCreator{c: c}
}

type casUseBuiltin struct {
	useBuiltinModel bool
}

func (c casUseBuiltin) apply(opt *options) {
	opt.useBuiltinModel = c.useBuiltinModel
}

func UseBuiltinRBACIfModelUnset(flag bool) Option {
	return casUseBuiltin{useBuiltinModel: flag}
}

// TODO:add flag to easy control enforcer policy load
