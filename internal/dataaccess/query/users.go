package query

import (
	"context"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/database"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/models"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/utils"
	"github.com/CATISNOTSODIUM/threadkeep-backend/prisma/db"
)


func GetUsers(currentDB * database.Database) ([]*models.User, error) {
	ctx := context.Background()
	userObjects, err := currentDB.Client.User.FindMany().Exec(ctx)

	if err != nil {
		return nil, err
	}

	users := []*models.User{}
	for _, userObject := range userObjects {
		users = append(users, &models.User {
			ID: userObject.ID,
			Name: userObject.Name,
		})
	}
	return users, nil
}

func GetUserByID(currentDB * database.Database, id string) (* models.User, error) {
	ctx := context.Background()
	userObject, err := currentDB.Client.User.FindUnique(db.User.ID.Equals(id)).Exec(ctx)

	if err != nil {
		return nil, err
	}

	user := &models.User {
		ID: userObject.ID,
		Name: userObject.Name,
	}
	return user, nil
}

func VerifyUser(currentDB * database.Database, name string, password string) (* models.User, bool) {
	ctx := context.Background()
	userObject, err := currentDB.Client.User.FindUnique(db.User.Name.Equals(name)).Exec(ctx)
	if err != nil {
		return nil, false
	}
	if (password == string(utils.Decode(userObject.Password))) {
		// verified
		JWTToken, err := utils.CreateToken(name)
		if err != nil {
			return nil, false
		}

		return &models.User {
			ID: userObject.ID,
			Name: userObject.Name,
			JWTToken: JWTToken,
		}, true
	}
	return nil, false
}
