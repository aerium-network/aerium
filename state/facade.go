package state

import (
	"time"

	"github.com/aerium-network/aerium/crypto"
	"github.com/aerium-network/aerium/crypto/bls"
	"github.com/aerium-network/aerium/crypto/hash"
	"github.com/aerium-network/aerium/genesis"
	"github.com/aerium-network/aerium/state/param"
	"github.com/aerium-network/aerium/store"
	"github.com/aerium-network/aerium/types/account"
	"github.com/aerium-network/aerium/types/amount"
	"github.com/aerium-network/aerium/types/block"
	"github.com/aerium-network/aerium/types/certificate"
	"github.com/aerium-network/aerium/types/protocol"
	"github.com/aerium-network/aerium/types/tx"
	"github.com/aerium-network/aerium/types/tx/payload"
	"github.com/aerium-network/aerium/types/validator"
	"github.com/aerium-network/aerium/types/vote"
)

type Stats struct {
	LastBlockHeight uint32
	LastBlockHash   hash.Hash
	LastBlockTime   time.Time

	TotalPower     int64
	CommitteePower int64

	TotalAccounts    int32
	TotalValidators  int32
	ActiveValidators int32

	IsPruned      bool
	PruningHeight uint32
}

type Facade interface {
	Genesis() *genesis.Genesis
	Params() *param.Params
	LastBlockHeight() uint32
	LastBlockHash() hash.Hash
	LastBlockTime() time.Time
	LastCertificate() *certificate.Certificate
	UpdateLastCertificate(v *vote.Vote) error
	ProposeBlock(valKey *bls.ValidatorKey, rewardAddr crypto.Address) (*block.Block, error)
	ValidateBlock(blk *block.Block, round int16) error
	CommitBlock(blk *block.Block, cert *certificate.Certificate) error
	CommitteeValidators() []*validator.Validator
	CommitteeProtocolVersions() map[protocol.Version]float64
	IsInCommittee(addr crypto.Address) bool
	Proposer(round int16) *validator.Validator
	IsProposer(addr crypto.Address, round int16) bool
	PendingTx(txID tx.ID) *tx.Tx
	AddPendingTx(trx *tx.Tx) error
	AddPendingTxAndBroadcast(trx *tx.Tx) error
	CommittedBlock(height uint32) *store.CommittedBlock
	CommittedTx(txID tx.ID) *store.CommittedTx
	BlockHash(height uint32) hash.Hash
	BlockHeight(h hash.Hash) uint32
	AccountByAddress(addr crypto.Address) *account.Account
	ValidatorByAddress(addr crypto.Address) *validator.Validator
	ValidatorByNumber(number int32) *validator.Validator
	ValidatorAddresses() []crypto.Address
	UpdateValidatorProtocolVersion(addr crypto.Address, ver protocol.Version)
	CalculateFee(amt amount.Amount, payloadType payload.Type) amount.Amount
	PublicKey(addr crypto.Address) (crypto.PublicKey, error)
	AvailabilityScore(valNum int32) float64
	AllPendingTxs() []*tx.Tx
	Stats() *Stats

	Close()
}
