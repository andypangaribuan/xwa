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
package getmyuser

import (
	"fmt"
	"xwa/app"

	"github.com/andypangaribuan/gmod/server"
)

func Exec(ctx server.FuseRContext) any {
	phone, user, err := app.WA.MyUser()
	if err != nil {
		return ctx.R400BadRequest(err)
	}

	return ctx.R200OK(fmt.Sprintf("%v: %v", phone, user), server.ResponseOpt{RawResponse: true})
}
