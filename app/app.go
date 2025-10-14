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

	RandomColors = []string{"blue", "white", "black", "yellow", "purple", "orange", "pink", "green", "red"}
	RandomNumbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	RandomWords = []string{"nova", "core", "arc", "pilot", "grid", "spark", "pulse", "forge", "loop", "echo", "mesh", "cell", "link", "path", "core", "helix", "layer"}
}
