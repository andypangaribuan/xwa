/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package message

import (
	"fmt"
	"xwa/app"

	"github.com/andypangaribuan/gmod/clog"
	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/server"
)

type stuHandler struct {
	ctx server.FuseRContext
	req *stuRequest
}

type stuRequest struct {
	body stuReqBody
}

type stuReqBody struct {
	ChatGroup  string `json:"chat_group"`
	ChatServer string `json:"chat_server"`
	Message    string `json:"message"`
}

func Exec(ctx server.FuseRContext) any {
	handler := &stuHandler{
		ctx: ctx,
		req: new(stuRequest),
	}

	return handler.main()
}

func (slf *stuHandler) controller() (server.FuseRContext, clog.Instance, *stuRequest) {
	return slf.ctx, slf.ctx.Clog(), slf.req
}

func (slf *stuHandler) main() any {
	ok, res := slf.validateRequest()
	if !ok {
		return res
	}

	return slf.process()
}
func (slf *stuHandler) validateRequest() (bool, any) {
	ctx, _, req := slf.controller()

	err := ctx.ReqParser(nil, &req.body)
	if err != nil {
		return false, ctx.R500InternalServerError(err)
	}

	nilOrEmptyKey := fm.FindNilOrEmptyString(map[string]any{
		"chat_group":  req.body.ChatGroup,
		"chat_server": req.body.ChatServer,
		"message":     req.body.Message,
	})

	if nilOrEmptyKey != "" {
		return false, ctx.R400BadRequest(fmt.Sprintf("%s on request body cannot be empty", nilOrEmptyKey))
	}

	return true, nil
}

func (slf *stuHandler) process() any {
	ctx, _, req := slf.controller()

	_, err := app.WA.SendMessage(req.body.ChatGroup, &req.body.ChatServer, &req.body.Message)
	if err != nil {
		return ctx.R400BadRequest(err)
	}

	return ctx.R200OK("ok", server.ResponseOpt{RawResponse: true})
}
