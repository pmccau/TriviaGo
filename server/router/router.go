package router

import (
	"github.com/pmccau/TriviaGo/server/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/test", middleware.Test)
	router.HandleFunc("/api/getQuestions", middleware.GetQuestions)
	router.HandleFunc("/api/getCategories", middleware.GetCategories)
	return router
}
