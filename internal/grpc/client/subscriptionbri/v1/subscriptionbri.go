package v1

import (
	"context"

	pool "github.com/processout/grpc-go-pool"
)

type (
	// GrpcClientManager interface
	GrpcClientManager interface {
		Client(ctx context.Context) (GrpcSubscriptionBriClient, error)
	}

	grpcClientManager struct {
		connPool *pool.Pool
	}

	// GrpcMitraintegrasiClient struct
	GrpcSubscriptionBriClient interface {
		Close()
		SubscriptionBriClient
	}

	grpcSubscriptionBriClient struct {
		conn *pool.ClientConn
		*subscriptionBriClient
	}
)

// Close
func (c *grpcSubscriptionBriClient) Close() {
	c.conn.Close()
}

// NewGrpcClientManager func
func NewGrpcClientManager(connPool *pool.Pool) GrpcClientManager {
	return &grpcClientManager{
		connPool: connPool,
	}
}

// Client func
func (c *grpcClientManager) Client(ctx context.Context) (GrpcSubscriptionBriClient, error) {
	cn, err := c.connPool.Get(ctx)
	if err != nil {
		return nil, err
	}
	client := &grpcSubscriptionBriClient{
		conn:                  cn,
		subscriptionBriClient: NewSubscriptionBriClient(cn).(*subscriptionBriClient),
	}
	return client, nil
}
