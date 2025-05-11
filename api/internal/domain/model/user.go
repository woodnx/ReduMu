package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserID uuid.UUID

type User struct {
	ID       UserID    `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Password string    `json:"password" db:"password"`
	Role     string    `json:"role" db:"role"`
	Created  time.Time `json:"created" db:"created"`
	Modified time.Time `json:"modified" db:"modified"`
}

func (u *User) ComparePassword(pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw))
}

type Users []*User
