package mutation

import (
	"context"
	"github.com/CATISNOTSODIUM/taggy-backend/prisma/db"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/database"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/models"
)


func CreateUser(currentDB * database.Database, name string) (* models.User, error) {
	ctx := context.Background()

	userObject, err := currentDB.Client.User.CreateOne(
		db.User.Name.Set(name),
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