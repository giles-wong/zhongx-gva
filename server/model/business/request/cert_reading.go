package request

import common "github.com/giles-wong/zhongx-gva/server/model/common/request"

type GetReadingList struct {
	common.PageInfo
	Name   string `json:"name" form:"name"`
	Number string `json:"number" form:"number"`
	Mobile string `json:"mobile" form:"mobile"`
	IdCard string `json:"idCard" form:"idCard"`
}

type AddReading struct {
	Name     string `json:"name" form:"name"`
	Number   string `json:"number" form:"number"`
	Mobile   string `json:"mobile" form:"mobile"`
	IdCard   string `json:"idCard" form:"idCard"`
	Remark   string `json:"remark" form:"remark"`
	Group    string `json:"group" form:"group"`
	Status   int    `json:"status" form:"status"`
	Referrer string `json:"referrer" form:"referrer"`
	Profile  string `json:"profile" form:"profile"`
	File     string `json:"file" form:"file"`
}

type EditReading struct {
	ID       string `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Number   string `json:"number" form:"number"`
	Mobile   string `json:"mobile" form:"mobile"`
	IdCard   string `json:"idCard" form:"idCard"`
	Remark   string `json:"remark" form:"remark"`
	Group    string `json:"group" form:"group"`
	Status   int    `json:"status" form:"status"`
	Referrer string `json:"referrer" form:"referrer"`
	Profile  string `json:"profile" form:"profile"`
	File     string `json:"file" form:"file"`
}

type ReadingId struct {
	ID uint `json:"id" form:"id"`
}
