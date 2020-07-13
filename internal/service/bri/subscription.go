package bri

import (
	"context"
	"log"
	"tapera/util/appcontext"
	"time"

	sbc "tapera.mitraintegrasi/grpc/client/subscriptionbri/v1"
)

func (s *service) Subscription(ctx context.Context, parm *SubscriptionParam) (*SubscriptionResponse, error) {
	logger := appcontext.Logger(ctx)
	start := time.Now()
	var response *SubscriptionResponse = new(SubscriptionResponse)
	logger.Debug().Msg("Start Subscription")

	// create grpc client from manager
	sbClient, err := s.sbClientMgr.Client(ctx)
	if err != nil {
		return nil, err
	}
	// defer the close function in order to return back the grpc connection into grpc connection pool
	defer sbClient.Close()

	var details []*sbc.Detail
	lenght := len(parm.Details)
	for i := 0; i < lenght; i++ {
		detail := sbc.Detail{
			NavDate:                parm.Details[i].NavDate,
			ReferenceNo:            parm.Details[i].ReferenceNo,
			SaCode:                 parm.Details[i].SaCode,
			InvestorFundUnitAcNo:   parm.Details[i].InvestorFundUnitAcNo,
			InvestorFundUnitAcName: parm.Details[i].InvestorFundUnitAcName,
			Sid:                    parm.Details[i].Sid,
			FundCode:               parm.Details[i].FundCode,
			AmountNominal:          parm.Details[i].AmountNominal,
		}
		details = append(details, &detail)
	}

	data := &sbc.SubscribData{
		PaymentPoolId:        parm.PaymentPoolID,
		PaymentBankBicCode:   parm.PaymentBankBicCode,
		JumlahTransaksiCount: parm.JumlahTransaksiCount,
		TotalAmountNominal:   parm.TotalAmountNominal,
		Detail:               details,
	}

	log.Println("Sending Data to gRPC Integrasi Mitra ......")
	subscribRequest := &sbc.SubscribRequest{Parm: data}
	result, err := sbClient.Subscription(ctx, subscribRequest)
	if err != nil {
		logger.Debug().Msgf("%v", err)
	}

	log.Println("response_code: ", result.ResponseCode)
	log.Println("response_description: ", result.ResponseDescription)
	log.Println("response_status: ", result.Status)

	response.ResponseCode = result.ResponseCode
	response.ResponseDescription = result.ResponseDescription
	response.ResponseStatus = result.Status

	elapsed := time.Since(start)
	log.Println("Subscription Time: ", elapsed)
	log.Println("End Subscription . . .")
	return response, err
}
