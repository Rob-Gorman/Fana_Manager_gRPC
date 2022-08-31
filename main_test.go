package main

// Just happy path integration tests to verify behavior during development

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
	Flags []*pb.FlagSubmit
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
	Flag1 := &pb.FlagSubmit{
		Key:         "Updated Flag from GoTest",
		DisplayName: "Flag1 in main_test.go",
		AudienceIDs: []string{"1"},
	}
	Flag2 := &pb.FlagSubmit{
		Key:         "GoTest Flag",
		DisplayName: "From inside main_test.go",
		AudienceIDs: []string{"1", "2"},
	}
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
	Aud2 := &pb.AudSubmit{
		Key:         "GoTest Audience TWO",
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
		Flags: []*pb.FlagSubmit{Flag1, Flag2},
		Conds: []*pb.CondSubmit{Cond1, Cond2},
		Auds:  []*pb.AudSubmit{Aud1, Aud2},
		Attrs: []*pb.AttrSubmit{Attr1},
	}
}

func outputHelper(t *testing.T, have interface{}, err error, f string) {
	t.Helper()
	if err != nil || have == nil {
		t.Fatalf("Failed %s: \nHave: %v\nError: %v\n", f, have, err)
	}
	fmt.Print(have, "\n\n")
}

func TestGetFlag(t *testing.T) {
	mc := *makeClient(t)
	have, err := mc.GetFlag(context.Background(), &pb.ID{ID: 1})
	outputHelper(t, have, err, "TestGetFlag")
}

func TestGetFlags(t *testing.T) {
	mc := *makeClient(t)
	have, err := mc.GetFlags(context.Background(), &pb.Empty{})
	outputHelper(t, have, err, "TestGetFlags")
}

func TestGetAudiences(t *testing.T) {
	mc := *makeClient(t)
	have, err := mc.GetAudiences(context.Background(), &pb.Empty{})
	outputHelper(t, have, err, "TestGetAudiences")
}

func TestGetAttributes(t *testing.T) {
	mc := *makeClient(t)
	have, err := mc.GetAttributes(context.Background(), &pb.Empty{})
	outputHelper(t, have, err, "TestGetAttributes")

}

func TestGetAttribute(t *testing.T) {
	mc := *makeClient(t)
	have, err := mc.GetAttribute(context.Background(), &pb.ID{ID: 1})
	outputHelper(t, have, err, "TestGetAttribute")
}

func TestGetAudience(t *testing.T) {
	mc := *makeClient(t)
	have, err := mc.GetAudience(context.Background(), &pb.ID{ID: 1})
	outputHelper(t, have, err, "TestGetAudience")
}

func TestCreateFlag(t *testing.T) {
	mc := *makeClient(t)
	req := data.Flags[1]
	have, err := mc.CreateFlag(context.Background(), req)
	outputHelper(t, have, err, "TestCreateFlag")
}

func TestCreateAudience(t *testing.T) {
	mc := *makeClient(t)
	req := data.Auds[0]
	have, err := mc.CreateAudience(context.Background(), req)
	outputHelper(t, have, err, "TestGetAudience")
}

func TestCreateAttribute(t *testing.T) {
	mc := *makeClient(t)
	req := data.Attrs[0]
	have, err := mc.CreateAttribute(context.Background(), req)
	outputHelper(t, have, err, "TestGetAttribute")
}

func TestUpdateFlag(t *testing.T) {
	mc := *makeClient(t)
	req := &pb.FlagUpdate{Updates: data.Flags[0], ID: 1}
	have, err := mc.UpdateFlag(context.Background(), req)
	outputHelper(t, have, err, "TestUpdateFlag")
}

func TestUpdateAudience(t *testing.T) {
	mc := *makeClient(t)
	req := &pb.AudUpdate{Updates: data.Auds[1], ID: 1}
	have, err := mc.UpdateAudience(context.Background(), req)
	outputHelper(t, have, err, "TestUpdateAudience")
}

func TestGetSDKs(t *testing.T) {
	mc := *makeClient(t)
	have, err := mc.GetSDKKeys(context.Background(), &pb.Empty{})
	outputHelper(t, have, err, "TestGetSDKs")
}

func TestRegenSDK(t *testing.T) {
	mc := *makeClient(t)
	have, err := mc.RegenerateSDK(context.Background(), &pb.ID{ID: 2})
	outputHelper(t, have, err, "TestRegenerateSDK")
}

func TestGetLogs(t *testing.T) {
	mc := *makeClient(t)
	have, err := mc.GetAuditLogs(context.Background(), &pb.Empty{})
	outputHelper(t, have, err, "TestGetLogs")
}

func TestDeleteFlag(t *testing.T) {
	mc := *makeClient(t)
	id := &pb.ID{ID: 1}
	ctx := context.Background()
	_, err := mc.DeleteFlag(ctx, id)
	// outputHelper(t, have, err, "DeleteFlag Delete phase")
	have2, err := mc.GetFlag(ctx, id)
	if err == nil || have2 != nil {
		t.Fatalf("Why does this resource still exist? %v", have2)
	}
}

func TestDeleteAudience(t *testing.T) {
	mc := *makeClient(t)
	id := &pb.ID{ID: 1}
	ctx := context.Background()
	_, err := mc.DeleteAudience(ctx, id)
	// outputHelper(t, have, err, "DeleteAudience Delete phase")
	have2, err := mc.GetAudience(ctx, id)
	if err == nil || have2 != nil {
		t.Fatalf("Why does this resource still exist? %v", have2)
	}
}

func TestDeleteAttribute(t *testing.T) {
	mc := *makeClient(t)
	id := &pb.ID{ID: 1}
	ctx := context.Background()
	_, err := mc.DeleteAttribute(ctx, id)
	// outputHelper(t, have, err, "DeleteAttribute Delete phase")
	have2, err := mc.GetAttribute(ctx, id)
	if err == nil || have2 != nil {
		t.Fatalf("Why does this resource still exist? %v", have2)
	}
}
