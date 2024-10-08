package user_auth

import (
	gin "github.com/gin-gonic/gin"
	user "github.com/go-home-admin/go-admin/generate/proto/user"
	http "github.com/go-home-admin/home/app/http"
)

// Login 登陆
func (receiver *Controller) Login(req *user.LoginRequest, ctx http.Context) (*user.LoginResponse, error) {
	// TODO 这里写业务
	return &user.LoginResponse{}, nil
}

// GinHandleLogin gin原始路由处理
// http.Post(/auth/login)
func (receiver *Controller) GinHandleLogin(ctx *gin.Context) {
	req := &user.LoginRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}
	
	resp, err := receiver.Login(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
