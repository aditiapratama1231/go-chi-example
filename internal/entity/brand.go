package entity

import "time"

type Brand struct {
	ID        uint64
	Name      string
	Products  []Product
	CreatedAt time.Time
	UpdatedAt *time.Time
}
