package auth

import (
	"context"
	"fmt"
	"github.com/ndtdat/social-network-monorepo/user-service/pkg/api/go/user/rpc"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(ctx context.Context, in *rpc.LoginRequest) (*rpc.LoginReply, error) {
	email := in.GetEmail()

	// Check email is existed
	userInfo, err := s.repo.FirstByFilters(ctx, email)
	if err != nil {
		return nil, err
	}
	if userInfo == nil {
		return nil, fmt.Errorf("user email [%v] not found", email)
	}

	// Check password is correct
	if err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(in.GetPassword())); err != nil {
		return nil, fmt.Errorf("password invalid")
	}

	// Gen access token
	accessToken, err := s.GenAccessToken(userInfo.ID)
	if err != nil {
		return nil, err
	}

	return &rpc.LoginReply{AccessToken: accessToken}, nil
}
