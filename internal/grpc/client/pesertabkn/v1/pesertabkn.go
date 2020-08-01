package v1

import (
	"context"

	pool "github.com/processout/grpc-go-pool"
)

type (
	// GrpcClientManager interface
	GrpcClientManager interface {
		Client(ctx context.Context) (GrpcPesertaBknClient, error)
	}

	grpcClientManager struct {
		connPool *pool.Pool
	}

	// GrpcPesertaBknClient struct
	GrpcPesertaBknClient interface {
		Close()
		PesertaBknClient
	}

	grpcPesertaBknClient struct {
		conn *pool.ClientConn
		*pesertaBknClient
	}
)

// Close
func (c *grpcPesertaBknClient) Close() {
	c.conn.Close()
}

// NewGrpcClientManager func
func NewGrpcClientManager(connPool *pool.Pool) GrpcClientManager {
	return &grpcClientManager{
		connPool: connPool,
	}
}

// Client func
func (c *grpcClientManager) Client(ctx context.Context) (GrpcPesertaBknClient, error) {
	cn, err := c.connPool.Get(ctx)
	if err != nil {
		return nil, err
	}
	client := &grpcPesertaBknClient{
		conn:             cn,
		pesertaBknClient: NewPesertaBknClient(cn).(*pesertaBknClient),
	}
	return client, nil
}
