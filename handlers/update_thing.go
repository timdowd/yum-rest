package handlers

import (
	"context"
	"net/http"

	"github.com/pizzahutdigital/phdmw/grpcmw"
	"github.com/pizzahutdigital/phdmw/phdlog"
	pb "github.com/pizzahutdigital/yum-rest/protobufs"
	"go.uber.org/zap"
)

// UpdateThing Updates a Thing
func (rs *RestServiceServer) UpdateThing(ctx context.Context, req *pb.UpdateThingReq) (*pb.UpdateThingRes, error) {

	// Set handler name for phdlog
	cid := grpcmw.HandlerStart(ctx, "UpdateThing")
	phdlog.Info(logMessage, cid, zap.String("Request", req.String()))

	// Return Thing
	return &pb.UpdateThingRes{
		Status:      http.StatusOK,
		Description: "Updated thing " + req.GetThing().GetId(),
	}, nil
}
