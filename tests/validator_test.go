package tests

import (
	"testing"

	"github.com/aerium-network/aerium/crypto"
	aerium "github.com/aerium-network/aerium/www/grpc/gen/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getValidator(t *testing.T, addr crypto.Address) *aerium.ValidatorInfo {
	t.Helper()

	res, err := tBlockchainClient.GetValidator(tCtx,
		&aerium.GetValidatorRequest{Address: addr.String()})
	if err != nil {
		return nil
	}

	return res.Validator
}

func TestGetValidator(t *testing.T) {
	val := getValidator(t, tValKeys[tNodeIdx2][0].Address())
	require.NotNil(t, val)
	assert.Equal(t, int32(1), val.Number)
}
