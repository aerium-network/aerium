package grpc

import (
	"context"
	"encoding/hex"

	"github.com/aerium-network/aerium/crypto"
	"github.com/aerium-network/aerium/crypto/bls"
	"github.com/aerium-network/aerium/crypto/hash"
	"github.com/aerium-network/aerium/types/amount"
	"github.com/aerium-network/aerium/types/tx"
	"github.com/aerium-network/aerium/types/tx/payload"
	"github.com/aerium-network/aerium/util/logger"
	aerium "github.com/aerium-network/aerium/www/grpc/gen/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type transactionServer struct {
	*Server
}

func newTransactionServer(server *Server) *transactionServer {
	return &transactionServer{
		Server: server,
	}
}

func (s *transactionServer) GetTransaction(_ context.Context,
	req *aerium.GetTransactionRequest,
) (*aerium.GetTransactionResponse, error) {
	id, err := hash.FromString(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid transaction ID: %v", err.Error())
	}

	committedTx := s.state.CommittedTx(id)
	if committedTx == nil {
		return nil, status.Errorf(codes.InvalidArgument, "transaction not found")
	}

	res := &aerium.GetTransactionResponse{
		BlockHeight: committedTx.Height,
		BlockTime:   committedTx.BlockTime,
	}

	switch req.Verbosity {
	case aerium.TransactionVerbosity_TRANSACTION_VERBOSITY_DATA:
		res.Transaction = &aerium.TransactionInfo{
			Id:   committedTx.TxID.String(),
			Data: hex.EncodeToString(committedTx.Data),
		}

	case aerium.TransactionVerbosity_TRANSACTION_VERBOSITY_INFO:
		trx, err := committedTx.ToTx()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "%s", err.Error())
		}
		res.Transaction = transactionToProto(trx)
	}

	return res, nil
}

func (s *transactionServer) BroadcastTransaction(_ context.Context,
	req *aerium.BroadcastTransactionRequest,
) (*aerium.BroadcastTransactionResponse, error) {
	b, err := hex.DecodeString(req.SignedRawTransaction)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid signed transaction")
	}

	trx, err := tx.FromBytes(b)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "couldn't decode transaction: %v", err.Error())
	}

	if err := trx.BasicCheck(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "couldn't verify transaction: %v", err.Error())
	}

	if err := s.state.AddPendingTxAndBroadcast(trx); err != nil {
		return nil, status.Errorf(codes.Canceled, "couldn't add to transaction pool: %v", err.Error())
	}

	return &aerium.BroadcastTransactionResponse{
		Id: trx.ID().String(),
	}, nil
}

func (s *transactionServer) CalculateFee(_ context.Context,
	req *aerium.CalculateFeeRequest,
) (*aerium.CalculateFeeResponse, error) {
	amt := amount.Amount(req.Amount)
	fee := s.state.CalculateFee(amt, payload.Type(req.PayloadType))

	if req.FixedAmount {
		amt -= fee
	}

	return &aerium.CalculateFeeResponse{
		Amount: amt.ToNanoAUM(),
		Fee:    fee.ToNanoAUM(),
	}, nil
}

func (s *transactionServer) GetRawTransferTransaction(_ context.Context,
	req *aerium.GetRawTransferTransactionRequest,
) (*aerium.GetRawTransactionResponse, error) {
	sender, err := crypto.AddressFromString(req.Sender)
	if err != nil {
		return nil, err
	}

	receiver, err := crypto.AddressFromString(req.Receiver)
	if err != nil {
		return nil, err
	}

	amt := amount.Amount(req.Amount)
	fee := s.getFee(req.Fee, amt)
	lockTime := s.getLockTime(req.LockTime)

	transferTx := tx.NewTransferTx(lockTime, sender, receiver, amt, fee, tx.WithMemo(req.Memo))
	rawTx, err := transferTx.Bytes()
	if err != nil {
		return nil, err
	}

	return &aerium.GetRawTransactionResponse{
		RawTransaction: hex.EncodeToString(rawTx),
	}, nil
}

func (s *transactionServer) GetRawBondTransaction(_ context.Context,
	req *aerium.GetRawBondTransactionRequest,
) (*aerium.GetRawTransactionResponse, error) {
	sender, err := crypto.AddressFromString(req.Sender)
	if err != nil {
		return nil, err
	}

	receiver, err := crypto.AddressFromString(req.Receiver)
	if err != nil {
		return nil, err
	}

	var publicKey *bls.PublicKey
	if req.PublicKey != "" {
		publicKey, err = bls.PublicKeyFromString(req.PublicKey)
		if err != nil {
			return nil, err
		}
	} else {
		publicKey = nil
	}

	amt := amount.Amount(req.Stake)
	fee := s.getFee(req.Fee, amt)
	lockTime := s.getLockTime(req.LockTime)

	bondTx := tx.NewBondTx(lockTime, sender, receiver, publicKey, amt, fee, tx.WithMemo(req.Memo))
	rawTx, err := bondTx.Bytes()
	if err != nil {
		return nil, err
	}

	return &aerium.GetRawTransactionResponse{
		RawTransaction: hex.EncodeToString(rawTx),
	}, nil
}

func (s *transactionServer) GetRawUnbondTransaction(_ context.Context,
	req *aerium.GetRawUnbondTransactionRequest,
) (*aerium.GetRawTransactionResponse, error) {
	validatorAddr, err := crypto.AddressFromString(req.ValidatorAddress)
	if err != nil {
		return nil, err
	}

	lockTime := s.getLockTime(req.LockTime)

	unbondTx := tx.NewUnbondTx(lockTime, validatorAddr, tx.WithMemo(req.Memo))
	rawTx, err := unbondTx.Bytes()
	if err != nil {
		return nil, err
	}

	return &aerium.GetRawTransactionResponse{
		RawTransaction: hex.EncodeToString(rawTx),
	}, nil
}

func (s *transactionServer) GetRawWithdrawTransaction(_ context.Context,
	req *aerium.GetRawWithdrawTransactionRequest,
) (*aerium.GetRawTransactionResponse, error) {
	validatorAddr, err := crypto.AddressFromString(req.ValidatorAddress)
	if err != nil {
		return nil, err
	}

	accountAddr, err := crypto.AddressFromString(req.AccountAddress)
	if err != nil {
		return nil, err
	}

	amt := amount.Amount(req.Amount)
	fee := s.getFee(req.Fee, amt)
	lockTime := s.getLockTime(req.LockTime)

	withdrawTx := tx.NewWithdrawTx(lockTime, validatorAddr, accountAddr, amt, fee, tx.WithMemo(req.Memo))
	rawTx, err := withdrawTx.Bytes()
	if err != nil {
		return nil, err
	}

	return &aerium.GetRawTransactionResponse{
		RawTransaction: hex.EncodeToString(rawTx),
	}, nil
}

func (s *transactionServer) GetRawBatchTransferTransaction(_ context.Context,
	req *aerium.GetRawBatchTransferTransactionRequest,
) (*aerium.GetRawTransactionResponse, error) {
	sender, err := crypto.AddressFromString(req.Sender)
	if err != nil {
		return nil, err
	}

	totalAmount := amount.Amount(0)

	recipients := make([]payload.BatchRecipient, 0, len(req.Recipients))
	for _, recipient := range req.Recipients {
		receiver, err := crypto.AddressFromString(recipient.Receiver)
		if err != nil {
			return nil, err
		}

		amt := amount.Amount(recipient.Amount)

		recipients = append(recipients, payload.BatchRecipient{
			To:     receiver,
			Amount: amt,
		})

		totalAmount += amt
	}

	fee := s.getFee(req.Fee, totalAmount)
	lockTime := s.getLockTime(req.LockTime)

	batchTransferTx := tx.NewBatchTransferTx(lockTime, sender, recipients, fee, tx.WithMemo(req.Memo))
	rawTx, err := batchTransferTx.Bytes()
	if err != nil {
		return nil, err
	}

	return &aerium.GetRawTransactionResponse{
		RawTransaction: hex.EncodeToString(rawTx),
	}, nil
}

func (s *transactionServer) getFee(f int64, amt amount.Amount) amount.Amount {
	fee := amount.Amount(f)
	if fee == 0 {
		fee = s.state.CalculateFee(amt, payload.TypeTransfer)
	}

	return fee
}

func (s *transactionServer) getLockTime(lockTime uint32) uint32 {
	if lockTime == 0 {
		lockTime = s.state.LastBlockHeight()
	}

	return lockTime
}

func transactionToProto(trx *tx.Tx) *aerium.TransactionInfo {
	trxInfo := &aerium.TransactionInfo{
		Id:          trx.ID().String(),
		Version:     int32(trx.Version()),
		LockTime:    trx.LockTime(),
		Fee:         trx.Fee().ToNanoAUM(),
		Value:       trx.Payload().Value().ToNanoAUM(),
		PayloadType: aerium.PayloadType(trx.Payload().Type()),
		Memo:        trx.Memo(),
	}

	if trx.PublicKey() != nil {
		trxInfo.PublicKey = trx.PublicKey().String()
	}

	if trx.Signature() != nil {
		trxInfo.Signature = trx.Signature().String()
	}

	switch trx.Payload().Type() {
	case payload.TypeTransfer:
		pld := trx.Payload().(*payload.TransferPayload)
		trxInfo.Payload = &aerium.TransactionInfo_Transfer{
			Transfer: &aerium.PayloadTransfer{
				Sender:   pld.From.String(),
				Receiver: pld.To.String(),
				Amount:   pld.Amount.ToNanoAUM(),
			},
		}
	case payload.TypeBond:
		pld := trx.Payload().(*payload.BondPayload)

		publicKeyStr := ""
		if pld.PublicKey != nil {
			publicKeyStr = pld.PublicKey.String()
		}

		trxInfo.Payload = &aerium.TransactionInfo_Bond{
			Bond: &aerium.PayloadBond{
				Sender:    pld.From.String(),
				Receiver:  pld.To.String(),
				Stake:     pld.Stake.ToNanoAUM(),
				PublicKey: publicKeyStr,
			},
		}

	case payload.TypeSortition:
		pld := trx.Payload().(*payload.SortitionPayload)
		trxInfo.Payload = &aerium.TransactionInfo_Sortition{
			Sortition: &aerium.PayloadSortition{
				Address: pld.Validator.String(),
				Proof:   hex.EncodeToString(pld.Proof[:]),
			},
		}
	case payload.TypeUnbond:
		pld := trx.Payload().(*payload.UnbondPayload)
		trxInfo.Payload = &aerium.TransactionInfo_Unbond{
			Unbond: &aerium.PayloadUnbond{
				Validator: pld.Validator.String(),
			},
		}
	case payload.TypeWithdraw:
		pld := trx.Payload().(*payload.WithdrawPayload)
		trxInfo.Payload = &aerium.TransactionInfo_Withdraw{
			Withdraw: &aerium.PayloadWithdraw{
				ValidatorAddress: pld.From.String(),
				AccountAddress:   pld.To.String(),
				Amount:           pld.Amount.ToNanoAUM(),
			},
		}

	case payload.TypeBatchTransfer:
		pld := trx.Payload().(*payload.BatchTransferPayload)
		recipients := make([]*aerium.Recipient, 0, len(pld.Recipients))
		for _, recipient := range pld.Recipients {
			recipients = append(recipients, &aerium.Recipient{
				Receiver: recipient.To.String(),
				Amount:   recipient.Amount.ToNanoAUM(),
			})
		}

		trxInfo.Payload = &aerium.TransactionInfo_BatchTransfer{
			BatchTransfer: &aerium.PayloadBatchTransfer{
				Sender:     pld.From.String(),
				Recipients: recipients,
			},
		}
	default:
		logger.Error("payload type not defined", "type", trx.Payload().Type())
	}

	return trxInfo
}

func (*transactionServer) DecodeRawTransaction(_ context.Context,
	req *aerium.DecodeRawTransactionRequest,
) (*aerium.DecodeRawTransactionResponse, error) {
	b, err := hex.DecodeString(req.RawTransaction)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid raw transaction")
	}

	trx, err := tx.FromBytes(b)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "couldn't decode transaction: %v", err.Error())
	}

	return &aerium.DecodeRawTransactionResponse{
		Transaction: transactionToProto(trx),
	}, nil
}
