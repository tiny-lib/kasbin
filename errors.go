/*
 * Copyright (c) 2023  czyt
 * All rights reserved.
 *
 *  Use of this source code is governed by a MIT License that can be
 * found in the LICENSE file.
 */

package kasbin

import "github.com/go-kratos/kratos/v2/errors"

const reason string = "FORBIDDEN"

var (
	ErrEnforcerContextCreatorMissing = errors.Forbidden(reason, "EnforcerContextCreator is required")
	ErrModelMissing                  = errors.Forbidden(reason, "Model is missing")
	ErrEnforcerMissing               = errors.Forbidden(reason, "Enforcer is missing")
	ErrParseContextFailed            = errors.Forbidden(reason, "Parse Context Failed")
	ErrUnauthorized                  = errors.Forbidden(reason, "Unauthorized Access")
)
