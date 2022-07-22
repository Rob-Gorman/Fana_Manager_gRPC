package handlers

import (
	"manager/api"
	"manager/utils"
	"net/http"
)

func (h Handler) GetFlagset(w http.ResponseWriter, r *http.Request) {
	fs := api.BuildFlagset(h.DB)
	utils.PayloadResponse(w, r, &fs)
}
