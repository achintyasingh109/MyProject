package Addition

import (
	"fmt"
)

func Sub(a int, b int) {
	c := b - a
	fmt.Println("The subtraction is ", c)

}

/*
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("book/{id}", controllers.DeleteBook).Methods("DELETE")
}
*/
