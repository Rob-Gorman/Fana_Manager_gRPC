package handlers

import (
	"encoding/json"
	"manager/cache"
	"net/http"
)

func (h Handler) GetFlagset(w http.ResponseWriter, r *http.Request) {
	fs := BuildFlagset(h.DB)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&fs)

	flagCache := cache.InitFlagCache()
	flagCache.FlushAllAsync()
	flagCache.Set("data", &fs)
}
