package common

type Pagination struct {
	Marker  *string // 记录标志
	Limit   *int    // 分页大小
	SortBy  *string // 排序字段, + - 号表示排序顺序
	PageNum *int    // 页码
}
