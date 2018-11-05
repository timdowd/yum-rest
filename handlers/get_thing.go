package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/pizzahutdigital/phdmw/grpcmw"
	"github.com/pizzahutdigital/phdmw/phdlog"
	pb "github.com/pizzahutdigital/yum-rest/protobufs"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetThing gets a Thing
func (rs *RestServiceServer) GetThing(ctx context.Context, req *pb.GetThingReq) (*pb.GetThingRes, error) {

	// Set handler name for phdlog
	cid := grpcmw.HandlerStart(ctx, "GetThing")
	phdlog.Info(logMessage, cid, zap.String("Request", req.String()))

	gtCtx, gtSpan := trace.StartSpan(ctx, "GetFakeThing")

	// Fake some latency so trace feels important
	time.Sleep(time.Millisecond * 500)
	_ = gtCtx

	gtSpan.End()

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
			ThingID: req.GetThingID(),
			Name:    "Tom",
			Object: &pb.Object{
				Name:  "Mini Tom",
				Value: 3,
			},
		},
	}, nil
}
