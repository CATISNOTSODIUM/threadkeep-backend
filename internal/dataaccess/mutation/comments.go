package mutation

import (
	"context"

	"github.com/CATISNOTSODIUM/taggy-backend/internal/dataaccess/query"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/database"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/models"
	"github.com/CATISNOTSODIUM/taggy-backend/prisma/db"
)

func CreateComment(currentDB * database.Database, userID string, threadID string, content string) (* models.Comment, error) {
	ctx := context.Background()
	commentObject, err := currentDB.Client.Comment.CreateOne(
		db.Comment.Content.Set(content),
		db.Comment.Parent.Link(db.Thread.ID.Equals(threadID),),
		db.Comment.User.Link(db.User.ID.Equals(userID),),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	user, err := query.GetUserByID(currentDB, commentObject.UserID)

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
	return &comment, nil
}
