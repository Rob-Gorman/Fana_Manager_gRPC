package handlers

import (
	"encoding/json"
	"io/ioutil"
	"manager/models"
	"manager/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func (h Handler) UpdateFlag(w http.ResponseWriter, r *http.Request) {
	// shape of request payload
	// the JSON tags identify what part of the incoming payload
	// to assign to the field in the `json.Unmarshal` method
	type flagUpdate struct {
		Key         string   `json:"key"`
		DisplayName string   `json:"displayName"`
		SdkKey      string   `json:"sdkKey"`
		Audiences   []string `json:"audiences"`
	}

	var flagReq flagUpdate
	var auds []models.Audience

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &flagReq)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	result := h.DB.Where("key in (?)", flagReq.Audiences).Find(&auds)
	if result.Error != nil {
		utils.BadRequestResponse(w, r, result.Error)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	utils.HandleErr(err, "string conv went south")

	var flag models.Flag
	h.DB.First(&flag, id)
	flag.Audiences = auds
	flag.DisplayName = flagReq.DisplayName
	flag.Key = flagReq.Key
	flag.Sdkkey = flagReq.SdkKey

	h.DB.Model(&flag).Association("Audiences")
	h.DB.Model(&flag).Session(&gorm.Session{FullSaveAssociations: true}).Updates(&flag)

	if result.Error != nil {
		utils.HandleErr(result.Error, "should put a failed to update")
		return
	}

	h.DB.Preload("Audiences").First(&flag, id)
	var respAuds []models.AudienceNoCondsResponse
	for ind, _ := range flag.Audiences {
		respAuds = append(respAuds, models.AudienceNoCondsResponse{Audience: &flag.Audiences[ind]})
	}
	response := models.FlagResponse{
		Flag:      &flag,
		Audiences: respAuds,
	}
	utils.UpdatedResponse(w, r, &response)
}

func (h Handler) ToggleFlag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	togglef := struct {
		Status bool `json:"status"`
	}{}

	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &togglef)

	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	var flag models.Flag
	update := map[string]interface{}{"status": togglef.Status}
	result := h.DB.Model(&flag).Where("id = ?", id).Updates(update)
	if result.Error != nil {
		utils.NoRecordResponse(w, r, result.Error)
		return
	}

	h.DB.First(&flag, id)
	response := models.FlagNoAudsResponse{Flag: &flag}

	utils.UpdatedResponse(w, r, &response)
}
