// gen for home toolset
package sys_queue

import (
	providers "github.com/go-home-admin/home/bootstrap/providers"
)

var _ControllerSingle *Controller

func GetAllProvider() []interface{} {
	return []interface{}{
		NewController(),
	}
}

func NewController() *Controller {
	if _ControllerSingle == nil {
		_ControllerSingle = &Controller{}
		providers.AfterProvider(_ControllerSingle, "")
	}
	return _ControllerSingle
}
