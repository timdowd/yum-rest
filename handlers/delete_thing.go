package handlers

import (
	"context"
	"net/http"

	"github.com/pizzahutdigital/phdmw/grpcmw"
	"github.com/pizzahutdigital/phdmw/phdlog"
	pb "github.com/pizzahutdigital/yum-rest/protobufs"
	"go.uber.org/zap"
)

// DeleteThing Deletes a Thing
func (rs *RestServiceServer) DeleteThing(ctx context.Context, req *pb.DeleteThingReq) (*pb.DeleteThingRes, error) {

	// Set handler name for phdlog
	cid := grpcmw.HandlerStart(ctx, "DeleteThing")
	phdlog.Info(logMessage, cid, zap.String("Request", req.String()))

	// Return Thing
	return &pb.DeleteThingRes{
		Status: http.StatusOK,
		// Big question: ID or id or Id ?
		Description: "Updated thing " + req.GetThingID(),
	}, nil
}
