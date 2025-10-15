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

import (
	"log"
	"xwa/app"

	"github.com/andypangaribuan/gmod/gm"
	"go.mau.fi/whatsmeow/types/events"
)

func waEventHandler(evt any) {
	switch v := evt.(type) {
	case *events.Message:
		if v.Info.IsFromMe {
			return
		}

		if v.Info.IsGroup {
			waGroupHandler(v)
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

func waGroupHandler(v *events.Message) {
	var (
		chatGroup  = v.Info.Chat.User
		chatServer = v.Info.Chat.Server
		senderUser = v.Info.Sender.User
		// jid         = types.NewJID(chatGroup, chatServer)
		convMessage = v.Message.Conversation
		// extMessage  = v.Message.GetExtendedTextMessage()
		// docMessage  = v.Message.GetDocumentMessage()
	)

	url, exists := app.Env.GroupMap[chatGroup]
	if !exists {
		return
	}

	if convMessage != nil {
		body := map[string]any{
			"chat_group":  chatGroup,
			"chat_server": chatServer,
			"sender_user": senderUser,
			"message":     *convMessage,
		}

		data, code, err := gm.Http.Post(nil, url).SetBody(body).Call()
		if err != nil {
			log.Printf("error when call '%v'. %v.\n", url, err)
			return
		}

		if code != 200 {
			log.Printf("not ok when call '%v'. code: %v, data: %v.\n", url, code, string(data))
			return
		}
	}
}
