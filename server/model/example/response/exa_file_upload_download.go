package response

import "github.com/giles-wong/zhongx-gva/server/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
