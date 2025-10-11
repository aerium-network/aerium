package grpc

import (
	"context"
	"errors"
	"fmt"

	"github.com/aerium-network/aerium/crypto"
	"github.com/aerium-network/aerium/crypto/bls"
	"github.com/aerium-network/aerium/crypto/ed25519"
	aerium "github.com/aerium-network/aerium/www/grpc/gen/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type utilServer struct {
	*Server
}

func newUtilsServer(server *Server) *utilServer {
	return &utilServer{
		Server: server,
	}
}

func (u *utilServer) SignMessageWithPrivateKey(_ context.Context,
	req *aerium.SignMessageWithPrivateKeyRequest,
) (*aerium.SignMessageWithPrivateKeyResponse, error) {
	prvKey, err := u.privateKeyFromString(req.PrivateKey)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	sig := prvKey.Sign([]byte(req.Message)).String()

	return &aerium.SignMessageWithPrivateKeyResponse{
		Signature: sig,
	}, nil
}

func (u *utilServer) VerifyMessage(_ context.Context,
	req *aerium.VerifyMessageRequest,
) (*aerium.VerifyMessageResponse, error) {
	pubKey, sig, err := u.publicKeyAndSigFromString(req.PublicKey, req.Signature)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err = pubKey.Verify([]byte(req.Message), sig); err == nil {
		return &aerium.VerifyMessageResponse{
			IsValid: true,
		}, nil
	}

	return &aerium.VerifyMessageResponse{
		IsValid: false,
	}, nil
}

func (*utilServer) PublicKeyAggregation(_ context.Context,
	req *aerium.PublicKeyAggregationRequest,
) (*aerium.PublicKeyAggregationResponse, error) {
	if len(req.PublicKeys) == 0 {
		return nil, status.Error(codes.InvalidArgument, "no public keys provided")
	}
	pubs := make([]*bls.PublicKey, len(req.PublicKeys))
	for i, pubKey := range req.PublicKeys {
		p, err := bls.PublicKeyFromString(pubKey)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid public key %s", pubKey))
		}
		pubs[i] = p
	}

	pk := bls.PublicKeyAggregate(pubs...)

	return &aerium.PublicKeyAggregationResponse{
		PublicKey: pk.String(),
		Address:   pk.AccountAddress().String(),
	}, nil
}

func (*utilServer) SignatureAggregation(_ context.Context,
	req *aerium.SignatureAggregationRequest,
) (*aerium.SignatureAggregationResponse, error) {
	if len(req.Signatures) == 0 {
		return nil, status.Error(codes.InvalidArgument, "no signatures provided")
	}
	sigs := make([]*bls.Signature, len(req.Signatures))
	for i, sig := range req.Signatures {
		s, err := bls.SignatureFromString(sig)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid signature %s", sig))
		}
		sigs[i] = s
	}

	return &aerium.SignatureAggregationResponse{
		Signature: bls.SignatureAggregate(sigs...).String(),
	}, nil
}

func (*utilServer) privateKeyFromString(prvStr string) (crypto.PrivateKey, error) {
	blsPrv, err := bls.PrivateKeyFromString(prvStr)
	if err == nil {
		return blsPrv, nil
	}

	ed25519Prv, err := ed25519.PrivateKeyFromString(prvStr)
	if err == nil {
		return ed25519Prv, nil
	}

	return nil, errors.New("invalid Private Key")
}

func (*utilServer) publicKeyAndSigFromString(pubStr, sigStr string) (crypto.PublicKey, crypto.Signature, error) {
	blsPub, err := bls.PublicKeyFromString(pubStr)
	if err == nil {
		blsSig, err := bls.SignatureFromString(sigStr)
		if err != nil {
			return nil, nil, errors.New("invalid BLS signature")
		}

		return blsPub, blsSig, nil
	}

	ed25519Pub, err := ed25519.PublicKeyFromString(pubStr)
	if err == nil {
		ed25519Sig, err := ed25519.SignatureFromString(sigStr)
		if err != nil {
			return nil, nil, errors.New("invalid Ed25519 signature")
		}

		return ed25519Pub, ed25519Sig, nil
	}

	return nil, nil, errors.New("invalid Public Key")
}
