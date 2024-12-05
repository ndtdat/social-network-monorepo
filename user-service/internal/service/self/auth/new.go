package auth

import (
	jwtbase "github.com/ndtdat/social-network-monorepo/gokit/pkg/jwt/base"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/repository/user"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/service/self/campaign"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/service/self/usercampaign"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	logger              *zap.Logger
	jm                  jwtbase.Manager
	db                  *gorm.DB
	repo                *user.Repository
	campaignService     *campaign.Service
	userCampaignService *usercampaign.Service
}

func NewService(
	logger *zap.Logger, jm jwtbase.Manager, db *gorm.DB, repo *user.Repository, campaignService *campaign.Service,
	userCampaignService *usercampaign.Service,
) *Service {
	return &Service{
		logger:              logger,
		jm:                  jm,
		db:                  db,
		repo:                repo,
		campaignService:     campaignService,
		userCampaignService: userCampaignService,
	}
}
