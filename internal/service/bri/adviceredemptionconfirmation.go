package bri

import (
	"context"
	"tapera/util/appcontext"

	arcbc "tapera.mitraintegrasi/grpc/client/adviceredemptionconfirmationbri/v1"
)

// AdviceRedemptionConfirmation func for cancel Redemption to bri
func (s *service) AdviceRedemptionConfirmation(ctx context.Context, parm *AdviceRedemptionConfirmationParam) (*AdviceRedemptionConfirmationResponse, error) {
	logger := appcontext.Logger(ctx)
	logger.Debug().Msg("Start Advice Redemption Confirmation")

	// create grpc client from manager
	arcbClient, err := s.arcbClientMgr.Client(ctx)
	if err != nil {
		return nil, err
	}

	defer arcbClient.Close()

	request := &arcbc.AdviceRedemptionConfirmationRequest{ReferenceNo: parm.ReferenceNo}

	result, err := arcbClient.AdviceRedemptionConfirmation(ctx, request)

	if err != nil {
		return nil, err
	}

	response := &AdviceRedemptionConfirmationResponse{
		Status:             result.Status,
		MessageCode:        result.MessageCode,
		MessageDescription: result.MessageDescription,
		Data: AdviceRedemptionConfirmationData{
			ReferenceNo: result.Data.ReferenceNo,
			NavDate:     result.Data.NavDate,
			Status:      result.Data.Status.String(),
		},
	}

	logger.Debug().Msg("End Advice Redemption Confirmation")

	return response, err
}
