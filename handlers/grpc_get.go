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
