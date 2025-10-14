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
	_ "github.com/andypangaribuan/gmod"

	"github.com/andypangaribuan/gmod/gm"
)

func init() {
	initEnv()

	gm.Conf.
		SetTimezone(Env.AppTimezone).
		SetClogAddress(Env.ClogAddress, Env.AppName, Env.AppVersion).
		Commit()
}
