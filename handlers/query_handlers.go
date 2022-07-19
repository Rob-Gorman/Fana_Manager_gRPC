package handlers

import (
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
