package business

import api "github.com/giles-wong/zhongx-gva/server/api/v1"

type RouterGroup struct {
	CertReadingRouter
}

var (
	certReadingApi = api.ApiGroupApp.BusinessApiGroup.CertReadingApi
)
