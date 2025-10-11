package txpool

import (
	"github.com/aerium-network/aerium/types/amount"
	"github.com/aerium-network/aerium/types/tx"
	"github.com/aerium-network/aerium/util/linkedmap"
)

type pool struct {
	list   *linkedmap.LinkedMap[tx.ID, *tx.Tx]
	minFee amount.Amount
}

func newPool(maxSize int, minFee amount.Amount) pool {
	return pool{
		list:   linkedmap.New[tx.ID, *tx.Tx](maxSize),
		minFee: minFee,
	}
}

func (p *pool) estimatedFee() amount.Amount {
	return p.minFee
}
