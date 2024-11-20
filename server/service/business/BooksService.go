package business

import (
	"errors"
	"github.com/giles-wong/zhongx-gva/server/global"
	"github.com/giles-wong/zhongx-gva/server/model/business"
	businessReq "github.com/giles-wong/zhongx-gva/server/model/business/request"
	businessRes "github.com/giles-wong/zhongx-gva/server/model/business/response"
	"gorm.io/gorm"
)

type BooksService struct {
}

func (bookService *BooksService) BookList(info businessReq.BooksList) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&business.Books{})
	var bookList []business.Books
	var bookResList []businessRes.Books

	//db = db.Where("is_deleted = ?", 0)
	if info.Name != "" {
		db = db.Where("book_name LIKE ?", "%"+info.Name+"%")
	}
	if info.Isbn != "" {
		db = db.Where("isbn = ?", info.Isbn)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&bookList).Error

	for _, item := range bookList {
		//把json 转为切片
		bookResList = append(bookResList, businessRes.Books{
			BookId:          item.BookId,
			BookName:        item.BookName,
			Author:          item.Author,
			Isbn:            item.Isbn,
			Price:           item.Price,
			Summary:         item.Summary,
			Category:        item.Category,
			CoverImage:      item.CoverImage,
			PublicationDate: item.PublicationDate,
			Publisher:       item.Publisher,
			Remark:          item.Remark,
			Status:          item.Status,
		})

	}

	return bookResList, total, err
}

func (bookService *BooksService) AddBook(book business.Books) (bookRes business.Books, err error) {
	var bookR business.Books
	if !errors.Is(global.GVA_DB.Where("isbn = ?", book.Isbn).First(&bookR).Error, gorm.ErrRecordNotFound) {
		return bookRes, errors.New("图书已经存在")
	}

	err = global.GVA_DB.Create(&book).Error
	return book, err
}

func (bookService *BooksService) EditBook(book business.Books) (bookRes business.Books, err error) {
	var bookR business.Books
	if !errors.Is(global.GVA_DB.Where("isbn = ?", book.Isbn).First(&bookR).Error, gorm.ErrRecordNotFound) {
		return bookRes, errors.New("图书已经存在")
	}

	err = global.GVA_DB.Create(&book).Error
	return book, err
}
