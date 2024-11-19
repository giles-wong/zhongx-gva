package business

import (
	"github.com/giles-wong/zhongx-gva/server/middleware"
	"github.com/gin-gonic/gin"
)

type BooksRouter struct {
}

func (b *BooksRouter) InitBooksRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	booksRouter := Router.Group("books").Use(middleware.OperationRecord())
	//booksRouterWithoutRecord := Router.Group("cert")

	{
		booksRouter.POST("createBooks", booksApi.CreateBooks) // 新建Books
		//booksRouter.DELETE("deleteBooks", booksApi.DeleteBooks)           // 删除Books
		//booksRouter.DELETE("deleteBooksByIds", booksApi.DeleteBooksByIds) // 批量删除Books
	}
}
