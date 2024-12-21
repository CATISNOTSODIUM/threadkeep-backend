package main

import (
	"fmt"
	"log"

	"github.com/CATISNOTSODIUM/taggy-backend/internal/dataaccess/mutation"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/database"
)

func main() {
	AddTags()
}


func AddTags() {

	db, err := database.Connect()
	if err != nil {
		log.Fatal("cannot retrieve database")
	}
	tags := []string{"Studies", "Romance", "Travel", "Discussion", "Rant", "Tech"}
    for _, tag := range tags {
		res, err := mutation.CreateTag(db, tag)
		if err != nil {
			log.Fatal("cannot create tag", err)
		} else {
			fmt.Println(res)
		}
	}
}
