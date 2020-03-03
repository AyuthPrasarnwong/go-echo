package models

import (
	"database/sql"
	"time"
)

// User model
type User struct {
	Base
	HasPassword
	Username         sql.NullString `gorm:"column:login" json:"-"`
	UsernameRes      interface{}    `gorm:"-" json:"username"`
	Firstname        sql.NullString `gorm:"column:first_name;type:varchar(100)" json:"-"`
	FirstnameRes     interface{}    `gorm:"-" json:"first_name"`
	Lastname         sql.NullString `gorm:"column:last_name;type:varchar(100)" json:"-"`
	LastnameRes      interface{}    `gorm:"-" json:"last_name"`
	Email            sql.NullString `gorm:"column:email" json:"-"`
	EmailRes         interface{}    `gorm:"-" json:"email"`
	ImageURL         sql.NullString `gorm:"column:image_url;size:255" json:"-"`
	ImageURLRes      interface{}    `gorm:"-" json:"image_url"`
	Activated        int            `gorm:"column:activated" json:"activated"`
	LangKey          sql.NullString `gorm:"column:lang_key;size:6" json:"-"`
	LangKeyRes       interface{}    `gorm:"-" json:"lang_key"`
	ActivationKey    sql.NullString `gorm:"column:activation_key;type:varchar(10)" json:"-"`
	ActivationKeyRes interface{}    `gorm:"-" json:"activation_key"`
	ResetKey         sql.NullString `gorm:"column:reset_key;type:varchar(10)" json:"-"`
	ResetKeyRes      interface{}    `gorm:"-" json:"reset_key"`
	ResetDate        time.Time      `gorm:"column:reset_date;type:varchar(10)" json:"reset_date"`
	CreatedBy        string         `gorm:"column:created_by" json:"created_by"`
	CreatedAt        time.Time      `gorm:"column:created_date" json:"createdAt"`
}

// BeforeCreate set created at
func (x *User) BeforeCreate() (err error) {
	x.CreatedAt = time.Now()
	x.ResetDate = time.Now()
	return
}

func (x *User) AfterFind() (err error) {
	var v interface{}
	v, err = x.Username.Value()
	x.UsernameRes = v
	v, err = x.Firstname.Value()
	x.FirstnameRes = v
	v, err = x.Lastname.Value()
	x.LastnameRes = v
	v, err = x.Email.Value()
	x.EmailRes = v
	v, err = x.ImageURL.Value()
	x.ImageURLRes = v
	v, err = x.LangKey.Value()
	x.LangKeyRes = v
	v, err = x.ActivationKey.Value()
	x.ActivationKeyRes = v
	v, err = x.ResetKey.Value()
	x.ResetKeyRes = v

	return
}

// TableName set table name for User
func (x *User) TableName() string {
	return "jhi_user"
}
