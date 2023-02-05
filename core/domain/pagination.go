package domain

type Pagination struct {
	Items []Product `json:"items"`
	Total int32     `json:"total"`
}
