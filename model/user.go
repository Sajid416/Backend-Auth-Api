package model

import "time"

type User struct {
	UserID    string    `json:"userId" gorm:"primaryKey;column:userId"`
	FullName  string    `json:"fullName" gorm:"column:fullName"`
	Email     string    `json:"email" gorm:"column:email;unique"`
	Password  string    `json:"password" gorm:"column:password"`
	Username  string    `json:"userName" gorm:"column:userName"`
	Phone     string    `json:"phone" gorm:"column:phone"`
	Gender    string    `json:"gender" gorm:"column:gender"`
	CreatedAt time.Time `json:"created_at" gorm:"column:Created_At;autoCreateTime"`
}
