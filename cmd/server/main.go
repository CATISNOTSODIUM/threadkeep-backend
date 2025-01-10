package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/router"
	"github.com/joho/godotenv"
)

func loadPortNumber() int {
	err := godotenv.Load()
	if err != nil {
		return 8080 // default port number
	} 
	_portNumber := os.Getenv("PORT")
	portNumber, err := strconv.Atoi(_portNumber)
	
	if err != nil {
		return 8080 // default port number
	} 

	return portNumber
}

func main() {
	r := router.Setup()
	portNumber := strconv.Itoa(loadPortNumber())
	fmt.Print("Listening on port " + portNumber)
	
	log.Fatalln(http.ListenAndServe(":" + portNumber, r))
}
