package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"manager/models"
	"manager/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h Handler) GetAllFlags(w http.ResponseWriter, r *http.Request) {
	var flags []models.Flag

	result := h.DB.Find(&flags)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	utils.PayloadResponse(w, r, flags)
	// w.Header().Add("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(flags)
}

// result := h.DB.Preload("Audiences").Find(&flags)
func (h Handler) GetAllAudiences(w http.ResponseWriter, r *http.Request) {
	var auds []models.Audience

	result := h.DB.Preload("Conditions").Find(&auds)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	utils.PayloadResponse(w, r, auds)
}

func (h Handler) GetAllAttributes(w http.ResponseWriter, r *http.Request) {
	var attrs []models.Attribute

	result := h.DB.Find(&attrs)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	utils.PayloadResponse(w, r, attrs)
}

func (h Handler) GetFlag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	utils.HandleErr(err, "string conv went south")

	var flag models.Flag

	result := h.DB.Preload("Audiences").First(&flag, id)

	if result.Error != nil {
		utils.NoRecordResponse(w, r, result.Error)
		return
	}

	utils.PayloadResponse(w, r, flag)
}

func (h Handler) GetAudience(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	utils.HandleErr(err, "string conv went south")

	var aud models.Audience

	result := h.DB.Preload("Conditions").First(&aud, id)

	if result.Error != nil {
		utils.NoRecordResponse(w, r, result.Error)
		return
	}

	utils.PayloadResponse(w, r, aud)
}

func (h Handler) GetAttribute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	utils.HandleErr(err, "string conv went south")

	var attr models.Attribute

	result := h.DB.First(&attr, id)

	if result.Error != nil {
		utils.NoRecordResponse(w, r, result.Error)
		return
	}

	utils.PayloadResponse(w, r, attr)
}

func (h Handler) CreateFlag(w http.ResponseWriter, r *http.Request) {
	// shape of request payload
	// the JSON tags identify what part of the incoming payload
	// to assign to the field in the `json.Unmarshal` method
	type flagPost struct {
		Name      string   `json:"name"`
		SdkKey    string   `json:"sdkKey"`
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
	flag.SDKkey = flagReq.SdkKey

	// Append to the Flags table
	// result := h.DB.Preload("Audiences").Create(&flag)
	result := h.DB.Save(&flag)

	if result.Error != nil {
		utils.HandleErr(result.Error, "should put a failed to create")
		return
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&flag)
}
