package model

type PageResult struct {
	Data       interface{} `json:"data"`       // 当前页的数据
	Total      int64       `json:"total"`      // 总记录数
	Page       int         `json:"page"`       // 当前页码
	PageSize   int         `json:"pageSize"`   // 每页数量
	TotalPages int         `json:"totalPages"` // 总页数
}
