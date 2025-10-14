/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package terminate

import (
	"os"

	"github.com/andypangaribuan/gmod/server"
)

func Exec(ctx server.FuseRContext) any {
	os.Exit(1)
	return ctx.R200OK("", server.ResponseOpt{RawResponse: true})
}
