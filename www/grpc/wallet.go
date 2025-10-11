package grpc

import (
	"context"
	"encoding/hex"
	"errors"

	"github.com/aerium-network/aerium/crypto"
	"github.com/aerium-network/aerium/wallet"
	aerium "github.com/aerium-network/aerium/www/grpc/gen/go"
)

//
// TODO: default_wallet should be loaded on starting the node.

type walletServer struct {
	*Server
	walletManager *wallet.Manager
}

func newWalletServer(server *Server, manager *wallet.Manager) *walletServer {
	return &walletServer{
		Server:        server,
		walletManager: manager,
	}
}

func (*walletServer) mapHistoryInfo(his []wallet.HistoryInfo) []*aerium.HistoryInfo {
	historyInfo := make([]*aerium.HistoryInfo, 0)
	for _, info := range his {
		historyInfo = append(historyInfo, &aerium.HistoryInfo{
			TransactionId: info.TxID,
			// Time:          uint32(hi.Time.Unix()),  // TODO: Fix me
			PayloadType: info.PayloadType,
			Description: info.Desc,
			Amount:      info.Amount.ToNanoAUM(),
		})
	}

	return historyInfo
}

func (s *walletServer) GetValidatorAddress(_ context.Context,
	req *aerium.GetValidatorAddressRequest,
) (*aerium.GetValidatorAddressResponse, error) {
	adr, err := s.walletManager.GetValidatorAddress(req.PublicKey)
	if err != nil {
		return nil, err
	}

	return &aerium.GetValidatorAddressResponse{
		Address: adr,
	}, nil
}

func (s *walletServer) CreateWallet(_ context.Context,
	req *aerium.CreateWalletRequest,
) (*aerium.CreateWalletResponse, error) {
	if req.WalletName == "" {
		return nil, errors.New("wallet name is required")
	}

	mnemonic, err := s.walletManager.CreateWallet(
		req.WalletName, req.Password,
	)
	if err != nil {
		return nil, err
	}

	return &aerium.CreateWalletResponse{
		Mnemonic: mnemonic,
	}, nil
}

func (s *walletServer) RestoreWallet(_ context.Context,
	req *aerium.RestoreWalletRequest,
) (*aerium.RestoreWalletResponse, error) {
	if req.WalletName == "" {
		return nil, errors.New("wallet name is required")
	}
	if req.Mnemonic == "" {
		return nil, errors.New("mnemonic is required")
	}

	if err := s.walletManager.RestoreWallet(
		req.WalletName, req.Mnemonic, req.Password,
	); err != nil {
		return nil, err
	}

	return &aerium.RestoreWalletResponse{
		WalletName: req.WalletName,
	}, nil
}

func (s *walletServer) LoadWallet(_ context.Context,
	req *aerium.LoadWalletRequest,
) (*aerium.LoadWalletResponse, error) {
	if err := s.walletManager.LoadWallet(req.WalletName, s.Address()); err != nil {
		return nil, err
	}

	return &aerium.LoadWalletResponse{
		WalletName: req.WalletName,
	}, nil
}

func (s *walletServer) UnloadWallet(_ context.Context,
	req *aerium.UnloadWalletRequest,
) (*aerium.UnloadWalletResponse, error) {
	if err := s.walletManager.UnloadWallet(req.WalletName); err != nil {
		return nil, err
	}

	return &aerium.UnloadWalletResponse{
		WalletName: req.WalletName,
	}, nil
}

func (s *walletServer) GetTotalBalance(_ context.Context,
	req *aerium.GetTotalBalanceRequest,
) (*aerium.GetTotalBalanceResponse, error) {
	balance, err := s.walletManager.TotalBalance(req.WalletName)
	if err != nil {
		return nil, err
	}

	return &aerium.GetTotalBalanceResponse{
		WalletName:   req.WalletName,
		TotalBalance: balance.ToNanoAUM(),
	}, nil
}

func (s *walletServer) SignRawTransaction(_ context.Context,
	req *aerium.SignRawTransactionRequest,
) (*aerium.SignRawTransactionResponse, error) {
	rawBytes, err := hex.DecodeString(req.RawTransaction)
	if err != nil {
		return nil, err
	}

	txID, data, err := s.walletManager.SignRawTransaction(
		req.WalletName, req.Password, rawBytes,
	)
	if err != nil {
		return nil, err
	}

	return &aerium.SignRawTransactionResponse{
		TransactionId:        hex.EncodeToString(txID),
		SignedRawTransaction: hex.EncodeToString(data),
	}, nil
}

func (s *walletServer) GetNewAddress(_ context.Context,
	req *aerium.GetNewAddressRequest,
) (*aerium.GetNewAddressResponse, error) {
	data, err := s.walletManager.GetNewAddress(
		req.WalletName,
		req.Label,
		req.Password,
		crypto.AddressType(req.AddressType),
	)
	if err != nil {
		return nil, err
	}

	return &aerium.GetNewAddressResponse{
		WalletName: req.WalletName,
		AddressInfo: &aerium.AddressInfo{
			Address:   data.Address,
			Label:     data.Label,
			PublicKey: data.PublicKey,
			Path:      data.Path,
		},
	}, nil
}

func (s *walletServer) GetAddressHistory(_ context.Context,
	req *aerium.GetAddressHistoryRequest,
) (*aerium.GetAddressHistoryResponse, error) {
	data, err := s.walletManager.AddressHistory(req.WalletName, req.Address)
	if err != nil {
		return nil, err
	}

	return &aerium.GetAddressHistoryResponse{
		HistoryInfo: s.mapHistoryInfo(data),
	}, nil
}

func (s *walletServer) SignMessage(_ context.Context,
	req *aerium.SignMessageRequest,
) (*aerium.SignMessageResponse, error) {
	sig, err := s.walletManager.SignMessage(req.Message, req.Password, req.Address, req.WalletName)
	if err != nil {
		return nil, err
	}

	return &aerium.SignMessageResponse{
		Signature: sig,
	}, nil
}

func (s *walletServer) GetTotalStake(_ context.Context,
	req *aerium.GetTotalStakeRequest,
) (*aerium.GetTotalStakeResponse, error) {
	stake, err := s.walletManager.TotalStake(req.WalletName)
	if err != nil {
		return nil, err
	}

	return &aerium.GetTotalStakeResponse{
		TotalStake: stake.ToNanoAUM(),
		WalletName: req.WalletName,
	}, nil
}

func (s *walletServer) GetAddressInfo(_ context.Context,
	req *aerium.GetAddressInfoRequest,
) (*aerium.GetAddressInfoResponse, error) {
	info, err := s.walletManager.GetAddressInfo(req.WalletName, req.Address)
	if err != nil {
		return nil, err
	}

	return &aerium.GetAddressInfoResponse{
		Address:    info.Address,
		Path:       info.Path,
		PublicKey:  info.PublicKey,
		Label:      info.Label,
		WalletName: req.WalletName,
	}, nil
}

func (s *walletServer) SetAddressLabel(_ context.Context,
	req *aerium.SetAddressLabelRequest,
) (*aerium.SetAddressLabelResponse, error) {
	return &aerium.SetAddressLabelResponse{}, s.walletMgr.SetAddressLabel(req.WalletName, req.Address, req.Label)
}

func (s *walletServer) ListWallet(_ context.Context,
	_ *aerium.ListWalletRequest,
) (*aerium.ListWalletResponse, error) {
	wallets, err := s.walletManager.ListWallet()
	if err != nil {
		return nil, err
	}

	return &aerium.ListWalletResponse{
		Wallets: wallets,
	}, nil
}

func (s *walletServer) GetWalletInfo(_ context.Context,
	req *aerium.GetWalletInfoRequest,
) (*aerium.GetWalletInfoResponse, error) {
	info, err := s.walletManager.WalletInfo(req.WalletName)
	if err != nil {
		return nil, err
	}

	return &aerium.GetWalletInfoResponse{
		WalletName: info.WalletName,
		Version:    info.Version,
		Network:    info.Network,
		Encrypted:  info.Encrypted,
		Uuid:       info.UUID,
		CreatedAt:  info.CreatedAt.Unix(),
	}, nil
}

func (s *walletServer) ListAddress(_ context.Context,
	req *aerium.ListAddressRequest,
) (*aerium.ListAddressResponse, error) {
	addrs, err := s.walletManager.ListAddress(req.WalletName)
	if err != nil {
		return nil, err
	}

	addrsPB := make([]*aerium.AddressInfo, 0, len(addrs))
	for _, addr := range addrs {
		addrsPB = append(addrsPB, &aerium.AddressInfo{
			Address:   addr.Address,
			Label:     addr.Label,
			PublicKey: addr.PublicKey,
			Path:      addr.Path,
		})
	}

	return &aerium.ListAddressResponse{
		Data: addrsPB,
	}, nil
}
