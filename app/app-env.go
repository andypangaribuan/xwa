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

import "github.com/andypangaribuan/gmod/gm"

func initEnv() {
	Env = &stuEnv{
		AppName:               gm.Util.Env.GetString("APP_NAME"),
		AppVersion:            gm.Util.Env.GetString("APP_VERSION", "0.0.0"),
		AppEnv:                gm.Util.Env.GetAppEnv("APP_ENV"),
		AppTimezone:           gm.Util.Env.GetString("APP_TIMEZONE"),
		AppRestPort:           gm.Util.Env.GetInt("APP_REST_PORT"),
		AppAutoRecover:        gm.Util.Env.GetBool("APP_AUTO_RECOVER"),
		AppServerPrintOnError: gm.Util.Env.GetBool("APP_SERVER_PRINT_ON_ERROR"),

		ZInAuthIps:  gm.Util.Env.GetStringSlice("ZIN_AUTH_IPS", ","),
		ClogAddress: gm.Util.Env.GetString("CLOG_ADDRESS", ""),

		WaLinkedDeviceName: gm.Util.Env.GetString("WA_LINKED_DEVICE_NAME", "xwa"),
		WaMyNumber:         gm.Util.Env.GetStringSlice("WA_MY_NUMBER", ","),
		WaSqlitePath:       gm.Util.Env.GetString("WA_SQLITE_PATH", "res/wa.db"),
		WaDefaultServer:    gm.Util.Env.GetString("WA_DEFAULT_SERVER", "s.whatsapp.net"),
		WaSuperUserPhone:   gm.Util.Env.GetStringSlice("WA_SUPER_USER_PHONE", ",", []string{}),
	}
}
