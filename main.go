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
	"xwa/cron"
	"xwa/event"
	"xwa/handler"
	"xwa/util"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/server"
)

func main() {
	util.ExitWithCtrlC()
	fm.CallOrderedInit()
	go event.WA()
	server.Cron(runCron)
	server.FuseR(app.Env.AppHttpPort, rest)
}

func runCron(router server.RouterC) {
	if app.Env.WaGroupDailyUpdateJid != "" && len(app.Env.WaGroupDailyUpdateAt) > 0 {
		for _, at := range app.Env.WaGroupDailyUpdateAt {
			router.EveryDay(at, cron.DailyUpdate)
		}
	}
}

func rest(router server.RouterR) {
	router.AutoRecover(app.Env.AppAutoRecover)
	router.PrintOnError(app.Env.AppServerPrintOnError)
	router.NoLog([]string{"GET: /healthz"})

	router.Endpoints(nil, nil, map[string][]func(server.FuseRContext) any{
		"GET: /healthz":           {handler.Private.Healthz},
		"GET: /private/client-ip": {handler.Private.GetClientIP},
		"GET: /private/my-ip":     {handler.Private.GetMyIp},
	})

	router.Endpoints(nil, handler.Auth.ValidateLocalIP, map[string][]func(server.FuseRContext) any{
		"GET: /private/terminate":     {handler.Private.Terminate},
		"GET: /private/joined-groups": {handler.Private.GetJoinedGroups},
		"GET: /private/my-user":       {handler.Private.GetMyUser},
	})

	router.Endpoints(nil, handler.Auth.ValidateIP, map[string][]func(server.FuseRContext) any{
		"POS: /send/message": {handler.Send.Message},
	})
}
