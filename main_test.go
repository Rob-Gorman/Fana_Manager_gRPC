package main

import (
	"context"
	"fmt"
	"log"
	"manager/api"
	"manager/configs"
	"manager/dev"
	"manager/pb"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

type sampleData struct {
	Conds []*pb.CondSubmit
	Auds  []*pb.AudSubmit
	Attrs []*pb.AttrSubmit
}

var data sampleData

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	configs.LoadDotEnv()
	s2 := api.NewServer()

	pb.RegisterFanaServer(s, api.NewDashServer(&s2.H))

	initSampleData()

	dev.RefreshSchema(s2.H.DB)

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

func initSampleData() {
	Cond1 := &pb.CondSubmit{
		AttributeID: 1,
		Operator:    "EQ",
		Negate:      true,
		Vals:        "Jersey",
	}
	Cond2 := &pb.CondSubmit{
		AttributeID: 2,
		Operator:    "EQ",
		Negate:      true,
		Vals:        "true",
	}
	Aud1 := &pb.AudSubmit{
		Key:         "GoTest Audience",
		DisplayName: "From inside main_test.go",
		Combine:     "ANY",
		Conditions:  []*pb.CondSubmit{Cond1},
	}
	Attr1 := &pb.AttrSubmit{
		Key:         "GoTest Attribute",
		DisplayName: "Attribute From main_test.go",
		Type:        "NUM",
	}

	data = sampleData{
		Conds: []*pb.CondSubmit{Cond1, Cond2},
		Auds:  []*pb.AudSubmit{Aud1},
		Attrs: []*pb.AttrSubmit{Attr1},
	}
}

func TestGetFlag(t *testing.T) {
	mc := *makeClient(t)
	have, err := mc.GetFlag(context.Background(), &pb.ID{ID: 1})
	if err != nil || have == nil {
		t.Fatalf("Failed TestGetFlag: \nHave: %v\nError: %v\n", have, err)
	}
	fmt.Print(have, "\n\n")
}

func TestGetFlags(t *testing.T) {
	mc := *makeClient(t)
	have, err := mc.GetFlags(context.Background(), &pb.Empty{})
	if err != nil || have == nil {
		t.Fatalf("Failed TestGetFlag: \nHave: %v\nError: %v\n", have, err)
	}
	fmt.Print(have, "\n\n")
}

func TestGetAudiences(t *testing.T) {
	mc := *makeClient(t)
	have, err := mc.GetAudiences(context.Background(), &pb.Empty{})
	if err != nil || have == nil {
		t.Fatalf("Failed TestGetAuds: \nHave: %v\nError: %v\n", have, err)
	}
	fmt.Print(have, "\n\n")
}

func TestGetAttributes(t *testing.T) {
	mc := *makeClient(t)
	have, err := mc.GetAttributes(context.Background(), &pb.Empty{})
	if err != nil || have == nil {
		t.Fatalf("Failed TestGetAttrs: \nHave: %v\nError: %v\n", have, err)
	}
	fmt.Print(have, "\n\n")
}

func TestGetAttribute(t *testing.T) {
	mc := *makeClient(t)
	have, err := mc.GetAttribute(context.Background(), &pb.ID{ID: 1})
	if err != nil || have == nil {
		t.Fatalf("Failed TestGetAttribute: \nHave: %v\nError: %v\n", have, err)
	}
	fmt.Print(have, "\n\n")
}

func TestGetAudience(t *testing.T) {
	mc := *makeClient(t)
	have, err := mc.GetAudience(context.Background(), &pb.ID{ID: 1})
	if err != nil || have == nil {
		t.Fatalf("Failed TestGetAudience: \nHave: %v\nError: %v\n", have, err)
	}
	fmt.Print(have, "\n\n")
}

func TestCreateFlag(t *testing.T) {
	mc := *makeClient(t)
	req := &pb.FlagSubmit{
		Key:         "GoTest Flag",
		DisplayName: "From inside main_test.go",
		AudienceIDs: []string{"1", "2"},
	}
	have, err := mc.CreateFlag(context.Background(), req)
	if err != nil || have == nil {
		t.Fatalf("Failed TestCreateFlag: \nHave: %v\nError: %v\n", have, err)
	}
	fmt.Print(have, "\n\n")
}

func TestCreateAudience(t *testing.T) {
	mc := *makeClient(t)
	req := data.Auds[0]
	have, err := mc.CreateAudience(context.Background(), req)
	if err != nil || have == nil {
		t.Fatalf("Failed TestCreateAudience: \nHave: %v\nError: %v\n", have, err)
	}
	fmt.Print(have, "\n\n")
}

func TestCreateAttribute(t *testing.T) {
	mc := *makeClient(t)
	req := data.Attrs[0]
	have, err := mc.CreateAttribute(context.Background(), req)
	if err != nil || have == nil {
		t.Fatalf("Failed TestCreateAttribute: \nHave: %v\nError: %v\n", have, err)
	}
	fmt.Print(have, "\n\n")
}
