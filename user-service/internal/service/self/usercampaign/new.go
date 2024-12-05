package usercampaign

import "github.com/ndtdat/social-network-monorepo/user-service/internal/repository/usercampaign"

type Service struct {
	repo *usercampaign.Repository
}

func NewService(repo *usercampaign.Repository) *Service {
	return &Service{repo: repo}
}
