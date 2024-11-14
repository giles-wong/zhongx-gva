package router

import (
	"github.com/giles-wong/zhongx-gva/server/router/example"
	"github.com/giles-wong/zhongx-gva/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}
