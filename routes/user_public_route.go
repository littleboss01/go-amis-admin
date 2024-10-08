// gen for home toolset
package routes

import (
	home_gin_1 "github.com/gin-gonic/gin"
	user_auth "github.com/go-home-admin/go-admin/app/http/user/user_auth"
	home_api_1 "github.com/go-home-admin/home/bootstrap/http/api"
)

// @Bean
type UserPublicRoutes struct {
	user_auth *user_auth.Controller `inject:""`
}

func (c *UserPublicRoutes) GetGroup() string {
	return "user-public"
}
func (c *UserPublicRoutes) GetRoutes() map[*home_api_1.Config]func(c *home_gin_1.Context) {
	return map[*home_api_1.Config]func(c *home_gin_1.Context){
		home_api_1.Post("/auth/login"):  c.user_auth.GinHandleLogin,
		home_api_1.Post("/auth/logout"): c.user_auth.GinHandleLogout,
		home_api_1.Get("/auth/menus"):   c.user_auth.GinHandleMyMenu,
	}
}
