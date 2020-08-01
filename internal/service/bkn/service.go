package bkn

import (
	"context"

	pbc "tapera.mitraintegrasi/grpc/client/pesertabkn/v1"
)

type (
	// Service interface of user service
	Service interface {
		GetPeserta(ctx context.Context, parm *PesertaParam) (*PesertaResponse, error)
	}

	// Service struct
	service struct {
		pbClientMgr pbc.GrpcClientManager
	}

	//PesertaParam struct
	PesertaParam struct {
		Nip string `json:"nip"`
	}

	// PesertaData struct
	PesertaData struct {
		ID                             string `json:"id"`
		NipBaru                        string `json:"nipBaru"`
		NipLama                        string `json:"nipLama"`
		Nama                           string `json:"nama"`
		GelarDepan                     string `json:"gelarDepan"`
		GelarBelakang                  string `json:"gelarBelakang"`
		TempatLahir                    string `json:"tempatLahir"`
		TempatLahirID                  string `json:"tempatLahirId"`
		TglLahir                       string `json:"tglLahir"`
		Agama                          string `json:"agama"`
		Email                          string `json:"email"`
		Nik                            string `json:"nik"`
		Alamat                         string `json:"alamat"`
		NoHp                           string `json:"noHp"`
		NoTelp                         string `json:"noTelp"`
		JenisPegawaiID                 string `json:"jenisPegawaiId"`
		JenisPegawaiNama               string `json:"jenisPegawaiNama"`
		KedudukanPnsID                 string `json:"kedudukanPnsId"`
		KedudukanPnsNama               string `json:"kedudukanPnsNama"`
		StatusPegawai                  string `json:"statusPegawai"`
		TmtPns                         string `json:"tmtPns"`
		TmtCpns                        string `json:"tmtCpns"`
		NoSeriKarpeg                   string `json:"noSeriKarpeg"`
		TkPendidikanTerakhirID         string `json:"tkPendidikanTerakhirId"`
		TkPendidikanTerakhir           string `json:"tkPendidikanTerakhir"`
		PendidikanTerakhirID           string `json:"pendidikanTerakhirId"`
		PendidikanTerakhirNama         string `json:"pendidikanTerakhirNama"`
		TahunLulus                     string `json:"tahunLulus"`
		InstansiIndukID                string `json:"instansiIndukId"`
		InstansiIndukNama              string `json:"instansiIndukNama"`
		SatuanKerjaIndukID             string `json:"satuanKerjaIndukId"`
		SatuanKerjaIndukNama           string `json:"satuanKerjaIndukNama"`
		KanregID                       string `json:"kanregId"`
		KanregNama                     string `json:"kanregNama"`
		InstansiKerjaID                string `json:"instansiKerjaId"`
		InstansiKerjaNama              string `json:"instansiKerjaNama"`
		InstansiKerjaKodeCepat         string `json:"instansiKerjaKodeCepat"`
		SatuanKerjaKerjaID             string `json:"satuanKerjaKerjaId"`
		SatuanKerjaKerjaNama           string `json:"satuanKerjaKerjaNama"`
		UnorID                         string `json:"unorId"`
		UnorNama                       string `json:"unorNama"`
		UnorIndukID                    string `json:"unorIndukId"`
		UnorIndukNama                  string `json:"unorIndukNama"`
		JenisJabatanID                 string `json:"jenisJabatanId"`
		JenisJabatan                   string `json:"jenisJabatan"`
		JabatanNama                    string `json:"jabatanNama"`
		JabatanStrukturalID            string `json:"jabatanStrukturalId"`
		JabatanStrukturalNama          string `json:"jabatanStrukturalNama"`
		JabatanFungsionalID            string `json:"jabatanFungsionalId"`
		JabatanFungsionalNama          string `json:"jabatanFungsionalNama"`
		JabatanFungsionalUmumID        string `json:"jabatanFungsionalUmumId"`
		JabatanFungsionalUmumNama      string `json:"jabatanFungsionalUmumNama"`
		TmtJabatan                     string `json:"tmtJabatan"`
		Eselon                         string `json:"eselon"`
		EselonID                       string `json:"eselonId"`
		EselonLevel                    string `json:"eselonLevel"`
		TmtEselon                      string `json:"tmtEselon"`
		LokasiKerjaID                  string `json:"lokasiKerjaId"` //Tidak ada di document BKN tp ada di response trial hit API
		LokasiKerja                    string `json:"lokasiKerja"`
		GolRuangAwal                   string `json:"golRuangAwal"`
		GolRuangAkhirID                string `json:"golRuangAkhirId"`
		GolRuangAkhir                  string `json:"golRuangAkhir"`
		TmtGolAkhir                    string `json:"tmtGolAkhir"`
		GajiPokok                      string `json:"gajiPokok"`
		MasaKerja                      string `json:"masaKerja"`
		NoSpmt                         string `json:"noSpmt"`
		TglSpmt                        string `json:"tglSpmt"`
		JenisKawinID                   string `json:"jenisKawinId"` //Tidak ada di document BKN tp ada di response trial hit API
		StatusPerkawinan               string `json:"statusPerkawinan"`
		JumlahIstriSuami               string `json:"jumlahIstriSuami"`
		JumlahAnak                     string `json:"jumlahAnak"`
		NoSuratKeteranganDokter        string `json:"noSuratKeteranganDokter"`
		TglSuratKeteranganDokter       string `json:"tglSuratKeteranganDokter"`
		NoSuratKeteranganBebasNarkoba  string `json:"noSuratKeteranganBebasNarkoba"`
		TglSuratKeteranganBebasNarkoba string `json:"tglSuratKeteranganBebasNarkoba"`
		Skck                           string `json:"skck"`
		TglSkck                        string `json:"tglSkck"`
		AkteKelahiran                  string `json:"akteKelahiran"`
		StatusHidup                    string `json:"statusHidup"`
		AkteMeninggal                  string `json:"akteMeninggal"`
		TglMeninggal                   string `json:"tglMeninggal"`
		NoAskes                        string `json:"noAskes"`
		NoTaspen                       string `json:"noTaspen"`
		NoNpwp                         string `json:"noNpwp"`
		TglNpwp                        string `json:"tglNpwp"`
		Bahasa                         string `json:"bahasa"`
		JenisKelamin                   string `json:"jenisKelamin"`
		KodePos                        string `json:"kodePos"`
	}

	// PesertaResponse struct
	PesertaResponse struct {
		Status             string       `json:"status"`
		MessageCode        string       `json:"message_code"`
		MessageDescription string       `json:"message_description"`
		Data               *PesertaData `json:"data"`
	}
)

// NewService func
func NewService(pbClientMgr pbc.GrpcClientManager) Service {
	return &service{pbClientMgr: pbClientMgr}
}
