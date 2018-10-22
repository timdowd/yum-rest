package grpcmw

import (
	"context"

	"github.com/pizzahutdigital/phdmw/phdcid"
	"google.golang.org/grpc"
)

// ConversationIDMiddleware - detect if a conversation ID is present and inject one if it is not present.
func ConversationIDMiddleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		// Get conversation from HTTP
		convID := phdcid.GetCIDFromContext(ctx)
		if convID == "" {
			// Opps conversation was not provided...
			convID = phdcid.CreateCID()
		}
		// set conversation ID in current context
		ctx = phdcid.SetCIDInContext(ctx, convID)

		resp, err := handler(ctx, req)

		return resp, err
	}
}
