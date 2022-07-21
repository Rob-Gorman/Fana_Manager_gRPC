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
	// TAKES AUDIENCE KEYS; NOT ID'S
	type flagPost struct {
		Key         string   `json:"key"`
		DisplayName string   `json:"displayName"`
		Sdkkey      string   `json:"sdkKey"`
		Audiences   []string `json:"audiences,omitempty"`
	}

	var auds []models.Audience
	var flagReq flagPost

	// Read to request body
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	// this translates the body into the flagPost form
	// using the json tags from the struct definition
	err := json.Unmarshal(body, &flagReq)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	h.DB.Where("key in ?", flagReq.Audiences).Find(&auds, flagReq.Audiences)

	flag := models.Flag{
		Audiences:   auds,
		Key:         flagReq.Key,
		DisplayName: flagReq.DisplayName,
		Sdkkey:      flagReq.Sdkkey,
	}
	flag.Audiences = auds
	flag.Key, flag.DisplayName = utils.ProcessNameToKeyDisplayName(flagReq.Key)
	flag.Sdkkey = flagReq.Sdkkey

	fmt.Printf("sdkkey req: %s\nsdkkey object: %s\n", flagReq.Sdkkey, flag.Sdkkey)

	// Append to the Flags table
	// result := h.DB.Preload("Audiences").Create(&flag)
	result := h.DB.Save(&flag)

	if result.Error != nil {
		utils.UnavailableResponse(w, r, err)
		return
	}

	response := models.FlagResponse{Flag: &flag}

	// Send a 201 created response
	utils.PayloadResponse(w, r, &response)
}

func (h Handler) CreateAttribute(w http.ResponseWriter, r *http.Request) {
	type attrPost struct {
		Key  string `json:"key"`
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
	attr.Key, attr.DisplayName = utils.ProcessNameToKeyDisplayName(attrReq.Key)
	attr.Type = attrReq.Type
	h.DB.Save(&attr)

	utils.CreatedResponse(w, r, &attr)
}

func (h Handler) CreateAudience(w http.ResponseWriter, r *http.Request) {
	type condPost struct {
		AttributeID uint   `json:"attributeID"`
		Operator    string `json:"operator"`
		Vals        string `json:"vals"`
	}

	type audPost struct {
		Key        string     `json:"key"`
		Combine    string     `json:"combine"`
		Conditions []condPost `json:"conditions"`
	}

	var audReq audPost
	var aud models.Audience
	// var conds []models.Condition

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.HandleErr(err, "should put a bad request error here")
		return
	}

	err = json.Unmarshal(body, &audReq)
	utils.HandleErr(err, "problem unmarshalling, what do?")

	// var aud models.Attribute
	aud.Key, aud.DisplayName = utils.ProcessNameToKeyDisplayName(aud.Key)
	// aud.Type = audReq.Type
	h.DB.Model(&aud).Save(&audReq)

	// utils.CreatedResponse(w, r, &models.AudienceResponse{Audience: &audReq})
	utils.CreatedResponse(w, r, &audReq)
}
