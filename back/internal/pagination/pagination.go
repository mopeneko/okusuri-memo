package pagination

type SortOrder int

const (
	SortOrderAsc  SortOrder = 1
	SortOrderDesc SortOrder = -1
)

type Pagination struct {
	SortKey   *string
	SortOrder *SortOrder
	Limit     *int64
	LastID    *string
}
