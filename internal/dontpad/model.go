package dontpad

type response struct {
	Ads          bool   `json:"ads"`
	LastModified int64  `json:"lastModified"`
	Changed      bool   `json:"changed"`
	Body         string `json:"body"`
}
