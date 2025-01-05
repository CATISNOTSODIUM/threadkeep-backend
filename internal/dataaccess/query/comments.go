package query

import (
	"context"
	"github.com/CATISNOTSODIUM/threadkeep-backend/prisma/db"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/database"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/models"
)


func GetCommentsByThreadID(currentDB * database.Database, threadID string) ([] * models.Comment, error) {
	ctx := context.Background()
	commentObjects, err := currentDB.Client.Comment.FindMany(db.Comment.ParentID.Equals(threadID)).Exec(ctx)

	if err != nil {
		return nil, err
	}

	comments := []*models.Comment{}
	for _, commentObject := range commentObjects {
		
		user, err := GetUserByID(currentDB, commentObject.UserID)

		if err != nil {
			return nil, err
		}

		comment := models.Comment {
			ID: commentObject.ID,
			Content: commentObject.Content,
			Likes: commentObject.Likes,
			Views: commentObject.Views,
			User: *user,
			ParentID: threadID,
			CreatedAt: commentObject.CreatedAt,
			UpdatedAt: commentObject.UpdatedAt,
		}	
		comments = append(comments, &comment)
	}
	if err != nil {
		return nil, err
	}
	

	return comments, nil
}
