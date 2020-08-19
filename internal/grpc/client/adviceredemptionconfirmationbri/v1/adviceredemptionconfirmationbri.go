package v1

import (
	"context"

	pool "github.com/processout/grpc-go-pool"
)

type (
	// GrpcClientManager interface
	GrpcClientManager interface {
		Client(ctx context.Context) (GrpcAdviceRedemptionConfirmationBriClient, error)
	}

	grpcClientManager struct {
		connPool *pool.Pool
	}

	// GrpcAdviceRedemptionConfirmationBriClient struct
	GrpcAdviceRedemptionConfirmationBriClient interface {
		Close()
		AdviceRedemptionConfirmationBriClient
	}

	grpcAdviceRedemptionConfirmationBriClient struct {
		conn *pool.ClientConn
		*adviceRedemptionConfirmationBriClient
	}
)

// Close
func (c *grpcAdviceRedemptionConfirmationBriClient) Close() {
	c.conn.Close()
}

// NewGrpcClientManager func
func NewGrpcClientManager(connPool *pool.Pool) GrpcClientManager {
	return &grpcClientManager{
		connPool: connPool,
	}
}

// Client func
func (c *grpcClientManager) Client(ctx context.Context) (GrpcAdviceRedemptionConfirmationBriClient, error) {
	cn, err := c.connPool.Get(ctx)
	if err != nil {
		return nil, err
	}
	client := &grpcAdviceRedemptionConfirmationBriClient{
		conn:                                  cn,
		adviceRedemptionConfirmationBriClient: NewAdviceRedemptionConfirmationBriClient(cn).(*adviceRedemptionConfirmationBriClient),
	}
	return client, nil
}
