package api

import (
	"context"
	"manager/pb"
	"manager/utils"
)

func (ds *DashServer) GetFlag(ctx context.Context, id *pb.ID) (*pb.FlagFullResp, error) {
	flag, err := ds.H.GetFlagR(int(id.ID))
	if err != nil {
		err = utils.NotFoundError(err)
		return nil, err
	}

	res := flag.ToFullResp()
	return res, nil
}

func (ds *DashServer) GetFlags(ctx context.Context, empty *pb.Empty) (res *pb.Flags, err error) {
	flags, err := ds.H.GetFlagsR()
	if err != nil {
		err = utils.NotFoundError(err)
		return nil, err
	}

	res = &pb.Flags{
		Flags: []*pb.FlagSparseResp{},
	}

	for ind := range flags {
		res.Flags = append(res.Flags, flags[ind].ToSparseResp())
	}
	return res, nil
}

func (ds *DashServer) GetAudience(ctx context.Context, id *pb.ID) (res *pb.AudienceFullResp, err error) {
	aud, err := ds.H.GetAudienceR(int(id.ID))
	if err != nil {
		err = utils.NotFoundError(err)
		return nil, err
	}

	res = aud.ToFullResp()

	res.Conditions = ds.H.BuildEmbeddedConds(aud.Conditions)
	res.Flags = ds.H.BuildEmbeddedFlags(aud.Flags)

	return res, nil
}

func (ds *DashServer) GetAudiences(ctx context.Context, empty *pb.Empty) (res *pb.Audiences, err error) {
	auds, err := ds.H.GetAudiencesR()
	if err != nil {
		err = utils.NotFoundError(err)
		return nil, err
	}

	res = &pb.Audiences{
		Audiences: []*pb.AudienceSparseResp{},
	}

	for ind := range auds {
		res.Audiences = append(res.Audiences, auds[ind].ToSparseResp())
	}
	return res, nil
}

func (ds *DashServer) GetAttribute(ctx context.Context, id *pb.ID) (*pb.AttributeResp, error) {
	attr, err := ds.H.GetAttributeR(int(id.ID))
	if err != nil {
		err = utils.NotFoundError(err)
		return nil, err
	}

	res := attr.ToSparseResp()

	res.Audiences = ds.H.BuildAttributeAudiences(attr.Conditions)

	return res, nil
}

func (ds *DashServer) GetAttributes(ctx context.Context, empty *pb.Empty) (res *pb.Attributes, err error) {
	attrs, err := ds.H.GetAttributesR()
	if err != nil {
		err = utils.NotFoundError(err)
		return nil, err
	}

	res = &pb.Attributes{
		Attributes: []*pb.AttributeResp{},
	}

	for ind := range attrs {
		res.Attributes = append(res.Attributes, attrs[ind].ToSparseResp())
	}
	return res, nil
}

func (ds *DashServer) GetSDKKeys(ctx context.Context, in *pb.Empty) (*pb.SDKKeys, error) {
	sdks, err := ds.H.GetSDKKeysR()
	if err != nil {
		err = utils.NotFoundError(err)
		return nil, err
	}

	resSDKs := []*pb.SDKKey{}
	for _, sdk := range sdks {
		resSDK := sdk.ToFullResp()
		resSDKs = append(resSDKs, resSDK)
	}

	res := &pb.SDKKeys{SDKs: resSDKs}
	return res, nil
}

func (ds *DashServer) GetAuditLogs(ctx context.Context, in *pb.Empty) (*pb.AuditLogResp, error) {
	logs, _ := ds.H.GetLogsR()

	res := logs.ToResp()

	return res, nil
}
