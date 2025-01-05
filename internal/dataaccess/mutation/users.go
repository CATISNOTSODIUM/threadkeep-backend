package mutation

import (
	"context"

	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/database"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/models"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/utils"
	"github.com/CATISNOTSODIUM/threadkeep-backend/prisma/db"
)


func CreateUser(currentDB * database.Database, name string, password string) (* models.User, error) {
	ctx := context.Background()
	userObject, err := currentDB.Client.User.CreateOne(
		db.User.Name.Set(name),
		db.User.Password.Set(utils.Encode([]byte(password))), 
	).Exec(ctx)

	if err != nil {
    	return nil, err
  	}
 
	user := &models.User {
		ID: userObject.ID,
		Name: userObject.Name,
	}
	
	return user, nil
}