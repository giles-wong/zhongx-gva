package request

import (
	"github.com/giles-wong/zhongx-gva/server/model/common/request"
	"github.com/giles-wong/zhongx-gva/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
