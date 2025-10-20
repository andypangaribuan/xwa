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
package getmyip

import (
	"fmt"
	"net/url"
	"xwa/app"

	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/server"
)

func Exec(ctx server.FuseRContext) any {
	jid, ok := ctx.ReqQuery()["jid"]
	if !ok {
		return ctx.R412PreconditionFailed("jid query needed")
	}

	pushUrl, ok := app.Env.GroupMap[jid]
	if !ok {
		return ctx.R406NotAcceptable("url not found")
	}

	parsed, err := url.Parse(pushUrl)
	if err != nil {
		return ctx.R400BadRequest("error when parse the url")
	}

	base := fmt.Sprintf("%s://%s", parsed.Scheme, parsed.Host)
	url := base + "/private/client-ip"

	res, code, err := gm.Http.Get(ctx.Clog(), url).Call()
	if err != nil {
		return ctx.R400BadRequest(err)
	}

	if code != 200 {
		return ctx.R406NotAcceptable(fmt.Sprintf("http code %v", code))
	}

	return ctx.R200OK(string(res), server.ResponseOpt{RawResponse: true})
}
