package models

import (
	"database/sql"
	"time"
)

// HasSoftDelete base model
type HasSoftDelete struct {
	DeletedAt sql.NullTime `gorm:"column:deleted_at" json:"-"`
	DeleteRes interface{}  `gorm:"-" json:"deleteAt"`
}

func (x *HasSoftDelete) BeforeCreate() (err error) {
	x.DeletedAt = sql.NullTime{Valid: false}
	return
}

func (x *HasSoftDelete) AfterFind() (err error) {
	if v, err := x.DeletedAt.Value(); err != nil {
		x.DeleteRes = v
	}
	return
}

// Stamp current time to model
func (x *HasSoftDelete) Stamp() {
	x.DeletedAt = sql.NullTime{Time: time.Now(), Valid: true}
}
