package utils

type PaginationMeta struct {
	Limit  int32 `query:"limit"`
	Offset int32 `query:"offset"`
}

const (
	DefaultLimit  = 100
	DefaultOffset = 0
)

func PaginationLimit(limit int32) int32 {
	if limit == 0 {
		return DefaultLimit
	}
	return limit
}

func PaginationOffset(offset int32) int32 {
	if offset == 0 {
		return DefaultOffset
	}
	return offset
}
