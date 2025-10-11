package hash_test

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/aerium-network/aerium/crypto/hash"
	"github.com/aerium-network/aerium/util/testsuite"
	"github.com/stretchr/testify/assert"
)

func TestHashFromString(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	hash1 := ts.RandHash()
	hash2, err := hash.FromString(hash1.String())
	assert.Contains(t, strings.ToUpper(hash1.String()), hash1.ShortString())
	assert.NoError(t, err)
	assert.Equal(t, hash1, hash2)

	_, err = hash.FromString("")
	assert.Error(t, err)

	_, err = hash.FromString("inv")
	assert.Error(t, err)

	_, err = hash.FromString("00")
	assert.Error(t, err)
}

func TestHashEmpty(t *testing.T) {
	_, err := hash.FromBytes(nil)
	assert.Error(t, err)

	_, err = hash.FromBytes([]byte{1})
	assert.Error(t, err)
}

func TestHash256(t *testing.T) {
	data := []byte("aerium")
	h1 := hash.Hash256(data)
	expected, _ := hex.DecodeString("993d37d2755bfd3efd4520b113c9216c02cd2ac352c576c969947433cae2c573")
	assert.Equal(t, expected, h1)
}

func TestHash160(t *testing.T) {
	data := []byte("aerium")
	h := hash.Hash160(data)
	expected, _ := hex.DecodeString("05f948a2cb316723a158adbe66bcbeee05512629")
	assert.Equal(t, expected, h)
}

func TestHashBasicCheck(t *testing.T) {
	h, err := hash.FromString("0000000000000000000000000000000000000000000000000000000000000000")
	assert.NoError(t, err)
	assert.True(t, h.IsUndef())
	assert.Equal(t, hash.UndefHash.Bytes(), h.Bytes())
}
