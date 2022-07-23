package handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"manager/cache"
	"manager/models"
	"manager/publisher"
	"manager/utils"
	"net/http"

	"gorm.io/gorm"
)

func (h Handler) CreateFlag(w http.ResponseWriter, r *http.Request) {
	var flagReq models.FlagSubmit

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &flagReq)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	flag := FlagReqToFlag(flagReq, h)

	err = h.DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(&flag).Error

	if err != nil {
		utils.UnprocessableEntityResponse(w, r, err)
		return
	}

	h.DB.Preload("Audiences").Find(&flag)
	respAuds := []models.AudienceNoCondsResponse{}

	for ind, _ := range flag.Audiences {
		respAuds = append(respAuds, models.AudienceNoCondsResponse{
			Audience: &flag.Audiences[ind],
		})
	}

	response := models.FlagResponse{
		Flag:      &flag,
		Audiences: respAuds,
	}

	byteArray, err := json.Marshal(&response)
	if err != nil {
		utils.HandleErr(err, "Unmarshalling error")
	}
	publisher.Redis.Publish(context.TODO(), "flag-update-channel", byteArray)

	utils.CreatedResponse(w, r, &response)

	flagCache := cache.InitFlagCache()
	fs := BuildFlagset(h.DB)
	flagCache.FlushAllAsync()
	flagCache.Set("data", &fs)
}

func (h Handler) CreateAttribute(w http.ResponseWriter, r *http.Request) {
	var attrReq models.Attribute
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.HandleErr(err, "should put a bad request error here")
		return
	}

	err = json.Unmarshal(body, &attrReq)
	if err != nil {

		utils.BadRequestResponse(w, r, err)
		return
	}

	err = h.DB.Create(&attrReq).Error
	if err != nil {
		utils.UnprocessableEntityResponse(w, r, err)
		return
	}
	// attr := models.Attribute{}
	h.DB.Find(&attrReq)

	utils.CreatedResponse(w, r, &attrReq)

	flagCache := cache.InitFlagCache()
	fs := BuildFlagset(h.DB)
	flagCache.FlushAllAsync()
	flagCache.Set("data", &fs)
}

func (h Handler) CreateAudience(w http.ResponseWriter, r *http.Request) {
	var aud models.Audience

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	err = json.Unmarshal(body, &aud)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	err = h.DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(&aud).Error
	if err != nil {
		utils.UnprocessableEntityResponse(w, r, err)
		return
	}

	h.DB.Model(&models.Audience{}).Preload("Conditions").Find(&aud)

	response := models.AudienceResponse{
		Audience:   &aud,
		Conditions: GetEmbeddedConds(aud, h.DB),
	}

	byteArray, err := json.Marshal(&aud)
	if err != nil {
		utils.HandleErr(err, "Unmarshalling error")
	}

	publisher.Redis.Publish(context.TODO(), "audience-update-channel", byteArray)

	utils.CreatedResponse(w, r, &response)
	flagCache := cache.InitFlagCache()
	fs := BuildFlagset(h.DB)
	flagCache.FlushAllAsync()
	flagCache.Set("data", &fs)
}
