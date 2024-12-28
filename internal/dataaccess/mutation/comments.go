package mutation

import (
	"context"
	"time"

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


func UpdateComment(currentDB * database.Database, id string, content string) (* models.Comment, error) {
	ctx := context.Background()
	commentObjectDB := currentDB.Client.Comment.FindUnique(db.Comment.ID.Equals(id))
	commentObject, err := commentObjectDB.Update(
		db.Comment.Content.Set(content),
		db.Comment.UpdatedAt.Set(time.Now()),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	comment := models.Comment {
		ID: commentObject.ID,
		Content: commentObject.Content,
		CreatedAt: commentObject.CreatedAt,
    	UpdatedAt: commentObject.UpdatedAt,
	}
	return &comment, nil
}

func DeleteComment(currentDB * database.Database, commentID string) (int, error) {
	ctx := context.Background()
	res, err := 
		currentDB.Client.Prisma.ExecuteRaw(`
			DELETE FROM "Comment"
			WHERE 
				"id" = $1
		`, commentID).Exec(ctx)
	if err != nil {
		return 0, err
	}
	return res.Count, nil
}