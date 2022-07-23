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

	err = h.DB.Delete(&models.Flag{}, id).Error
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

	err = h.DB.Delete(&models.Audience{}, id).Error
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

	err = h.DB.Delete(&models.Attribute{}, id).Error
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
