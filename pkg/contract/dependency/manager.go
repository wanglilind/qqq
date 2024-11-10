package dependency

import (
	"context"
	"sync"
	"time"

	"github.com/wanglilind/qqq/pkg/contract"
)

// ä¾èµç®¡çå?
type DependencyManager struct {
	dependencies map[string][]Dependency
	versions     map[string][]Version
	mu           sync.RWMutex
	config       *Config
}

type Dependency struct {
	ContractAddr string
	MinVersion   string
	MaxVersion   string
	Required     bool
	Permissions  []string
}

type Version struct {
	Version     string
	Timestamp   time.Time
	Hash        string
	Dependencies []Dependency
	Status      string
}

func NewDependencyManager(config *Config) *DependencyManager {
	return &DependencyManager{
		dependencies: make(map[string][]Dependency),
		versions:     make(map[string][]Version),
		config:      config,
	}
}

// æ·»å ä¾èµ
func (dm *DependencyManager) AddDependency(contractAddr string, dep Dependency) error {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	// éªè¯ä¾èµ
	if err := dm.validateDependency(dep); err != nil {
		return err
	}

	// æ£æ¥å¾ªç¯ä¾èµ?
	if dm.hasCircularDependency(contractAddr, dep.ContractAddr) {
		return ErrCircularDependency
	}

	// å­å¨ä¾èµ
	if _, exists := dm.dependencies[contractAddr]; !exists {
		dm.dependencies[contractAddr] = make([]Dependency, 0)
	}
	dm.dependencies[contractAddr] = append(dm.dependencies[contractAddr], dep)

	return nil
}

// æ£æ¥ä¾èµ?
func (dm *DependencyManager) CheckDependencies(ctx context.Context, contractAddr string) error {
	dm.mu.RLock()
	defer dm.mu.RUnlock()

	deps := dm.dependencies[contractAddr]
	for _, dep := range deps {
		// æ£æ¥çæ¬å¼å®¹æ?
		if err := dm.checkVersionCompatibility(dep); err != nil {
			return err
		}

		// æ£æ¥æé?
		if err := dm.checkPermissions(dep); err != nil {
			return err
		}
	}

	return nil
}

// è·åä¾èµå?
func (dm *DependencyManager) GetDependencyGraph(contractAddr string) (map[string][]string, error) {
	dm.mu.RLock()
	defer dm.mu.RUnlock()

	graph := make(map[string][]string)
	visited := make(map[string]bool)

	if err := dm.buildDependencyGraph(contractAddr, graph, visited); err != nil {
		return nil, err
	}

	return graph, nil
}

// éªè¯ä¾èµ
func (dm *DependencyManager) validateDependency(dep Dependency) error {
	// å®ç°ä¾èµéªè¯é»è¾
	return nil
}

// æ£æ¥å¾ªç¯ä¾èµ?
func (dm *DependencyManager) hasCircularDependency(from, to string) bool {
	visited := make(map[string]bool)
	return dm.detectCircularDependency(from, to, visited)
}

// æ£æµå¾ªç¯ä¾èµ?
func (dm *DependencyManager) detectCircularDependency(current, target string, visited map[string]bool) bool {
	if current == target {
		return true
	}

	if visited[current] {
		return false
	}

	visited[current] = true
	for _, dep := range dm.dependencies[current] {
		if dm.detectCircularDependency(dep.ContractAddr, target, visited) {
			return true
		}
	}

	return false
}

// æ£æ¥çæ¬å¼å®¹æ?
func (dm *DependencyManager) checkVersionCompatibility(dep Dependency) error {
	// å®ç°çæ¬å¼å®¹æ§æ£æ¥é»è¾
	return nil
}

// æ£æ¥æé?
func (dm *DependencyManager) checkPermissions(dep Dependency) error {
	// å®ç°æéæ£æ¥é»è¾
	return nil
}

// æå»ºä¾èµå?
func (dm *DependencyManager) buildDependencyGraph(contractAddr string, graph map[string][]string, visited map[string]bool) error {
	// å®ç°ä¾èµå¾æå»ºé»è¾
	return nil
} 
