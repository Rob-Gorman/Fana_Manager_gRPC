package controllers

import (
	"encoding/json"
	"net/http"
	"sovereign/configs"
	"sovereign/responses"
	"sovereign/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetFlags(db *mongo.Client) http.HandlerFunc {
	flagsCollection := configs.GetCollection(db, "flags")
	ctx, _ := utils.StandardContext()

	return func(w http.ResponseWriter, r *http.Request) {
		cur, err := flagsCollection.Find(ctx, bson.D{})
		utils.HandleErr(err, "Cannot find Flags Collection")

		var flags []responses.FlagResponse
		cur.All(ctx, &flags)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(flags); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
