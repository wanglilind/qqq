package p2p

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type Node struct {
	host       host.Host
	pubsub     *pubsub.PubSub
	topics     map[string]*pubsub.Topic
	subs       map[string]*pubsub.Subscription
	peers      map[peer.ID]bool
	mu         sync.RWMutex
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewNode(listenAddr string) (*Node, error) {
	ctx, cancel := context.WithCancel(context.Background())
	
	// åå»ºlibp2pèç¹
	h, err := libp2p.New(
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/%s/tcp/0", listenAddr)),
		libp2p.EnableAutoRelay(),
		libp2p.EnableNATService(),
	)
	if err != nil {
		cancel()
		return nil, err
	}

	// åå»ºpubsubç³»ç»
	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		cancel()
		return nil, err
	}

	return &Node{
		host:   h,
		pubsub: ps,
		topics: make(map[string]*pubsub.Topic),
		subs:   make(map[string]*pubsub.Subscription),
		peers:  make(map[peer.ID]bool),
		ctx:    ctx,
		cancel: cancel,
	}, nil
}

func (n *Node) Subscribe(topic string, handler func([]byte)) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	// å å¥topic
	t, err := n.pubsub.Join(topic)
	if err != nil {
		return err
	}
	n.topics[topic] = t

	// è®¢éæ¶æ¯
	sub, err := t.Subscribe()
	if err != nil {
		return err
	}
	n.subs[topic] = sub

	// å¤çæ¶æ¯
	go func() {
		for {
			msg, err := sub.Next(n.ctx)
			if err != nil {
				return
			}
			if msg.ReceivedFrom == n.host.ID() {
				continue
			}
			handler(msg.Data)
		}
	}()

	return nil
}

func (n *Node) Publish(topic string, data []byte) error {
	n.mu.RLock()
	t, exists := n.topics[topic]
	n.mu.RUnlock()

	if !exists {
		return fmt.Errorf("topic %s not found", topic)
	}

	return t.Publish(n.ctx, data)
}

func (n *Node) Connect(addr string) error {
	// è¿æ¥å°å¯¹ç­èç?
	peerInfo, err := peer.AddrInfoFromString(addr)
	if err != nil {
		return err
	}

	if err := n.host.Connect(n.ctx, *peerInfo); err != nil {
		return err
	}

	n.mu.Lock()
	n.peers[peerInfo.ID] = true
	n.mu.Unlock()

	return nil
}

func (n *Node) Close() error {
	n.cancel()
	return n.host.Close()
} 
