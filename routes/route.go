package routes

import (
	"github.com/gorilla/mux"
	"github.com/theshubhamy/cloverapi/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/getAllMovies", controllers.GetMovies).Methods("GET")
	router.HandleFunc("/api/createMovie", controllers.CreateMovies).Methods("POST")
	router.HandleFunc("/api/updateMovie/{id}", controllers.MarkedAsWatched).Methods("PUT")
	router.HandleFunc("/api/deleteMovie/{id}", controllers.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteAllMovies", controllers.DeleteAllMovie).Methods("DELETE")

	return router
}
