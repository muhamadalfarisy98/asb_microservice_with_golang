package main

import (
	"asb_microservice_with_golang/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	// -- routing --

	router.GET("/books/price", controllers.QueryHitungPrice)

	// routes books
	router.GET("/books", controllers.GetBooks)
	router.GET("/categories/:id/books", controllers.GetBooksByIDCategory)
	router.POST("/books", controllers.PostBooks)
	router.PUT("/books/:id", (controllers.UpdateBooks))
	router.DELETE("/books/:id", controllers.DeleteBooks)

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
