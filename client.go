package quokka

import (
	"net/http"
	"sync"
	"time"
)

type Client struct {
	httpc       *http.Client
	mu          sync.RWMutex
	visitorData string
}

func NewClient(httpc *http.Client) *Client {
	if httpc == nil {
		httpc = &http.Client{
			Timeout: 15 * time.Second,
		}
	}

	return &Client{
		httpc: httpc,
		mu:    sync.RWMutex{},
	}
}
