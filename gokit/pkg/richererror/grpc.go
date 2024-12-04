package richererror

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Aborted(reason string) error {
	return status.New(codes.Aborted, reason).Err()
}

func Ok(reason string) error {
	return status.New(codes.OK, reason).Err()
}

func Cancelled(reason string) error {
	return status.New(codes.Canceled, reason).Err()
}

func Unknown(reason string) error {
	return status.New(codes.Unknown, reason).Err()
}

func InvalidArgument(reason string) error {
	return status.New(codes.InvalidArgument, reason).Err()
}

func DeadlineExceeded(reason string) error {
	return status.New(codes.DeadlineExceeded, reason).Err()
}

func NotFound(reason string) error {
	return status.New(codes.NotFound, reason).Err()
}

func AlreadyExists(reason string) error {
	return status.New(codes.AlreadyExists, reason).Err()
}

func PermissionDenied(reason string) error {
	return status.New(codes.PermissionDenied, reason).Err()
}

func ResourceExhausted(reason string) error {
	return status.New(codes.ResourceExhausted, reason).Err()
}

func FailedPrecondition(reason string) error {
	return status.New(codes.FailedPrecondition, reason).Err()
}

func OutOfRange(reason string) error {
	return status.New(codes.OutOfRange, reason).Err()
}

func Unimplemented(reason string) error {
	return status.New(codes.Unimplemented, reason).Err()
}

func Internal(reason string) error {
	return status.New(codes.Internal, reason).Err()
}

func Unavailable(reason string) error {
	return status.New(codes.Unavailable, reason).Err()
}

func DataLoss(reason string) error {
	return status.New(codes.DataLoss, reason).Err()
}

func Unauthenticated(reason string) error {
	return status.New(codes.Unauthenticated, reason).Err()
}
