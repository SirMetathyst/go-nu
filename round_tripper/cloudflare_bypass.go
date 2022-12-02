// original source: https://github.com/DaRealFreak/cloudflare-bp-go

// Package round_tripper provides a round tripper to not get detected by CloudFlare directly on the first HTTP request
// The round tripper will add required/validated request headers and updates the client TLS configuration
// It'll NOT solve challenges provided by CloudFlare, just prevent from being detected on the first request
package round_tripper

import (
	"crypto/tls"
	//modification: removed fake-useragent because I don't think I need it
	"net/http"
)

// cloudFlareRoundTripper is a custom round tripper add the validated request headers.
type cloudFlareRoundTripper struct {
	inner   http.RoundTripper
	options Options
}

// Options the option to set custom headers
type Options struct {
	AddMissingHeaders bool
	Headers           map[string]string
}

// NewCloudFlareBypassRoundTripper returns a round tripper adding the required headers for the CloudFlare checks
// and updates the TLS configuration of the passed inner transport.
func NewCloudFlareBypassRoundTripper(inner http.RoundTripper, options ...Options) http.RoundTripper {
	if trans, ok := inner.(*http.Transport); ok {
		trans.TLSClientConfig = getCloudFlareTLSConfiguration()
	}

	roundTripper := &cloudFlareRoundTripper{
		inner: inner,
	}

	if options != nil {
		roundTripper.options = options[0]
	} else {
		roundTripper.options = GetDefaultOptions()
	}

	return roundTripper
}

// RoundTrip adds the required request headers to pass CloudFlare checks.
func (ug *cloudFlareRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	if ug.options.AddMissingHeaders {
		for header, value := range ug.options.Headers {
			if _, ok := r.Header[header]; !ok {
				r.Header.Set(header, value)
			}
		}
	}

	// in case we don't have an inner transport layer from the round tripper
	if ug.inner == nil {
		return (&http.Transport{
			TLSClientConfig: getCloudFlareTLSConfiguration(),
		}).RoundTrip(r)
	}

	return ug.inner.RoundTrip(r)
}

// getCloudFlareTLSConfiguration returns an accepted client TLS configuration to not get detected by CloudFlare directly
// in case the configuration needs to be updated later on: https://wiki.mozilla.org/Security/Server_Side_TLS .
func getCloudFlareTLSConfiguration() *tls.Config {
	return &tls.Config{
		CurvePreferences: []tls.CurveID{tls.CurveP256, tls.CurveP384, tls.CurveP521, tls.X25519},
	}
}

// GetDefaultOptions returns the options set by default
func GetDefaultOptions() Options {
	return Options{
		AddMissingHeaders: true,
		Headers: map[string]string{
			"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
			"Accept-Language": "en-US,en;q=0.5",
			//modification: removed fake-useragent
			//"User-Agent":      browser.Firefox(),
			"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36",
		},
	}
}
