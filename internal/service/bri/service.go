package bri

import (
	"context"

	mic "tapera.mitraintegrasi/grpc/client/mitraintegrasi/v1"
)

type (
	// Service interface of user service
	Service interface {
		PendaftaranPeserta(ctx context.Context, parm *PendaftaranPesertaParam) (*PendaftaranPesertaResponse, error)
	}

	// Service struct
	service struct {
		miClientMgr mic.GrpcClientManager
	}

	// PendaftaranPesertaParam struct
	PendaftaranPesertaParam struct {
		Data []*PendaftaranPesertaData `json:"data"`
	}

	// PendaftaranPesertaData struct
	PendaftaranPesertaData struct {
		ProcessingDate       string `json:"processing_date"`
		ProcessingTime       string `json:"processing_time"`
		SaCode               string `json:"sa_code"`
		Sid                  string `json:"sid"`
		SidGenerationStatus  string `json:"sid_generation_status"`
		NewSid               string `json:"new_sid"`
		FirstName            string `json:"first_name"`
		MiddleName           string `json:"middle_name"`
		LastName             string `json:"last_name"`
		Nationality          string `json:"nationality"`
		Identification       string `json:"identification"`
		IDNo                 string `json:"id_no"`
		IDExpDate            string `json:"id_exp_date"`
		NpwpNo               string `json:"npwp_no"`
		NpwpDate             string `json:"npwp_date"`
		BirthCountry         string `json:"birth_country"`
		BirthPlace           string `json:"birth_place"`
		BirthDate            string `json:"birth_date"`
		Gender               string `json:"gender"`
		EducationBackground  string `json:"education_background"`
		MotherName           string `json:"mother_name"`
		Religion             string `json:"religion"`
		Occupation           string `json:"occupation"`
		IncomeLevel          string `json:"income_level"`
		MaritalStatus        string `json:"marital_status"`
		SpouseName           string `json:"spouse_name"`
		RiskProfile          string `json:"risk_profile"`
		InvestmentObjective  string `json:"investment_objective"`
		SourceOfFund         string `json:"source_of_fund"`
		AssetOwner           string `json:"asset_owner"`
		KtpAddress           string `json:"ktp_address"`
		KtpCityCode          string `json:"ktp_city_code"`
		KtpPostal            string `json:"ktp_postal"`
		CorrAddress          string `json:"corr_address"`
		CorrCityCode         string `json:"corr_city_code"`
		CorrCityName         string `json:"corr_city_name"`
		CorrPostal           string `json:"corr_postal"`
		CorrCountry          string `json:"corr_country"`
		DomAddress           string `json:"dom_address"`
		DomCityCode          string `json:"dom_city_code"`
		DomCityName          string `json:"dom_city_name"`
		DomPostal            string `json:"dom_postal"`
		DomCountry           string `json:"dom_country"`
		HomePhone            string `json:"home_phone"`
		MobilePhone          string `json:"mobile_phone"`
		Facsimile            string `json:"facsimile"`
		Email                string `json:"email"`
		StatementType        string `json:"statement_type"`
		Fatca                string `json:"fatca"`
		Tin                  string `json:"tin"`
		TinCountry           string `json:"tin_country"`
		ExternalCifNo        string `json:"external_cif_no"`
		Type                 string `json:"type"`
		InvestorType         string `json:"investor_type"`
		RedPayBankBic1       string `json:"red_pay_bank_bic_1"`
		RedPayBankBi1        string `json:"red_pay_bank_bi_1"`
		RedPayBankName1      string `json:"red_pay_bank_name_1"`
		RedPayBankCountry1   string `json:"red_pay_bank_country_1"`
		RedPayBankBranch1    string `json:"red_pay_bank_branch_1"`
		RedPayBankAccCcy1    string `json:"red_pay_bank_acc_ccy_1"`
		REDMPaymentACNo1     string `json:"REDM_Payment_ACNo_1"`
		RedPayBankAccName1   string `json:"red_pay_bank_acc_name_1"`
		RedPayBankBic2       string `json:"red_pay_bank_bic_2"`
		RedPayBankBi2        string `json:"red_pay_bank_bi_2"`
		RedPayBankName2      string `json:"red_pay_bank_name_2"`
		RedPayBankCountry2   string `json:"red_pay_bank_country_2"`
		RedPayBankBranch2    string `json:"red_pay_bank_branch_2"`
		RedPayBankAccCcy2    string `json:"red_pay_bank_acc_ccy_2"`
		RedPayBankAccNo2     string `json:"red_pay_bank_acc_no_2"`
		RedPayBankAccName2   string `json:"red_pay_bank_acc_name_2"`
		RedPayBankBic3       string `json:"red_pay_bank_bic_3"`
		RedPayBankBi3        string `json:"red_pay_bank_bi_3"`
		RedPayBankName3      string `json:"red_pay_bank_name_3"`
		RedPayBankCountry3   string `json:"red_pay_bank_country_3"`
		RedPayBankBranch3    string `json:"red_pay_bank_branch_3"`
		RedPayBankAccCcy3    string `json:"red_pay_bank_acc_ccy_3"`
		RedPayBankAccNo3     string `json:"red_pay_bank_acc_no_3"`
		RedPayBankAccName3   string `json:"red_pay_bank_acc_name_3"`
		ClientCode           string `json:"client_code"`
		InvestorFundUnitAcNo string `json:"investor_fund_unit_ac_no"`
		InvestorNumber       string `json:"investor_number"`
	}

	// PendaftaranPesertaResponse struct
	PendaftaranPesertaResponse struct {
		Status             string `json:"status"`
		MessageCode        string `json:"message_code"`
		MessageDescription string `json:"message_description"`
	}
)

// NewService func
func NewService(miClientMgr mic.GrpcClientManager) Service {
	return &service{miClientMgr: miClientMgr}
}
