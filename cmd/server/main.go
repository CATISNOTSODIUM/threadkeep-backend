package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/router"
)

func main() {
	r := router.Setup()
	fmt.Print("Listening on port 8080")
	
	log.Fatalln(http.ListenAndServe(":8080", r))
}
