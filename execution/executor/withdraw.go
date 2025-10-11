package executor

import (
	"github.com/aerium-network/aerium/sandbox"
	"github.com/aerium-network/aerium/types/account"
	"github.com/aerium-network/aerium/types/amount"
	"github.com/aerium-network/aerium/types/tx"
	"github.com/aerium-network/aerium/types/tx/payload"
	"github.com/aerium-network/aerium/types/validator"
)

type WithdrawExecutor struct {
	sbx      sandbox.Sandbox
	pld      *payload.WithdrawPayload
	fee      amount.Amount
	sender   *validator.Validator
	receiver *account.Account
}

func newWithdrawExecutor(trx *tx.Tx, sbx sandbox.Sandbox) (*WithdrawExecutor, error) {
	pld := trx.Payload().(*payload.WithdrawPayload)

	sender := sbx.Validator(pld.From)
	if sender == nil {
		return nil, ValidatorNotFoundError{Address: pld.From}
	}

	receiver := sbx.Account(pld.To)
	if receiver == nil {
		receiver = sbx.MakeNewAccount(pld.To)
	}

	return &WithdrawExecutor{
		sbx:      sbx,
		pld:      pld,
		fee:      trx.Fee(),
		sender:   sender,
		receiver: receiver,
	}, nil
}

func (e *WithdrawExecutor) Check(_ bool) error {
	if e.sender.Stake() < e.pld.Value()+e.fee {
		return ErrInsufficientFunds
	}

	if !e.sender.IsUnbonded() {
		return ErrValidatorBonded
	}

	if e.sbx.CurrentHeight() < e.sender.UnbondingHeight()+e.sbx.Params().UnbondInterval {
		return ErrUnbondingPeriod
	}

	return nil
}

func (e *WithdrawExecutor) Execute() {
	e.sender.SubtractFromStake(e.pld.Amount + e.fee)
	e.receiver.AddToBalance(e.pld.Amount)

	e.sbx.UpdateValidator(e.sender)
	e.sbx.UpdateAccount(e.pld.To, e.receiver)
}
