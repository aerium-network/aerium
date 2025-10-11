package message

import (
	"testing"

	"github.com/aerium-network/aerium/types/tx"
	"github.com/aerium-network/aerium/util/testsuite"
	"github.com/stretchr/testify/assert"
)

func TestTransactionsType(t *testing.T) {
	msg := &TransactionsMessage{}
	assert.Equal(t, TypeTransaction, msg.Type())
}

func TestTransactionsMessage(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	t.Run("No transactions", func(t *testing.T) {
		msg := NewTransactionsMessage(nil)

		err := msg.BasicCheck()
		assert.ErrorIs(t, err, BasicCheckError{Reason: "no transaction"})
	})

	t.Run("OK", func(t *testing.T) {
		trx := ts.GenerateTestTransferTx()
		msg := NewTransactionsMessage([]*tx.Tx{trx})

		assert.NoError(t, msg.BasicCheck())
		assert.Contains(t, msg.String(), trx.ID().ShortString())
	})
}
