package httpmw

import (
	"net/http"

	"github.com/pizzahutdigital/phdmw/phdcid"
)

// ConversationIDMiddleware returns a hanlder that will
func ConversationIDMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get conversation from HTTP
		convID := phdcid.GetCIDFromHTTP(r.Header)
		if convID == "" {
			// Opps conversation was not provided...
			convID = phdcid.CreateCID()
			r.Header["Grpc-Metadata-"+phdcid.GetCIDKeyName()] = []string{convID}
		}
		// set conversation ID in current context
		ctx := phdcid.SetCIDInContext(r.Context(), convID)
		// associated context with response
		nr := r.WithContext(ctx)

		// set response header with conversation ID
		phdcid.SetCIDInHTTP(w.Header(), convID)
		// continue in chain...
		h.ServeHTTP(w, nr)
	})
}
