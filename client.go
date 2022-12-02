package nu

import (
	"github.com/SirMetathyst/go-nu/round_tripper"
	"net/http"
	"time"
)

var DefaultClient = New()

type Client struct {
	client *http.Client
}

func New() *Client {

	client := &http.Client{}
	client.Transport = round_tripper.NewCloudFlareBypassRoundTripper(client.Transport)
	client.Transport = round_tripper.NewThrottledRoundTripper(1*time.Second, 2, client.Transport)

	return &Client{client: client}
}

func NewWithClient(client *http.Client) *Client {
	return &Client{client: client}
}
