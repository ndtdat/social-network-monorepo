package grpc

type ClientManager interface {
	Init()
	Close()
	GRPCClientManager()
}
