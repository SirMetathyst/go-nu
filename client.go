package nu

import (
	"fmt"
	"github.com/SirMetathyst/go-nu/round_tripper"
	"golang.org/x/net/html"
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

func (s *Client) request(url string) (*html.Node, error) {

	response, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	doc, err := html.Parse(response.Body)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return doc, nil
}
