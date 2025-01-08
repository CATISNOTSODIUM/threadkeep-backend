package query

import (
	"context"
	"strconv"

	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/database"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/models"
	"github.com/CATISNOTSODIUM/threadkeep-backend/prisma/db"
)


func GetThreads(currentDB * database.Database, skip int, max_per_page int, name string, tags []string) ([]*models.Thread, error) {
	ctx := context.Background()


	var TagQuery []db.ThreadWhereParam
	if (name != "") {
		TagQuery = append(TagQuery, db.Thread.Title.Contains(name))
		TagQuery = append(TagQuery, db.Thread.Title.Mode(db.QueryModeInsensitive))
	} 

	if (len((tags)) > 0) {
		var TagIDQuery []db.TagsOnThreadsWhereParam
		for _, tag := range tags {
			TagIDQuery = append(TagIDQuery, db.TagsOnThreads.And(db.TagsOnThreads.TagID.Equals(tag)))
		}
		TagQuery = append(TagQuery, db.Thread.Tags.Some(db.TagsOnThreads.Or(TagIDQuery...))) 
	}


	filteredObject := currentDB.Client.Thread.FindMany()
	if (len(TagQuery) > 0) {
		filteredObject =  currentDB.Client.Thread.FindMany(TagQuery...)
	}
	threadObjects, err := filteredObject.Skip(skip).Take(max_per_page).OrderBy(
		db.Thread.CreatedAt.Order(db.SortOrderDesc),
	).With(
		db.Thread.Tags.Fetch().With( 
			db.TagsOnThreads.Tag.Fetch(), // get tag name
		),
	).With(
		db.Thread.User.Fetch(),
	).With(
		db.Thread.Likes.Fetch(),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}


	threads := []*models.Thread{}
	for _, threadObject := range threadObjects {
		userObject := threadObject.User()
		user := models.User {
			Name: userObject.Name,
			ID: userObject.ID,
		}

		tags := [] models.Tag{}

		tagObjects := threadObject.Tags()
		for _, tagObject := range tagObjects { 
			tag := models.Tag {
				ID: tagObject.TagID,
				Name: tagObject.Tag().Name,
			}	
			tags = append(tags, tag)
		} 

		thread := models.Thread {
			ID: threadObject.ID,
			Title: threadObject.Title,
			Content: threadObject.Content,
			Likes: len(threadObject.Likes()),
			Views: threadObject.Views,
			User: user,
			Tags: tags,
			CreatedAt: threadObject.CreatedAt,
			UpdatedAt: threadObject.UpdatedAt,
		}	
		threads = append(threads, &thread)
	}
	
	return threads, nil
}

func GetSavedThreads(currentDB * database.Database, userID string) ([]*models.Thread, error) {
	ctx := context.Background()
	savedObject, err := currentDB.Client.Saved.FindMany(
		db.Saved.UserID.Equals(userID),
	).With(
		db.Saved.Thread.Fetch(),
	).Exec(ctx)
	
	if err != nil {
		return nil, err
	}


	threads := []*models.Thread{}
	for _, savedObject := range savedObject {
		threadObject := savedObject.Thread()
		// many fields have been omitted
		thread := models.Thread {
			ID: threadObject.ID,
			Title: threadObject.Title,
			Content: threadObject.Content,
		}	
		threads = append(threads, &thread)
	}
	
	return threads, nil
}

func CountThreads(currentDB * database.Database) (int, error) {
	ctx := context.Background()
	var res []struct{
		NumberOfThreads  string     `json:"numberOfThreads"`
	}
	err := currentDB.Client.Prisma.QueryRaw(`SELECT count(*) as numberOfThreads FROM "Thread"`).Exec(ctx, &res)
	if err != nil {
		return -1, err
	}
	count, err := strconv.Atoi(res[0].NumberOfThreads)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func GetThreadByID(currentDB * database.Database, id string) (* models.Thread, error) {
	ctx := context.Background()
	threadObject, err := currentDB.Client.Thread.FindUnique(db.Thread.ID.Equals(id)).With(
		db.Thread.Tags.Fetch().With( 
			db.TagsOnThreads.Tag.Fetch(), // get tag name
		),
	).With(
		db.Thread.User.Fetch(),
	).With(
		db.Thread.Likes.Fetch(),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	userObject := threadObject.User()
	user := models.User {
		Name: userObject.Name,
		ID: userObject.ID,
	}

	tags := [] models.Tag{}
	tagObjects := threadObject.Tags()
	for _, tagObject := range tagObjects { 
		tag := models.Tag {
			ID: tagObject.TagID,
			Name: tagObject.Tag().Name,
		}	
		tags = append(tags, tag)
	} 

	thread := &models.Thread {
		ID: threadObject.ID,
		Title: threadObject.Title,
		Content: threadObject.Content,
		Likes: len(threadObject.Likes()),
		Views: threadObject.Views,
		User: user,
		Tags: tags,
		CreatedAt: threadObject.CreatedAt,
    	UpdatedAt: threadObject.UpdatedAt,
	}
	return thread, nil
}

func GetThreadTagsByID(currentDB * database.Database, id string) ([] models.Tag, error) {
	ctx := context.Background()
	tagOnThreadObjects, err := currentDB.Client.TagsOnThreads.FindMany(
		db.TagsOnThreads.ThreadID.Equals(id),
	).With(
		db.TagsOnThreads.Tag.Fetch(),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	tags := [] models.Tag{}
	
	for _, tagOnThreadObject := range tagOnThreadObjects {

		tag := models.Tag {
			ID: tagOnThreadObject.Tag().ID,
			Name: tagOnThreadObject.Tag().Name,
		}	
		tags = append(tags, tag)
	}

	return tags, nil
}

// check if userID has liked the thread
func IsLikeThread(currentDB * database.Database, userID string, threadID string) (bool, error) {
	ctx := context.Background()
	var res []struct {
		IsLike db.RawBoolean `json:"is_like"`
	}
	err := currentDB.Client.Prisma.QueryRaw(
		`SELECT EXISTS (
			SELECT * FROM "Likes" 
			WHERE "userID" = $1
			AND   "threadID" = $2
		) AS "is_like"`, userID, threadID, 
	).Exec(ctx, &res)

	if err != nil {
		return false, err
	}
	return bool(res[0].IsLike), err
}