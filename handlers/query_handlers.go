package handlers

import (
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
	// TODO TODO TODO: *********************************************
	// need to populate the audiences with a _response_ object
	var flags []models.Flag

	// SELECT * FROM flags;
	result := h.DB.Find(&flags)

	if result.Error != nil {
		fmt.Println(result.Error)
	}
	
	var response []models.FlagNoAudsResponse
	for ind, _ := range flags {
		response = append(response, models.FlagNoAudsResponse{Flag: &flags[ind]})
	}

	utils.PayloadResponse(w, r, &response)
	// w.Header().Add("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(flags) // need a new target for w

	// ****~~~ CACHING WORKFLOW ****~~~
	// Flush cache 
	// flagCache.FlushAllAsync()
	// Store copy of data
	// flagCache.Set("data", value) // `value` needs to match struct, not sure what it will be
}

// result := h.DB.Preload("Audiences").Find(&flags)
func (h Handler) GetAllAudiences(w http.ResponseWriter, r *http.Request) {
	// TODO TODO TODO: *********************************************
	// need to populate the conds with a _response_ object (need to populate attrs?)
	// need also to populate flag id's?
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
	utils.HandleErr(err, "string conv went south")

	var flag models.Flag
	var auds []models.AudienceNoCondsResponse

	h.DB.Preload("Audiences").Find(&flag, id)
	for ind, _ := range flag.Audiences {
		auds = append(auds, models.AudienceNoCondsResponse{Audience: &flag.Audiences[ind]})
	}

	utils.PayloadResponse(w, r, &models.FlagResponse{Flag: &flag, Audiences: auds})
}

func (h Handler) GetAudience(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	utils.HandleErr(err, "string conv went south")

	var aud models.Audience
	var conds []models.ConditionEmbedded

	result := h.DB.Preload("Conditions").First(&aud, id)

	for ind, _ := range aud.Conditions {
		cond := aud.Conditions[ind]
		var attr models.Attribute
		h.DB.Find(&attr, cond.AttributeID)
		h.DB.Find(&cond)
		cond.Attribute = attr
		conds = append(conds, models.ConditionEmbedded{
			Condition: &cond,
			Attribute: models.AttributeEmbedded{Attribute: &attr},
		})
	}

	if result.Error != nil {
		utils.NoRecordResponse(w, r, result.Error)
		return
	}

	response := models.AudienceResponse{
		Audience:   &aud,
		Conditions: conds,
	}

	utils.PayloadResponse(w, r, &response)
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
