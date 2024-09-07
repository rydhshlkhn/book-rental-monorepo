package userservice

import (
	"net/http"
	"time"

	"github.com/golangid/candi/candiutils"
)

type userserviceRESTImpl struct {
	host    string
	authKey string
	httpReq candiutils.HTTPRequest
}

// NewUserserviceServiceREST constructor
func NewUserserviceServiceREST(host string, authKey string) Userservice {

	return &userserviceRESTImpl{
		host:    host,
		authKey: authKey,
		httpReq: candiutils.NewHTTPRequest(
			candiutils.HTTPRequestSetRetries(5),
			candiutils.HTTPRequestSetSleepBetweenRetry(500*time.Millisecond),
			candiutils.HTTPRequestSetHTTPErrorCodeThreshold(http.StatusBadRequest),
			candiutils.HTTPRequestSetBreakerName("userservice"),
		),
	}
}
