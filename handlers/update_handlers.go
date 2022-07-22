package handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"manager/models"
	"manager/publisher"
	"manager/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func (h Handler) UpdateFlag(w http.ResponseWriter, r *http.Request) {
	var flagReq models.FlagSubmit

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &flagReq)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	fr := FlagReqToFlag(flagReq, h)
	var flag models.Flag
	h.DB.First(&flag, id)
	flag.Audiences = fr.Audiences
	flag.DisplayName = fr.DisplayName
	flag.Key = fr.Key
	flag.Sdkkey = fr.Sdkkey

	h.DB.Model(&flag).Association("Audiences").Replace(flag.Audiences)
	err = h.DB.Model(&flag).Session(&gorm.Session{FullSaveAssociations: true}).Updates(&flag).Error

	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	response := FlagToFlagResponse(flag, h)
  
		byteArray, err := json.Marshal(&response)
		if err != nil {
			utils.HandleErr(err, "our unmarshalling sucks")
		}

		publisher.Redis.Publish(context.TODO(), "flag-update-channel", byteArray)

	utils.UpdatedResponse(w, r, &response)
}

func (h Handler) ToggleFlag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	togglef := struct {
		Status bool `json:"status"`
	}{}

	body, _ := ioutil.ReadAll(r.Body)

	err = json.Unmarshal(body, &togglef)

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

	byteArray, err := json.Marshal(&response)
	if err != nil {
		utils.HandleErr(err, "our unmarshalling sucks")
	}

	publisher.Redis.Publish(context.TODO(), "flag-toggle-channel", byteArray)

	utils.UpdatedResponse(w, r, &response)
}

func (h Handler) UpdateAudience(w http.ResponseWriter, r *http.Request) {
	var req models.Audience

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	aud := BuildAudUpdate(req, id, h)

	h.DB.Model(&aud).Association("Conditions").Replace(aud.Conditions)
	err = h.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&aud).Error
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	h.DB.Model(&models.Audience{}).Preload("Conditions").Find(&aud)

	response := models.AudienceResponse{
		Audience:   &aud,
		Conditions: GetEmbeddedConds(aud, h.DB),
	}

	utils.CreatedResponse(w, r, &response)
}
