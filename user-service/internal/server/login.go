package server

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/user-service/pkg/api/go/user/rpc"
)

func (u *UserServer) Login(ctx context.Context, in *rpc.LoginRequest) (*rpc.LoginReply, error) {
	return u.authService.Login(ctx, in)
}
