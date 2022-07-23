package handlers

import (
	"manager/models"

	"gorm.io/gorm"
)

func GetEmbeddedConds(aud models.Audience, db *gorm.DB) []models.ConditionEmbedded {
	conds := []models.ConditionEmbedded{}
	for ind, _ := range aud.Conditions {
		cond := aud.Conditions[ind]
		var attr models.Attribute
		db.Find(&attr, cond.AttributeID)
		db.Find(&cond)
		cond.Attribute = attr
		conds = append(conds, models.ConditionEmbedded{
			Condition: &cond,
			Attribute: attr.Key,
		})
	}
	return conds
}

func FlagReqToFlag(flagReq models.FlagSubmit, h Handler) (flag models.Flag) {
	auds := []models.Audience{}

	h.DB.Where("key in (?)", flagReq.Audiences).Find(&auds)

	flag = models.Flag{
		Audiences:   auds,
		Key:         flagReq.Key,
		DisplayName: flagReq.DisplayName,
		Sdkkey:      flagReq.SdkKey,
	}

	return flag
}

func FlagToFlagResponse(flag models.Flag, h Handler) models.FlagResponse {
	h.DB.Preload("Audiences").First(&flag)
	respAuds := []models.AudienceNoCondsResponse{}
	for ind, _ := range flag.Audiences {
		respAuds = append(respAuds, models.AudienceNoCondsResponse{Audience: &flag.Audiences[ind]})
	}
	return models.FlagResponse{
		Flag:      &flag,
		Audiences: respAuds,
	}
}

func BuildAudUpdate(req models.Audience, id int, h Handler) (aud models.Audience) {
	h.DB.Find(&aud, id)
	aud.Conditions = req.Conditions
	aud.Combine = req.Combine
	aud.DisplayName = req.DisplayName
	return aud
}

func GetEmbeddedFlags(flags []models.Flag) []models.FlagNoAudsResponse {
	fr := []models.FlagNoAudsResponse{}
	for i, _ := range flags {
		fr = append(fr, models.FlagNoAudsResponse{Flag: &flags[i]})
	}

	return fr
}
