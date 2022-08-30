package api

import (
	"context"
	"manager/handlers"
	"manager/pb"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type DashServer struct {
	pb.UnimplementedFanaServer
	H *handlers.Handler
}

func Init(h *handlers.Handler) {
	gsrv := grpc.NewServer()
	msrv := NewDashServer(h)

	pb.RegisterFanaServer(gsrv, msrv)
	reflection.Register(gsrv)

	listener, _ := net.Listen("tcp", ":9090")
	gsrv.Serve(listener)
}

func NewDashServer(h *handlers.Handler) *DashServer {
	return &DashServer{H: h}
}

func (ds *DashServer) GetFlag(ctx context.Context, id *pb.ID) (*pb.FlagFullResp, error) {
	flag, err := ds.H.GetFlagR(int(id.ID))
	if err != nil {
		return nil, err
	}
	resp := flag.ToFullResp()
	return resp, nil
}

func (ds *DashServer) GetFlags(ctx context.Context, empty *pb.Empty) (resp *pb.Flags, err error) {
	flags, err := ds.H.GetFlagsR()
	if err != nil {
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

func (ds *DashServer) GetAudiences(ctx context.Context, empty *pb.Empty) (resp *pb.Audiences, err error) {
	auds, err := ds.H.GetAudiencesR()
	if err != nil {
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

func (ds *DashServer) GetAttributes(ctx context.Context, empty *pb.Empty) (resp *pb.Attributes, err error) {
	attrs, err := ds.H.GetAttributesR()
	if err != nil {
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
