package voteset

import (
	"testing"

	"github.com/aerium-network/aerium/types/vote"
	"github.com/aerium-network/aerium/util/testsuite"
	"github.com/stretchr/testify/assert"
)

func TestDoubleVote(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	hash := ts.RandHash()
	height := ts.RandHeight()
	round := ts.RandRound()
	signer := ts.RandValAddress()
	power := ts.RandInt64(1000)

	v := vote.NewPrecommitVote(hash, height, round, signer)

	vb := newVoteBox()

	vb.addVote(v, power)
	vb.addVote(v, power)

	assert.Equal(t, power, vb.votedPower)
}
