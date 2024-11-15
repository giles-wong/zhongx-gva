package response

import (
	"github.com/giles-wong/zhongx-gva/server/model/business"
)

type ReadingResponse struct {
	Reading business.CertReading `json:"reading"`
}
