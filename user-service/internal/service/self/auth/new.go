package auth

import (
	"github.com/ndtdat/social-network-monorepo/user-service/internal/repository/user"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/service/self/campaign"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	logger          *zap.Logger
	db              *gorm.DB
	userRepo        *user.Repository
	campaignService *campaign.Service
}

func NewService(
	logger *zap.Logger, db *gorm.DB, userRepo *user.Repository, campaignService *campaign.Service,
) *Service {
	return &Service{
		logger:          logger,
		db:              db,
		userRepo:        userRepo,
		campaignService: campaignService,
	}
}
