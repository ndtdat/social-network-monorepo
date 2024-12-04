package util

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/common"
	"google.golang.org/protobuf/proto"
)

func UnmarshalAndValidateMsg(_ context.Context, data []byte, req common.ProtoMsgWrapper) (proto.Message, error) {
	err := proto.Unmarshal(data, req)
	if err != nil {
		return nil, err
	}

	if err = req.Validate(); err != nil {
		return nil, err
	}

	return req, nil
}
