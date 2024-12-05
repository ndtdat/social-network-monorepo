package grpc

import (
	"context"
	"crypto/x509"
	"fmt"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/common"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/credentials/xds"
	"google.golang.org/grpc/metadata"
	//noline:revive
	_ "google.golang.org/grpc/resolver"
	_ "google.golang.org/grpc/xds"
)

type Client struct {
	Client *grpc.ClientConn
	host   string
	port   string
	opts   []grpc.DialOption
}

func NewClient(
	ctx context.Context, host, port string, tls bool, opts ...grpc.DialOption,
) (*Client, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	gRPCClient := &Client{
		host: host,
		port: port,
		opts: opts,
	}

	unaryClientInterceptors := []grpc.UnaryClientInterceptor{
		gRPCClient.injectRequestMetadata,
	}

	options := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithChainUnaryInterceptor(unaryClientInterceptors...),
	}

	var (
		creds credentials.TransportCredentials
		err   error
	)
	if tls {
		pool, _ := x509.SystemCertPool()
		creds = credentials.NewClientTLSFromCert(pool, "")
	} else {
		creds, _ = xds.NewClientCredentials(xds.ClientOptions{
			FallbackCreds: insecure.NewCredentials(),
		})
	}

	options = append(options, grpc.WithTransportCredentials(creds))
	options = append(options, opts...)

	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%s", host, port), options...)
	if err != nil {
		return nil, err
	}

	gRPCClient.Client = conn

	return gRPCClient, nil
}

func (c *Client) injectRequestMetadata(
	ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	return invoker(c.GenRequestMetadata(ctx), method, req, reply, cc, opts...)
}

func (c *Client) GenRequestMetadata(ctx context.Context) context.Context {
	var outMD metadata.MD
	inMd, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		outMD = metadata.New(nil)
	} else {
		outMD = inMd.Copy()
	}

	outMD.Set(common.InternalCallHeader, "true")

	return metadata.NewOutgoingContext(ctx, outMD)
}
