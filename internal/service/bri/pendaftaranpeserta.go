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

	ppbc "tapera.mitraintegrasi/grpc/client/pendaftaranpesertabri/v1"
)

const javascriptISOString = "2006-01-02T150405"

// PendaftaranPeserta func for pendaftaran peserta to bri
func (s *service) PendaftaranPeserta(ctx context.Context, parm *PendaftaranPesertaParam) (*PendaftaranPesertaResponse, error) {
	logger := appcontext.Logger(ctx)
	var response *PendaftaranPesertaResponse = new(PendaftaranPesertaResponse)
	logger.Debug().Msg("Start PendaftaranPeserta")
	curtime := time.Now().UTC().Format(javascriptISOString)

	// create grpc client from manager
	miClient, err := s.ppbClientMgr.Client(ctx)
	if err != nil {
		return nil, err
	}
	// defer the close function in order to return back the grpc connection into grpc connection pool
	defer miClient.Close()

	// req := &PendaftaranPesertaRequest{
	// 	Data: []*PendaftaranPesertaData{
	// 		&PendaftaranPesertaData{
	// 			processing_date:       "20200309",
	// 			processing_time:       "180500",
	// 			sa_code:               "DBSI1",
	// 			sid:                   "GBU115",
	// 			sid_generation_status: "1",
	// 			new_sid:               "",
	// 			first_name:            "GBU",
	// 			middle_name:           "NASABAH",
	// 			last_name:             "DUA",
	// 			nationality:           "ID",
	// 			identification:        "1",
	// 			id_no:                 "1774647877",
	// 			id_exp_date:           "20200405",
	// 			npwp_no:               "775757889999999",
	// 			npwp_date:             "20200405",
	// 			birth_country:         "ID",
	// 			birth_place:           "SIGUMPAR",
	// 			birth_date:            "19880101",
	// 			gender:                "1",
	// 			education_background:  "1",
	// 			mother_name:           "MOTHERNAME",
	// 			religion:              "1",
	// 			occupation:            "1",
	// 			income_level:          "1",
	// 			marital_status:        "1",
	// 			spouse_name:           "GBU NASABAH SATU",
	// 			risk_profile:          "1",
	// 			investment_objective:  "1",
	// 			source_of_fund:        "1",
	// 			asset_owner:           "1",
	// 			ktp_address:           "CIKOKO",
	// 			ktp_city_code:         "3171",
	// 			ktp_postal:            "12345",
	// 			corr_address:          "TEBET",
	// 			corr_city_code:        "3171",
	// 			corr_city_name:        "JAKARTA SELATAN",
	// 			corr_postal:           "12345",
	// 			corr_country:          "ID",
	// 			dom_address:           "JAKARTA SELATAN",
	// 			dom_city_code:         "3171",
	// 			dom_city_name:         "JAKARTA SELATAN",
	// 			dom_postal:            "12345",
	// 			dom_country:           "ID",
	// 			home_phone:            "65446679",
	// 			mobile_phone:          "081111111111",
	// 			facsimile:             "775890",
	// 			email:                 "kame@gmail.com",
	// 			statement_type:        "1",
	// 			fatca:                 "1",
	// 			tin:                   "123456789012345678901234567890",
	// 			tin_country:           "ID",
	// 			external_cif_no:       "",
	// 			type: "1",
	// 			investor_type:            "1",
	// 			red_pay_bank_bic_1:       "BIC001",
	// 			red_pay_bank_bi_1:        "MC001",
	// 			red_pay_bank_name_1:      "MAYBANK",
	// 			red_pay_bank_country_1:   "ID",
	// 			red_pay_bank_branch_1:    "THAMRIN",
	// 			red_pay_bank_acc_ccy_1:   "IDR",
	// 			REDM_Payment_ACNo_1:      "1234567",
	// 			red_pay_bank_acc_name_1:  "GBU NASABAH DUA BANK SATU",
	// 			red_pay_bank_bic_2:       "BIC002",
	// 			red_pay_bank_bi_2:        "MC02",
	// 			red_pay_bank_name_2:      "BANK RAKYAT INDONESIA",
	// 			red_pay_bank_country_2:   "ID",
	// 			red_pay_bank_branch_2:    "THAMRIN",
	// 			red_pay_bank_acc_ccy_2:   "USD",
	// 			red_pay_bank_acc_no_2:    "2345678",
	// 			red_pay_bank_acc_name_2:  "GBU NASABAH DUA BANK DUA",
	// 			red_pay_bank_bic_3:       "BIC003",
	// 			red_pay_bank_bi_3:        "MC03",
	// 			red_pay_bank_name_3:      "BANK RAKYAT INDONESIA",
	// 			red_pay_bank_country_3:   "ID",
	// 			red_pay_bank_branch_3:    "THAMRIN",
	// 			red_pay_bank_acc_ccy_3:   "USD",
	// 			red_pay_bank_acc_no_3:    "3456789",
	// 			red_pay_bank_acc_name_3:  "GBU NASABAH DUA BANK TIGA",
	// 			client_code:              "CL003",
	// 			investor_fund_unit_ac_no: "CL015",
	// 			investor_number:          "81020710015",
	// 		},
	// 	},
	// }

	var datas []*ppbc.PendaftaranPesertaData
	for i := 0; i < len(parm.Data); i++ {
		data := &ppbc.PendaftaranPesertaData{
			ProcessingDate:       parm.Data[i].ProcessingDate,
			ProcessingTime:       parm.Data[i].ProcessingTime,
			Sid:                  parm.Data[i].Sid,
			SaCode:               parm.Data[i].SaCode,
			SidGenerationStatus:  parm.Data[i].SidGenerationStatus,
			NewSid:               parm.Data[i].NewSid,
			FirstName:            parm.Data[i].FirstName,
			MiddleName:           parm.Data[i].MiddleName,
			LastName:             parm.Data[i].LastName,
			Nationality:          parm.Data[i].Nationality,
			Identification:       parm.Data[i].Identification,
			IdNo:                 parm.Data[i].IDNo,
			IdExpDate:            parm.Data[i].IDExpDate,
			NpwpNo:               parm.Data[i].NpwpNo,
			NpwpDate:             parm.Data[i].NpwpDate,
			BirthCountry:         parm.Data[i].BirthCountry,
			BirthPlace:           parm.Data[i].BirthPlace,
			BirthDate:            parm.Data[i].BirthDate,
			Gender:               parm.Data[i].Gender,
			EducationBackground:  parm.Data[i].EducationBackground,
			MotherName:           parm.Data[i].MotherName,
			Religion:             parm.Data[i].Religion,
			Occupation:           parm.Data[i].Occupation,
			IncomeLevel:          parm.Data[i].IncomeLevel,
			MaritalStatus:        parm.Data[i].MaritalStatus,
			SpouseName:           parm.Data[i].SpouseName,
			RiskProfile:          parm.Data[i].RiskProfile,
			InvestmentObjective:  parm.Data[i].InvestmentObjective,
			SourceOfFund:         parm.Data[i].SourceOfFund,
			AssetOwner:           parm.Data[i].AssetOwner,
			KtpAddress:           parm.Data[i].KtpAddress,
			KtpCityCode:          parm.Data[i].KtpCityCode,
			KtpPostal:            parm.Data[i].KtpPostal,
			CorrAddress:          parm.Data[i].CorrAddress,
			CorrCityCode:         parm.Data[i].CorrCityCode,
			CorrCityName:         parm.Data[i].CorrCityName,
			CorrPostal:           parm.Data[i].CorrPostal,
			CorrCountry:          parm.Data[i].CorrCountry,
			DomAddress:           parm.Data[i].DomAddress,
			DomCityCode:          parm.Data[i].DomCityCode,
			DomCityName:          parm.Data[i].DomCityName,
			DomPostal:            parm.Data[i].DomPostal,
			DomCountry:           parm.Data[i].DomCountry,
			HomePhone:            parm.Data[i].HomePhone,
			MobilePhone:          parm.Data[i].MobilePhone,
			Facsimile:            parm.Data[i].Facsimile,
			Email:                parm.Data[i].Email,
			StatementType:        parm.Data[i].StatementType,
			Fatca:                parm.Data[i].Fatca,
			Tin:                  parm.Data[i].Tin,
			TinCountry:           parm.Data[i].TinCountry,
			ExternalCifNo:        parm.Data[i].ExternalCifNo,
			Type:                 parm.Data[i].Type,
			InvestorType:         parm.Data[i].InvestorType,
			RedPayBankBic_1:      parm.Data[i].RedPayBankBic1,
			RedPayBankBi_1:       parm.Data[i].RedPayBankBi1,
			RedPayBankName_1:     parm.Data[i].RedPayBankName1,
			RedPayBankCountry_1:  parm.Data[i].RedPayBankCountry1,
			RedPayBankBranch_1:   parm.Data[i].RedPayBankBranch1,
			RedPayBankAccCcy_1:   parm.Data[i].RedPayBankAccCcy1,
			REDM_Payment_ACNo_1:  parm.Data[i].REDMPaymentACNo1,
			RedPayBankAccName_1:  parm.Data[i].RedPayBankAccName1,
			RedPayBankBic_2:      parm.Data[i].RedPayBankBic2,
			RedPayBankBi_2:       parm.Data[i].RedPayBankBi2,
			RedPayBankName_2:     parm.Data[i].RedPayBankName2,
			RedPayBankCountry_2:  parm.Data[i].RedPayBankCountry2,
			RedPayBankBranch_2:   parm.Data[i].RedPayBankBranch2,
			RedPayBankAccCcy_2:   parm.Data[i].RedPayBankAccCcy2,
			RedPayBankAccNo_2:    parm.Data[i].RedPayBankAccNo2,
			RedPayBankAccName_2:  parm.Data[i].RedPayBankAccName2,
			RedPayBankBic_3:      parm.Data[i].RedPayBankBic3,
			RedPayBankBi_3:       parm.Data[i].RedPayBankBi3,
			RedPayBankName_3:     parm.Data[i].RedPayBankName3,
			RedPayBankCountry_3:  parm.Data[i].RedPayBankCountry3,
			RedPayBankBranch_3:   parm.Data[i].RedPayBankBranch3,
			RedPayBankAccCcy_3:   parm.Data[i].RedPayBankAccCcy3,
			RedPayBankAccNo_3:    parm.Data[i].RedPayBankAccNo3,
			RedPayBankAccName_3:  parm.Data[i].RedPayBankAccName3,
			ClientCode:           parm.Data[i].ClientCode,
			InvestorFundUnitAcNo: parm.Data[i].InvestorFundUnitAcNo,
			InvestorNumber:       parm.Data[i].InvestorNumber,
		}
		datas = append(datas, data)
	}

	errorCount := 0
	successCount := 0
	failedCount := 0
	status := "Failed"

	// // var rsTrial *v1.PendaftaranPesertaResponse
	request := &ppbc.PendaftaranPesertaRequest{Data: datas}
	stream, err := miClient.PendaftaranPeserta(ctx, request)
	if err != nil {
		logger.Debug().Msgf("%v", err)
	}
	var results []*ppbc.PendaftaranPesertaResponse
	for {
		result, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Debug().Msgf("%v.ListFeatures(_) = _, %v", miClient)
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
		logger.Debug().Msgf(result.Status + " for SID " + result.Data.Sid + ", the reason is " + result.MessageDescription)
	}
	jsonString, _ := json.Marshal(results)
	ioutil.WriteFile("../output/"+curtime+".json", jsonString, os.ModePerm)
	response.Status = status
	response.MessageDescription = "Error " + strconv.Itoa(errorCount) + ", Success " + strconv.Itoa(successCount) + ", Failed " + strconv.Itoa(failedCount)

	logger.Debug().Msg("End PendaftaranPeserta")

	return response, err
}
