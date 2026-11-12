package paginations

import (
	"math"

	"portal-berita/internal/core/domains/entity"
)

// PaginationService calculates pagination metadata for API responses.
type PaginationService interface {
	Calculate(totalItems, page, perPage int) (*entity.PaginationMeta, error)
}

// Paginator implements PaginationService.
// It holds no state; configuration is passed per-call.
type Paginator struct{}

// Calculate computes pagination metadata given total items, page number, and page size.
// Returns PaginationMeta with:
// - CurrentPage: the requested page (clamped to valid range)
// - PerPage: items per page (default 10, min 1)
// - TotalPages: total pages needed for totalItems
// - TotalItems: same as input totalItems
// - FirstItem: 1-based index of first item on this page
// - LastItem: 1-based index of last item on this page
func (p *Paginator) Calculate(totalItems, page, perPage int) (*entity.PaginationMeta, error) {
	// Validate page: 0 is empty, less than 1 is invalid
	if page == 0 {
		return nil, ErrPageEmpty
	}
	if page < 1 {
		return nil, ErrPageInvalid
	}

	// Validate perPage
	if perPage < 1 {
		return nil, ErrPerPageInvalid
	}

	// Calculate total pages
	totalPages := int(math.Ceil(float64(totalItems) / float64(perPage)))
	if totalPages == 0 {
		totalPages = 1
	}

	// Reject page beyond total pages
	if page > totalPages {
		return nil, ErrMaxPage
	}

	// Calculate 1-based item offsets for current page
	firstItem := (page-1)*perPage + 1
	lastItem := min(page*perPage, totalItems)

	// If no items, return empty pagination
	if totalItems == 0 {
		return &entity.PaginationMeta{
			CurrentPage: page,
			PerPage:     perPage,
			TotalPages:  1,
			TotalItems:  0,
			FirstItem:   0,
			LastItem:    0,
		}, nil
	}

	return &entity.PaginationMeta{
		CurrentPage: page,
		PerPage:     perPage,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
		FirstItem:   firstItem,
		LastItem:    lastItem,
	}, nil
}

// NewPaginator creates a new PaginationService.
func NewPaginator() PaginationService {
	return &Paginator{}
}
