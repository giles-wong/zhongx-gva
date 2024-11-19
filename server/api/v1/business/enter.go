package business

import "github.com/giles-wong/zhongx-gva/server/service"

type ApiGroup struct {
	CertReadingApi
	BooksApi
}

var (
	readingService = service.ServiceGroupApp.BusinessServiceGroup.ReadingService
	bookService    = service.ServiceGroupApp.BusinessServiceGroup.BooksService
)
