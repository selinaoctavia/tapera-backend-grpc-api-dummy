package v1

import (
	"context"

	pool "github.com/processout/grpc-go-pool"
)

type (
	// GrpcClientManager interface
	GrpcClientManager interface {
		Client(ctx context.Context) (GrpcCancelSubscribeBriClient, error)
	}

	grpcClientManager struct {
		connPool *pool.Pool
	}

	// GrpcCancelSubcribeBriClient struct
	GrpcCancelSubscribeBriClient interface {
		Close()
		CancelSubscribeBriClient
	}

	grpcCancelSubscribeBriClient struct {
		conn *pool.ClientConn
		*cancelSubscribeBriClient
	}
)

// Close
func (c *grpcCancelSubscribeBriClient) Close() {
	c.conn.Close()
}

// NewGrpcClientManager func
func NewGrpcClientManager(connPool *pool.Pool) GrpcClientManager {
	return &grpcClientManager{
		connPool: connPool,
	}
}

// Client func
func (c *grpcClientManager) Client(ctx context.Context) (GrpcCancelSubscribeBriClient, error) {
	cn, err := c.connPool.Get(ctx)
	if err != nil {
		return nil, err
	}
	client := &grpcCancelSubscribeBriClient{
		conn:                     cn,
		cancelSubscribeBriClient: NewCancelSubscribeBriClient(cn).(*cancelSubscribeBriClient),
	}
	return client, nil
}
