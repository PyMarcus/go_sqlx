package entities

import "time"

type User struct {
	ID       int        `json:"id" gorm:"primary_key" db:"id"`
	NickName string     `json:"nick_name" gorm:"nick_name" db:"nick_name"`
	Name     string     `json:"name" gorm:"name" db:"name"`
	Email    string     `json:"email" gorm:"email" db:"email"`
	CreateAt *time.Time `json:"create_at" gorm:"create_at" db:"create_at"`
	UpdateAt *time.Time `json:"update_at" gorm:"update_at" db:"update_at"`
}
