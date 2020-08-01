package bkn

import (
	"context"

	"tapera/util/appcontext"

	pbc "tapera.mitraintegrasi/grpc/client/pesertabkn/v1"
)

func (s *service) GetPeserta(ctx context.Context, parm *PesertaParam) (*PesertaResponse, error) {
	logger := appcontext.Logger(ctx)
	logger.Debug().Msg("Start Get Peserta From BKN")

	// create grpc client from manager
	pbClient, err := s.pbClientMgr.Client(ctx)
	if err != nil {
		return nil, err
	}

	defer pbClient.Close()

	request := &pbc.PesertaRequest{Nip: parm.Nip}

	result, err := pbClient.GetPeserta(ctx, request)

	if err != nil {
		return nil, err
	}

	response := &PesertaResponse{
		Status:             result.Status,
		MessageCode:        result.MessageCode,
		MessageDescription: result.MessageDescription,
	}

	if result.Data == nil {
		return response, err
	}

	resposeData := &PesertaData{
		ID:                             result.Data.Id,
		NipBaru:                        result.Data.NipBaru,
		NipLama:                        result.Data.NipLama,
		Nama:                           result.Data.Nama,
		GelarDepan:                     result.Data.GelarDepan,
		GelarBelakang:                  result.Data.GelarBelakang,
		TempatLahir:                    result.Data.TempatLahir,
		TempatLahirID:                  result.Data.TempatLahirId,
		TglLahir:                       result.Data.TglLahir,
		Agama:                          result.Data.Agama,
		Email:                          result.Data.Email,
		Nik:                            result.Data.Nik,
		Alamat:                         result.Data.Alamat,
		NoHp:                           result.Data.NoHp,
		NoTelp:                         result.Data.NoTelp,
		JenisPegawaiID:                 result.Data.JenisPegawaiId,
		JenisPegawaiNama:               result.Data.JenisPegawaiNama,
		KedudukanPnsID:                 result.Data.KedudukanPnsId,
		KedudukanPnsNama:               result.Data.KedudukanPnsNama,
		StatusPegawai:                  result.Data.StatusPegawai,
		TmtPns:                         result.Data.TmtPns,
		TmtCpns:                        result.Data.TmtCpns,
		NoSeriKarpeg:                   result.Data.NoSeriKarpeg,
		TkPendidikanTerakhirID:         result.Data.TkPendidikanTerakhirId,
		TkPendidikanTerakhir:           result.Data.TkPendidikanTerakhir,
		PendidikanTerakhirID:           result.Data.PendidikanTerakhirId,
		PendidikanTerakhirNama:         result.Data.PendidikanTerakhirNama,
		TahunLulus:                     result.Data.TahunLulus,
		InstansiIndukID:                result.Data.InstansiIndukId,
		InstansiIndukNama:              result.Data.InstansiIndukNama,
		SatuanKerjaIndukID:             result.Data.SatuanKerjaIndukId,
		SatuanKerjaIndukNama:           result.Data.SatuanKerjaIndukNama,
		KanregID:                       result.Data.KanregId,
		KanregNama:                     result.Data.KanregNama,
		InstansiKerjaID:                result.Data.InstansiKerjaId,
		InstansiKerjaNama:              result.Data.InstansiKerjaNama,
		InstansiKerjaKodeCepat:         result.Data.InstansiKerjaKodeCepat,
		SatuanKerjaKerjaID:             result.Data.SatuanKerjaKerjaId,
		SatuanKerjaKerjaNama:           result.Data.SatuanKerjaKerjaNama,
		UnorID:                         result.Data.UnorId,
		UnorNama:                       result.Data.UnorNama,
		UnorIndukID:                    result.Data.UnorIndukId,
		UnorIndukNama:                  result.Data.UnorIndukNama,
		JenisJabatanID:                 result.Data.JenisJabatanId,
		JenisJabatan:                   result.Data.JenisJabatan,
		JabatanNama:                    result.Data.JabatanNama,
		JabatanStrukturalID:            result.Data.JabatanStrukturalId,
		JabatanStrukturalNama:          result.Data.JabatanStrukturalNama,
		JabatanFungsionalID:            result.Data.JabatanFungsionalId,
		JabatanFungsionalNama:          result.Data.JabatanFungsionalNama,
		JabatanFungsionalUmumID:        result.Data.JabatanFungsionalUmumId,
		JabatanFungsionalUmumNama:      result.Data.JabatanFungsionalUmumNama,
		TmtJabatan:                     result.Data.TmtJabatan,
		Eselon:                         result.Data.Eselon,
		EselonID:                       result.Data.EselonId,
		EselonLevel:                    result.Data.EselonLevel,
		TmtEselon:                      result.Data.TmtEselon,
		LokasiKerjaID:                  result.Data.LokasiKerjaId,
		LokasiKerja:                    result.Data.LokasiKerja,
		GolRuangAwal:                   result.Data.GolRuangAwal,
		GolRuangAkhirID:                result.Data.GolRuangAkhirId,
		GolRuangAkhir:                  result.Data.GolRuangAkhir,
		TmtGolAkhir:                    result.Data.TmtGolAkhir,
		GajiPokok:                      result.Data.GajiPokok,
		MasaKerja:                      result.Data.MasaKerja,
		NoSpmt:                         result.Data.NoSpmt,
		TglSpmt:                        result.Data.TglSpmt,
		JenisKawinID:                   result.Data.JenisKawinId,
		StatusPerkawinan:               result.Data.StatusPerkawinan,
		JumlahIstriSuami:               result.Data.JumlahIstriSuami,
		JumlahAnak:                     result.Data.JumlahAnak,
		NoSuratKeteranganDokter:        result.Data.NoSuratKeteranganDokter,
		TglSuratKeteranganDokter:       result.Data.TglSuratKeteranganDokter,
		NoSuratKeteranganBebasNarkoba:  result.Data.NoSuratKeteranganBebasNarkoba,
		TglSuratKeteranganBebasNarkoba: result.Data.TglSuratKeteranganBebasNarkoba,
		Skck:                           result.Data.Skck,
		TglSkck:                        result.Data.TglSkck,
		AkteKelahiran:                  result.Data.AkteKelahiran,
		StatusHidup:                    result.Data.StatusHidup,
		AkteMeninggal:                  result.Data.AkteMeninggal,
		TglMeninggal:                   result.Data.TglMeninggal,
		NoAskes:                        result.Data.NoAskes,
		NoTaspen:                       result.Data.NoTaspen,
		NoNpwp:                         result.Data.NoNpwp,
		TglNpwp:                        result.Data.TglNpwp,
		Bahasa:                         result.Data.Bahasa,
		JenisKelamin:                   result.Data.JenisKelamin,
		KodePos:                        result.Data.KodePos,
	}

	response.Data = resposeData

	logger.Debug().Msg("End Get Peserta From BKN")

	return response, err
}
