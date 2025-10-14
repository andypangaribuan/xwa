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

import "github.com/andypangaribuan/gmod/ice"

type stuEnv struct {
	AppName               string
	AppVersion            string
	AppEnv                ice.AppEnv
	AppTimezone           string
	AppRestPort           int
	AppAutoRecover        bool
	AppServerPrintOnError bool

	ZInAuthIps  []string
	ClogAddress string

	WaLinkedDeviceName string
	WaMyNumber         []string
	WaSqlitePath       string
	WaDefaultServer    string
	WaSuperUserPhone   []string
}
