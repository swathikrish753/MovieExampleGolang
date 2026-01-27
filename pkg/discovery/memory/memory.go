package memory

import (
	"sync"
	"time"
)

type serviceName string
type instanceID string

type Registry struct {
	sync.RWMutex
	serviceAddrs map[serviceName]map[instanceID]*serviceInstance
}
type serviceInstance struct {
	hostPort   string
	lastActive time.Time
}

func NewRegistry() *Registry {
	return &Registry{
		serviceAddrs: map[serviceName]map[instanceID]*serviceInstance{},
	}
}
