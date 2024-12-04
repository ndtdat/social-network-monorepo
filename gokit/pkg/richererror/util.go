package richererror

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/common"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/log"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/util"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"strings"

	//nolint:staticcheck
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func GRPCWebIOSErrorWrapper(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}
	if isGRPCWebRequest(ctx) || isIOSRequest(ctx) {
		return wrapStatus(ctx, status.Convert(err))
	}

	return err
}

func wrapStatus(ctx context.Context, status *status.Status) error {
	setErrorToTrailer(ctx, status)

	return status.Err()
}

func setErrorToTrailer(ctx context.Context, status *status.Status) {
	details := status.Details()
	if len(details) == 0 {
		return
	}

	if detailStr, err := json.Marshal(details); err != nil {
		log.WithContext(ctx).Debug("Cannot decode status, therefore cannot set error to trailers")
	} else {
		trailer := metadata.Pairs(common.RicherErrorDetailsHeader, base64.StdEncoding.EncodeToString(detailStr))
		if err := grpc.SetTrailer(ctx, trailer); err != nil {
			log.WithContext(ctx).Debug(fmt.Sprintf("Cannot set trailer %v to context", trailer))
		}
	}
}

func isGRPCWebRequest(ctx context.Context) bool {
	return util.FieldFromIncomingCtx(ctx, common.GRPCWebHeader) == "1"
}

func isIOSRequest(ctx context.Context) bool {
	return strings.Contains(util.FieldFromIncomingCtx(ctx, common.UserAgentHeader), "grpc-swift")
}

func GetGRPCErrorDetails(err error) []*errdetails.ErrorInfo {
	var results = make([]*errdetails.ErrorInfo, 0)

	if e, ok := status.FromError(err); ok {
		for _, detail := range e.Details() {
			errInfo := detail.(*errdetails.ErrorInfo)
			if errInfo != nil {
				results = append(results, errInfo)
			}
		}
	}

	return results
}

func GetGRPCErrorDetailByCode(err error, code ErrorCode) *errdetails.ErrorInfo {
	for _, errInfo := range GetGRPCErrorDetails(err) {
		val, ok := errInfo.GetMetadata()["code"]
		if ok && val == code.ToString() {
			return errInfo
		}
	}

	return nil
}

func CheckGRPCErrorByCode(err error, code ErrorCode) bool {
	for _, errInfo := range GetGRPCErrorDetails(err) {
		val, ok := errInfo.GetMetadata()["code"]
		if ok && val == code.ToString() {
			return true
		}
	}

	return false
}
