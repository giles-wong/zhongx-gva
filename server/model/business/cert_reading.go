package business

import (
	"gorm.io/gorm"
	"time"
)

type CertReading struct {
	ID        uint           `json:"ID" gorm:"primarykey"`
	Name      string         `json:"name" gorm:"comment:姓名"`
	Number    string         `json:"number" gorm:"index;comment:证书编号"`
	IdCard    string         `json:"idCard" gorm:"comment:身份证号"`
	Mobile    string         `json:"mobile" gorm:"comment:手机号"`
	Profile   string         `json:"profile" gorm:"comment:头像照片"`
	File      string         `json:"file" gorm:"type:json;comment:证书文件"`
	Group     string         `json:"group" gorm:"comment:所属组"`
	Referrer  string         `json:"referrer" gorm:"comment:推荐人/单位"`
	Status    int            `json:"status" gorm:"default:1;comment:证书状态 1正常 2冻结"`
	Remark    string         `json:"remark" gorm:"comment:备注"`
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

func (CertReading) TableName() string {
	return "cert_reading"
}
