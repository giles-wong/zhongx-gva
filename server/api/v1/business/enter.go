package business

import "github.com/giles-wong/zhongx-gva/server/service"

type ApiGroup struct {
	CertReadingApi
}

var (
	readingService = service.ServiceGroupApp.BusinessServiceGroup.ReadingService
)
