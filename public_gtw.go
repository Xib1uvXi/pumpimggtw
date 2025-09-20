package pumpimggtw

import (
	"context"
	"fmt"

	"resty.dev/v3"
)

type PublicGateway struct {
	url    string
	client *resty.Client
}

func NewPublicGateway(url string) *PublicGateway {
	return &PublicGateway{
		url:    url,
		client: resty.New(),
	}
}

func (g *PublicGateway) Get(ctx context.Context, cid string) ([]byte, error) {
	resp, err := g.client.R().SetContext(ctx).Get(fmt.Sprintf("%s/ipfs/%s", g.url, cid))
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
