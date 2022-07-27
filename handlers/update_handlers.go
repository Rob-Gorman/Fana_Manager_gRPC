package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"manager/models"
	"manager/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// var ctx = utils.StandardContext()

func (h Handler) UpdateFlag(w http.ResponseWriter, r *http.Request) {
	var flagReq models.FlagSubmit

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &flagReq)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid resource ID."))
		return
	}

	fr := FlagReqToFlag(flagReq, h)
	var flag models.Flag
	h.DB.First(&flag, id)
	flag.Audiences = fr.Audiences
	flag.DisplayName = fr.DisplayName
	flag.Sdkkey = fr.Sdkkey

	if flagReq.Audiences != nil {
		h.DB.Model(&flag).Association("Audiences").Replace(flag.Audiences)
	}

	err = h.DB.Model(&flag).Session(&gorm.Session{
		FullSaveAssociations: true,
		SkipHooks:            true,
	}).Updates(&flag).Error

	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	response := FlagToFlagResponse(flag, h)

	pub := FlagUpdateForPublisher(h.DB, []models.Flag{flag})
	PublishContent(&pub, "flag-update-channel")
	RefreshCache(h.DB)

	utils.UpdatedResponse(w, r, &response)
}

func (h Handler) ToggleFlag(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nGot a flag toggle!")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid resource ID."))
		return
	}

	togglef := struct {
		Status bool `json:"status"`
	}{}

	body, _ := ioutil.ReadAll(r.Body)

	err = json.Unmarshal(body, &togglef)

	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	fmt.Println("Looking for flag...")
	var flag models.Flag
	h.DB.Find(&flag, id)
	fmt.Println("done!")
	flag.Status = togglef.Status
	flag.DisplayName = fmt.Sprintf("__%v", flag.Status) // hacky way to clue it's a toggle action, see flag update hook
	fmt.Println("Updating flag status...")
	err = h.DB.Select("status").Updates(&flag).Error
	if err != nil {
		utils.NoRecordResponse(w, r, err)
		return
	}
	fmt.Println("done!")
	
	fmt.Println("Retrieving flag again...")
	h.DB.First(&flag, id)
	response := models.FlagNoAudsResponse{Flag: &flag}
	fmt.Println("done!")
	
	fmt.Println("Making pub/sub message...")
	pub := FlagUpdateForPublisher(h.DB, []models.Flag{flag})
	fmt.Println("done!")
	fmt.Println("Publishing message...")
	PublishContent(&pub, "flag-toggle-channel")
	fmt.Println("done!")

	fmt.Println("refreshing cache...")
	RefreshCache(h.DB)
	fmt.Println("done refreshing cache!")
	
	fmt.Println("WRiting response...")
	utils.UpdatedResponse(w, r, &response)
	fmt.Println("Sent response!")
}

func (h Handler) UpdateAudience(w http.ResponseWriter, r *http.Request) {
	var req models.Audience

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid resource ID."))
		return
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	aud := BuildAudUpdate(req, id, h)

	if req.Conditions != nil {
		h.DB.Model(&aud).Association("Conditions").Replace(aud.Conditions)
	}

	err = h.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&aud).Error
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	h.DB.Model(&models.Audience{}).Preload("Flags").Preload("Conditions").Find(&aud)

	response := models.AudienceResponse{
		Audience:   &aud,
		Conditions: GetEmbeddedConds(aud, h.DB),
		Flags:      GetEmbeddedFlags(aud.Flags),
	}

	pub := FlagUpdateForPublisher(h.DB, aud.Flags)
	PublishContent(&pub, "flag-update-channel")
	RefreshCache(h.DB)

	utils.CreatedResponse(w, r, &response)
}
