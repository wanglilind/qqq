package discovery

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

type ServiceDiscovery struct {
	client *api.Client
	config *Config
}

func NewServiceDiscovery(config *Config) (*ServiceDiscovery, error) {
	// åå»ºConsulå®¢æ·ç«?
	consulConfig := api.DefaultConfig()
	consulConfig.Address = config.ConsulAddress

	client, err := api.NewClient(consulConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create consul client: %v", err)
	}

	return &ServiceDiscovery{
		client: client,
		config: config,
	}, nil
}

func (sd *ServiceDiscovery) RegisterService(serviceName string, port int) error {
	// æ³¨åæå¡
	registration := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%d", serviceName, port),
		Name:    serviceName,
		Port:    port,
		Tags:    []string{"gfc", "v1"},
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://localhost:%d/health", port),
			Interval: "10s",
			Timeout:  "5s",
		},
	}

	if err := sd.client.Agent().ServiceRegister(registration); err != nil {
		return fmt.Errorf("failed to register service: %v", err)
	}

	return nil
} 
