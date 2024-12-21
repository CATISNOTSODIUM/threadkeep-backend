package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/router"
)

func main() {
	r := router.Setup()
	fmt.Print("Listening on port 5000 at http://localhost:5000!")
	
	log.Fatalln(http.ListenAndServe(":5000", r))
}
