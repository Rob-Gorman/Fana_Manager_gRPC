package api

import (
	"sovereign/configs"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	*mux.Router
	DB *mongo.Client
}

func NewServer() *Server {
	s := &Server{
		DB:     configs.ConnectClient(),
		Router: mux.NewRouter(),
	}

	s.routes()
	return s
}

func (s *Server) routes() {
	s.FlagRoutes()
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
