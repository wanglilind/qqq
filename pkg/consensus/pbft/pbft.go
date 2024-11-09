package pbft

import (
	"context"
	"sync"
	"time"

	"github.com/wanglilind/qqq/pkg/blockchain"
)

type State int

const (
	PrePrepare State = iota
	Prepare
	Commit
	Finalize
)

type ConsensusMessage struct {
	Type      State
	BlockHash string
	NodeID    string
	Signature []byte
}

type PBFT struct {
	mu              sync.RWMutex
	nodeID          string
	validators      map[string]bool
	currentState    State
	prepareVotes    map[string]map[string]bool  // blockHash -> nodeID -> vote
	commitVotes     map[string]map[string]bool
	pendingBlocks   map[string]*blockchain.Block
	config          *Config
	msgChan         chan ConsensusMessage
}

func NewPBFT(nodeID string, config *Config) *PBFT {
	return &PBFT{
		nodeID:        nodeID,
		validators:    make(map[string]bool),
		prepareVotes:  make(map[string]map[string]bool),
		commitVotes:   make(map[string]map[string]bool),
		pendingBlocks: make(map[string]*blockchain.Block),
		config:        config,
		msgChan:      make(chan ConsensusMessage, 1000),
	}
}

func (p *PBFT) Start(ctx context.Context) error {
	go p.processMessages(ctx)
	return nil
}

func (p *PBFT) ProposeBlock(block *blockchain.Block) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	blockHash := block.CalculateHash()
	p.pendingBlocks[blockHash] = block
	
	// å¹¿æ­PrePrepareæ¶æ¯
	msg := ConsensusMessage{
		Type:      PrePrepare,
		BlockHash: blockHash,
		NodeID:    p.nodeID,
	}
	
	return p.broadcast(msg)
}

func (p *PBFT) processMessages(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-p.msgChan:
			p.handleMessage(msg)
		}
	}
}

func (p *PBFT) handleMessage(msg ConsensusMessage) {
	p.mu.Lock()
	defer p.mu.Unlock()

	switch msg.Type {
	case PrePrepare:
		p.handlePrePrepare(msg)
	case Prepare:
		p.handlePrepare(msg)
	case Commit:
		p.handleCommit(msg)
	}
}

func (p *PBFT) handlePrePrepare(msg ConsensusMessage) {
	// éªè¯æ¶æ¯
	if !p.validateMessage(msg) {
		return
	}

	// åå§åæç¥¨è®°å½?
	if _, exists := p.prepareVotes[msg.BlockHash]; !exists {
		p.prepareVotes[msg.BlockHash] = make(map[string]bool)
	}

	// åéPrepareæ¶æ¯
	prepareMsg := ConsensusMessage{
		Type:      Prepare,
		BlockHash: msg.BlockHash,
		NodeID:    p.nodeID,
	}
	p.broadcast(prepareMsg)
}

func (p *PBFT) handlePrepare(msg ConsensusMessage) {
	// è®°å½Prepareæç¥¨
	p.prepareVotes[msg.BlockHash][msg.NodeID] = true

	// æ£æ¥æ¯å¦è¾¾å°åå¤éå?
	if p.checkPrepareQuorum(msg.BlockHash) {
		commitMsg := ConsensusMessage{
			Type:      Commit,
			BlockHash: msg.BlockHash,
			NodeID:    p.nodeID,
		}
		p.broadcast(commitMsg)
	}
}

func (p *PBFT) handleCommit(msg ConsensusMessage) {
	// è®°å½Commitæç¥¨
	if _, exists := p.commitVotes[msg.BlockHash]; !exists {
		p.commitVotes[msg.BlockHash] = make(map[string]bool)
	}
	p.commitVotes[msg.BlockHash][msg.NodeID] = true

	// æ£æ¥æ¯å¦è¾¾å°æäº¤éå?
	if p.checkCommitQuorum(msg.BlockHash) {
		p.finalizeBlock(msg.BlockHash)
	}
}

func (p *PBFT) checkPrepareQuorum(blockHash string) bool {
	return len(p.prepareVotes[blockHash]) >= p.getQuorumSize()
}

func (p *PBFT) checkCommitQuorum(blockHash string) bool {
	return len(p.commitVotes[blockHash]) >= p.getQuorumSize()
}

func (p *PBFT) getQuorumSize() int {
	return (len(p.validators) * 2 / 3) + 1
}

func (p *PBFT) validateMessage(msg ConsensusMessage) bool {
	// éªè¯æ¶æ¯ç­¾åååéèèº«ä»?
	return true
}

func (p *PBFT) broadcast(msg ConsensusMessage) error {
	// å¹¿æ­æ¶æ¯å°ææéªè¯èç?
	return nil
}

func (p *PBFT) finalizeBlock(blockHash string) {
	// æç»ç¡®è®¤åºå?
	block := p.pendingBlocks[blockHash]
	// æ§è¡åºåå¹¶æ´æ°ç¶æ?
	delete(p.pendingBlocks, blockHash)
	delete(p.prepareVotes, blockHash)
	delete(p.commitVotes, blockHash)
} 
