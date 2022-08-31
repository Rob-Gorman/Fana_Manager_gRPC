package api

import (
	"context"
	"manager/pb"
	"manager/utils"
)

func (ds *DashServer) UpdateFlag(ctx context.Context, in *pb.FlagUpdate) (*pb.FlagFullResp, error) {
	flag, err := ds.H.UpdateFlagR(in.Updates, uint(in.ID))
	if err != nil {
		err = utils.InternalError(err)
		return nil, err
	}
	res := flag.ToFullResp()
	return res, nil
}

func (ds *DashServer) UpdateAudience(ctx context.Context, in *pb.AudUpdate) (*pb.AudienceFullResp, error) {
	audience, err := ds.H.UpdateAudienceR(in.Updates, uint(in.ID))
	if err != nil {
		err = utils.InternalError(err)
		return nil, err
	}
	res := audience.ToFullResp()
	return res, nil
}

func (ds *DashServer) RegenerateSDK(ctx context.Context, in *pb.ID) (res *pb.SDKKey, err error) {
	sdk, err := ds.H.RegenSDK(uint(in.ID))
	if err != nil {
		err = utils.InternalError(err)
		return nil, err
	}

	res = sdk.ToFullResp()
	return res, nil
}
