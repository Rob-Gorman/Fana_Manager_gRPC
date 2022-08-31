package api

import (
	"context"
	"manager/pb"
	"manager/utils"
)

func (ds *DashServer) CreateFlag(ctx context.Context, in *pb.FlagSubmit) (res *pb.FlagFullResp, err error) {
	flag, err := ds.H.CreateFlagR(in)
	if err != nil {
		err = utils.InternalError(err)
		return nil, err
	}
	res = flag.ToFullResp()
	return res, nil
}

func (ds *DashServer) CreateAudience(ctx context.Context, in *pb.AudSubmit) (res *pb.AudienceFullResp, err error) {
	aud, err := ds.H.CreateAudienceR(in)
	if err != nil {
		err = utils.InternalError(err)
		return nil, err
	}
	res = aud.ToFullResp()
	res.Conditions = ds.H.BuildEmbeddedConds(aud.Conditions)

	return res, nil
}

func (ds *DashServer) CreateAttribute(ctx context.Context, in *pb.AttrSubmit) (res *pb.AttributeResp, err error) {
	attr, err := ds.H.CreateAttributeR(in)
	if err != nil {
		err = utils.InternalError(err)
		return nil, err
	}

	res = attr.ToSparseResp()
	return res, nil
}
