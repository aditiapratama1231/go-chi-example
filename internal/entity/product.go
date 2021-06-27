package entity

import "time"

type Product struct {
	ID        uint64
	Name      string
	Stock     int64
	BrandID   uint64
	CreatedAt time.Time
	UpdatedAt *time.Time
}
