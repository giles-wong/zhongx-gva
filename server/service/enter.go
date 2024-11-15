package service

import (
	"github.com/giles-wong/zhongx-gva/server/service/business"
	"github.com/giles-wong/zhongx-gva/server/service/example"
	"github.com/giles-wong/zhongx-gva/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	BusinessServiceGroup business.ServiceGroup
}
