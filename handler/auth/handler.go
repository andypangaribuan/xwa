/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

/* spell-checker: disable */
package auth

import (
	validateip "xwa/handler/auth/validate-ip"
	validatelocalip "xwa/handler/auth/validate-local-ip"

	"github.com/andypangaribuan/gmod/server"
)

type Handler struct{}

func (*Handler) ValidateIP(ctx server.FuseRContext) any {
	return validateip.Exec(ctx)
}

func (*Handler) ValidateLocalIP(ctx server.FuseRContext) any {
	return validatelocalip.Exec(ctx)
}
