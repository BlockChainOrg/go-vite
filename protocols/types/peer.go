package types

import (
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/p2p"
	"log"
	"math/big"
	"sync"
)

// @section Peer for protocol handle, not p2p Peer.
type Peer struct {
	*p2p.Peer
	ID      string
	Head    types.Hash
	Height  *big.Int
	Version int
	RW      MsgReadWriter
	Lock    sync.RWMutex
	// use this channel to ensure that only one goroutine send msg simultaneously.
	Sending chan struct{}
}

func (p *Peer) Update(status *StatusMsg) {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	p.Height = status.Height
	p.Head = status.CurrentBlock
	log.Printf("peer %s update status: height %d Head %s\n", p.ID, p.Height, p.Head)
}
