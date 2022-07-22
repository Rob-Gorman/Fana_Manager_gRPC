package handlers

import (
	"encoding/json"
	"fmt"
	"manager/models"
	"manager/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"manager/cache"
)

var (
	flagCache cache.FlagCache = cache.NewRedisCache("localhost:6379", 0, 1000000)
)

func (h Handler) GetAllFlags(w http.ResponseWriter, r *http.Request) {
	var flags []models.Flag

	result := h.DB.Find(&flags)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	response := []models.FlagNoAudsResponse{}

	for ind, _ := range flags {
		response = append(response, models.FlagNoAudsResponse{Flag: &flags[ind]})
	}

	// checking conversion of []byte to string. i think it's p easy

	fmt.Printf("before marshal %v\n", result)
	payload, err := json.Marshal(&flags)
	if err != nil {
		panic(err)
	}
	fmt.Printf("after marshal %v\n", payload)
	// publisher.Pub.PublishTo("flag-toggle-channel", string(payload))

	utils.PayloadResponse(w, r, &response)

	// ****~~~ CACHING WORKFLOW ****~~~
	// Flush cache 
	// flagCache.FlushAllAsync()
	// Store copy of data
	// flagCache.Set("data", value) // `value` needs to match struct, not sure what it will be

}

func (h Handler) GetAllAudiences(w http.ResponseWriter, r *http.Request) {
	var auds []models.Audience
	var respAuds []models.AudienceNoCondsResponse

	result := h.DB.Preload("Conditions").Find(&auds)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	for ind, _ := range auds {
		respAuds = append(respAuds, models.AudienceNoCondsResponse{Audience: &auds[ind]})
	}

	utils.PayloadResponse(w, r, respAuds)
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
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	var flag models.Flag
	auds := []models.AudienceNoCondsResponse{}

	h.DB.Preload("Audiences").Find(&flag, id)
	for ind, _ := range flag.Audiences {
		auds = append(auds, models.AudienceNoCondsResponse{Audience: &flag.Audiences[ind]})
	}

	utils.PayloadResponse(w, r, &models.FlagResponse{Flag: &flag, Audiences: auds})
}

func (h Handler) GetAudience(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	var aud models.Audience

	result := h.DB.Preload("Conditions").First(&aud, id)

	if result.Error != nil {
		utils.NoRecordResponse(w, r, result.Error)
		return
	}

	conds := GetEmbeddedConds(aud, h.DB)

	response := models.AudienceResponse{
		Audience:   &aud,
		Conditions: conds,
	}

	utils.PayloadResponse(w, r, &response)
}

func (h Handler) GetAttribute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	var attr models.Attribute

	result := h.DB.First(&attr, id)

	if result.Error != nil {
		utils.NoRecordResponse(w, r, result.Error)
		return
	}

	utils.PayloadResponse(w, r, attr)
}
