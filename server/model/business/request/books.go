package request

import (
	common "github.com/giles-wong/zhongx-gva/server/model/common/request"
	"time"
)

type BooksList struct {
	common.PageInfo
}

type Books struct {
	BookId          int64     `json:"bookId" form:"bookId"`
	BookName        string    `json:"bookName" form:"bookName"`
	Author          string    `json:"author" form:"author"`
	Price           string    `json:"price" form:"price"`
	Publish         string    `json:"publish" form:"publish"`
	PublicationDate time.Time `json:"publicationDate" form:"publicationDate"`
	Isbn            string    `json:"isbn" form:"isbn"`
	CoverImage      string    `json:"coverImage" form:"coverImage"`
	Summary         string    `json:"summary" form:"summary"`
	Category        string    `json:"category" form:"category"`
	Remark          string    `json:"remark" form:"remark"`
	Status          int       `json:"status" form:"status"`
}
