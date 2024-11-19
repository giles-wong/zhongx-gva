package response

import "github.com/giles-wong/zhongx-gva/server/model/business/request"

type ReadingResponse struct {
	ID       uint               `json:"ID"`
	Name     string             `json:"name"`
	Number   string             `json:"number"`
	IdCard   string             `json:"idCard"`
	Mobile   string             `json:"mobile"`
	Profile  string             `json:"profile"`
	File     []request.CertFile `json:"file"`
	Group    string             `json:"group"`
	Referrer string             `json:"referrer"`
	Status   int                `json:"status"`
	Remark   string             `json:"remark"`
}
