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

func findSenderServer(sender string) *string {
	waMxSenderServer.Lock()
	defer waMxSenderServer.Unlock()

	server, ok := waSenderServer[sender]
	if ok {
		return &server
	}

	if waClient == nil {
		return nil
	}

	groups, err := waClient.GetJoinedGroups()
	if err != nil {
		return nil
	}

	for _, group := range groups {
		if group.JID.User == sender {
			waSenderServer[sender] = group.JID.Server
			return &group.JID.Server
		}
	}

	return nil
}
