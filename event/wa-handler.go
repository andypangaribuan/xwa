/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package event

import "go.mau.fi/whatsmeow/types/events"

func waEventHandler(evt any) {
	switch v := evt.(type) {
	case *events.Message:
		if v.Info.IsFromMe {
			return
		}

		// if !v.Info.IsFromMe {
		// 	if v.Info.IsGroup {
		// 		waGroupHandler(v)
		// 	} else {
		// 		waDirectHandler(v)
		// 	}
		// }
	}
}
