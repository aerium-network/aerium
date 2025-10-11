package sync

import (
	"time"

	"github.com/aerium-network/aerium/sync/peerset"
	"github.com/aerium-network/aerium/sync/peerset/peer"
	"github.com/aerium-network/aerium/sync/peerset/peer/service"
)

type Synchronizer interface {
	Start() error
	Stop()
	Moniker() string
	SelfID() peer.ID
	PeerSet() *peerset.PeerSet
	Services() service.Services
	ClockOffset() (time.Duration, error)
	IsClockOutOfSync() bool
}
