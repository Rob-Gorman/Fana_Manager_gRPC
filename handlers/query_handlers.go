package handlers

import (
	"encoding/json"
	"fmt"
	"manager/models"
	"manager/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h Handler) GetAllFlags(w http.ResponseWriter, r *http.Request) {
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

	// checking conversion of []byte to string. i think it's p easy
	
		fmt.Printf("before marshal %v\n", result)
		payload, err := json.Marshal(&flags)
		if err != nil {
			panic(err)
		}
		fmt.Printf("after marshal %v\n", payload)
		// publisher.Pub.PublishTo("flag-toggle-channel", string(payload))

	utils.PayloadResponse(w, r, &response)
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
