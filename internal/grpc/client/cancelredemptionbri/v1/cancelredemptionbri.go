package v1

import (
	"context"

	pool "github.com/processout/grpc-go-pool"
)

type (
	// GrpcClientManager interface
	GrpcClientManager interface {
		Client(ctx context.Context) (GrpcCancelRedemptionBriClient, error)
	}

	grpcClientManager struct {
		connPool *pool.Pool
	}

	// GrpcCancelSubcribeBriClient struct
	GrpcCancelRedemptionBriClient interface {
		Close()
		CancelRedemptionBriClient
	}

	grpcCancelRedemptionBriClient struct {
		conn *pool.ClientConn
		*cancelRedemptionBriClient
	}
)

// Close
func (c *grpcCancelRedemptionBriClient) Close() {
	c.conn.Close()
}

// NewGrpcClientManager func
func NewGrpcClientManager(connPool *pool.Pool) GrpcClientManager {
	return &grpcClientManager{
		connPool: connPool,
	}
}

// Client func
func (c *grpcClientManager) Client(ctx context.Context) (GrpcCancelRedemptionBriClient, error) {
	cn, err := c.connPool.Get(ctx)
	if err != nil {
		return nil, err
	}
	client := &grpcCancelRedemptionBriClient{
		conn:                      cn,
		cancelRedemptionBriClient: NewCancelRedemptionBriClient(cn).(*cancelRedemptionBriClient),
	}
	return client, nil
}
