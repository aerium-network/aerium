package manager

import (
	"github.com/aerium-network/aerium/crypto/bls"
	"github.com/aerium-network/aerium/crypto/hash"
	"github.com/aerium-network/aerium/types/proposal"
	"github.com/aerium-network/aerium/types/vote"
)

type Reader interface {
	ConsensusKey() *bls.PublicKey
	AllVotes() []*vote.Vote
	HandleQueryVote(height uint32, round int16) *vote.Vote
	HandleQueryProposal(height uint32, round int16) *proposal.Proposal
	Proposal() *proposal.Proposal
	HasVote(h hash.Hash) bool
	HeightRound() (uint32, int16)
	IsActive() bool
	IsProposer() bool
}

type Consensus interface {
	Reader

	MoveToNewHeight()
	AddVote(vote *vote.Vote)
	SetProposal(prop *proposal.Proposal)
	IsDeprecated() bool
}

type ManagerReader interface {
	Instances() []Reader
	HandleQueryVote(height uint32, round int16) *vote.Vote
	HandleQueryProposal(height uint32, round int16) *proposal.Proposal
	Proposal() *proposal.Proposal
	HeightRound() (uint32, int16)
	HasActiveInstance() bool
}

type Manager interface {
	ManagerReader

	MoveToNewHeight()
	AddVote(vote *vote.Vote)
	SetProposal(prop *proposal.Proposal)
	IsDeprecated() bool
}
