package response

import "github.com/giles-wong/zhongx-gva/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
