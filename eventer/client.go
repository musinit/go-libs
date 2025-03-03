package eventer

import (
	"context"
	"github.com/trustwallet/go-libs/client"
	"github.com/trustwallet/go-libs/middleware"
	"net/http"
)

type Client struct {
	client.Request
}

type Status struct {
	Status bool `json:"status"`
}

type Event struct {
	Name      string            `json:"name"`
	CreatedAt int64             `json:"created_at"`
	Params    map[string]string `json:"params"`
}

var senderClient *Client
var batchLimit = 100

func Init(url string, limit int) {
	senderClient = &Client{client.InitJSONClient(url, middleware.SentryErrorHandler)}
	batchLimit = limit
}

func (c Client) SendBatch(events []Event) (status Status, err error) {
	_, err = senderClient.Execute(context.Background(), client.NewReqBuilder().
		Method(http.MethodPost).
		WriteTo(&status).
		PathStatic("").
		Body(events).
		Build())
	return
}
