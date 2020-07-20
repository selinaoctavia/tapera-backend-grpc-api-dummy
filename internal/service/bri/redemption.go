package bri

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"tapera/util/appcontext"
	"time"

	rbc "tapera.mitraintegrasi/grpc/client/redemptionbri/v1"
)

func (s *service) Redemption(ctx context.Context, parm *RedemptionParam) (*RedemptionResponse, error) {
	logger := appcontext.Logger(ctx)
	var response *RedemptionResponse = new(RedemptionResponse)
	logger.Debug().Msg("Start Redemption")

	// create grpc client from manager
	rbClient, err := s.rbClientMgr.Client(ctx)
	curtime := time.Now().UTC().Format(javascriptISOString)
	if err != nil {
		return nil, err
	}
	// defer the close function in order to return back the grpc connection into grpc connection pool
	defer rbClient.Close()

	var datas []*rbc.RedemptionData
	lenght := len(parm.Data)
	for i := 0; i < lenght; i++ {
		data := &rbc.RedemptionData{
			NavDate:                parm.Data[i].NavDate,
			ReferenceNo:            parm.Data[i].ReferenceNo,
			SaCode:                 parm.Data[i].SaCode,
			InvestorFundUnitAcNo:   parm.Data[i].InvestorFundUnitAcNo,
			InvestorFundUnitAcName: parm.Data[i].InvestorFundUnitAcName,
			FundCode:               parm.Data[i].FundCode,
		}
		datas = append(datas, data)
	}

	errorCount := 0
	successCount := 0
	failedCount := 0
	status := "Failed"

	request := &rbc.RedemptionRequest{Data: datas}
	stream, err := rbClient.Redemption(ctx, request)
	if err != nil {
		logger.Debug().Msgf("%v", err)
	}
	var results []*rbc.RedemptionResponse
	for {
		result, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Debug().Msgf("%v.ListFeatures(_) = _, %v", rbClient)
			logger.Debug().Msgf(err.Error())
		}
		if result.Status == "Error" {
			errorCount++
		}
		if result.Status == "Success" {
			successCount++
			status = "Success"
		}
		if result.Status == "Failed" {
			failedCount++
		}
		results = append(results, result)
		logger.Debug().Msgf(result.Status + " for Reference Number " + result.Data.ReferenceNo + ", the reason is " + result.MessageDescription)
	}
	jsonString, _ := json.Marshal(results)
	ioutil.WriteFile("../output/redemption_"+curtime+".json", jsonString, os.ModePerm)
	response.Status = status
	response.MessageDescription = "Error " + strconv.Itoa(errorCount) + ", Success " + strconv.Itoa(successCount) + ", Failed " + strconv.Itoa(failedCount)

	return response, err
}
