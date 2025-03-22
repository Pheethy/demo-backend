package repository

import (
	"context"
	"demo-backend/models"
	"demo-backend/service/user"
	"encoding/json"
	"errors"

	"github.com/jmoiron/sqlx"
)

// Adapter Users repository
type userRepository struct {
	psqlDB *sqlx.DB
}

// Constructor
func NewUserRepository(psqlDB *sqlx.DB) user.IUserRepository {
	return &userRepository{psqlDB: psqlDB}
}

// Method
func (u *userRepository) FetchUsers(ctx context.Context) ([]*models.User, error) {
	scriptSQL := `
    SELECT
      COALESCE(array_to_json(array_agg("us")), '[]'::json)
    FROM (
      SELECT
        users.id,
        users.username,
        users.password,
        users.email,
        users.created_at,
        users.updated_at,
        (
          SELECT
            COALESCE(array_to_json(array_agg("im")), '[]'::json)
          FROM (
            SELECT
              images.id,
              images.user_id,
              images.filename,
              images.url,
              images.created_at,
              images.updated_at
            FROM
              images
            WHERE
              images.user_id = users.id
          ) AS "im"
        ) AS "images"
      FROM
        users
    ) AS "us"
  `
	users := make([]*models.User, 0)
	stmt, err := u.psqlDB.PreparexContext(ctx, scriptSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var userInstant []byte
	if err = stmt.GetContext(ctx, &userInstant); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(userInstant, &users); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("no users found")
	}

	return users, nil
}
