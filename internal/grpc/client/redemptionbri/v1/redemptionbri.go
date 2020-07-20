package v1

import (
	"context"

	pool "github.com/processout/grpc-go-pool"
)

type (
	// GrpcClientManager interface
	GrpcClientManager interface {
		Client(ctx context.Context) (GrpcRedemptionBriClient, error)
	}

	grpcClientManager struct {
		connPool *pool.Pool
	}

	// GrpcRedemptionBriClient struct
	GrpcRedemptionBriClient interface {
		Close()
		RedemptionBriClient
	}

	grpcRedemptionBriClient struct {
		conn *pool.ClientConn
		*redemptionBriClient
	}
)

// Close
func (c *grpcRedemptionBriClient) Close() {
	c.conn.Close()
}

// NewGrpcClientManager func
func NewGrpcClientManager(connPool *pool.Pool) GrpcClientManager {
	return &grpcClientManager{
		connPool: connPool,
	}
}

// Client func
func (c *grpcClientManager) Client(ctx context.Context) (GrpcRedemptionBriClient, error) {
	cn, err := c.connPool.Get(ctx)
	if err != nil {
		return nil, err
	}
	client := &grpcRedemptionBriClient{
		conn:                cn,
		redemptionBriClient: NewRedemptionBriClient(cn).(*redemptionBriClient),
	}
	return client, nil
}
