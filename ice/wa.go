/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package ice

import (
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
)

type WA interface {
	MyUser() (string, string, error)
	GetJoinedGroups() ([]*types.GroupInfo, error)
	SendMessage(user string, server *string, conversation *string) (*whatsmeow.SendResponse, error)
}
