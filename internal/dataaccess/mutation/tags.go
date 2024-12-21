package mutation

import (
	"context"

	"github.com/CATISNOTSODIUM/taggy-backend/internal/database"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/models"
	"github.com/CATISNOTSODIUM/taggy-backend/prisma/db"
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