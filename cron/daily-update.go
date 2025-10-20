/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package cron

import (
	"fmt"
	"strconv"
	"time"
	"xwa/app"

	"github.com/andypangaribuan/gmod/gm"
)

func DailyUpdate() {
	mxDailyUpdate.Lock()
	defer mxDailyUpdate.Unlock()

	delaySecond, err := strconv.Atoi(gm.Util.GetRandom(1, "0123456789"))
	if err != nil {
		return
	}

	delay := time.Second * time.Duration(delaySecond)
	if delay > 0 {
		time.Sleep(delay)
	}

	code := "```"
	message := fmt.Sprintf(`%v
status : active
service: %v
device : %v
%v`, code, app.Env.AppName, app.Env.WaLinkedDeviceName, code)
	_, _ = app.WA.SendMessage(app.Env.WaGroupDailyUpdateJid, nil, &message)
}
