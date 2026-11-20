package entity

import "time"

type CategoryEntity struct {
	Id          int64
	Title       string
	Slug        string
	User        UserEntity
	CreatedById int64
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}
