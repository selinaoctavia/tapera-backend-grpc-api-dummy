package bri

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"tapera/util/httpclient"
	"testing"
)

func TestPendaftaranPeserta(t *testing.T) {
	ctx := context.Background()

	responseSuccessMock := &PendaftaranPesertaResultSuccess{
		ResponseCode:        "00000",
		ResponseDescription: "success",
		ResponseStatus:      "true",
		Data:                "",
	}

	responseFailedMock := &PendaftaranPesertaResultSuccess{
		ResponseCode:        "10176",
		ResponseDescription: "Investor No length incorect (11!n)",
		ResponseStatus:      "false",
		Data:                "",
	}

	responseErrorMock := &PendaftaranPesertaResultError{
		Status: Status{
			Code: "0601",
			Desc: "Invalid token",
		},
	}

	resultSuccessMock := &PendaftaranPesertaResponse{
		Status:             "Success",
		MessageCode:        "00000",
		MessageDescription: "success",
	}

	resultFailedMock := &PendaftaranPesertaResponse{
		Status:             "Failed",
		MessageCode:        "10176",
		MessageDescription: "Investor No length incorect (11!n)",
	}

	resultErrorMock := &PendaftaranPesertaResponse{
		Status:             "Error",
		MessageCode:        "0601",
		MessageDescription: "Invalid token",
	}

	paramMock := []PendaftaranPesertaParam{
		PendaftaranPesertaParam{
			Data: []*PendaftaranPesertaData{
				&PendaftaranPesertaData{
					ProcessingDate:      "20200309",
					ProcessingTime:      "180500",
					SaCode:              "DBSI1",
					Sid:                 "GBU909",
					SidGenerationStatus: "1",
					NewSid:              "",
					FirstName:           "XXXX",
				},
			},
		},
	}

	paramsMock := &paramMock

	dataSuccessMock, _ := json.MarshalIndent(responseSuccessMock, "", "  ")
	dataFailedMock, _ := json.MarshalIndent(responseFailedMock, "", "  ")
	dataErrorMock, _ := json.MarshalIndent(responseErrorMock, "", "  ")

	clientSuccessMock := httpclient.NewFactoryMock(func(*http.Request) *http.Response {
		json := string(dataSuccessMock)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(json))),
			Header:     make(http.Header),
		}
	})

	clientFailedMock := httpclient.NewFactoryMock(func(*http.Request) *http.Response {
		json := string(dataFailedMock)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(json))),
			Header:     make(http.Header),
		}
	})

	clientErrorMock := httpclient.NewFactoryMock(func(*http.Request) *http.Response {
		json := string(dataErrorMock)
		return &http.Response{
			StatusCode: 401,
			Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(json))),
			Header:     make(http.Header),
		}
	})

	srv := NewService("https://sandbox.partner.api.bri.co.id", clientSuccessMock)
	var result *PendaftaranPesertaResponse
	result, err := srv.PendaftaranPeserta(ctx, "xxx6Ndws4uXlSFhYrRB1lcn5xxx", "xxHQK8srvOPBkOxx", paramsMock)

	if err != nil {
		t.Error(err)
	}

	if *resultSuccessMock != *result {
		t.Errorf("Response Success Invalid!")
	}

	srv = NewService("https://sandbox.partner.api.bri.co.id", clientFailedMock)
	result, err = srv.PendaftaranPeserta(ctx, "xxx6Ndws4uXlSFhYrRB1lcn5xxx", "xxHQK8srvOPBkOxx", paramsMock)

	if err != nil {
		t.Error(err)
	}

	if *resultFailedMock != *result {
		t.Errorf("Response Failed Invalid!")
	}

	srv = NewService("https://sandbox.partner.api.bri.co.id", clientErrorMock)
	result, err = srv.PendaftaranPeserta(ctx, "xxx6Ndws4uXlSFhYrRB1lcn5xxx", "xxHQK8srvOPBkOxx", paramsMock)

	if err != nil {
		t.Error(err)
	}

	if *resultErrorMock != *result {
		t.Errorf("Response Error Invalid!")
	}
}
