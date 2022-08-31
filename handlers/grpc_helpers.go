package handlers

import (
	"manager/models"
	"manager/pb"
)

func (h *Handler) BuildAttributeAudiences(conds []models.Condition) (resp []*pb.AudienceSparseResp) {
	audids := []uint{}
	for _, cond := range conds {
		audids = append(audids, cond.AudienceID)
	}

	auds := []models.Audience{}

	if len(audids) > 0 {
		h.DB.Find(&auds, audids)
	}

	resp = []*pb.AudienceSparseResp{}
	for i := range auds {
		audience := auds[i].ToSparseResp()
		resp = append(resp, audience)
	}

	return resp
}

func (h *Handler) BuildEmbeddedConds(conds []models.Condition) (embConds []*pb.ConditionEmbedded) {
	embConds = []*pb.ConditionEmbedded{}
	for _, cond := range conds {
		embCond := cond.ToEmbeddedResp()
		embConds = append(embConds, embCond)
	}

	return embConds
}

func (h *Handler) BuildEmbeddedFlags(flags []models.Flag) (embFlags []*pb.FlagSparseResp) {
	embFlags = []*pb.FlagSparseResp{}
	for _, flag := range flags {
		embFlag := flag.ToSparseResp()
		embFlags = append(embFlags, embFlag)
	}

	return embFlags
}

func (h *Handler) FlagFromReq(in *pb.FlagSubmit) (*models.Flag, error) {
	auds := []models.Audience{}

	err := h.DB.Where("key in (?)", in.AudienceIDs).Find(&auds).Error
	if err != nil {
		return nil, err
	}

	flag := &models.Flag{
		Audiences:   auds,
		Key:         in.Key,
		DisplayName: in.DisplayName,
	}

	return flag, nil
}

func (h *Handler) AudienceFromReq(ar *pb.AudSubmit) (aud *models.Audience, err error) {
	conds := h.ConditionsFromReq(ar.Conditions)
	aud = &models.Audience{
		Key:         ar.Key,
		DisplayName: ar.DisplayName,
		Combine:     ar.Combine,
		Conditions:  conds,
	}

	return aud, nil
}

func (h *Handler) ConditionsFromReq(crs []*pb.CondSubmit) (conds []models.Condition) {
	conds = []models.Condition{}
	for _, cr := range crs {
		cond := models.Condition{
			AttributeID: uint(cr.AttributeID),
			Operator:    cr.Operator,
			Negate:      cr.Negate,
			Vals:        cr.Vals,
		}
		conds = append(conds, cond)
	}

	return conds
}

func (h *Handler) AttributeFromReq(in *pb.AttrSubmit) (attr *models.Attribute, err error) {
	return &models.Attribute{
		Key:         in.Key,
		DisplayName: in.DisplayName,
		Type:        in.Type,
	}, nil
}
