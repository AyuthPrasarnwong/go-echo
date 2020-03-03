package models

// Base default Primary Key
type Base struct {
	ID int64 `gorm:"primary_key;column:id" json:"id"`
}

// SetID set id
func (x *Base) SetID(v int64) *Base {
	x.ID = v
	return x
}
