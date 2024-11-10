package service

import (
	"context"
	"sync"

	"github.com/wanglilind/qqq/pkg/config"
	"github.com/wanglilind/qqq/pkg/consensus/pbft"
)

type ConsensusService struct {
	nodeManager    *node.Manager
	consensusAlgo  *algorithm.Consensus
	networkManager *network.Manager
	config         *config.Config
	mu            sync.RWMutex
	isRunning     bool
}

func NewConsensusService(cfg *config.Config) *ConsensusService {
	return &ConsensusService{
		nodeManager:    node.NewManager(cfg),
		consensusAlgo:  algorithm.NewConsensus(cfg),
		networkManager: network.NewManager(cfg),
		config:        cfg,
	}
}

// StartP2PNetwork å¯å¨P2Pç½ç»
func (s *ConsensusService) StartP2PNetwork() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isRunning {
		return nil
	}

	// åå§åç½ç»?
	if err := s.networkManager.Initialize(); err != nil {
		return err
	}

	// å¯å¨èç¹åç°
	if err := s.nodeManager.StartDiscovery(); err != nil {
		return err
	}

	// å¯å¨å±è¯ç®æ³
	if err := s.consensusAlgo.Start(); err != nil {
		return err
	}

	s.isRunning = true
	return nil
}

// StopP2PNetwork åæ­¢P2Pç½ç»
func (s *ConsensusService) StopP2PNetwork() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.isRunning {
		return nil
	}

	// åæ­¢å±è¯ç®æ³
	s.consensusAlgo.Stop()

	// åæ­¢èç¹åç°
	s.nodeManager.StopDiscovery()

	// å³é­ç½ç»è¿æ¥
	s.networkManager.Shutdown()

	s.isRunning = false
	return nil
}

// ProposeBlock æè®®æ°åºå?
func (s *ConsensusService) ProposeBlock(ctx context.Context, req *ProposeBlockRequest) (*ProposeBlockResponse, error) {
	// éªè¯æè®®èèº«ä»?
	if err := s.validateProposer(req.ProposerId); err != nil {
		return nil, err
	}

	// åå»ºåºåæè®®
	block := &Block{
		Transactions: req.Transactions,
		Timestamp:   req.Timestamp,
		ProposerId:  req.ProposerId,
	}

	// è¿è¡å±è¯
	consensus, err := s.consensusAlgo.ProcessBlock(block)
	if err != nil {
		return nil, err
	}

	// å¹¿æ­åºå
	if err := s.networkManager.BroadcastBlock(block); err != nil {
		return nil, err
	}

	return &ProposeBlockResponse{
		BlockId:     block.Id,
		ConsensusId: consensus.Id,
		Status:      "proposed",
	}, nil
}

// GetConsensusStatus è·åå±è¯ç¶æ?
func (s *ConsensusService) GetConsensusStatus(ctx context.Context, req *GetConsensusStatusRequest) (*GetConsensusStatusResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// è·åå½åå±è¯ç¶æ?
	status := s.consensusAlgo.GetStatus()
	
	// è·åæ´»è·èç¹åè¡¨
	activeNodes := s.nodeManager.GetActiveNodes()

	return &GetConsensusStatusResponse{
		CurrentRound:    status.CurrentRound,
		ActiveNodes:     len(activeNodes),
		ConsensusState:  status.State,
		LastBlockId:     status.LastBlockId,
		LastUpdateTime:  status.LastUpdateTime,
	}, nil
}

// validateProposer éªè¯æè®®èèº«ä»?
func (s *ConsensusService) validateProposer(proposerId string) error {
	// æ£æ¥æè®®èæ¯å¦æ¯æ´»è·èç¹
	if !s.nodeManager.IsActiveNode(proposerId) {
		return ErrInvalidProposer
	}

	// æ£æ¥æè®®èæé?
	if !s.consensusAlgo.HasProposerRights(proposerId) {
		return ErrNoProposerRights
	}

	return nil
} 
