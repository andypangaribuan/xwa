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
	"errors"
	"xwa/app"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

func (slf *stuWAClient) GetJoinedGroups() ([]*types.GroupInfo, error) {
	return slf.client.GetJoinedGroups(slf.ctx)
}

func (slf *stuWAClient) SendMessage(user string, server *string, conversation *string) (*whatsmeow.SendResponse, error) {
	senderServer := app.Env.WaDefaultServer
	if server != nil {
		senderServer = *server
	}

	jid := types.NewJID(user, senderServer)

	if conversation != nil {
		resp, err := slf.client.SendMessage(slf.ctx, jid, &waE2E.Message{
			Conversation: proto.String(*conversation),
		})

		return &resp, err
	}

	return nil, errors.New("doesn't meet condition")
}
