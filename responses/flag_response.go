package responses

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlagResponse struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Key         string             `json:"key"`
	DisplayName string             `json:"displayName"`
	SDKKey      string             `json:"sdkKey"`
	Status      bool               `json:"status"`
	Audiences   []string           `json:"audiences,omitempty"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
	Aggregated  []bson.M           `json:"aggregated,omitempty"`
}
