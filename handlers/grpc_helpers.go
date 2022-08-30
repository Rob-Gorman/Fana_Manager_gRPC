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
