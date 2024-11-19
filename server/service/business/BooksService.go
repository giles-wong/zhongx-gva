package business

import (
	"errors"
	"github.com/giles-wong/zhongx-gva/server/global"
	"github.com/giles-wong/zhongx-gva/server/model/business"
	"gorm.io/gorm"
)

type BooksService struct {
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
