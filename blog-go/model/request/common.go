package request

type PageInfo struct {
	Page     int `json:"page" from:"page"`
	PageSize int `json:"page_size" from:"page_size"`
}
