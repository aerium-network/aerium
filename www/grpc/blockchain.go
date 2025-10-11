package grpc

import (
	"context"
	"encoding/hex"

	"github.com/aerium-network/aerium/crypto"
	"github.com/aerium-network/aerium/crypto/hash"
	"github.com/aerium-network/aerium/types/account"
	"github.com/aerium-network/aerium/types/validator"
	"github.com/aerium-network/aerium/types/vote"
	aerium "github.com/aerium-network/aerium/www/grpc/gen/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type blockchainServer struct {
	*Server
}

func newBlockchainServer(server *Server) *blockchainServer {
	return &blockchainServer{
		Server: server,
	}
}

func (s *blockchainServer) GetBlockchainInfo(_ context.Context,
	_ *aerium.GetBlockchainInfoRequest,
) (*aerium.GetBlockchainInfoResponse, error) {
	vals := s.state.CommitteeValidators()
	valInfos := make([]*aerium.ValidatorInfo, 0, len(vals))
	for _, val := range vals {
		valInfos = append(valInfos, s.validatorToProto(val))
	}

	committeeProtocolVersions := make(map[int32]float64)
	for k, v := range s.state.CommitteeProtocolVersions() {
		committeeProtocolVersions[int32(k)] = v
	}

	stats := s.state.Stats()

	return &aerium.GetBlockchainInfoResponse{
		LastBlockHeight:           stats.LastBlockHeight,
		LastBlockHash:             stats.LastBlockHash.String(),
		TotalAccounts:             stats.TotalAccounts,
		TotalValidators:           stats.TotalValidators,
		ActiveValidators:          stats.ActiveValidators,
		TotalPower:                stats.TotalPower,
		CommitteePower:            stats.CommitteePower,
		IsPruned:                  stats.IsPruned,
		PruningHeight:             stats.PruningHeight,
		LastBlockTime:             stats.LastBlockTime.Unix(),
		CommitteeValidators:       valInfos,
		CommitteeProtocolVersions: committeeProtocolVersions,
	}, nil
}

func (s *blockchainServer) GetConsensusInfo(_ context.Context,
	_ *aerium.GetConsensusInfoRequest,
) (*aerium.GetConsensusInfoResponse, error) {
	instances := make([]*aerium.ConsensusInfo, 0)

	for _, cons := range s.consMgr.Instances() {
		height, round := cons.HeightRound()
		votes := cons.AllVotes()
		voteInfos := make([]*aerium.VoteInfo, 0, len(votes))
		for _, v := range votes {
			voteInfos = append(voteInfos, s.voteToProto(v))
		}

		instances = append(instances,
			&aerium.ConsensusInfo{
				Address: cons.ConsensusKey().ValidatorAddress().String(),
				Active:  cons.IsActive(),
				Height:  height,
				Round:   int32(round),
				Votes:   voteInfos,
			})
	}

	var proposalInfo *aerium.ProposalInfo
	prop := s.consMgr.Proposal()
	if prop != nil {
		var blockData string
		data, err := prop.Block().Bytes()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		blockData = hex.EncodeToString(data)

		proposalInfo = &aerium.ProposalInfo{
			Height:    prop.Height(),
			Round:     int32(prop.Round()),
			BlockData: blockData,
			Signature: prop.Signature().String(),
		}
	}

	return &aerium.GetConsensusInfoResponse{
		Instances: instances,
		Proposal:  proposalInfo,
	}, nil
}

func (s *blockchainServer) GetBlockHash(_ context.Context,
	req *aerium.GetBlockHashRequest,
) (*aerium.GetBlockHashResponse, error) {
	height := req.GetHeight()
	h := s.state.BlockHash(height)
	if h.IsUndef() {
		return nil, status.Errorf(codes.NotFound, "block not found with this height")
	}

	return &aerium.GetBlockHashResponse{
		Hash: h.String(),
	}, nil
}

func (s *blockchainServer) GetBlockHeight(_ context.Context,
	req *aerium.GetBlockHeightRequest,
) (*aerium.GetBlockHeightResponse, error) {
	h, err := hash.FromString(req.GetHash())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid hash: %v", err)
	}
	height := s.state.BlockHeight(h)
	if height == 0 {
		return nil, status.Errorf(codes.NotFound, "block not found with this hash")
	}

	return &aerium.GetBlockHeightResponse{
		Height: height,
	}, nil
}

func (s *blockchainServer) GetBlock(_ context.Context,
	req *aerium.GetBlockRequest,
) (*aerium.GetBlockResponse, error) {
	height := req.GetHeight()
	cBlk := s.state.CommittedBlock(height)
	if cBlk == nil {
		return nil, status.Errorf(codes.NotFound, "block not found")
	}
	res := &aerium.GetBlockResponse{
		Height: cBlk.Height,
		Hash:   cBlk.BlockHash.String(),
	}

	switch req.Verbosity {
	case aerium.BlockVerbosity_BLOCK_VERBOSITY_DATA:
		res.Data = hex.EncodeToString(cBlk.Data)

	case aerium.BlockVerbosity_BLOCK_VERBOSITY_INFO,
		aerium.BlockVerbosity_BLOCK_VERBOSITY_TRANSACTIONS:
		block, err := cBlk.ToBlock()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		blockTime := block.Header().UnixTime()
		seed := block.Header().SortitionSeed()
		cert := block.PrevCertificate()
		var prevCert *aerium.CertificateInfo

		if cert != nil {
			committers := make([]int32, len(cert.Committers()))
			for i, n := range cert.Committers() {
				committers[i] = n
			}
			absentees := make([]int32, len(cert.Absentees()))
			for i, n := range cert.Absentees() {
				absentees[i] = n
			}
			prevCert = &aerium.CertificateInfo{
				Hash:       cert.Hash().String(),
				Round:      int32(cert.Round()),
				Committers: committers,
				Absentees:  absentees,
				Signature:  cert.Signature().String(),
			}
		}
		header := &aerium.BlockHeaderInfo{
			Version:         int32(block.Header().Version()),
			PrevBlockHash:   block.Header().PrevBlockHash().String(),
			StateRoot:       block.Header().StateRoot().String(),
			SortitionSeed:   hex.EncodeToString(seed[:]),
			ProposerAddress: block.Header().ProposerAddress().String(),
		}

		trxs := make([]*aerium.TransactionInfo, 0, block.Transactions().Len())
		for _, trx := range block.Transactions() {
			if req.Verbosity == aerium.BlockVerbosity_BLOCK_VERBOSITY_INFO {
				data, _ := trx.Bytes()
				trxs = append(trxs, &aerium.TransactionInfo{
					Id:   trx.ID().String(),
					Data: hex.EncodeToString(data),
				})
			} else {
				trxs = append(trxs, transactionToProto(trx))
			}
		}

		res.BlockTime = blockTime
		res.Header = header
		res.Txs = trxs
		res.PrevCert = prevCert
	}

	return res, nil
}

func (s *blockchainServer) GetAccount(_ context.Context,
	req *aerium.GetAccountRequest,
) (*aerium.GetAccountResponse, error) {
	addr, err := crypto.AddressFromString(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address: %v", err)
	}
	acc := s.state.AccountByAddress(addr)
	if acc == nil {
		return nil, status.Errorf(codes.NotFound, "account not found")
	}
	res := &aerium.GetAccountResponse{
		Account: s.accountToProto(addr, acc),
	}

	return res, nil
}

func (s *blockchainServer) GetValidatorByNumber(_ context.Context,
	req *aerium.GetValidatorByNumberRequest,
) (*aerium.GetValidatorResponse, error) {
	val := s.state.ValidatorByNumber(req.Number)
	if val == nil {
		return nil, status.Errorf(codes.NotFound, "validator not found")
	}

	return &aerium.GetValidatorResponse{
		Validator: s.validatorToProto(val),
	}, nil
}

func (s *blockchainServer) GetValidator(_ context.Context,
	req *aerium.GetValidatorRequest,
) (*aerium.GetValidatorResponse, error) {
	addr, err := crypto.AddressFromString(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid validator address: %v", err.Error())
	}
	val := s.state.ValidatorByAddress(addr)
	if val == nil {
		return nil, status.Errorf(codes.NotFound, "validator not found")
	}

	return &aerium.GetValidatorResponse{
		Validator: s.validatorToProto(val),
	}, nil
}

func (s *blockchainServer) GetValidatorAddresses(_ context.Context,
	_ *aerium.GetValidatorAddressesRequest,
) (*aerium.GetValidatorAddressesResponse, error) {
	addresses := s.state.ValidatorAddresses()
	addressesPB := make([]string, 0, len(addresses))
	for _, address := range addresses {
		addressesPB = append(addressesPB, address.String())
	}

	return &aerium.GetValidatorAddressesResponse{Addresses: addressesPB}, nil
}

func (s *blockchainServer) GetPublicKey(_ context.Context,
	req *aerium.GetPublicKeyRequest,
) (*aerium.GetPublicKeyResponse, error) {
	addr, err := crypto.AddressFromString(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid account address: %v", err.Error())
	}

	publicKey, err := s.state.PublicKey(addr)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "public key not found")
	}

	return &aerium.GetPublicKeyResponse{PublicKey: publicKey.String()}, nil
}

func (s *blockchainServer) GetTxPoolContent(_ context.Context,
	req *aerium.GetTxPoolContentRequest,
) (*aerium.GetTxPoolContentResponse, error) {
	result := make([]*aerium.TransactionInfo, 0)

	for _, t := range s.state.AllPendingTxs() {
		if req.PayloadType == aerium.PayloadType_PAYLOAD_TYPE_UNSPECIFIED ||
			req.PayloadType == aerium.PayloadType(t.Payload().Type()) {
			result = append(result, transactionToProto(t))
		}
	}

	return &aerium.GetTxPoolContentResponse{
		Txs: result,
	}, nil
}

func (s *blockchainServer) validatorToProto(val *validator.Validator) *aerium.ValidatorInfo {
	data, _ := val.Bytes()

	return &aerium.ValidatorInfo{
		Hash:                val.Hash().String(),
		Data:                hex.EncodeToString(data),
		PublicKey:           val.PublicKey().String(),
		Address:             val.Address().String(),
		Number:              val.Number(),
		Stake:               val.Stake().ToNanoAUM(),
		LastBondingHeight:   val.LastBondingHeight(),
		LastSortitionHeight: val.LastSortitionHeight(),
		UnbondingHeight:     val.UnbondingHeight(),
		AvailabilityScore:   s.state.AvailabilityScore(val.Number()),
		ProtocolVersion:     int32(val.ProtocolVersion()),
	}
}

func (*blockchainServer) accountToProto(addr crypto.Address, acc *account.Account) *aerium.AccountInfo {
	data, _ := acc.Bytes()

	return &aerium.AccountInfo{
		Hash:    acc.Hash().String(),
		Data:    hex.EncodeToString(data),
		Number:  acc.Number(),
		Balance: acc.Balance().ToNanoAUM(),
		Address: addr.String(),
	}
}

func (*blockchainServer) voteToProto(vte *vote.Vote) *aerium.VoteInfo {
	cpRound := int32(0)
	cpValue := int32(0)
	if vte.IsCPVote() {
		cpRound = int32(vte.CPRound())
		cpValue = int32(vte.CPValue())
	}

	return &aerium.VoteInfo{
		Type:      aerium.VoteType(vte.Type()),
		Voter:     vte.Signer().String(),
		BlockHash: vte.BlockHash().String(),
		Round:     int32(vte.Round()),
		CpRound:   cpRound,
		CpValue:   cpValue,
	}
}
