package campaign

import "github.com/ndtdat/social-network-monorepo/user-service/internal/repository/campaign"

type Service struct {
	repo *campaign.Repository
}

func NewService(repo *campaign.Repository) *Service {
	return &Service{repo: repo}
}
