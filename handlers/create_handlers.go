package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"manager/models"
	"manager/utils"
	"net/http"
)

func (h Handler) CreateFlag(w http.ResponseWriter, r *http.Request) {
	// shape of request payload
	// the JSON tags identify what part of the incoming payload
	// to assign to the field in the `json.Unmarshal` method
	type flagPost struct {
		Name      string   `json:"name"`
		Sdkkey    string   `json:"sdkKey"`
		Audiences []string `json:"audiences"`
	}

	var flagReq flagPost
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.HandleErr(err, "should put a bad request error here")
		return
	}

	// this translates the body into the flagPost form
	// using the json tags from the struct definition
	err = json.Unmarshal(body, &flagReq)
	utils.HandleErr(err, "our unmarshalling sucks")

	// get audience objects to insert join reference
	// (GORM model for flags expects Audience objects, not key strings)
	var dbAuds []models.Audience
	h.DB.Where("key in ?", flagReq.Audiences).Find(&dbAuds)

	// h.DB.Association("Audiences")
	var flag models.Flag
	flag.Audiences = dbAuds
	flag.Key, flag.DisplayName = utils.ProcessNameToKeyDisplayName(flagReq.Name)
	flag.Sdkkey = flagReq.Sdkkey

	fmt.Printf("sdkkey req: %s\nsdkkey object: %s\n", flagReq.Sdkkey, flag.Sdkkey)

	// Append to the Flags table
	// result := h.DB.Preload("Audiences").Create(&flag)
	result := h.DB.Save(&flag)

	if result.Error != nil {
		utils.HandleErr(result.Error, "should put a failed to create")
		return
	}

	response := models.FlagResponse{Flag: &flag}

	// Send a 201 created response
	utils.PayloadResponse(w, r, &response)
}

func (h Handler) CreateAttribute(w http.ResponseWriter, r *http.Request) {
	type attrPost struct {
		Name string `json:"name"`
		Type string `json:"attrType"`
	}

	var attrReq attrPost
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.HandleErr(err, "should put a bad request error here")
		return
	}

	err = json.Unmarshal(body, &attrReq)
	utils.HandleErr(err, "problem unmarshalling, what do?")

	var attr models.Attribute
	attr.Key, attr.DisplayName = utils.ProcessNameToKeyDisplayName(attrReq.Name)
	attr.Type = attrReq.Type
	h.DB.Save(&attr)

	utils.CreatedResponse(w, r, &attr)
}

// func (h Handler) CreateAudience(w http.ResponseWriter, r *http.Request) {
// 	type attrPost struct {
// 		Name string `json:"name"`
// 		Type string `json:"attrType"`
// 	}

// 	var attrReq attrPost
// 	defer r.Body.Close()
// 	body, err := ioutil.ReadAll(r.Body)

// 	if err != nil {
// 		utils.HandleErr(err, "should put a bad request error here")
// 		return
// 	}

// 	err = json.Unmarshal(body, &attrReq)
// 	utils.HandleErr(err, "problem unmarshalling, what do?")

// 	var attr models.Attribute
// 	attr.Key, attr.DisplayName = utils.ProcessNameToKeyDisplayName(attrReq.Name)
// 	attr.Type = attrReq.Type
// 	h.DB.Save(&attr)

// 	utils.CreatedResponse(w, r, &attr)
// }
