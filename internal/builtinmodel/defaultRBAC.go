/*
 * Copyright (c) 2023  czyt
 * All rights reserved.
 *
 *  Use of this source code is governed by a MIT License that can be
 * found in the LICENSE file.
 */

package builtinmodel

import "github.com/casbin/casbin/v2/model"

const (
	defaultRBACModelDefine = `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && (r.act == p.act || p.act == "*")
`
)

func LoadDefaultRBACModel() (model.Model, error) {
	return model.NewModelFromString(defaultRBACModelDefine)
}
