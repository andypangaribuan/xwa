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
package getjoinedgroups

import (
	"fmt"
	"xwa/app"

	"github.com/andypangaribuan/gmod/server"
)

func Exec(ctx server.FuseRContext) any {
	ls := make([]string, 0)
	groups, err := app.WA.GetJoinedGroups()
	if err == nil {
		for _, group := range groups {
			ls = append(ls, fmt.Sprintf("%v: %v", group.JID.User, group.GroupName.Name))
		}
	}

	return ctx.R200OK(ls)
}
