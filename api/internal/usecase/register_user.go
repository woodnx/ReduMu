package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/woodnx/ReduMu/api/internal/domain/model"
	"github.com/woodnx/ReduMu/api/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	repo repository.IUserRepo
}

type RegisterUserDTO struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Password string    `json:"password" db:"password"`
	Role     string    `json:"role" db:"role"`
	Created  time.Time `json:"created" db:"created"`
	Modified time.Time `json:"modified" db:"modified"`
}

func NewRegisterUser(repo repository.IUserRepo) *RegisterUser {
	return &RegisterUser{repo: repo}
}

func (ru *RegisterUser) Exec(ctx context.Context, name, password, role string) (*RegisterUserDTO, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}
	user := &model.User{
		Name:     name,
		Password: string(pw),
		Role:     role,
	}
	if err := ru.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return nil, nil
}
