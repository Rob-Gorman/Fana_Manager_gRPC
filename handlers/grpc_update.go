package handlers

import (
	"manager/models"
	"manager/pb"

	"gorm.io/gorm"
)

func (h *Handler) UpdateFlagR(in *pb.FlagSubmit, id uint) (*models.Flag, error) {
	flag, err := h.FlagFromReq(in)
	flag.ID = id
	if err != nil {
		return nil, err
	}

	if in.AudienceIDs != nil {
		h.DB.Model(&flag).Omit("Audiences.*").Association("Audiences").Replace(flag.Audiences)
	}

	err = h.DB.Omit("Audiences").Session(&gorm.Session{
		SkipHooks: true,
	}).Updates(&flag).Error
	if err != nil {
		return nil, err
	}

	return flag, nil
}

func (h *Handler) UpdateAudienceR(in *pb.AudSubmit, id uint) (*models.Audience, error) {
	aud, err := h.AudienceFromReq(in)
	aud.ID = id
	if err != nil {
		return nil, err
	}

	if in.Conditions != nil {
		h.DB.Model(&aud).Association("Conditions").Replace(aud.Conditions)
	}

	err = h.DB.Session(&gorm.Session{
		FullSaveAssociations: true,
		SkipHooks:            true,
	}).Updates(&aud).Error

	if err != nil {
		return nil, err
	}

	h.DB.Model(&models.Audience{}).Preload("Flags").Preload("Conditions").Find(&aud)

	return aud, nil
}

func (h *Handler) RegenSDK(id uint) (*models.Sdkkey, error) {
	sdk := models.Sdkkey{}
	h.DB.Find(&sdk, id)

	newSDK := &models.Sdkkey{
		Key:  NewSDKKey(sdk.Key),
		Type: sdk.Type,
	}

	err := h.DB.Create(&newSDK).Error
	if err != nil {
		return nil, err
	}

	h.DB.Unscoped().Delete(&sdk)

	h.DB.Find(&newSDK)
	return newSDK, nil
}
