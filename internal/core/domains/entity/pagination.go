package entity

// PaginationMeta represents pagination metadata for API responses.
// It contains current page info, total pages, and item offsets.
type PaginationMeta struct {
	CurrentPage  int // current page number (1-based)
	PerPage      int // items per page
	TotalPages   int // total number of pages
	TotalItems   int // total number of items across all pages
	FirstItem    int // 1-based index of first item on current page
	LastItem     int // 1-based index of last item on current page
}