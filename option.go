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

type CasModel struct {
	m model.Model
}

func (c CasModel) apply(opt *options) {
	opt.model = c.m
}

func WithModel(m model.Model) Option {
	return CasModel{m: m}
}

type CasPolicy struct {
	p persist.Adapter
}

func (c CasPolicy) apply(opt *options) {
	opt.policy = c.p
}

func WithPolicy(p persist.Adapter) Option {
	return CasPolicy{p: p}
}

type CasEnforcer struct {
	e *casbin.SyncedEnforcer
}

func (c CasEnforcer) apply(opt *options) {
	opt.enforcer = c.e
}

func WithEnforcer(e *casbin.SyncedEnforcer) Option {
	return CasEnforcer{e: e}
}

type CasEnforcerCtxCreator struct {
	c EnforcerContextCreator
}

func (c CasEnforcerCtxCreator) apply(opt *options) {
	opt.enforcerContextCreator = c.c
}

func WithEnforcerContextCreator(c EnforcerContextCreator) Option {
	return CasEnforcerCtxCreator{c: c}
}

type CasUseBuiltin struct {
	useBuiltinModel bool
}

func (c CasUseBuiltin) apply(opt *options) {
	opt.useBuiltinModel = c.useBuiltinModel
}

func UseBuiltinModelAsDefault(use bool) Option {
	return CasUseBuiltin{useBuiltinModel: use}
}
