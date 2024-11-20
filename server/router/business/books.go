package business

import (
	"github.com/giles-wong/zhongx-gva/server/middleware"
	"github.com/gin-gonic/gin"
)

type BooksRouter struct {
}

func (b *BooksRouter) InitBooksRouter(Router *gin.RouterGroup) {
	booksRouter := Router.Group("books").Use(middleware.OperationRecord())
	booksRouterWithoutRecord := Router.Group("books")

	{
		booksRouter.POST("addBook", booksApi.CreateBooks) // 新建Books
		booksRouter.POST("editBook", booksApi.EditBooks)
		booksRouter.DELETE("deleteBook", booksApi.DeleteBooks) // 删除Books
		//booksRouter.DELETE("deleteBooksByIds", booksApi.DeleteBooksByIds) // 批量删除Books
	}
	{
		booksRouterWithoutRecord.POST("getBookList", booksApi.BookList)
	}
}
