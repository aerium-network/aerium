package consensus

import (
	"github.com/aerium-network/aerium/types/proposal"
	"github.com/aerium-network/aerium/types/vote"
)

type consState interface {
	enter()
	decide()
	onAddVote(v *vote.Vote)
	onSetProposal(p *proposal.Proposal)
	onTimeout(t *ticker)
	name() string
}
