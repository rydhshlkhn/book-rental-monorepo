package paymentservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golangid/candi/candihelper"
	"github.com/golangid/candi/candiutils"
	"github.com/golangid/candi/tracer"
)

type paymentserviceRESTImpl struct {
	host    string
	authKey string
	httpReq candiutils.HTTPRequest
}

// NewPaymentserviceServiceREST constructor
func NewPaymentserviceServiceREST(host string, authKey string) Paymentservice {

	return &paymentserviceRESTImpl{
		host:    host,
		authKey: authKey,
		httpReq: candiutils.NewHTTPRequest(
			candiutils.HTTPRequestSetRetries(5),
			candiutils.HTTPRequestSetSleepBetweenRetry(500*time.Millisecond),
			candiutils.HTTPRequestSetHTTPErrorCodeThreshold(http.StatusBadRequest),
			candiutils.HTTPRequestSetBreakerName("paymentservice"),
		),
	}
}

func (us *paymentserviceRESTImpl) Pay(ctx context.Context, req RequestPayment) (res *Response, code int, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "AuthServiceSDK:GenerateToken")
	defer trace.Finish()

	headers := map[string]string{
		candihelper.HeaderContentType:   candihelper.HeaderMIMEApplicationJSON,
		candihelper.HeaderAuthorization: us.authKey,
	}

	uri := fmt.Sprintf("http://localhost:8005/v1/payment")
	body, code, err := candiutils.NewHTTPRequest(
		candiutils.HTTPRequestSetHTTPErrorCodeThreshold(http.StatusBadRequest),
	).Do(ctx, http.MethodGet, uri, candihelper.ToBytes(req), headers)
	// if err != nil {
	// 	trace.SetError(err)
	// 	return claim, code, err
	// }

	json.Unmarshal(body, &res)
	if err != nil {
		err = errors.New(res.Message)
	}
	return
}
