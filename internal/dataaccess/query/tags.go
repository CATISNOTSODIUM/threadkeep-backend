package query

import (
	"context"

	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/database"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/models"
)


func GetTags(currentDB * database.Database) ([]*models.Tag, error) {
	ctx := context.Background()
	tagObjects, err := currentDB.Client.Tag.FindMany().Exec(ctx)

	if err != nil {
		return nil, err
	}

	tags := []*models.Tag{}
	for _, tagObject := range tagObjects {


		tag := models.Tag {
			ID: tagObject.ID,
			Name: tagObject.Name,
		}	
		tags = append(tags, &tag)
	}
	
	return tags, nil
}