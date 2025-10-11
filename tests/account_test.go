package tests

import (
	"testing"

	"github.com/aerium-network/aerium/crypto"
	aerium "github.com/aerium-network/aerium/www/grpc/gen/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getAccount(t *testing.T, addr crypto.Address) *aerium.AccountInfo {
	t.Helper()

	res, err := tBlockchainClient.GetAccount(tCtx,
		&aerium.GetAccountRequest{Address: addr.String()})
	if err != nil {
		return nil
	}

	return res.Account
}

func TestGetAccount(t *testing.T) {
	acc := getAccount(t, crypto.TreasuryAddress)
	require.NotNil(t, acc)
	assert.Equal(t, int32(0), acc.Number)
}
