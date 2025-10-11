package txpool

import (
	"github.com/aerium-network/aerium/sandbox"
	"github.com/aerium-network/aerium/types/amount"
	"github.com/aerium-network/aerium/types/block"
	"github.com/aerium-network/aerium/types/tx"
	"github.com/aerium-network/aerium/types/tx/payload"
)

type Reader interface {
	PrepareBlockTransactions() block.Txs
	PendingTx(txID tx.ID) *tx.Tx
	HasTx(txID tx.ID) bool
	Size() int
	EstimatedFee(amt amount.Amount, payloadType payload.Type) amount.Amount
	AllPendingTxs() []*tx.Tx
}

type TxPool interface {
	Reader

	SetNewSandboxAndRecheck(sbx sandbox.Sandbox)
	AppendTxAndBroadcast(trx *tx.Tx) error
	AppendTx(trx *tx.Tx) error
	HandleCommittedBlock(blk *block.Block)
}
