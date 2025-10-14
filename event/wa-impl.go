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

func (slf *stuWAClient) isClientAvailable() error {
	if slf.client == nil {
		return errors.New("client is nil")
	}

	return nil
}

func (slf *stuWAClient) MyUser() (string, error) {
	err := slf.isClientAvailable()
	if err != nil {
		return "", err
	}

	return slf.client.Store.LID.User, nil
}

func (slf *stuWAClient) GetJoinedGroups() ([]*types.GroupInfo, error) {
	err := slf.isClientAvailable()
	if err != nil {
		return nil, err
	}

	return slf.client.GetJoinedGroups(slf.ctx)
}

func (slf *stuWAClient) SendMessage(user string, server *string, conversation *string) (*whatsmeow.SendResponse, error) {
	err := slf.isClientAvailable()
	if err != nil {
		return nil, err
	}

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
