package business

import api "github.com/giles-wong/zhongx-gva/server/api/v1"

type RouterGroup struct {
	CertReadingRouter
	BooksRouter
}

var (
	certReadingApi = api.ApiGroupApp.BusinessApiGroup.CertReadingApi
	booksApi       = api.ApiGroupApp.BusinessApiGroup.BooksApi
)
