package user_auth

import (
	gin "github.com/gin-gonic/gin"
	user "github.com/go-home-admin/go-admin/generate/proto/user"
	http "github.com/go-home-admin/home/app/http"
)

// Logout 退出登陆
func (receiver *Controller) Logout(req *user.LoginRequest, ctx http.Context) (*user.LoginResponse, error) {
	// TODO 这里写业务
	return &user.LoginResponse{}, nil
}

// GinHandleLogout gin原始路由处理
// http.Post(/auth/logout)
func (receiver *Controller) GinHandleLogout(ctx *gin.Context) {
	req := &user.LoginRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}
	
	resp, err := receiver.Logout(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
