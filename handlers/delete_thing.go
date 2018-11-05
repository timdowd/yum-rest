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

// DeleteThing Deletes a Thing
func (rs *RestServiceServer) DeleteThing(ctx context.Context, req *pb.DeleteThingReq) (*pb.DeleteThingRes, error) {

	// Set handler name for phdlog
	cid := grpcmw.HandlerStart(ctx, "DeleteThing")
	phdlog.Info(logMessage, cid, zap.String("Request", req.String()))

	if req.GetThingId() == failID {
		return nil, status.Errorf(codes.NotFound, "Thing `%s` does not exist", req.GetThingId())
	}

	// Return Thing
	return &pb.DeleteThingRes{
		Status:      http.StatusOK,
		Description: "Deleted thing " + req.GetThingId(),
	}, nil
}
