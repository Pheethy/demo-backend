package usecase

import (
	"context"
	"demo-backend/models"
	"demo-backend/service/user"
)

// Adapter
type userUsecase struct {
	userRepo user.IUserRepository
}

// Constructor
func NewUserUsecase(userRepo user.IUserRepository) user.IUserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) FetchUsers(ctx context.Context) ([]*models.User, error) {
	return u.userRepo.FetchUsers(ctx)
}
