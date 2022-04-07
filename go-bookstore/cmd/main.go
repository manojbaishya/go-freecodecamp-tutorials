package main

import (
	"fmt"
	"go-bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "go-bookstore/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Bookstore Application
// @version 1.0
// @description This is a Bookstore Application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9010
// @BasePath /
func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	router.PathPrefix("/documentation/").Handler(httpSwagger.WrapHandler)

	http.Handle("/", router)
	fmt.Println("Running server on http://localhost:9010.")
	log.Fatal(http.ListenAndServe("localhost:9010", router))
}
