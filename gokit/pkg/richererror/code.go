package richererror

import (
	//nolint:staticcheck
	"github.com/golang/protobuf/proto"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/util"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

type ErrorCode int

func (e ErrorCode) ToInt() int {
	return int(e)
}

func (e ErrorCode) ToString() string {
	return util.IntToString(e.ToInt())
}

const ErrorCodeBase = 1000_0000

const (
	ErrorAccessTokenExpired ErrorCode = iota + ErrorCodeBase
)

func NewRicherError(grpcCode codes.Code, errorCode ErrorCode, msg string, kps ...string) error {
	return NewRicherCode(grpcCode, msg, NewErrorDetail(errorCode.ToInt(), msg, kps...))
}

func NewErrorDetail(code int, msg string, kps ...string) *errdetails.ErrorInfo {
	metadata := map[string]string{
		"code": strconv.Itoa(code),
		"msg":  msg,
	}

	nKeyPair := len(kps) / 2
	for i := 0; i < nKeyPair; i++ {
		keyIdx := i * 2
		metadata[kps[keyIdx]] = kps[keyIdx+1]
	}

	return &errdetails.ErrorInfo{Metadata: metadata}
}

func NewRicherCode(code codes.Code, msg string, details ...proto.Message) error {
	status := status.New(code, msg)

	for _, d := range details {
		status, _ = status.WithDetails(d)
	}

	return status.Err()
}
