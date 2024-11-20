package business

import (
	"fmt"
	"github.com/giles-wong/general/snowflake"
	"github.com/giles-wong/zhongx-gva/server/global"
	"github.com/giles-wong/zhongx-gva/server/model/business"
	businessReq "github.com/giles-wong/zhongx-gva/server/model/business/request"
	"github.com/giles-wong/zhongx-gva/server/model/common/response"
	"github.com/giles-wong/zhongx-gva/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type BooksApi struct{}

func (b BooksApi) BookList(ctx *gin.Context) {
	var pageInfo businessReq.BooksList
	err := ctx.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	list, total, err := bookService.BookList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", ctx)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", ctx)
}

func (b *BooksApi) CreateBooks(c *gin.Context) {
	var r businessReq.Books
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 生成BookId
	if r.BookId == "" {
		node, err := snowflake.NewNode(1)

		if err != nil {
			fmt.Println(err)
			return
		}
		//生成一个雪花Id
		r.BookId = strconv.FormatInt(node.GetId(), 10)
	}

	book := &business.Books{
		BookName:        r.BookName,
		Author:          r.Author,
		Isbn:            r.Isbn,
		BookId:          r.BookId,
		Category:        r.Category,
		CoverImage:      r.CoverImage,
		Price:           r.Price,
		PublicationDate: r.PublicationDate,
		Publisher:       r.Publisher,
		Remark:          r.Remark,
		Status:          r.Status,
		Summary:         r.Summary,
	}

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

	book := &business.Books{
		BookName:        r.BookName,
		Author:          r.Author,
		Isbn:            r.Isbn,
		BookId:          r.BookId,
		Category:        r.Category,
		CoverImage:      r.CoverImage,
		Price:           r.Price,
		PublicationDate: r.PublicationDate,
		Publisher:       r.Publisher,
		Remark:          r.Remark,
		Status:          r.Status,
		Summary:         r.Summary,
	}

	err = bookService.EditBook(*book)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithDetailed(r, "添加失败", c)
		return
	}
	response.OkWithDetailed(r, "添加成功", c)
}

func (b *BooksApi) DeleteBooks(c *gin.Context) {
	var r businessReq.BookId
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = bookService.DeleteBook(r.BookId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
