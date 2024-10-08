package account_list

import (
	gin "github.com/gin-gonic/gin"
	mail "github.com/go-home-admin/go-admin/generate/proto/mail"
	http "github.com/go-home-admin/home/app/http"
)

// GetInfo1 登陆信息
func (receiver *Controller) GetInfo1(req *mail.GetInfoRequest, ctx http.Context) (*mail.GetInfoResponse, error) {
	// TODO 这里写业务
	return &mail.GetInfoResponse{}, nil
}

// GinHandleGetInfo1 gin原始路由处理
// http.Get(/mail/list)
func (receiver *Controller) GinHandleGetInfo1(ctx *gin.Context) {
	req := &mail.GetInfoRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}
	
	resp, err := receiver.GetInfo1(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
