// gen for home toolset
package providers

import (
	web "github.com/go-home-admin/go-admin/web"
	providers_0 "github.com/go-home-admin/home/bootstrap/providers"
	services "github.com/go-home-admin/home/bootstrap/services"
	filesystem "github.com/go-home-admin/home/bootstrap/services/filesystem"
)

var _AppSingle *App
var _AuthSingle *Auth
var _ResponseSingle *Response
var _RouteSingle *Route

func GetAllProvider() []interface{} {
	return []interface{}{
		NewApp(),
		NewAuth(),
		NewResponse(),
		NewRoute(),
	}
}

func NewApp() *App {
	if _AppSingle == nil {
		_AppSingle = &App{}
		_AppSingle.Container = services.NewContainer()
		_AppSingle.FrameworkProvider = providers_0.NewFrameworkProvider()
		_AppSingle.MysqlProvider = providers_0.NewMysqlProvider()
		_AppSingle.Route = NewRoute()
		_AppSingle.Response = NewResponse()
		_AppSingle.Provider = web.NewProvider()
		_AppSingle.Local = filesystem.NewLocal()
		providers_0.AfterProvider(_AppSingle, "")
	}
	return _AppSingle
}
func NewAuth() *Auth {
	if _AuthSingle == nil {
		_AuthSingle = &Auth{}
		providers_0.AfterProvider(_AuthSingle, "")
	}
	return _AuthSingle
}
func NewResponse() *Response {
	if _ResponseSingle == nil {
		_ResponseSingle = &Response{}
		providers_0.AfterProvider(_ResponseSingle, "")
	}
	return _ResponseSingle
}
func NewRoute() *Route {
	if _RouteSingle == nil {
		_RouteSingle = &Route{}
		_RouteSingle.RouteProvider = providers_0.NewRouteProvider()
		providers_0.AfterProvider(_RouteSingle, "")
	}
	return _RouteSingle
}
