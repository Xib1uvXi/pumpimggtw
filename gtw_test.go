package pumpimggtw_test

import (
	"testing"

	"github.com/Xib1uvXi/pumpimggtw"
	"github.com/stretchr/testify/assert"
)

func TestNewGateway(t *testing.T) {
	ipfsNode, err := pumpimggtw.NewIPFSNode()
	assert.NoError(t, err)
	gtw, err := pumpimggtw.NewGateway(ipfsNode, pumpimggtw.NewPublicGateway("https://ipfs.io"))
	assert.NoError(t, err)

	body, err := gtw.Get("QmPwZqadGRwxJ1oraZM62YHkzfcD2Tqwu86U9BCEXABraw")
	assert.NoError(t, err)
	assert.True(t, len(body) > 0)
	t.Logf("body length: %d", len(body))
}
