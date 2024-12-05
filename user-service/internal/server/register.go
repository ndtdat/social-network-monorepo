package server

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/user-service/pkg/api/go/user/rpc"
)

func (u *UserServer) Register(ctx context.Context, in *rpc.RegisterRequest) (*rpc.RegisterReply, error) {
	return u.authService.Register(ctx, in)
}
