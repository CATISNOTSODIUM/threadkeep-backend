package mutation

import (
	"context"
	"github.com/CATISNOTSODIUM/taggy-backend/prisma/db"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/database"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/models"
)

func CreateComment(currentDB * database.Database, user * models.User, thread * models.Thread, content string) (* models.Comment, error) {
	ctx := context.Background()
	commentObject, err := currentDB.Client.Comment.CreateOne(
		db.Comment.Content.Set(content),
		db.Comment.Parent.Link(db.Thread.ID.Equals(thread.ID),),
		db.Comment.User.Link(db.User.ID.Equals(user.ID),),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	comment := models.Comment {
		ID: commentObject.ID,
		Content: commentObject.Content,
		Likes: commentObject.Likes,
		Views: commentObject.Views,
		User: *user,
		Parent: *thread,
		CreatedAt: commentObject.CreatedAt,
    	UpdatedAt: commentObject.UpdatedAt,
	}
	return &comment, nil
}
