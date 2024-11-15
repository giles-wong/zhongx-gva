package request

type CertFile struct {
	Name string `json:"name" form:"name"`
	URL  string `json:"url" form:"url"`
}
