package storage

import (
	"errors"
	"exoplanet-service/models"
	"fmt"
	"github.com/sirupsen/logrus"
	"sync"
)

type InMemoryStore struct {
	mu         sync.RWMutex
	exoplanets map[string]models.Exoplanet
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		exoplanets: make(map[string]models.Exoplanet),
	}
}

var Store = NewInMemoryStore()

func (s *InMemoryStore) AddExoplanet(exoplanet models.Exoplanet) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.exoplanets[exoplanet.ID] = exoplanet
}

func (s *InMemoryStore) ListExoplanets() []models.Exoplanet {
	s.mu.RLock()
	defer s.mu.RUnlock()
	exoplanets := make([]models.Exoplanet, 0, len(s.exoplanets))
	for _, exoplanet := range s.exoplanets {
		exoplanets = append(exoplanets, exoplanet)
	}
	return exoplanets
}

func (s *InMemoryStore) GetExoplanetByID(id string) (*models.Exoplanet, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	exoplanet, exists := s.exoplanets[id]
	if !exists {
		logrus.Errorf("exoplanet with %s not found", id)
		return nil, errors.New(fmt.Sprintf("exoplanet with %s not found", id))
	}
	return &exoplanet, nil
}

func (s *InMemoryStore) UpdateExoplanet(id string, exoplanet models.Exoplanet) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.exoplanets[id]
	if !exists {
		logrus.Errorf("exoplanet with %s not found", id)
		return errors.New(fmt.Sprintf("exoplanet with %s not found", id))
	}
	s.exoplanets[id] = exoplanet
	return nil
}

func (s *InMemoryStore) DeleteExoplanet(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.exoplanets[id]
	if !exists {
		logrus.Errorf("exoplanet with %s not found", id)
		return errors.New(fmt.Sprintf("exoplanet with %s not found", id))
	}
	delete(s.exoplanets, id)
	return nil
}
