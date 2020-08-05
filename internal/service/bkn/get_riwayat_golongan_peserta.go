package bkn

import (
	"context"

	"tapera/util/appcontext"

	rgpbc "tapera.mitraintegrasi/grpc/client/riwayatgolonganpesertabkn/v1"
)

func (s *service) GetRiwayatGolonganPeserta(ctx context.Context, parm *RiwayatGolonganPesertaParam) (*RiwayatGolonganPesertaResponse, error) {
	logger := appcontext.Logger(ctx)
	logger.Debug().Msg("Start Get Riwayat Golongan Peserta From BKN")

	// create grpc client from manager
	rgpbClient, err := s.rgpbClientMgr.Client(ctx)
	if err != nil {
		return nil, err
	}

	defer rgpbClient.Close()

	request := &rgpbc.RiwayatGolonganPesertaRequest{Nip: parm.Nip}

	result, err := rgpbClient.GetRiwayatGolonganPeserta(ctx, request)

	if err != nil {
		return nil, err
	}

	response := &RiwayatGolonganPesertaResponse{
		Status:             result.Status,
		MessageCode:        result.MessageCode,
		MessageDescription: result.MessageDescription,
	}

	if result.Data == nil {
		return response, err
	}

	dataLenght := len(result.Data)
	var datas []*RiwayatGolonganPesertaData
	for l := 0; l < dataLenght; l++ {
		data := &RiwayatGolonganPesertaData{
			ID:                     result.Data[l].Id,
			IDPns:                  result.Data[l].IdPns,
			GolonganID:             result.Data[l].GolonganId,
			Golongan:               result.Data[l].Golongan,
			SkNomor:                result.Data[l].SkNomor,
			SkTanggal:              result.Data[l].SkTanggal,
			TmtGolongan:            result.Data[l].TmtGolongan,
			NoPertekBkn:            result.Data[l].NoPertekBkn,
			TglPertekBkn:           result.Data[l].TglPertekBkn,
			JumlahKreditUtama:      result.Data[l].JumlahKreditUtama,
			JumlahKreditTambahan:   result.Data[l].JumlahKreditTambahan,
			JenisKPId:              result.Data[l].JenisKpId,
			JenisKPNama:            result.Data[l].JenisKpNama,
			MasaKerjaGolonganTahun: result.Data[l].MasaKerjaGolonganTahun,
			MasaKerjaGolonganBulan: result.Data[l].MasaKerjaGolonganBulan,
		}
		datas = append(datas, data)
	}
	response.Data = datas

	logger.Debug().Msg("End Get Riwayat Golongan Peserta From BKN")

	return response, err
}
