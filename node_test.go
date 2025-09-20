package pumpimggtw_test

import (
	"testing"

	"github.com/Xib1uvXi/pumpimggtw"
	"github.com/stretchr/testify/assert"
)

func TestNewIPFSNode(t *testing.T) {
	node, err := pumpimggtw.NewIPFSNode()
	assert.NoError(t, err)

	cidStr := "bafybeig6gucsv5kqgv5agk6u6ccar3gcrxkggl2bmgysve5qtcce2mnvuy"

	err = node.Pin(cidStr)
	assert.NoError(t, err)

	// resp, err := node.Get(cidStr)
	// assert.NoError(t, err)
	// assert.True(t, len(resp) > 0)
	// t.Logf("res length: %d", len(resp))
}
