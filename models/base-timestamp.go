package models

import "time"

// HasTimestamp base model
type HasTimestamp struct {
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// Stamp current time to model
func (x *HasTimestamp) Stamp() {
	x.UpdatedAt = time.Now()
	if x.CreatedAt.IsZero() {
		x.CreatedAt = x.UpdatedAt
	}
}
