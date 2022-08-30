package main

import (
	"context"
	"fmt"
	"log"
	"manager/api"
	"manager/configs"
	"manager/pb"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	configs.LoadDotEnv()
	s2 := api.NewServer()

	pb.RegisterFanaServer(s, api.NewDashServer(&s2.H))
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func makeClient(t *testing.T) *pb.FanaClient {
	t.Helper()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	// defer conn.Close()
	client := pb.NewFanaClient(conn)
	return &client
}

func TestGetFlag(t *testing.T) {
	mc := *makeClient(t)
	have, err := mc.GetFlag(context.Background(), &pb.ID{ID: 1})
	if err != nil || have == nil {
		t.Fatalf("Failed TestGetFlag: \nHave: %v\nError: %v\n", have, err)
	}
	fmt.Println(have)
}

func TestGetFlags(t *testing.T) {
	mc := *makeClient(t)
	have, err := mc.GetFlags(context.Background(), &pb.Empty{})
	if err != nil || have == nil {
		t.Fatalf("Failed TestGetFlag: \nHave: %v\nError: %v\n", have, err)
	}
	fmt.Println(have)
}

func TestGetAudiences(t *testing.T) {
	mc := *makeClient(t)
	have, err := mc.GetAudiences(context.Background(), &pb.Empty{})
	if err != nil || have == nil {
		t.Fatalf("Failed TestGetAttrs: \nHave: %v\nError: %v\n", have, err)
	}
	fmt.Println(have)
}

func TestGetAttributes(t *testing.T) {
	mc := *makeClient(t)
	have, err := mc.GetAttributes(context.Background(), &pb.Empty{})
	if err != nil || have == nil {
		t.Fatalf("Failed TestGetAttrs: \nHave: %v\nError: %v\n", have, err)
	}
	fmt.Println(have)
}
