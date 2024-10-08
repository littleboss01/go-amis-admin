package user_auth

import (
	gin "github.com/gin-gonic/gin"
	user "github.com/go-home-admin/go-admin/generate/proto/user"
	http "github.com/go-home-admin/home/app/http"
)

// MyMenu 我的菜单
func (receiver *Controller) MyMenu(req *user.MyMenuRequest, ctx http.Context) (*user.MyMenuResponse, error) {
	// TODO 这里写业务
	return &user.MyMenuResponse{}, nil
}

// GinHandleMyMenu gin原始路由处理
// http.Get(/auth/menus)
func (receiver *Controller) GinHandleMyMenu(ctx *gin.Context) {
	req := &user.MyMenuRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}
	
	resp, err := receiver.MyMenu(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
