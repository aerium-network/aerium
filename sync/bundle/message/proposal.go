package message

import (
	"github.com/aerium-network/aerium/network"
	"github.com/aerium-network/aerium/types/proposal"
	"github.com/aerium-network/aerium/types/protocol"
)

type ProposalMessage struct {
	Proposal        *proposal.Proposal `cbor:"1,keyasint"`
	ProtocolVersion protocol.Version   `cbor:"2,keyasint"`
}

func NewProposalMessage(p *proposal.Proposal) *ProposalMessage {
	return &ProposalMessage{
		Proposal:        p,
		ProtocolVersion: protocol.ProtocolVersionLatest,
	}
}

func (*ProposalMessage) BasicCheck() error {
	// Basic checks for the proposal are deferred to the consensus phase
	// to avoid unnecessary validation for validators outside the committee.
	return nil
}

func (*ProposalMessage) Type() Type {
	return TypeProposal
}

func (*ProposalMessage) TopicID() network.TopicID {
	return network.TopicIDConsensus
}

func (*ProposalMessage) ShouldBroadcast() bool {
	return true
}

func (m *ProposalMessage) ConsensusHeight() uint32 {
	return m.Height()
}

func (m *ProposalMessage) Height() uint32 {
	return m.Proposal.Height()
}

func (m *ProposalMessage) String() string {
	return m.Proposal.String()
}
