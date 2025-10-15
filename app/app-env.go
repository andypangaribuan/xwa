/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package app

import (
	"strings"

	"github.com/andypangaribuan/gmod/gm"
)

func initEnv() {
	var (
		groupKey     = "WA_GROUP_"
		groupProcess = make(map[string]any, 0)
		groupMap     = make(map[string]string, 0)
		groups       = gm.Util.Env.GetKeysByPrefix(groupKey)
	)

	for _, group := range groups {
		group = strings.ReplaceAll(group, groupKey, "")
		ls := strings.Split(group, "_")
		if len(ls) > 1 {
			key := groupKey + ls[0] + "_"
			_, exists := groupProcess[key]
			if !exists {
				groupProcess[key] = nil

				jid := gm.Util.Env.GetString(key+"JID", "")
				url := gm.Util.Env.GetString(key+"PUSH_URL", "")
				groupMap[jid] = url
			}
		}
	}

	Env = &stuEnv{
		AppName:               gm.Util.Env.GetString("APP_NAME"),
		AppVersion:            gm.Util.Env.GetString("APP_VERSION", "0.0.0"),
		AppEnv:                gm.Util.Env.GetAppEnv("APP_ENV"),
		AppTimezone:           gm.Util.Env.GetString("APP_TIMEZONE"),
		AppHttpPort:           gm.Util.Env.GetInt("APP_HTTP_PORT"),
		AppAutoRecover:        gm.Util.Env.GetBool("APP_AUTO_RECOVER"),
		AppServerPrintOnError: gm.Util.Env.GetBool("APP_SERVER_PRINT_ON_ERROR"),

		ZInAuthIps:  gm.Util.Env.GetStringSlice("ZIN_AUTH_IPS", ","),
		ClogAddress: gm.Util.Env.GetString("CLOG_ADDRESS", ""),

		WaLinkedDeviceName: gm.Util.Env.GetString("WA_LINKED_DEVICE_NAME", "xwa"),
		WaMyNumber:         gm.Util.Env.GetStringSlice("WA_MY_NUMBER", ","),
		WaSqlitePath:       gm.Util.Env.GetString("WA_SQLITE_PATH", "res/wa.db"),
		WaDefaultServer:    gm.Util.Env.GetString("WA_DEFAULT_SERVER", "s.whatsapp.net"),
		WaSuperUserPhone:   gm.Util.Env.GetStringSlice("WA_SUPER_USER_PHONE", ",", []string{}),

		GroupMap: groupMap,
	}
}
