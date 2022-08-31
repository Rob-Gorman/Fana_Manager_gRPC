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

	resp := flag.ToFullResp()
	return resp, nil
}

func (ds *DashServer) GetFlags(ctx context.Context, empty *pb.Empty) (resp *pb.Flags, err error) {
	flags, err := ds.H.GetFlagsR()
	if err != nil {
		err = utils.NotFoundError(err)
		return nil, err
	}

	resp = &pb.Flags{
		Flags: []*pb.FlagSparseResp{},
	}

	for ind := range flags {
		resp.Flags = append(resp.Flags, flags[ind].ToSparseResp())
	}
	return resp, nil
}

func (ds *DashServer) GetAudience(ctx context.Context, id *pb.ID) (resp *pb.AudienceFullResp, err error) {
	aud, err := ds.H.GetAudienceR(int(id.ID))
	if err != nil {
		err = utils.NotFoundError(err)
		return nil, err
	}

	resp = aud.ToFullResp()

	resp.Conditions = ds.H.BuildEmbeddedConds(aud.Conditions)
	resp.Flags = ds.H.BuildEmbeddedFlags(aud.Flags)

	return resp, nil
}

func (ds *DashServer) GetAudiences(ctx context.Context, empty *pb.Empty) (resp *pb.Audiences, err error) {
	auds, err := ds.H.GetAudiencesR()
	if err != nil {
		err = utils.NotFoundError(err)
		return nil, err
	}

	resp = &pb.Audiences{
		Audiences: []*pb.AudienceSparseResp{},
	}

	for ind := range auds {
		resp.Audiences = append(resp.Audiences, auds[ind].ToSparseResp())
	}
	return resp, nil
}

func (ds *DashServer) GetAttribute(ctx context.Context, id *pb.ID) (*pb.AttributeResp, error) {
	attr, err := ds.H.GetAttributeR(int(id.ID))
	if err != nil {
		err = utils.NotFoundError(err)
		return nil, err
	}

	resp := attr.ToSparseResp()

	resp.Audiences = ds.H.BuildAttributeAudiences(attr.Conditions)

	return resp, nil
}

func (ds *DashServer) GetAttributes(ctx context.Context, empty *pb.Empty) (resp *pb.Attributes, err error) {
	attrs, err := ds.H.GetAttributesR()
	if err != nil {
		err = utils.NotFoundError(err)
		return nil, err
	}

	resp = &pb.Attributes{
		Attributes: []*pb.AttributeResp{},
	}

	for ind := range attrs {
		resp.Attributes = append(resp.Attributes, attrs[ind].ToSparseResp())
	}
	return resp, nil
}
