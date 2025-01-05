package database

import (
	"github.com/CATISNOTSODIUM/threadkeep-backend/prisma/db"
	"log"
)

type Database struct {
	Client * db.PrismaClient
}

func Connect() (* Database, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}
	return & Database {
		client,
	}, nil
}

func (currentDB Database) Close() {
	if err := currentDB.Client.Prisma.Disconnect(); err != nil {
      log.Fatalln(err)
  }
}


