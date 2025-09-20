package pumpimggtw

import (
	"context"
	"log"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
)

type Gateway struct {
	IPFSNode      *IPFSNode
	PublicGateway *PublicGateway
	Cache         *lru.Cache[string, []byte]
}

func NewGateway(ipfsNode *IPFSNode, publicGateway *PublicGateway) (*Gateway, error) {
	cache, err := lru.New[string, []byte](1000)
	if err != nil {
		return nil, err
	}
	return &Gateway{
		IPFSNode:      ipfsNode,
		PublicGateway: publicGateway,
		Cache:         cache,
	}, nil
}

func (g *Gateway) Get(cid string) ([]byte, error) {
	if val, ok := g.Cache.Get(cid); ok {
		return val, nil
	}

	// async pin
	go func() {
		err := g.IPFSNode.Pin(cid)
		if err != nil {
			return
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Millisecond)
	defer cancel()

	result, err := g.PublicGateway.Get(ctx, cid)
	if err != nil {
		log.Printf("Failed to get from public gateway: %v", err)

		result, err := g.IPFSNode.Get(cid)
		if err != nil {
			log.Printf("Failed to get from ipfs node: %v", err)
			return nil, err
		}
		g.Cache.Add(cid, result)
		return result, nil
	}

	g.Cache.Add(cid, result)
	return result, nil
}
