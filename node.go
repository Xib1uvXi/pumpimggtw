package pumpimggtw

import (
	"context"
	"io"
	"time"

	"github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/kubo/client/rpc"
)

type IPFSNode struct {
	Node *rpc.HttpApi
}

func NewIPFSNode() (*IPFSNode, error) {
	node, err := rpc.NewLocalApi()
	if err != nil {
		return nil, err
	}
	return &IPFSNode{
		Node: node,
	}, nil
}

// block get
func (n *IPFSNode) GetBlock(cidStr string) ([]byte, error) {
	c, err := cid.Decode(cidStr)
	if err != nil {
		return nil, err
	}

	resp, err := n.Node.Block().Get(context.Background(), path.FromCid(c))
	if err != nil {
		return nil, err
	}
	return io.ReadAll(resp)
}

// unixfs get
func (n *IPFSNode) Get(cidStr string) ([]byte, error) {
	c, err := cid.Decode(cidStr)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := n.Node.Unixfs().Get(ctx, path.FromCid(c))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return io.ReadAll(r.(io.Reader))
}

// pin
func (n *IPFSNode) Pin(cidStr string) error {
	c, err := cid.Decode(cidStr)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return n.Node.Pin().Add(ctx, path.FromCid(c))
}
