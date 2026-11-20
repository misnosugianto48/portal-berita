package responses

import "time"

// CategoryResponse represents single category data in HTTP responses.
type CategoryResponse struct {
	Id          int64        `json:"id"`
	Title       string       `json:"title"`
	Slug        string       `json:"slug"`
	CreatedById int64        `json:"created_by_id"`
	User        UserResponse `json:"user,omitzero"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   *time.Time   `json:"updated_at,omitempty"`
}

// UserResponse represents nested author info inside category response.
type UserResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// CategoryDetailResponse represents response envelope for single category.
type CategoryDetailResponse struct {
	Meta
	Data CategoryResponse `json:"data"`
}

// CategoryListResponse represents response envelope for category list.
type CategoryListResponse struct {
	Meta
	Data []CategoryResponse `json:"data"`
}
