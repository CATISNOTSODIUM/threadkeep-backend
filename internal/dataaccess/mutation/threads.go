package mutation

import (
	"context"
	"github.com/CATISNOTSODIUM/taggy-backend/prisma/db"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/database"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/models"
)

func CreateThread(currentDB * database.Database, user * models.User, title string, content string) (* models.Thread, error) {
	ctx := context.Background()
	threadObject, err := currentDB.Client.Thread.CreateOne(
		db.Thread.Title.Set(title),
		db.Thread.Content.Set(content),
		db.Thread.User.Link(
			db.User.ID.Equals(user.ID),
		),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	thread := models.Thread {
		ID: threadObject.ID,
		Title: threadObject.Title,
		Content: threadObject.Content,
		Likes: threadObject.Likes,
		Views: threadObject.Views,
		User: *user,
		CreatedAt: threadObject.CreatedAt,
    	UpdatedAt: threadObject.UpdatedAt,
	}
	return &thread, nil
}
