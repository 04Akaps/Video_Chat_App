package types

type Paging struct {
	Pagination bool  `form:"pagination"`
	Page       int64 `form:"page"`
	PageSize   int64 `form:"pageSize"`
}
