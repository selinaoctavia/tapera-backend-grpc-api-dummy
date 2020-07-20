package bri

import (
	"context"
	"tapera/util/appcontext"

	crbc "tapera.mitraintegrasi/grpc/client/cancelredemptionbri/v1"
)

// CancelRedemption func for cancel Redemption to bri
func (s *service) CancelRedemption(ctx context.Context, parm *CancelRedemptionParam) (*CancelRedemptionResponse, error) {
	logger := appcontext.Logger(ctx)
	logger.Debug().Msg("Start Cancel Redemption")

	// create grpc client from manager
	crbClient, err := s.crbClientMgr.Client(ctx)
	if err != nil {
		return nil, err
	}

	defer crbClient.Close()

	request := &crbc.CancelRedemptionRequest{ReferenceNo: parm.ReferenceNo}

	result, err := crbClient.CancelRedemption(ctx, request)

	if err != nil {
		return nil, err
	}

	response := &CancelRedemptionResponse{
		Status:             result.Status,
		MessageCode:        result.MessageCode,
		MessageDescription: result.MessageDescription,
		Data: &CancelRedemptionParam{
			ReferenceNo: result.Data.ReferenceNo,
		},
	}

	logger.Debug().Msg("End Cancel Redemption")

	return response, err
}
