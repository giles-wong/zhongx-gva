package request

import (
	common "github.com/giles-wong/zhongx-gva/server/model/common/request"
)

type BooksList struct {
	common.PageInfo
	Name string `json:"name" form:"name"`
	Isbn string `json:"isbn" form:"isbn"`
}

type Books struct {
	BookId          int64  `json:"bookId" form:"bookId"`
	BookName        string `json:"bookName" form:"bookName"`
	Author          string `json:"author" form:"author"`
	Price           string `json:"price" form:"price"`
	Publisher       string `json:"publisher" form:"publisher"`
	PublicationDate string `json:"publicationDate" form:"publicationDate"`
	Isbn            string `json:"isbn" form:"isbn"`
	CoverImage      string `json:"coverImage" form:"coverImage"`
	Summary         string `json:"summary" form:"summary"`
	Category        int8   `json:"category" form:"category"`
	Remark          string `json:"remark" form:"remark"`
	Status          int8   `json:"status" form:"status"`
}
