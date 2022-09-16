package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"manager/cache"
	"manager/models"
	"manager/publisher"
	"manager/utils"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

func PublishContent(data interface{}, channel string) {
	if publisher.Redis == nil {
		utils.HandleErr(nil, "No connection established with Redis; cannot publish")
		return
	}

	byteArray, err := json.Marshal(data)
	if err != nil {
		utils.HandleErr(err, "Unmarshalling error")
	}
	fmt.Println("Manager trying to publish (helpers.go)", string(byteArray))

	err = publisher.Redis.Publish(context.TODO(), channel, byteArray).Err()
	if err != nil {
		utils.HandleErr(err, " : Error trying to publish to redis")
		return
	}
}

func RefreshCache(db *gorm.DB) {
	flagCache := cache.InitFlagCache()
	fs := BuildFlagset(db)
	flagCache.FlushAllAsync()
	flagCache.Set("data", &fs)
}

func OrphanedAud(aud *models.Audience) bool {
	return len((*aud).Flags) == 0
}

func OrphanedAttr(attr *models.Attribute, h *Handler) bool {
	asscs := h.DB.Model(attr).Association("Conditions").Count()
	return asscs == 0
}

func NewSDKKey(s string) string {
	digits := []byte("0123456789abcdefghijkm")
	newKey := []byte{}
	rand.Seed(time.Now().UnixNano())

	for _, char := range s {
		if char == '-' {
			newKey = append(newKey, '-')
		} else {
			randInd := rand.Intn(len(digits))
			newKey = append(newKey, digits[randInd])
		}
	}

	return string(newKey)
}
