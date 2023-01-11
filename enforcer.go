/*
 * Copyright (c) 2023  czyt
 * All rights reserved.
 *
 *  Use of this source code is governed by a MIT License that can be
 * found in the LICENSE file.
 */

package kasbin

import "context"

type EnforcerContextProvider func() interface{}

type EnforcerContextCreator interface {
	// ParseContext Parse Context info from http
	ParseContext(ctx context.Context) error
	// WithProvider Provide the function to GetEnforcerContext
	WithProvider(providers ...EnforcerContextProvider)
	// GetEnforcerContext Call the EnforcerContextProviders by order and return its result
	GetEnforcerContext() []interface{}
}
