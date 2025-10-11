package grpc

import (
	"context"
	"testing"

	"github.com/aerium-network/aerium/crypto/bls"
	"github.com/aerium-network/aerium/util/testsuite"
	aerium "github.com/aerium-network/aerium/www/grpc/gen/go"
	"github.com/stretchr/testify/assert"
)

func TestSignMessageWithPrivateKey(t *testing.T) {
	conf := testConfig()
	td := setup(t, conf)
	conn, client := td.utilClient(t)

	msg := "aerium"
	prvStr := "SECRET1PDRWTLP5PX0FAHDX39GXZJP7FKZFALML0D5U9TT9KVQHDUC99CMGQQJVK67"
	invalidPrvStr := "INVSECRET1PDRWTLP5PX0FAHDX39GXZJP7FKZFALML0D5U9TT9KVQHDUC99CMGQQJVK67"
	expectedSig := "b84f476cb5114560056a110db2951f5a0eb2efb2830e684779b9697d70736add77b8fcc2debd0f0b1f953e589d2f7646"

	t.Run("", func(t *testing.T) {
		res, err := client.SignMessageWithPrivateKey(context.Background(),
			&aerium.SignMessageWithPrivateKeyRequest{
				Message:    msg,
				PrivateKey: prvStr,
			})

		assert.Nil(t, err)
		assert.Equal(t, expectedSig, res.Signature)
	})

	t.Run("", func(t *testing.T) {
		res, err := client.SignMessageWithPrivateKey(context.Background(),
			&aerium.SignMessageWithPrivateKeyRequest{
				Message:    msg,
				PrivateKey: invalidPrvStr,
			})

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	assert.Nil(t, conn.Close(), "Error closing connection")
	td.StopServer()
}

func TestSignMessageWithED25519PrivateKey(t *testing.T) {
	conf := testConfig()
	td := setup(t, conf)
	conn, client := td.utilClient(t)

	msg := "aerium"
	prvStr := "SECRET1RYY62A96X25ZAL4DPL5Z63G83GCSFCCQ7K0CMQD3MFNLYK3A6R26QUUK3Y0"
	invalidPrvStr := "INVSECRET1RYY62A96X25ZAL4DPL5Z63G83GCSFCCQ7K0CMQD3MFNLYK3A6R26QUUK3Y0"
	expectedSig := "db867d6dcf0a7f1c8731ca26cbfe5510901a87fbe76eba3a881a9cf8b4cbf22a29530d3dca8b9665edc7f171fdc568ebf0564f67b374178953d925f816be2b0c"

	t.Run("", func(t *testing.T) {
		res, err := client.SignMessageWithPrivateKey(context.Background(),
			&aerium.SignMessageWithPrivateKeyRequest{
				Message:    msg,
				PrivateKey: prvStr,
			})

		assert.Nil(t, err)
		assert.Equal(t, expectedSig, res.Signature)
	})

	t.Run("", func(t *testing.T) {
		res, err := client.SignMessageWithPrivateKey(context.Background(),
			&aerium.SignMessageWithPrivateKeyRequest{
				Message:    msg,
				PrivateKey: invalidPrvStr,
			})

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	assert.Nil(t, conn.Close(), "Error closing connection")
	td.StopServer()
}

func TestVerifyMessage(t *testing.T) {
	conf := testConfig()
	td := setup(t, conf)
	conn, client := td.utilClient(t)

	msg := "aerium"
	pubStr := "public1p4u8hfytl2pj6l9rj0t54gxcdmna4hq52ncqkkqjf3arha5mlk3x4mzpyjkhmdl20jae7f65aamjr" +
		"vqcvf4sudcapz52ctcwc8r9wz3z2gwxs38880cgvfy49ta5ssyjut05myd4zgmjqstggmetyuyg7v5jhx47a"
	sigStr := "b84f476cb5114560056a110db2951f5a0eb2efb2830e684779b9697d70736add77b8fcc2debd0f0b1f953e589d2f7646"
	invalidSigStr := "113d67a8624cbb7972b29328e15ec76cc846076ccf00a9e94d991c677846f334ae4ba4551396fbcd6d1cab7593baf3c9"

	t.Run("valid message", func(t *testing.T) {
		res, err := client.VerifyMessage(context.Background(),
			&aerium.VerifyMessageRequest{
				Message:   msg,
				Signature: sigStr,
				PublicKey: pubStr,
			})
		assert.Nil(t, err)
		assert.True(t, res.IsValid)
	})

	t.Run("invalid message", func(t *testing.T) {
		res, err := client.VerifyMessage(context.Background(),
			&aerium.VerifyMessageRequest{
				Message:   msg,
				Signature: invalidSigStr,
				PublicKey: pubStr,
			})

		assert.Nil(t, err)
		assert.False(t, res.IsValid)
	})

	assert.Nil(t, conn.Close(), "Error closing connection")
	td.StopServer()
}

func TestVerifyED25519Message(t *testing.T) {
	conf := testConfig()
	td := setup(t, conf)
	conn, client := td.utilClient(t)

	msg := "aerium"
	pubStr := "public1rvqxnpfph8tnc3ck55z85w285t5jetylmmktr9wlzs0zvx7kx500szxfudh"
	sigStr := "db867d6dcf0a7f1c8731ca26cbfe5510901a87fbe76eba3a881a9cf8b4cbf22a29530d3dca8b9665edc7f171fdc568ebf0564f67b374178953d925f816be2b0c"
	invalidSigStr := "001aaa09c408bfcf7e79dd90c583eeeaefe7c732ca5643cfb2ea7a6d22105b" +
		"874a412080525a855bbd5df94110a7d0083d6e386e016ecf8b7f522c339f79d305"

	t.Run("valid message", func(t *testing.T) {
		res, err := client.VerifyMessage(context.Background(),
			&aerium.VerifyMessageRequest{
				Message:   msg,
				Signature: sigStr,
				PublicKey: pubStr,
			})
		assert.Nil(t, err)
		assert.True(t, res.IsValid)
	})

	t.Run("invalid message", func(t *testing.T) {
		res, err := client.VerifyMessage(context.Background(),
			&aerium.VerifyMessageRequest{
				Message:   msg,
				Signature: invalidSigStr,
				PublicKey: pubStr,
			})

		assert.Nil(t, err)
		assert.False(t, res.IsValid)
	})

	assert.Nil(t, conn.Close(), "Error closing connection")
	td.StopServer()
}

func TestBLSPublicKeyAggregation(t *testing.T) {
	ts := testsuite.NewTestSuite(t)
	conf := testConfig()
	td := setup(t, conf)
	conn, client := td.utilClient(t)

	pub1, _ := ts.RandBLSKeyPair()
	pub2, _ := ts.RandBLSKeyPair()
	pub3, _ := ts.RandBLSKeyPair()
	aggPub := bls.PublicKeyAggregate(pub1, pub2, pub3)
	invalidPub := "invalidpub"

	t.Run("zero public keys", func(t *testing.T) {
		res, err := client.PublicKeyAggregation(context.Background(),
			&aerium.PublicKeyAggregationRequest{
				PublicKeys: []string{},
			})

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("only one public key", func(t *testing.T) {
		res, err := client.PublicKeyAggregation(context.Background(),
			&aerium.PublicKeyAggregationRequest{
				PublicKeys: []string{pub1.String()},
			})

		assert.Nil(t, err)
		assert.Equal(t, pub1.String(), res.PublicKey)
	})

	t.Run("invalid public key", func(t *testing.T) {
		res, err := client.PublicKeyAggregation(context.Background(),
			&aerium.PublicKeyAggregationRequest{
				PublicKeys: []string{pub1.String(), pub2.String(), invalidPub, pub3.String()},
			})

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("valid public keys", func(t *testing.T) {
		res, err := client.PublicKeyAggregation(context.Background(),
			&aerium.PublicKeyAggregationRequest{
				PublicKeys: []string{pub1.String(), pub2.String(), pub3.String()},
			})

		assert.Nil(t, err)
		assert.Equal(t, aggPub.String(), res.PublicKey)
	})

	assert.Nil(t, conn.Close(), "Error closing connection")
	td.StopServer()
}

func TestBLSSignatureAggregation(t *testing.T) {
	ts := testsuite.NewTestSuite(t)
	conf := testConfig()
	td := setup(t, conf)
	conn, client := td.utilClient(t)

	sig1 := ts.RandBLSSignature()
	sig2 := ts.RandBLSSignature()
	sig3 := ts.RandBLSSignature()
	aggSig := bls.SignatureAggregate(sig1, sig2, sig3)
	invalidSig := "invalidsig"

	t.Run("zero signatures", func(t *testing.T) {
		res, err := client.SignatureAggregation(context.Background(),
			&aerium.SignatureAggregationRequest{
				Signatures: []string{},
			})

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("only one signature", func(t *testing.T) {
		res, err := client.SignatureAggregation(context.Background(),
			&aerium.SignatureAggregationRequest{
				Signatures: []string{sig1.String()},
			})

		assert.Nil(t, err)
		assert.Equal(t, sig1.String(), res.Signature)
	})

	t.Run("invalid signature", func(t *testing.T) {
		res, err := client.SignatureAggregation(context.Background(),
			&aerium.SignatureAggregationRequest{
				Signatures: []string{sig1.String(), sig2.String(), invalidSig, sig3.String()},
			})

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("valid signatures", func(t *testing.T) {
		res, err := client.SignatureAggregation(context.Background(),
			&aerium.SignatureAggregationRequest{
				Signatures: []string{sig1.String(), sig2.String(), sig3.String()},
			})

		assert.Nil(t, err)
		assert.Equal(t, aggSig.String(), res.Signature)
	})

	assert.Nil(t, conn.Close(), "Error closing connection")
	td.StopServer()
}
