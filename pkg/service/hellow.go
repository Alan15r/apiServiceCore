package service

import (
	"github.com/Alan15r/apiServiceCore/pkg/repository"
)

type HellowService struct {
	repo *repository.Repository
}

func NewHellowService(repo *repository.Repository) *HellowService {
	return &HellowService{repo: repo}
}

func (hs *HellowService) ServiceHellow(name string) error {
	return hs.repo.Hellow.RepositoryHellow(name)
}
