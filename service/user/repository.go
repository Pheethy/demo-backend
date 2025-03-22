package user

import (
	"context"
	"demo-backend/models"
)

// Port Users repository
type IUserRepository interface {
	FetchUsers(ctx context.Context) ([]*models.User, error)
}
