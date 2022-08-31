package models

import (
	"manager/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (f *Flag) ToFullResp() (res *pb.FlagFullResp) {
	audiences := []*pb.AudienceSparseResp{}
	for i := range f.Audiences {
		aud := f.Audiences[i].ToSparseResp()
		audiences = append(audiences, aud)
	}

	res = &pb.FlagFullResp{
		ID:          int32(f.ID),
		Key:         f.Key,
		DisplayName: f.DisplayName,
		Status:      f.Status,
		CreatedAt:   timestamppb.New(f.CreatedAt),
		UpdatedAt:   timestamppb.New(f.UpdatedAt),
		Audiences:   audiences,
	}

	return res
}

func (f *Flag) ToSparseResp() (res *pb.FlagSparseResp) {
	res = &pb.FlagSparseResp{
		ID:          int32(f.ID),
		Key:         f.Key,
		DisplayName: f.DisplayName,
		Status:      f.Status,
		CreatedAt:   timestamppb.New(f.CreatedAt),
		UpdatedAt:   timestamppb.New(f.UpdatedAt),
	}
	return res
}

func (aud *Audience) ToSparseResp() (res *pb.AudienceSparseResp) {
	res = &pb.AudienceSparseResp{
		ID:          int32(aud.ID),
		Key:         aud.Key,
		DisplayName: aud.DisplayName,
		CreatedAt:   timestamppb.New(aud.CreatedAt),
		UpdatedAt:   timestamppb.New(aud.UpdatedAt),
	}

	return res
}

func (aud *Audience) ToFullResp() (res *pb.AudienceFullResp) {
	res = &pb.AudienceFullResp{
		ID:          int32(aud.ID),
		Key:         aud.Key,
		DisplayName: aud.DisplayName,
		Combine:     aud.Combine,
		Conditions:  nil, // filled by helper
		Flags:       nil, // filled by helper
		CreatedAt:   timestamppb.New(aud.CreatedAt),
		UpdatedAt:   timestamppb.New(aud.UpdatedAt),
	}

	return res
}

func (attr *Attribute) ToSparseResp() (res *pb.AttributeResp) {
	res = &pb.AttributeResp{
		ID:          int32(attr.ID),
		Key:         attr.Key,
		DisplayName: attr.DisplayName,
		Type:        attr.Type,
		CreatedAt:   timestamppb.New(attr.CreatedAt),
	}

	return res
}

func (cond *Condition) ToEmbeddedResp() (res *pb.ConditionEmbedded) {
	res = &pb.ConditionEmbedded{
		Attribute: cond.Attribute.ToSparseResp(),
		Operator:  cond.Operator,
		Negate:    cond.Negate,
		Vals:      cond.Vals,
	}

	return res
}

func (sdk *Sdkkey) ToFullResp() (res *pb.SDKKey) {
	res = &pb.SDKKey{
		ID:        int32(sdk.ID),
		Key:       sdk.Key,
		Status:    sdk.Status,
		Type:      sdk.Type,
		CreatedAt: timestamppb.New(sdk.CreatedAt),
		UpdatedAt: timestamppb.New(sdk.UpdatedAt),
	}

	return res
}

func (fl *FlagLog) ToResp() (res *pb.LogMsg) {
	return &pb.LogMsg{
		LogID:     int32(fl.ID),
		ID:        int32(fl.FlagID),
		Key:       fl.FlagKey,
		Action:    fl.EventDesc,
		CreatedAt: timestamppb.New(fl.CreatedAt),
	}
}

func (aul *AudienceLog) ToResp() (res *pb.LogMsg) {
	return &pb.LogMsg{
		LogID:     int32(aul.ID),
		ID:        int32(aul.AudienceID),
		Key:       aul.AudienceKey,
		Action:    aul.EventDesc,
		CreatedAt: timestamppb.New(aul.CreatedAt),
	}
}

func (atl *AttributeLog) ToResp() (res *pb.LogMsg) {
	return &pb.LogMsg{
		LogID:     int32(atl.ID),
		ID:        int32(atl.AttributeID),
		Key:       atl.AttributeKey,
		Action:    atl.EventDesc,
		CreatedAt: timestamppb.New(atl.CreatedAt),
	}
}

func (ls *AuditLogs) ToResp() *pb.AuditLogResp {
	fLogs := []*pb.LogMsg{}
	for _, fl := range ls.FlagLogs {
		fLogs = append(fLogs, fl.ToResp())
	}

	auLogs := []*pb.LogMsg{}
	for _, au := range ls.AudienceLogs {
		auLogs = append(auLogs, au.ToResp())
	}

	attrLogs := []*pb.LogMsg{}
	for _, attr := range ls.AttributeLogs {
		attrLogs = append(attrLogs, attr.ToResp())
	}

	return &pb.AuditLogResp{
		FlagLogs:      fLogs,
		AudienceLogs:  auLogs,
		AttributeLogs: attrLogs,
	}
}
