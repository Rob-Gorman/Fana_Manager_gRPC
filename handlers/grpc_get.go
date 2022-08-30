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
		fmt.Printf("Problem getting flag: %v", err)
		return nil, err
	}
	return flags, nil
}

func (h *Handler) GetAudiencesR() (auds []*models.Audience, err error) {
	err = h.DB.Find(&auds).Error
	if err != nil {
		fmt.Printf("Problem getting attribute: %v", err)
		return nil, err
	}
	return auds, nil
}

func (h *Handler) GetAttributesR() (attrs []*models.Attribute, err error) {
	err = h.DB.Find(&attrs).Error
	if err != nil {
		fmt.Printf("Problem getting attribute: %v", err)
		return nil, err
	}
	return attrs, nil
}
