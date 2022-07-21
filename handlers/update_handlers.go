package handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"manager/models"
	"manager/publisher"
	"manager/utils"
	"net/http"
)

func (h Handler) UpdateFlag(w http.ResponseWriter, r *http.Request) {

	// shape of request payload
	// the JSON tags identify what part of the incoming payload
	// to assign to the field in the `json.Unmarshal` method
	type flagUpdate struct {
		Name   string `json:"name"`
		SdkKey string `json:"sdkKey"`
		Status bool   `json:"status"`
		// Audiences []string `json:"audiences"`
	}

	var flagReq flagUpdate
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.HandleErr(err, "should put a bad request error here")
		return
	}

	err = json.Unmarshal(body, &flagReq)
	utils.HandleErr(err, "our unmarshalling sucks")

	// get audience objects to insert join reference
	// (GORM model for flags expects Audience objects, not key strings)
	// var dbAuds []models.Audience
	// h.DB.Where("key in ?", flagReq.Audiences).Find(&dbAuds)

	// h.DB.Association("Audiences")
	var updatedFlag models.Flag

	// update the table `updatedFlag` belongs to with the
	// fields existent in flagReq object (these have to map accurately)
	result := h.DB.Model(&updatedFlag).Updates(&flagReq)

	if result.Error != nil {
		utils.HandleErr(result.Error, "should put a failed to update")
		return
	}

	byteArray, err := json.MarshalIndent(&updatedFlag, "", "  ")
	if err != nil {
		utils.HandleErr(err, "our unmarshalling sucks")
	}
	publisher.Redis.Publish(context.TODO(), "flag-update-channel", byteArray)

	// Send a 201 created response
	utils.UpdatedResponse(w, r, &updatedFlag)
}
