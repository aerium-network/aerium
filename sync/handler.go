package sync

import (
	"github.com/aerium-network/aerium/sync/bundle"
	"github.com/aerium-network/aerium/sync/bundle/message"
	"github.com/aerium-network/aerium/sync/peerset/peer"
)

type messageHandler interface {
	ParseMessage(message.Message, peer.ID)
	PrepareBundle(message.Message) *bundle.Bundle
}
