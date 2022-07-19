package api

import (
	"manager/db"
	"manager/handlers"

	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router // this is our express-router
	H           handlers.Handler
	// this Handler is a wrapper around our DB to allow us to define methods on it
	// those methods being our controller functions (handlers)
}

func NewServer() *Server {
	s := &Server{
		H:      handlers.New(db.Init()),
		Router: mux.NewRouter(),
	}

	s.routes()
	return s
}

func (s *Server) routes() {
	s.dashboardRoutes()
	// s.providerRoutes()
}

// func (s *Server) getMongoData() http.HandlerFunc {
// 	query := configs.ConnectClient()
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		if err := json.NewEncoder(w).Encode(query); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 	}
// }

// func (s *Server) getFlagData() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		if err := json.NewEncoder(w).Encode(*(s.Ruleset)); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 	}
// }

// func GetRuleset(db *mongo.Database) (results []bson.M) {
// 	ctx, _ := utils.StandardContext()
// 	flags := db.Collection("flags")
// 	query, err := flags.Find(ctx, bson.D{})
// 	utils.HandleErr(err, "Cannot query the collection")

// 	query.All(ctx, &results)
// 	return results
// }
