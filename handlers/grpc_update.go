package handlers

import (
	"fmt"
	"manager/models"
	"manager/pb"

	"gorm.io/gorm"
)

func (h *Handler) ToggleFlagR(in *pb.FlagToggle) error {
	id := in.ID
	status := in.Status

	var flag models.Flag
	h.DB.Find(&flag, id)
	flag.Status = status
	flag.DisplayName = fmt.Sprintf("__%v", flag.Status) // hacky way to clue it's a toggle action, see flag update hook
	err := h.DB.Select("status").Updates(&flag).Error
	if err != nil {
		return err
	}

	pub := FlagUpdateForPublisher(h.DB, []models.Flag{flag})
	PublishContent(&pub, "flag-toggle-channel")
	RefreshCache(h.DB)

	return nil
}

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

	pub := FlagUpdateForPublisher(h.DB, []models.Flag{*flag})
	PublishContent(&pub, "flag-update-channel")
	RefreshCache(h.DB)

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

	pub := FlagUpdateForPublisher(h.DB, aud.Flags)
	PublishContent(&pub, "flag-update-channel")
	RefreshCache(h.DB)

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

	RefreshCache(h.DB)

	h.DB.Find(&newSDK)
	return newSDK, nil
}
