package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"manager/models"
	"manager/utils"
	"net/http"

	"gorm.io/gorm"
)

func (h Handler) CreateFlag(w http.ResponseWriter, r *http.Request) {
	// TAKES AUDIENCE KEYS; NOT ID'S
	type flagPost struct {
		Key         string   `json:"key"`
		DisplayName string   `json:"displayName"`
		Sdkkey      string   `json:"sdkKey"`
		Audiences   []string `json:"audiences,omitempty"`
	}

	var auds []models.Audience
	var flagReq flagPost

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &flagReq)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	h.DB.Where("key in (?)", flagReq.Audiences).Find(&auds)

	flag := models.Flag{
		Audiences:   auds,
		Key:         flagReq.Key,
		DisplayName: flagReq.DisplayName,
		Sdkkey:      flagReq.Sdkkey,
	}
	result := h.DB.Save(&flag)

	if result.Error != nil {
		utils.UnavailableResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h Handler) CreateAttribute(w http.ResponseWriter, r *http.Request) {
	type attrPost struct {
		Key         string `json:"key"`
		DisplayName string `json:"displayName"`
		Type        string `json:"attrType"`
	}

	var attrReq attrPost
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.HandleErr(err, "should put a bad request error here")
		return
	}

	err = json.Unmarshal(body, &attrReq)
	utils.HandleErr(err, "problem unmarshalling, what do?")

	attr := models.Attribute{
		Key:         attrReq.Key,
		DisplayName: attrReq.DisplayName,
		Type:        attrReq.Type,
	}
	h.DB.Save(&attr)

	utils.CreatedResponse(w, r, &attr)
}

func (h Handler) CreateAudience(w http.ResponseWriter, r *http.Request) {
	type condPost struct {
		AttributeID uint   `json:"attributeID"`
		Operator    string `json:"operator"`
		Vals        string `json:"vals"`
		Negate      bool   `json:"negate"`
	}

	type audPost struct {
		Key         string     `json:"key"`
		DisplayName string     `json:"displayName"`
		Combine     string     `json:"combine"`
		Conditions  []condPost `json:"conditions"`
	}

	var audReq audPost
	var aud models.Audience
	// var conds []models.Condition

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.HandleErr(err, "should put a bad request error here")
		return
	}

	err = json.Unmarshal(body, &audReq) // THIS WORKS
	utils.HandleErr(err, "problem unmarshalling, what do?")

	printable, err := json.Marshal(&audReq)
	fmt.Println(string(printable))

	// for ind, _ := range audReq.Conditions {
	// 	var attr models.Attribute
	// 	h.DB.Find(&attr, audReq.Conditions[ind].AttributeID)
	// 	cond := model.Condition{}
	// }

	h.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&aud)

	// var aud models.Attribute
	// aud.Type = audReq.Type
	h.DB.Model(&aud).Save(&audReq)

	// utils.CreatedResponse(w, r, &models.AudienceResponse{Audience: &audReq})
	utils.CreatedResponse(w, r, &audReq)
}
