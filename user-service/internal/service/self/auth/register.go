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
		campaignService = s.campaignService
		email           = in.GetEmail()
		password        = in.GetPassword()
		campaignCode    = in.GetCampaignCode()
		now             = uint64(util.CurrentUnix())
	)
	// TODO: Check email is exist
	userInfo, err := s.repo.FirstByFilters(ctx, email)
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

	userID := suid.New()
	userInfo = &model.User{
		ID:       userID,
		Email:    email,
		Password: string(passwordHash),
	}

	accessToken, err := s.GenAccessToken(userID)
	if err != nil {
		return nil, err
	}

	if err = s.db.Transaction(func(tx *gorm.DB) error {
		tx = tx.WithContext(ctx)

		if campaignCode != "" {
			campaignInfo, err := campaignService.SelectValidCampaignForUpdate(tx, campaignCode, now)
			if err != nil {
				return err
			}
			if campaignInfo != nil {
				campaignInfo.JoinedQty++

				if err = campaignService.UpdateWithTx(tx, campaignInfo); err != nil {
					return err
				}

				if err = s.userCampaignService.CreateOrNothingWithTx(tx, &model.UserCampaign{
					UserID:     userID,
					CampaignID: campaignInfo.ID,
				}); err != nil {
					return err
				}
			} else {
				s.logger.Info(
					fmt.Sprintf("User email [%v] register with invaid comapign code [%v]", email, campaignCode),
				)
			}
		}

		if err = s.repo.CreateOrNothingWithTx(tx, userInfo); err != nil {
			return err
		}

		// TODO: Trigger to allocate voucher

		return nil
	}); err != nil {
		return nil, err
	}

	return &rpc.RegisterReply{AccessToken: accessToken}, nil
}
