/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package main

import (
	"xwa/app"
	"xwa/event"
	"xwa/handler"
	"xwa/util"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/server"
)

func main() {
	util.ExitWithCtrlC()
	fm.CallOrderedInit()
	event.WA()
	server.FuseR(app.Env.AppRestPort, rest)
}

func rest(router server.RouterR) {
	router.AutoRecover(app.Env.AppAutoRecover)
	router.PrintOnError(app.Env.AppServerPrintOnError)
	router.NoLog([]string{"GET: /healthz"})

	router.Endpoints(nil, nil, map[string][]func(server.FuseRContext) any{
		"GET: /healthz":           {handler.Private.Healthz},
		"GET: /private/client-ip": {handler.Private.GetClientIP},
	})

	router.Endpoints(nil, handler.Auth.ValidateLocalIP, map[string][]func(server.FuseRContext) any{
		"GET: /private/terminate": {handler.Private.Terminate},
	})

	// if app.Env.AppType == "api" {
	// 	router.Endpoints(nil, handler.Auth.ValidateIP, map[string][]func(server.FuseRContext) any{
	// 		"GET: /wa/groups":       {handler.WA.GetGroups},
	// 		"POS: /wa/send-message": {handler.WA.SendMessage},
	// 		"POS: /wa/send-otp":     {handler.WA.SendOtp},
	// 		"GET: /wa/otp-code":     {handler.WA.GetOtpCode},
	// 	})
	// }
}
