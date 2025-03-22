package user

import (
	"context"
	"demo-backend/models"
)

// Port User Usecase
type IUserUsecase interface {
	FetchUsers(ctx context.Context) ([]*models.User, error)
}
