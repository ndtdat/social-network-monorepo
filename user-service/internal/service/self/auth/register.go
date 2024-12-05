package auth

import (
	"context"
	"fmt"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/suid"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/util"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/model"
	"github.com/ndtdat/social-network-monorepo/user-service/pkg/api/go/user/rpc"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *Service) Register(ctx context.Context, in *rpc.RegisterRequest) (*rpc.RegisterReply, error) {
	var (
		logger       = s.logger
		email        = in.GetEmail()
		password     = in.GetPassword()
		campaignCode = in.GetCampaignCode()
		now          = uint64(util.CurrentUnix())
	)
	// TODO: Check email is exist
	userInfo, err := s.userRepo.FirstByFilters(ctx, email)
	if err != nil {
		return nil, err
	}
	if userInfo != nil {
		return nil, fmt.Errorf("email [%v] is taken", email)
	}

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, fmt.Errorf("cannot create password hash due to %v", err)
	}

	userInfo = &model.User{
		ID:       suid.New(),
		Email:    email,
		Password: string(passwordHash),
	}

	if err = s.db.Transaction(func(tx *gorm.DB) error {
		tx = tx.WithContext(ctx)

		campaignInfo, err := s.campaignService.SelectValidCampaignForUpdate(tx, campaignCode, now)
		if err != nil {
			return err
		}
		if campaignInfo != nil {
			campaignInfo.JoinedQty++
		}

		return nil
	}); err != nil {
		return nil, err
	}

	// TODO: Get CODE param and check it's existed

	// TODO: If available => Call purchase-service to allocate voucher => Can ignore if failed

	// TODO: Use transaction to update data

	return &rpc.RegisterReply{}, nil
}
