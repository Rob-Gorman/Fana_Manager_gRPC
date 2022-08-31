package handlers

import (
	"manager/models"
	"manager/pb"

	"gorm.io/gorm"
)

func (h *Handler) CreateFlagR(in *pb.FlagSubmit) (flag *models.Flag, err error) {
	flag, err = h.FlagFromReq(in)
	if err != nil {
		return nil, err
	}

	err = h.DB.Omit("Audiences.*").Session(&gorm.Session{FullSaveAssociations: true}).Create(&flag).Error
	if err != nil {
		return nil, err
	}

	h.DB.Preload("Audiences").Find(&flag)
	return flag, nil
}

func (h *Handler) CreateAudienceR(in *pb.AudSubmit) (audience *models.Audience, err error) {
	audience, err = h.AudienceFromReq(in)
	if err != nil {
		return nil, err
	}

	err = h.DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(&audience).Error
	if err != nil {
		return nil, err
	}

	h.DB.Model(&models.Audience{}).Preload("Conditions").Find(&audience)

	return audience, nil
}

func (h *Handler) CreateAttributeR(in *pb.AttrSubmit) (attr *models.Attribute, err error) {
	attr, _ = h.AttributeFromReq(in)

	err = h.DB.Create(&attr).Error
	if err != nil {
		return nil, err
	}

	h.DB.Find(&attr)
	return attr, nil
}
