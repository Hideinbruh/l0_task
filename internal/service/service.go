package service

import (
	"awesomeProject/internal/model"
	"awesomeProject/internal/repository"
	"encoding/json"
	"errors"
)

type Service struct {
	repo  *repository.Repository
	cache *repository.OrderCache
}

func NewService(repo *repository.Repository, cache *repository.OrderCache) *Service {
	return &Service{repo: repo, cache: cache}
}

func (s *Service) Save(msg []byte) error {
	var data *model.Order
	if !json.Valid(msg) {
		return errors.New("Invalid json")
	}
	err := json.Unmarshal(msg, data)
	if err != nil {
		return err
	}
	if err = s.repo.Save(data); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetOrder(orderUid string) (*model.Order, error) {
	modelOrder, err := s.repo.GetDataById(orderUid)
	if err != nil {
		return nil, err
	}
	return modelOrder, nil
}

func (s *Service) CreateOrderCache(order *model.Order) error {
	err := s.repo.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetModelCache(orderUid string) (*model.Order, error) {
	return s.cache.GetOrderCache(orderUid)
}
