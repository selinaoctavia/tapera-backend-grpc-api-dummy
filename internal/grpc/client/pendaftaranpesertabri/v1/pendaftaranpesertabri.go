package v1

import (
	"context"

	pool "github.com/processout/grpc-go-pool"
)

type (
	// GrpcClientManager interface
	GrpcClientManager interface {
		Client(ctx context.Context) (GrpcPendaftaranPesertaBriClient, error)
	}

	grpcClientManager struct {
		connPool *pool.Pool
	}

	// GrpcPendaftaranPesertaBriClient struct
	GrpcPendaftaranPesertaBriClient interface {
		Close()
		PendaftaranPesertaBriClient
	}

	grpcPendaftaranPesertaBriClient struct {
		conn *pool.ClientConn
		*pendaftaranPesertaBriClient
	}
)

// Close
func (c *grpcPendaftaranPesertaBriClient) Close() {
	c.conn.Close()
}

// NewGrpcClientManager func
func NewGrpcClientManager(connPool *pool.Pool) GrpcClientManager {
	return &grpcClientManager{
		connPool: connPool,
	}
}

// Client func
func (c *grpcClientManager) Client(ctx context.Context) (GrpcPendaftaranPesertaBriClient, error) {
	cn, err := c.connPool.Get(ctx)
	if err != nil {
		return nil, err
	}
	client := &grpcPendaftaranPesertaBriClient{
		conn:                        cn,
		pendaftaranPesertaBriClient: NewPendaftaranPesertaBriClient(cn).(*pendaftaranPesertaBriClient),
	}
	return client, nil
}
