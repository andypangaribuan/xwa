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
	"context"
	"fmt"
	"xwa/app"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

func WA() {
	dbLog := waLog.Stdout("Database", "WARN", true)
	ctx := context.Background()
	container, err := sqlstore.New(ctx, "sqlite3", fmt.Sprintf("file:%v?_foreign_keys=on", app.Env.WaSqlitePath), dbLog)
	if err != nil {
		panic(err)
	}

	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	store.DeviceProps.Os = proto.String(app.Env.WaLinkedDeviceName)
	deviceStore, err := container.GetFirstDevice(ctx)
	if err != nil {
		panic(err)
	}

	waClient = &stuWAClient{ctx: context.Background()}
	app.WA = waClient
	clientLog := waLog.Stdout("Client", "WARN", true)
	waClient.client = whatsmeow.NewClient(deviceStore, clientLog)
	waClient.client.AddEventHandler(waEventHandler)
}
