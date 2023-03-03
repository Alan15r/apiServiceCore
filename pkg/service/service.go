package service

import (
	"github.com/Alan15r/apiServiceCore/pkg/repository"
)

type Service struct {
	Hellow
}

type Hellow interface {
	ServiceHellow(name string) error
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Hellow: NewHellowService(repos),
	}
}
