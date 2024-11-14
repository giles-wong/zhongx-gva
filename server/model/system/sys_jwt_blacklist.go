package system

import (
	"github.com/giles-wong/zhongx-gva/server/global"
)

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
