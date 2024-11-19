package business

import (
	"gorm.io/gorm"
	"math/big"
	"time"
)

type Books struct {
	ID              uint           `json:"id" gorm:"primarykey"`
	BookId          big.Int        `json:"book_id" gorm:"common:图书ID"`
	BookName        string         `json:"book_name" gorm:"common:图书名称"`
	Author          string         `json:"author" gorm:"common:作者"`
	Price           string         `json:"price" gorm:"common:价格"`
	Publish         string         `json:"publish" gorm:"common:出版社"`
	PublicationDate time.Time      `json:"publication_date" gorm:"common:出版日期"`
	Isbn            string         `json:"isbn" gorm:"common:ISBN号"`
	CoverImage      string         `json:"cover_image" gorm:"common:封面图片"`
	Summary         string         `json:"summary" gorm:"common:简介"`
	Category        string         `json:"category" gorm:"common:分类"`
	Remark          string         `json:"remark" gorm:"common:备注"`
	Status          int            `json:"status" gorm:"default:1;comment:证书状态 1正常 2冻结"`
	CreatedAt       time.Time      // 创建时间
	UpdatedAt       time.Time      // 更新时间
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

func (Books) TableName() string {
	return "books"
}
