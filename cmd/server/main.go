package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/router"
)

func main() {
	r := router.Setup()
	fmt.Print("Listening on port 8000 at http://localhost:8000!")
	
	log.Fatalln(http.ListenAndServe(":8000", r))
}
