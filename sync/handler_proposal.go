package sync

import (
	"github.com/aerium-network/aerium/sync/bundle"
	"github.com/aerium-network/aerium/sync/bundle/message"
	"github.com/aerium-network/aerium/sync/peerset/peer"
)

type proposalHandler struct {
	*synchronizer
}

func newProposalHandler(sync *synchronizer) messageHandler {
	return &proposalHandler{
		sync,
	}
}

func (handler *proposalHandler) ParseMessage(m message.Message, _ peer.ID) {
	msg := m.(*message.ProposalMessage)
	handler.logger.Trace("parsing Proposal message", "msg", msg)

	handler.consMgr.SetProposal(msg.Proposal)

	handler.state.UpdateValidatorProtocolVersion(
		msg.Proposal.Block().Header().ProposerAddress(),
		msg.ProtocolVersion,
	)
}

func (*proposalHandler) PrepareBundle(m message.Message) *bundle.Bundle {
	return bundle.NewBundle(m)
}
