/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package handler

import (
	"xwa/handler/auth"
	"xwa/handler/private"
)

var (
	Private *private.Handler
	Auth    *auth.Handler
)
