package infra

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/woodnx/ReduMu/api/clock"
	"github.com/woodnx/ReduMu/api/gen/db"
	"github.com/woodnx/ReduMu/api/internal/domain/model"
	"github.com/woodnx/ReduMu/api/internal/domain/repository"
)

type userRepository struct {
	clocker clock.Clocker
	q       *db.Queries
}

func NewUserInfra(sqldb *sql.DB) repository.IUserRepo {
	return &userRepository{
		clocker: clock.RealClocker{},
		q:       db.New(sqldb),
	}
}

func (r userRepository) Create(ctx context.Context, u *model.User) error {
	u.Created = r.clocker.Now()
	u.Modified = r.clocker.Now()
	id := uuid.New()

	arg := db.AddUserParams{
		ID:       id,
		Name:     u.Name,
		Password: u.Password,
		Role:     u.Role,
		Created:  u.Created,
		Modified: u.Modified,
	}
	_, err := r.q.AddUser(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}

func (r userRepository) FindByID(ctx context.Context, id model.UserID) (*model.User, error) {
	u, err := r.q.GetUserByID(ctx, uuid.UUID(id))
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:       model.UserID(u.ID),
		Name:     u.Name,
		Role:     u.Role,
		Created:  u.Created,
		Modified: u.Modified,
	}
	return user, nil
}

func (r userRepository) GetAll(ctx context.Context) (model.Users, error) {
	us, err := r.q.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	users := model.Users{}
	for _, u := range us {
		users = append(users, &model.User{
			ID:       model.UserID(u.ID),
			Name:     u.Name,
			Role:     u.Role,
			Created:  u.Created,
			Modified: u.Modified,
		})
	}
	return users, nil
}
