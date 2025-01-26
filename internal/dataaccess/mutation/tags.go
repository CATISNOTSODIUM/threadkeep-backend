package mutation

import (
	"context"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/database"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/models"
	"github.com/CATISNOTSODIUM/threadkeep-backend/prisma/db"
)


func CreateTag(currentDB * database.Database, name string) (* models.Tag, error) {
	ctx := context.Background()
	userObject, err := currentDB.Client.Tag.CreateOne(
		db.Tag.Name.Set(name),
	).Exec(ctx)

	if err != nil {
    	return nil, err
  	}
 
	tag := &models.Tag {
		ID: userObject.ID,
		Name: userObject.Name,
	}
	
	return tag, nil
}