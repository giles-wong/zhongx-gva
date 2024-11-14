package response

import (
	"github.com/giles-wong/zhongx-gva/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
