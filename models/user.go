package models

import "github.com/gofrs/uuid"

type User struct {
	Id        *uuid.UUID `json:"id" db:"id"`
	Username  string     `json:"username" db:"username"`
	Password  string     `json:"password" db:"password"`
	Email     string     `json:"email" db:"email"`
	CreatedAt string     `json:"created_at" db:"created_at"`
	UpdatedAt string     `json:"updated_at" db:"updated_at"`
	Images    []*Image   `json:"images" db:"images"`
}

type Image struct {
	Id        *uuid.UUID `json:"id" db:"id"`
	UserId    *uuid.UUID `json:"user_id" db:"user_id"`
	Filename  string     `json:"filename" db:"filename"`
	URL       string     `json:"url" db:"url"`
	CreatedAt string     `json:"created_at" db:"created_at"`
	UpdatedAt string     `json:"updated_at" db:"updated_at"`
}

func (user *User) NewId() {
	id, _ := uuid.NewV4()
	user.Id = &id
}
