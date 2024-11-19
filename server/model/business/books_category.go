package business

import (
	"gorm.io/gorm"
	"time"
)

type BooksCategory struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Name      string         `json:"name" gorm:"comment:分类名称"`
	Pid       uint           `json:"pid" gorm:"comment:父级ID"`
	Level     int            `json:"level" gorm:"comment:分类级别"`
	Remark    string         `json:"remark" gorm:"comment:备注"`
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

func (BooksCategory) TableName() string {
	return "books_category"
}
