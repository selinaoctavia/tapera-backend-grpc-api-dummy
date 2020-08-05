package v1

import (
	"context"

	pool "github.com/processout/grpc-go-pool"
)

type (
	// GrpcClientManager interface
	GrpcClientManager interface {
		Client(ctx context.Context) (GrpcRiwayatGolonganPesertaBknClient, error)
	}

	grpcClientManager struct {
		connPool *pool.Pool
	}

	// GrpcRiwayatGolonganPesertaBknClient struct
	GrpcRiwayatGolonganPesertaBknClient interface {
		Close()
		RiwayatGolonganPesertaBknClient
	}

	grpcRiwayatGolonganPesertaBknClient struct {
		conn *pool.ClientConn
		*riwayatGolonganPesertaBknClient
	}
)

// Close
func (c *grpcRiwayatGolonganPesertaBknClient) Close() {
	c.conn.Close()
}

// NewGrpcClientManager func
func NewGrpcClientManager(connPool *pool.Pool) GrpcClientManager {
	return &grpcClientManager{
		connPool: connPool,
	}
}

// Client func
func (c *grpcClientManager) Client(ctx context.Context) (GrpcRiwayatGolonganPesertaBknClient, error) {
	cn, err := c.connPool.Get(ctx)
	if err != nil {
		return nil, err
	}
	client := &grpcRiwayatGolonganPesertaBknClient{
		conn:                            cn,
		riwayatGolonganPesertaBknClient: NewRiwayatGolonganPesertaBknClient(cn).(*riwayatGolonganPesertaBknClient),
	}
	return client, nil
}
