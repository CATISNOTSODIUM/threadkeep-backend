package users

import (
	"github.com/CATISNOTSODIUM/taggy-backend/internal/database"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/models"
)

func List(db *database.Database) ([]models.User, error) {
	users := []models.User{
		{
			ID:   1,
			Name: "CVWO",
		},
	}
	return users, nil
}
