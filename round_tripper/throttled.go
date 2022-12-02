// original source: https://gist.github.com/zdebra/10f0e284c4672e99f0cb767298f20c11

package round_tripper

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// throttledRoundTripper Rate Limited HTTP Client
type throttledRoundTripper struct {
	roundTripperWrap http.RoundTripper
	rateLimiter      *rate.Limiter
}

func (c *throttledRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {

	err := c.rateLimiter.Wait(r.Context()) // This is a blocking call. Honors the rate limit
	if err != nil {
		return nil, err
	}

	return c.roundTripperWrap.RoundTrip(r)
}

// NewThrottledRoundTripper wraps transportWrap with a rate limitter
// examle usage:
// client := http.DefaultClient
// client.Transport = NewThrottledRoundTripper(10*time.Seconds, 60, http.DefaultTransport) allows 1 requests every 10 seconds with a burst of up to 60 requests
func NewThrottledRoundTripper(limitPeriod time.Duration, requestCount int, transportWrap http.RoundTripper) http.RoundTripper {
	return &throttledRoundTripper{
		roundTripperWrap: transportWrap,
		rateLimiter:      rate.NewLimiter(rate.Every(limitPeriod), requestCount),
	}
}
