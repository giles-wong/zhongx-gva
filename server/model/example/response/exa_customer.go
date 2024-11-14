package response

import "github.com/giles-wong/zhongx-gva/server/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
