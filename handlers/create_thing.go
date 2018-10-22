package handlers

import (
	"context"
	"net/http"

	"github.com/pizzahutdigital/phdmw/grpcmw"
	"github.com/pizzahutdigital/phdmw/phdlog"
	pb "github.com/pizzahutdigital/yum-rest/protobufs"
	"go.uber.org/zap"
)

// CreateThing Creates a Thing
func (rs *RestServiceServer) CreateThing(ctx context.Context, req *pb.CreateThingReq) (*pb.CreateThingRes, error) {

	// Set handler name for phdlog
	cid := grpcmw.HandlerStart(ctx, "CreateThing")
	phdlog.Info(logMessage, cid, zap.String("Request", req.String()))

	// Return Thing
	return &pb.CreateThingRes{
		Status:      http.StatusOK,
		Description: http.StatusText(http.StatusOK),
		ThingID:     req.GetThing().GetId(),
	}, nil
}
