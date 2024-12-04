package util

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func FieldFromIncomingCtx(ctx context.Context, field string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	return FieldFromMD(md, field)
}

func FieldFromMD(md metadata.MD, field string) string {
	values := md.Get(field)
	if len(values) < 1 {
		return ""
	}

	return values[0]
}
