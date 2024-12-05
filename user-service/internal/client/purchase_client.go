package client

import (
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/client/grpc"
	"github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase"
	grpcconn "google.golang.org/grpc"
)

func (m *MicroservicesManager) newPurchaseClient(
	host, port string, tls bool,
) (purchase.PurchaseClient, *grpcconn.ClientConn, error) {
	stub, err := grpc.NewClient(m.ctx, host, port, tls)
	if err != nil {
		return nil, nil, err
	}

	client := stub.Client

	return purchase.NewPurchaseClient(client), client, nil
}
