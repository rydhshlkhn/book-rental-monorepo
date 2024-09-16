package libraryservice

import (
	"net/http"
	"time"

	"github.com/golangid/candi/candiutils"
)

type libraryserviceRESTImpl struct {
	host    string
	authKey string
	httpReq candiutils.HTTPRequest
}

// NewLibraryserviceServiceREST constructor
func NewLibraryserviceServiceREST(host string, authKey string) Libraryservice {

	return &libraryserviceRESTImpl{
		host:    host,
		authKey: authKey,
		httpReq: candiutils.NewHTTPRequest(
			candiutils.HTTPRequestSetRetries(5),
			candiutils.HTTPRequestSetSleepBetweenRetry(500*time.Millisecond),
			candiutils.HTTPRequestSetHTTPErrorCodeThreshold(http.StatusBadRequest),
			candiutils.HTTPRequestSetBreakerName("libraryservice"),
		),
	}
}
