package v1

import (
	"github.com/giles-wong/zhongx-gva/server/api/v1/business"
	"github.com/giles-wong/zhongx-gva/server/api/v1/example"
	"github.com/giles-wong/zhongx-gva/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	BusinessApiGroup business.ApiGroup
}
