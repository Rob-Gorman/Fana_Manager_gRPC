package api

import (
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
