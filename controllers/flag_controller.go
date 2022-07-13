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
	lookup := flagAudPopulate()

	return func(w http.ResponseWriter, r *http.Request) {
		cur, err := flagsCollection.Aggregate(ctx, mongo.Pipeline{lookup})
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

func flagAudPopulate() (lookup bson.D) {
	lookup = bson.D{
		{"$lookup", bson.D{
			{"from", "audiences"},
			{"localField", "audiences"},
			{"foreignField", "_id"},
			{"as", "aggregated"},
		}},
	}

	return lookup
}
