package handlers

import (
	"manager/utils"
	"net/http"
)

func (h Handler) GetFlagset(w http.ResponseWriter, r *http.Request) {
	fs := BuildFlagset(h.DB)
	utils.PayloadResponse(w, r, &fs)
}
