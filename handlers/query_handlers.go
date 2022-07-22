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
}

// result := h.DB.Preload("Audiences").Find(&flags)
func (h Handler) GetAllAudiences(w http.ResponseWriter, r *http.Request) {
	// TODO TODO TODO: *********************************************
	// need to populate the conds with a _response_ object (need to populate attrs?)
	// need also to populate flag id's?
	var auds []models.Audience
	var respAuds []models.AudienceResponse

	result := h.DB.Preload("Conditions").Find(&auds)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	for ind, _ := range auds {
		respAuds = append(respAuds, models.AudienceResponse{Audience: &auds[ind]})
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

	result := h.DB.Preload("Audiences").First(&flag, id)

	if result.Error != nil {
		utils.NoRecordResponse(w, r, result.Error)
		return
	}

	// Store flag in cache, flag id as the key
	fmt.Printf("\nAdding flag id %v to the cache:\n %v\n\n", vars["id"], flag)
	flagCache.Set(vars["id"], &flag) // THIS LINE IS PROBLEMATIC ? trying to marshal then add a type flag to the cache. 

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
