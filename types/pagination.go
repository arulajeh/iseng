package types

type Pagination struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalPage int `json:"total_page"`
	TotalData int `json:"total_data"`
	Offset    int `json:"-"`
}
