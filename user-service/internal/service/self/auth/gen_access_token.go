package auth

import (
	"fmt"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/util"
)

func (s *Service) GenAccessToken(userID uint64) (string, error) {
	jm := s.jm
	idStr := util.Uint64ToString(userID)

	// TODO: Hardcode Role for implementation quickly
	accessToken, err := jm.GenerateAccessToken(idStr, "", "", "", []string{"USER"})
	if err != nil {
		return "", fmt.Errorf("cannot generate access token due to: %v", err)
	}

	return accessToken, nil
}
