package initialize

import (
	_ "github.com/giles-wong/zhongx-gva/server/source/example"
	_ "github.com/giles-wong/zhongx-gva/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
