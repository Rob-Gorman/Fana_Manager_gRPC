package api

import (
	"context"
	"manager/pb"
)

func (ds *DashServer) DeleteFlag(ctx context.Context, in *pb.ID) (res *pb.Empty, err error) {
	err = ds.H.DeleteFlagR(int(in.ID))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ds *DashServer) DeleteAudience(ctx context.Context, in *pb.ID) (res *pb.Empty, err error) {
	err = ds.H.DeleteAudienceR(int(in.ID))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ds *DashServer) DeleteAttribute(ctx context.Context, in *pb.ID) (res *pb.Empty, err error) {
	err = ds.H.DeleteAttributeR(int(in.ID))
	if err != nil {
		return nil, err
	}

	return res, nil
}
