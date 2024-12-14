package query

import (
	"context"
	"github.com/CATISNOTSODIUM/taggy-backend/prisma/db"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/database"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/models"
)

func GetThreads(currentDB * database.Database) ([]*models.Thread, error) {
	ctx := context.Background()
	threadObjects, err := currentDB.Client.Thread.FindMany().Exec(ctx)

	if err != nil {
		return nil, err
	}

	threads := []*models.Thread{}
	for _, threadObject := range threadObjects {
		user, err := GetUserByID(currentDB, threadObject.UserID)

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
		threads = append(threads, &thread)
	}
	
	return threads, nil
}

func GetThreadByID(currentDB * database.Database, id string) (* models.Thread, error) {
	ctx := context.Background()
	threadObject, err := currentDB.Client.Thread.FindUnique(db.Thread.ID.Equals(id)).Exec(ctx)

	if err != nil {
		return nil, err
	}

	user, err := GetUserByID(currentDB, threadObject.UserID)

	if err != nil {
		return nil, err
	}
	
	thread := &models.Thread {
		ID: threadObject.ID,
		Title: threadObject.Title,
		Content: threadObject.Content,
		Likes: threadObject.Likes,
		Views: threadObject.Views,
		User: *user,
		CreatedAt: threadObject.CreatedAt,
    	UpdatedAt: threadObject.UpdatedAt,
	}
	return thread, nil
}
