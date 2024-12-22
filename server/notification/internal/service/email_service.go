package service

import (
	"context"
	"github.com/oqamase/ozon/notification/pkg/notification"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Service) SendEmailMessage(ctx context.Context, request *notification.SendEmailMessageRequest) (*emptypb.Empty, error) {
	log.Print(request.Content)
	log.Print(request.Receiver)
	return &emptypb.Empty{}, nil
}
