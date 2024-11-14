package request

import (
	"github.com/giles-wong/zhongx-gva/server/model/common/request"
	"github.com/giles-wong/zhongx-gva/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
