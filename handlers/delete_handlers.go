package handlers

import (
	"manager/models"
	"manager/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h Handler) DeleteFlag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid flag ID."))
		return
	}

	flag := &models.Flag{}
	err = h.DB.Preload("Audiences").First(&flag, id).Error
	if err != nil {
		utils.NoRecordResponse(w, r, err)
		return
	}
	h.DB.Model(&flag).Association("Audiences").Delete(flag.Audiences)
	err = h.DB.Unscoped().Delete(&flag).Error
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) DeleteAudience(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid audience ID."))
		return
	}

	aud := &models.Audience{}
	err = h.DB.Preload("Flags").First(&aud, id).Error
	if err != nil {
		utils.NoRecordResponse(w, r, err)
		return
	}
	h.DB.Model(&aud).Association("Flags").Delete(aud.Flags)
	err = h.DB.Unscoped().Delete(&models.Audience{}, id).Error
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) DeleteAttribute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid attribute ID."))
		return
	}

	err = h.DB.Unscoped().Delete(&models.Attribute{}, id).Error
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
