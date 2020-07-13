package bri

import (
	"context"
	"tapera/util/appcontext"

	csbc "tapera.mitraintegrasi/grpc/client/cancelsubscribebri/v1"
)

// CancelSubscribe func for cancel subscribe to bri
func (s *service) CancelSubscribe(ctx context.Context, parm *CancelSubscribeParam) (*CancelSubscribeResponse, error) {
	logger := appcontext.Logger(ctx)
	// var response *PendaftaranPesertaResponse = new(PendaftaranPesertaResponse)
	logger.Debug().Msg("Start Cancel Subscribe")
	// curtime := time.Now().UTC().Format(javascriptISOString)

	// create grpc client from manager
	csbClient, err := s.csbClientMgr.Client(ctx)
	if err != nil {
		return nil, err
	}

	defer csbClient.Close()

	request := &csbc.CancelSubscribeRequest{PaymentPoolId: parm.PaymentPoolID}

	result, err := csbClient.CancelSubscribe(ctx, request)

	if err != nil {
		return nil, err
	}

	response := &CancelSubscribeResponse{
		Status:             result.Status,
		MessageCode:        result.MessageCode,
		MessageDescription: result.MessageDescription,
		Data: &Data{
			PaymentPoolID: result.Data.PaymentPoolId,
		},
	}

	logger.Debug().Msg("End Cancel Subscribe")

	return response, err
}
