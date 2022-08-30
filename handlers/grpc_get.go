package handlers

import (
	"fmt"
	"manager/models"
)

func (h *Handler) GetFlagR(id int) (flag *models.Flag, err error) {
	err = h.DB.Preload("Audiences").First(&flag, id).Error
	if err != nil {
		fmt.Printf("Problem getting flag: %v", err)
		return nil, err
	}
	return flag, nil
}

func (h *Handler) GetFlagsR() (flags []*models.Flag, err error) {
	err = h.DB.Find(&flags).Error
	if err != nil {
		fmt.Printf("Problem getting flags: %v", err)
		return nil, err
	}
	return flags, nil
}

func (h *Handler) GetAudienceR(id int) (aud *models.Audience, err error) {
	err = h.DB.Preload("Flags").Preload("Conditions").First(&aud, id).Error
	if err != nil {
		fmt.Printf("Problem getting audiences: %v", err)
		return nil, err
	}
	return aud, nil
}

func (h *Handler) GetAudiencesR() (auds []*models.Audience, err error) {
	err = h.DB.Find(&auds).Error
	if err != nil {
		fmt.Printf("Problem getting audiences: %v", err)
		return nil, err
	}
	return auds, nil
}

func (h *Handler) GetAttributeR(id int) (attr *models.Attribute, err error) {
	err = h.DB.Preload("Conditions").First(&attr, id).Error
	if err != nil {
		fmt.Printf("Problem getting attr: %v", err)
		return nil, err
	}
	return attr, nil
}

func (h *Handler) GetAttributesR() (attrs []*models.Attribute, err error) {
	err = h.DB.Find(&attrs).Error
	if err != nil {
		fmt.Printf("Problem getting attrs: %v", err)
		return nil, err
	}
	return attrs, nil
}
