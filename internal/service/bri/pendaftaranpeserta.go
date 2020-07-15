package bri

import (
	"context"
	"encoding/json"
	"errors"
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
	if parm.Data == nil {
		return nil, errors.New("Data nil")
	}
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
	var sidGenerationStatus ppbc.SIDGenerationStatus
	var identification ppbc.Identification
	var gender ppbc.Gender
	var educationBackground ppbc.EducationalBackground
	var religion ppbc.Religion
	var occupation ppbc.Occupation
	var incomeLevel ppbc.IncomeLevel
	var maritalStatus ppbc.MaritalStatus
	var riskProfile ppbc.RiskProfile
	var investmentObjective ppbc.InvestmentObjective
	var sourceOfFund ppbc.SourceOfFund
	var assetOwner ppbc.AssetOwner
	var statementType ppbc.StatementType
	var fatca ppbc.Fatca
	var typeValue ppbc.Type
	var investorType ppbc.InvestorType

	for i := 0; i < len(parm.Data); i++ {

		switch parm.Data[i].SidGenerationStatus {
		case "1":
			sidGenerationStatus = ppbc.SIDGenerationStatus_COMPLETED
		case "2":
			sidGenerationStatus = ppbc.SIDGenerationStatus_NA
		case "3":
			sidGenerationStatus = ppbc.SIDGenerationStatus_FAILED
		default:
			sidGenerationStatus = ppbc.SIDGenerationStatus_UNKNOWN_SID_GENERATION_STATUS
		}

		switch parm.Data[i].Identification {
		case "1":
			identification = ppbc.Identification_KTP
		case "2":
			identification = ppbc.Identification_PASSPORT
		default:
			identification = ppbc.Identification_UNKNOWN_IDENTIFICATION
		}

		switch parm.Data[i].Gender {
		case "1":
			gender = ppbc.Gender_MALE
		case "2":
			gender = ppbc.Gender_FEMALE
		default:
			gender = ppbc.Gender_UNKNOWN_GENDER
		}

		switch parm.Data[i].EducationBackground {
		case "1":
			educationBackground = ppbc.EducationalBackground_PRIMARY_SCHOOL
		case "2":
			educationBackground = ppbc.EducationalBackground_JUNIOR_HIGH_SCHOOL
		case "3":
			educationBackground = ppbc.EducationalBackground_SENIOR_HIGH_SCHOOL
		case "4":
			educationBackground = ppbc.EducationalBackground_DIPLOMA
		case "5":
			educationBackground = ppbc.EducationalBackground_UNDERGRADUATE
		case "6":
			educationBackground = ppbc.EducationalBackground_POST_GRADUATE
		case "7":
			educationBackground = ppbc.EducationalBackground_DOCTORAL_PROGRAM
		case "8":
			educationBackground = ppbc.EducationalBackground_OTHER_EDUCATIONAL_BACKGROUND
		default:
			educationBackground = ppbc.EducationalBackground_UNKNOWN_EDUCATIONAL_BACKGROUND
		}

		switch parm.Data[i].Religion {
		case "1":
			religion = ppbc.Religion_ISLAM
		case "2":
			religion = ppbc.Religion_PROTESTANT
		case "3":
			religion = ppbc.Religion_CATHOLIC
		case "4":
			religion = ppbc.Religion_HINDUISM
		case "5":
			religion = ppbc.Religion_BUDHISM
		case "6":
			religion = ppbc.Religion_CONFUCIANISM
		case "7":
			religion = ppbc.Religion_OTHER_RELIGION
		default:
			religion = ppbc.Religion_UNKNOWN_RELIGION
		}

		switch parm.Data[i].Occupation {
		case "1":
			occupation = ppbc.Occupation_STUDENT
		case "2":
			occupation = ppbc.Occupation_HOUSEWIFE
		case "3":
			occupation = ppbc.Occupation_ENTERPRENEUR
		case "4":
			occupation = ppbc.Occupation_CIVIL_SERVANT
		case "5":
			occupation = ppbc.Occupation_INDONESIA_NATIONAL_ARMED_FORCE_OR_POLICE
		case "6":
			occupation = ppbc.Occupation_RETIREMENT
		case "7":
			occupation = ppbc.Occupation_LECTURER_OR_TEACHER
		case "8":
			occupation = ppbc.Occupation_PRIVATE_EMPLOYEE
		case "9":
			occupation = ppbc.Occupation_OTHER_OCCUPATION
		default:
			occupation = ppbc.Occupation_UNKNOWN_OCCUPATION
		}

		switch parm.Data[i].IncomeLevel {
		case "1":
			incomeLevel = ppbc.IncomeLevel_LESS_THAN_10_MILLION_PER_YEAR
		case "2":
			incomeLevel = ppbc.IncomeLevel_BETWEEN_10_UNTIL_50_MILLION_PER_YEAR
		case "3":
			incomeLevel = ppbc.IncomeLevel_BETWEEN_50_UNTIL_100_MILLION_PER_YEAR
		case "4":
			incomeLevel = ppbc.IncomeLevel_BETWEEN_100_UNTIL_500_MILLION_PER_YEAR
		case "5":
			incomeLevel = ppbc.IncomeLevel_BETWEEN_500_UNTIL_1_BILLION_PER_YEAR
		case "6":
			incomeLevel = ppbc.IncomeLevel_MORE_THAN_1_BILLION_PER_YEAR
		default:
			incomeLevel = ppbc.IncomeLevel_UNKNOWN_INCOME_LEVEL
		}

		switch parm.Data[i].MaritalStatus {
		case "1":
			maritalStatus = ppbc.MaritalStatus_SINGLE
		case "2":
			maritalStatus = ppbc.MaritalStatus_MARRIED
		case "3":
			maritalStatus = ppbc.MaritalStatus_DIVORCE
		default:
			maritalStatus = ppbc.MaritalStatus_UNKNOWN_MARITALSTATUS
		}

		switch parm.Data[i].RiskProfile {
		case "1":
			riskProfile = ppbc.RiskProfile_LOW
		case "2":
			riskProfile = ppbc.RiskProfile_LOW_TO_MODERATE
		case "3":
			riskProfile = ppbc.RiskProfile_MODERATE
		case "4":
			riskProfile = ppbc.RiskProfile_MODERATE_TOHIGH
		case "5":
			riskProfile = ppbc.RiskProfile_HIGH
		default:
			riskProfile = ppbc.RiskProfile_UNKNOWN_RISK_PROFILE
		}

		switch parm.Data[i].InvestmentObjective {
		case "1":
			investmentObjective = ppbc.InvestmentObjective_GAIN_FORM_PRICE_MARGIN
		case "2":
			investmentObjective = ppbc.InvestmentObjective_INVESTMENT
		case "3":
			investmentObjective = ppbc.InvestmentObjective_SPECULATION
		case "4":
			investmentObjective = ppbc.InvestmentObjective_OBTAIN_THE_REVENUE
		case "5":
			investmentObjective = ppbc.InvestmentObjective_OTHER_INVESTMENT_OBJECTIVE
		default:
			investmentObjective = ppbc.InvestmentObjective_UNKNOWN_INVESTMENT_OBJECTIVE
		}

		switch parm.Data[i].SourceOfFund {
		case "1":
			sourceOfFund = ppbc.SourceOfFund_REVENUE
		case "2":
			sourceOfFund = ppbc.SourceOfFund_BUSINESS_PROFIT
		case "3":
			sourceOfFund = ppbc.SourceOfFund_SAVING_INTEREST
		case "4":
			sourceOfFund = ppbc.SourceOfFund_LEGACY
		case "5":
			sourceOfFund = ppbc.SourceOfFund_FUND_FROM_PARENTS_OR_CHILDREN
		case "6":
			sourceOfFund = ppbc.SourceOfFund_GRANT
		case "7":
			sourceOfFund = ppbc.SourceOfFund_FUND_FROM_HUSBAND_OR_WIFE
		case "8":
			sourceOfFund = ppbc.SourceOfFund_DRAWING
		case "9":
			sourceOfFund = ppbc.SourceOfFund_INVESTMENT_GAIN
		case "10":
			sourceOfFund = ppbc.SourceOfFund_OTHER_SOURCE_OF_FUND
		default:
			sourceOfFund = ppbc.SourceOfFund_UNKNOWN_SOURCE_OF_FUND
		}

		switch parm.Data[i].AssetOwner {
		case "1":
			assetOwner = ppbc.AssetOwner_MYSELF
		case "2":
			assetOwner = ppbc.AssetOwner_REPRESENTING_OTHER_PARTY
		default:
			assetOwner = ppbc.AssetOwner_UNKNOWN_ASSET_OWNER
		}

		switch parm.Data[i].StatementType {
		case "1":
			statementType = ppbc.StatementType_HARDCOPY
		case "2":
			statementType = ppbc.StatementType_ESTATEMENT
		default:
			statementType = ppbc.StatementType_UNKNOWN_STATEMENT_TYPE
		}

		switch parm.Data[i].Fatca {
		case "1":
			fatca = ppbc.Fatca_US_PERSON
		case "2":
			fatca = ppbc.Fatca_NON_US_PERSON
		case "3":
			fatca = ppbc.Fatca_RECALCITRANT_ACCOUNT_HOLDER_WITH_US_INDICIA
		case "4":
			fatca = ppbc.Fatca_RECALCITRANT_ACCOUNT_HOLDER_WITHOUT_US_INDICIA
		case "5":
			fatca = ppbc.Fatca_DORMANT_ACCOUNT
		default:
			fatca = ppbc.Fatca_UNKNOWN_FATCHA
		}

		switch parm.Data[i].Type {
		case "1":
			typeValue = ppbc.Type_INPUT
		case "2":
			typeValue = ppbc.Type_AMENDMENT
		default:
			typeValue = ppbc.Type_UNKNOWN_TYPE
		}

		switch parm.Data[i].InvestorType {
		case "1":
			investorType = ppbc.InvestorType_INDIVIDUAL
		case "2":
			investorType = ppbc.InvestorType_INSTITUTIONAL
		default:
			investorType = ppbc.InvestorType_UNKNOWN_INVESTOR_TYPE
		}

		data := &ppbc.PendaftaranPesertaData{
			ProcessingDate:       parm.Data[i].ProcessingDate,
			ProcessingTime:       parm.Data[i].ProcessingTime,
			Sid:                  parm.Data[i].Sid,
			SaCode:               parm.Data[i].SaCode,
			SidGenerationStatus:  sidGenerationStatus,
			NewSid:               parm.Data[i].NewSid,
			FirstName:            parm.Data[i].FirstName,
			MiddleName:           parm.Data[i].MiddleName,
			LastName:             parm.Data[i].LastName,
			Nationality:          parm.Data[i].Nationality,
			Identification:       identification,
			IdNo:                 parm.Data[i].IDNo,
			IdExpDate:            parm.Data[i].IDExpDate,
			NpwpNo:               parm.Data[i].NpwpNo,
			NpwpDate:             parm.Data[i].NpwpDate,
			BirthCountry:         parm.Data[i].BirthCountry,
			BirthPlace:           parm.Data[i].BirthPlace,
			BirthDate:            parm.Data[i].BirthDate,
			Gender:               gender,
			EducationBackground:  educationBackground,
			MotherName:           parm.Data[i].MotherName,
			Religion:             religion,
			Occupation:           occupation,
			IncomeLevel:          incomeLevel,
			MaritalStatus:        maritalStatus,
			SpouseName:           parm.Data[i].SpouseName,
			RiskProfile:          riskProfile,
			InvestmentObjective:  investmentObjective,
			SourceOfFund:         sourceOfFund,
			AssetOwner:           assetOwner,
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
			StatementType:        statementType,
			Fatca:                fatca,
			Tin:                  parm.Data[i].Tin,
			TinCountry:           parm.Data[i].TinCountry,
			ExternalCifNo:        parm.Data[i].ExternalCifNo,
			Type:                 typeValue,
			InvestorType:         investorType,
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
