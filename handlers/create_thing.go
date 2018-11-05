package handlers

import (
	"context"
	"net/http"

	"github.com/pizzahutdigital/phdmw/grpcmw"
	"github.com/pizzahutdigital/phdmw/phdlog"
	pb "github.com/pizzahutdigital/yum-rest/protobufs"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// CreateThing Creates a Thing
func (rs *RestServiceServer) CreateThing(ctx context.Context, req *pb.CreateThingReq) (*pb.CreateThingRes, error) {

	// Set handler name for phdlog
	cid := grpcmw.HandlerStart(ctx, "CreateThing")
	phdlog.Info(logMessage, cid, zap.String("Request", req.String()))

	if req.GetThing().GetThingID() == failID {
		return nil, status.Errorf(codes.InvalidArgument, "Thing `%s` already exists", req.GetThing().GetThingID())
	}

	if req.GetThing().GetName() == "Todd" {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot name Thing %s", req.GetThing().GetName())
	}

	if req.GetThing().GetThingID() == "uuid" {
		return &pb.CreateThingRes{
			Status:      http.StatusCreated,
			Description: http.StatusText(http.StatusOK),
			ThingID:     "09963975-06b8-4e59-aa61-3514b5dd22b3",
		}, nil
	}

	header := metadata.Pairs("Location", "/Things/"+req.GetThing().GetThingID())
	_ = grpc.SetHeader(ctx, header)

	// Return Thing
	return &pb.CreateThingRes{
		Status:      http.StatusCreated,
		Description: http.StatusText(http.StatusOK),
		ThingID:     req.GetThing().GetThingID(),
	}, nil
}
