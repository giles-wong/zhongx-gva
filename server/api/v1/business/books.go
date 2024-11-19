package business

import (
	"github.com/giles-wong/zhongx-gva/server/global"
	"github.com/giles-wong/zhongx-gva/server/model/business"
	businessReq "github.com/giles-wong/zhongx-gva/server/model/business/request"
	"github.com/giles-wong/zhongx-gva/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BooksApi struct{}

func (b *BooksApi) CreateBooks(c *gin.Context) {
	var r businessReq.Books
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	book := &business.Books{}

	_, err = bookService.AddBook(*book)
	if err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Error(err))
		response.FailWithDetailed(r, "添加失败", c)
		return
	}
	response.OkWithDetailed(r, "添加成功", c)
}

func (b *BooksApi) EditBooks(c *gin.Context) {
	var r businessReq.Books
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	book := &business.Books{}

	_, err = bookService.AddBook(*book)
	if err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Error(err))
		response.FailWithDetailed(r, "添加失败", c)
		return
	}
	response.OkWithDetailed(r, "添加成功", c)
}
