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

// GetThing gets a Thing
func (rs *RestServiceServer) GetThing(ctx context.Context, req *pb.GetThingReq) (*pb.GetThingRes, error) {

	// Set handler name for phdlog
	cid := grpcmw.HandlerStart(ctx, "GetThing")
	phdlog.Info(logMessage, cid, zap.String("Request", req.String()))

	if req.GetThingID() == failID {
		return nil, status.Errorf(codes.NotFound, "Thing `%s` was not found", req.GetThingID())
	}

	if req.GetThingID() == "dberror" {
		return nil, status.Errorf(codes.Internal, "Database Error: {{err}}")
	}

	// Return Thing
	return &pb.GetThingRes{
		Status:      http.StatusOK,
		Description: http.StatusText(http.StatusOK),
		Thing: &pb.Thing{
			Id:   req.GetThingID(),
			Name: "Tom",
			Object: &pb.Object{
				Name:  "Mini Tom",
				Value: 3,
			},
		},
	}, nil
}
