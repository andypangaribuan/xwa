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
package private

import (
	getclientip "xwa/handler/private/get-client-ip"
	gethealth "xwa/handler/private/get-health"
	"xwa/handler/private/terminate"

	"github.com/andypangaribuan/gmod/server"
)

type Handler struct{}

func (*Handler) Healthz(ctx server.FuseRContext) any {
	return gethealth.Exec(ctx)
}

func (*Handler) GetClientIP(ctx server.FuseRContext) any {
	return getclientip.Exec(ctx)
}

func (*Handler) Terminate(ctx server.FuseRContext) any {
	return terminate.Exec(ctx)
}
