package models

import "time"

type SoftDelete struct {
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
