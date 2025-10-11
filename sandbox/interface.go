package sandbox

import (
	"github.com/aerium-network/aerium/committee"
	"github.com/aerium-network/aerium/crypto"
	"github.com/aerium-network/aerium/crypto/bls"
	"github.com/aerium-network/aerium/sortition"
	"github.com/aerium-network/aerium/state/param"
	"github.com/aerium-network/aerium/types/account"
	"github.com/aerium-network/aerium/types/amount"
	"github.com/aerium-network/aerium/types/tx"
	"github.com/aerium-network/aerium/types/validator"
)

type Sandbox interface {
	Account(crypto.Address) *account.Account
	MakeNewAccount(crypto.Address) *account.Account
	UpdateAccount(crypto.Address, *account.Account)

	CommitTransaction(trx *tx.Tx)
	RecentTransaction(txID tx.ID) bool
	IsBanned(crypto.Address) bool

	Validator(crypto.Address) *validator.Validator
	MakeNewValidator(*bls.PublicKey) *validator.Validator
	UpdateValidator(*validator.Validator)
	JoinedToCommittee(crypto.Address)
	IsJoinedCommittee(crypto.Address) bool
	UpdatePowerDelta(delta int64)
	PowerDelta() int64
	AccumulatedFee() amount.Amount

	VerifyProof(uint32, sortition.Proof, *validator.Validator) bool
	Committee() committee.Reader

	Params() *param.Params
	CurrentHeight() uint32
	IsMainnet() bool

	IterateAccounts(consumer func(crypto.Address, *account.Account, bool))
	IterateValidators(consumer func(*validator.Validator, bool, bool))
}
