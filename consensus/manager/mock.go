package manager

import (
	"github.com/aerium-network/aerium/consensus"
	"github.com/aerium-network/aerium/crypto/bls"
	"github.com/aerium-network/aerium/state"
	"github.com/aerium-network/aerium/types/proposal"
	"github.com/aerium-network/aerium/types/vote"
	"github.com/aerium-network/aerium/util/testsuite"
)

func MockingManager(ts *testsuite.TestSuite, state *state.MockState,
	valKeys []*bls.ValidatorKey,
) (Manager, []*consensus.MockConsensus) {
	mocks := make([]*consensus.MockConsensus, len(valKeys))
	instances := make([]Consensus, len(valKeys))
	for i, key := range valKeys {
		cons := consensus.MockingConsensus(ts, state, key)
		mocks[i] = cons
		instances[i] = cons
	}

	return &manager{
		instances:         instances,
		upcomingVotes:     make([]*vote.Vote, 0),
		upcomingProposals: make([]*proposal.Proposal, 0),
	}, mocks
}
