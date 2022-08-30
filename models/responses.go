package models

import (
	"manager/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (f *Flag) ToFullResp() (resp *pb.FlagFullResp) {
	audiences := []*pb.AudienceSparseResp{}
	for i := range f.Audiences {
		aud := f.Audiences[i].ToSparseResp()
		audiences = append(audiences, aud)
	}

	resp = &pb.FlagFullResp{
		ID:          int32(f.ID),
		Key:         f.Key,
		DisplayName: f.DisplayName,
		Status:      f.Status,
		CreatedAt:   timestamppb.New(f.CreatedAt),
		UpdatedAt:   timestamppb.New(f.UpdatedAt),
		Audiences:   audiences,
	}

	return resp
}

func (f *Flag) ToSparseResp() (resp *pb.FlagSparseResp) {
	resp = &pb.FlagSparseResp{
		ID:          int32(f.ID),
		Key:         f.Key,
		DisplayName: f.DisplayName,
		Status:      f.Status,
		CreatedAt:   timestamppb.New(f.CreatedAt),
		UpdatedAt:   timestamppb.New(f.UpdatedAt),
	}
	return resp
}

func (aud *Audience) ToSparseResp() (resp *pb.AudienceSparseResp) {
	resp = &pb.AudienceSparseResp{
		ID:          int32(aud.ID),
		Key:         aud.Key,
		DisplayName: aud.DisplayName,
		CreatedAt:   timestamppb.New(aud.CreatedAt),
		UpdatedAt:   timestamppb.New(aud.UpdatedAt),
	}

	return resp
}
