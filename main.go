package main

import (
	"Shubham/Routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	Routes.UserRoutes(r)
	Routes.ProductRoutes(r)
	Routes.CartRoutes(r)
	Routes.WishRoutes(r)
	fmt.Println("Server is Started")
	log.Fatal(http.ListenAndServe(":9006", r))
}
