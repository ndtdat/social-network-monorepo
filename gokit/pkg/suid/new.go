package suid

import (
	"fmt"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/random"
	"time"
)

type Service struct {
	generators  []*SUID
	numInstance int
}

func NewService(numInstance int) *Service {
	var generators []*SUID
	machineID, err := lower16BitPrivateIP(defaultInterfaceAddrs)
	if err != nil {
		panic(fmt.Sprintf("cannot get private ip to init suid due to: %v", err))
	}

	for i := 0; i < numInstance; i++ {
		gen, err := NewSUID(time.Now(), byte(machineID))
		if err != nil {
			panic(fmt.Errorf("cannot init suid due to %v", err))
		}
		generators = append(generators, gen)

		time.Sleep(time.Duration(random.Int64(minInitSleep, maxInitSleep)) * time.Microsecond)
	}

	return &Service{
		generators:  generators,
		numInstance: numInstance,
	}
}

func (s *Service) NextID() (uint64, error) {
	id, err := s.generators[random.Int64(0, int64(s.numInstance-1))].NextID()

	return id, err
}
