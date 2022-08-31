package handlers

import (
	"errors"
	"manager/models"
	"manager/utils"
)

func (h *Handler) DeleteFlagR(id int) error {
	flag := &models.Flag{}
	err := h.DB.Preload("Audiences").First(&flag, id).Error
	if err != nil {
		return err
	}

	h.DB.Model(&flag).Association("Audiences").Delete(flag.Audiences)
	err = h.DB.Unscoped().Delete(&flag).Error
	return err
}

func (h *Handler) DeleteAudienceR(id int) error {
	audience := &models.Audience{}
	err := h.DB.Preload("Flags").First(&audience, id).Error
	if err != nil {
		return utils.NotFoundError(err)
	}

	if !OrphanedAud(audience) {
		msg := "Cannot delete Audience assigned to Flags."
		return utils.InvalidArgumentError(errors.New("audience with dependencies"), msg)
	}

	h.DB.Model(&audience).Association("Flags").Delete(audience.Flags)
	err = h.DB.Unscoped().Delete(&audience).Error
	if err != nil {
		return utils.InternalError(err)
	}
	return nil
}

func (h *Handler) DeleteAttributeR(id int) error {
	attr := &models.Attribute{}
	err := h.DB.First(attr, id).Error
	if err != nil {
		return utils.NotFoundError(err)
	}

	if !OrphanedAttr(attr, h) {
		msg := "Cannot delete Attribute assigned to Audiences."
		utils.InvalidArgumentError(errors.New("attribute with dependencies"), msg)
	}

	err = h.DB.Unscoped().Delete(&attr).Error
	if err != nil {
		return utils.InternalError(err)
	}
	return nil
}
