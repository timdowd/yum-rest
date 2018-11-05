package handlers

const (
	failID = "xyz"
)

var logMessage = "yum-rest"

// RestServiceServer proxy between rest and rpc
type RestServiceServer struct {
}

// NewRest creates new service for testing
func NewRest() (bs *RestServiceServer, err error) {

	rs := &RestServiceServer{}
	return rs, nil
}
