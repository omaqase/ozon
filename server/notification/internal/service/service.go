package service

import (
	"errors"
	"github.com/oqamase/ozon/notification/pkg/notification"
)

var (
	ErrFailedInitializeService = errors.New("failed to initialize service")
)

type Configuration func(s *Service) error

type Service struct {
	notification.UnimplementedNotificationServiceServer
}

func NewService(configs ...Configuration) (*Service, error) {
	service := &Service{}

	for _, config := range configs {
		if err := config(service); err != nil {
			return nil, ErrFailedInitializeService
		}
	}

	return service, nil
}
