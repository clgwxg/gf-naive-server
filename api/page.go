package api

type Page struct {
	PageNum  int `p:"pageNum" d:"1"`
	PageSize int `p:"pageSize" v:"min:10|max:50#分页数量最小10条|" d:"10"`
}

type PageResult[T any] struct {
	Total int `json:"total"`
	List  []T `json:"list"`
}
