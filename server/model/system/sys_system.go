package system

import (
	"github.com/giles-wong/zhongx-gva/server/config"
)

// 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}
