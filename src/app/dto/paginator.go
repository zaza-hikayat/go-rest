package dto

type Pagination struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Limit    int    `json:"limit"`
	Search   string `json:"search"`
	SearchBy string `json:"searchBy"`
}
