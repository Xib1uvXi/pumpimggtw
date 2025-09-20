package pumpimggtw_test

import (
	"context"
	"testing"

	"github.com/Xib1uvXi/pumpimggtw"
	"github.com/stretchr/testify/assert"
)

func TestNewPublicGateway(t *testing.T) {
	pubGtw := pumpimggtw.NewPublicGateway("https://ipfs.io")
	body, err := pubGtw.Get(context.Background(), "QmPwZqadGRwxJ1oraZM62YHkzfcD2Tqwu86U9BCEXABraw")
	if err != nil {
		t.Fatalf("Failed to get body: %v", err)
	}
	assert.True(t, len(body) > 0)
	t.Logf("body length: %d", len(body))
}
