/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

/* cspell: disable-next-line */
package getclientip

import "github.com/andypangaribuan/gmod/server"

func Exec(ctx server.FuseRContext) any {
	return ctx.R200OK(ctx.GetClientIP(), server.ResponseOpt{RawResponse: true})
}
