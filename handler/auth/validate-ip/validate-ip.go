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
package validateip

import (
	"strings"
	"xwa/app"

	"github.com/andypangaribuan/gmod/server"
)

func Exec(ctx server.FuseRContext) any {
	clientIp1, clientIp2, clientIp3, clientIp4 := getClientIp(ctx)
	authorized := false

	for _, ip := range app.Env.ZInAuthIps {
		if ip == "*" {
			authorized = true
			break
		}

		length := len(strings.Split(ip, "."))
		switch {
		case length == 1 && clientIp1 == ip:
			authorized = true
		case length == 2 && clientIp2 == ip:
			authorized = true
		case length == 3 && clientIp3 == ip:
			authorized = true
		case length == 4 && clientIp4 == ip:
			authorized = true
		}

		if authorized {
			break
		}
	}

	if !authorized {
		return ctx.R401Unauthorized("Unauthorized IP address")
	}

	return ctx.R200OK("", server.ResponseOpt{RawResponse: true})
}

func getClientIp(ctx server.FuseRContext) (string, string, string, string) {
	var (
		clientIp1 = ""
		clientIp2 = ""
		clientIp3 = ""
		clientIp4 = ""
	)

	ls := strings.Split(ctx.GetClientIP(), ".")
	for i, v := range ls {
		switch i {
		case 0:
			clientIp1 = v
		case 1:
			clientIp2 = clientIp1 + "." + v
		case 2:
			clientIp3 = clientIp2 + "." + v
		case 3:
			clientIp4 = clientIp3 + "." + v
		}
	}

	return clientIp1, clientIp2, clientIp3, clientIp4
}
