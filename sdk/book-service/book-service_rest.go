package bookservice

import (
	"net/http"
	"time"

	"github.com/golangid/candi/candiutils"
)

type bookserviceRESTImpl struct {
	host    string
	authKey string
	httpReq candiutils.HTTPRequest
}

// NewBookserviceServiceREST constructor
func NewBookserviceServiceREST(host string, authKey string) Bookservice {

	return &bookserviceRESTImpl{
		host:    host,
		authKey: authKey,
		httpReq: candiutils.NewHTTPRequest(
			candiutils.HTTPRequestSetRetries(5),
			candiutils.HTTPRequestSetSleepBetweenRetry(500*time.Millisecond),
			candiutils.HTTPRequestSetHTTPErrorCodeThreshold(http.StatusBadRequest),
			candiutils.HTTPRequestSetBreakerName("bookservice"),
		),
	}
}
