package handlers

import (
	"context"
	"net/http"

	"github.com/pizzahutdigital/phdmw/grpcmw"
	"github.com/pizzahutdigital/phdmw/phdlog"
	pb "github.com/pizzahutdigital/yum-rest/protobufs"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdateThing Updates a Thing
func (rs *RestServiceServer) UpdateThing(ctx context.Context, req *pb.UpdateThingReq) (*pb.UpdateThingRes, error) {

	// Set handler name for phdlog
	cid := grpcmw.HandlerStart(ctx, "UpdateThing")
	phdlog.Info(logMessage, cid, zap.String("Request", req.String()))

	if req.GetThing().GetThingId() == failID {
		return nil, status.Errorf(codes.NotFound, "Thing `%s` does not exist", req.GetThing().GetThingId())
	}

	if req.GetThing().GetThingId() == "upsert" {
		return &pb.UpdateThingRes{
			Status:      http.StatusCreated,
			Description: "Updated thing " + req.GetThing().GetThingId() + ", we would want a new field for new id?",
		}, nil
	}

	// Return Thing
	return &pb.UpdateThingRes{
		Status:      http.StatusOK,
		Description: "Updated thing " + req.GetThing().GetThingId(),
	}, nil
}
