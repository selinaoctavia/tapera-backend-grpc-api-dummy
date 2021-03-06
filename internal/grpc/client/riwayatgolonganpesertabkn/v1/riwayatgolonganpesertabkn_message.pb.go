// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: riwayatgolonganpesertabkn_message.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RiwayatGolonganPesertaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IdPns                  string `protobuf:"bytes,2,opt,name=id_pns,json=idPns,proto3" json:"id_pns,omitempty"`
	GolonganId             string `protobuf:"bytes,3,opt,name=golongan_id,json=golonganId,proto3" json:"golongan_id,omitempty"`
	Golongan               string `protobuf:"bytes,4,opt,name=golongan,proto3" json:"golongan,omitempty"`
	SkNomor                string `protobuf:"bytes,5,opt,name=sk_nomor,json=skNomor,proto3" json:"sk_nomor,omitempty"`
	SkTanggal              string `protobuf:"bytes,6,opt,name=sk_tanggal,json=skTanggal,proto3" json:"sk_tanggal,omitempty"`
	TmtGolongan            string `protobuf:"bytes,7,opt,name=tmt_golongan,json=tmtGolongan,proto3" json:"tmt_golongan,omitempty"`
	NoPertekBkn            string `protobuf:"bytes,8,opt,name=no_pertek_bkn,json=noPertekBkn,proto3" json:"no_pertek_bkn,omitempty"`
	TglPertekBkn           string `protobuf:"bytes,9,opt,name=tgl_pertek_bkn,json=tglPertekBkn,proto3" json:"tgl_pertek_bkn,omitempty"`
	JumlahKreditUtama      string `protobuf:"bytes,10,opt,name=jumlah_kredit_utama,json=jumlahKreditUtama,proto3" json:"jumlah_kredit_utama,omitempty"`
	JumlahKreditTambahan   string `protobuf:"bytes,11,opt,name=jumlah_kredit_tambahan,json=jumlahKreditTambahan,proto3" json:"jumlah_kredit_tambahan,omitempty"`
	JenisKpId              string `protobuf:"bytes,12,opt,name=jenis_kp_id,json=jenisKpId,proto3" json:"jenis_kp_id,omitempty"`
	JenisKpNama            string `protobuf:"bytes,13,opt,name=jenis_kp_nama,json=jenisKpNama,proto3" json:"jenis_kp_nama,omitempty"`
	MasaKerjaGolonganTahun string `protobuf:"bytes,14,opt,name=masa_kerja_golongan_tahun,json=masaKerjaGolonganTahun,proto3" json:"masa_kerja_golongan_tahun,omitempty"`
	MasaKerjaGolonganBulan string `protobuf:"bytes,15,opt,name=masa_kerja_golongan_bulan,json=masaKerjaGolonganBulan,proto3" json:"masa_kerja_golongan_bulan,omitempty"`
}

func (x *RiwayatGolonganPesertaData) Reset() {
	*x = RiwayatGolonganPesertaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_riwayatgolonganpesertabkn_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RiwayatGolonganPesertaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RiwayatGolonganPesertaData) ProtoMessage() {}

func (x *RiwayatGolonganPesertaData) ProtoReflect() protoreflect.Message {
	mi := &file_riwayatgolonganpesertabkn_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RiwayatGolonganPesertaData.ProtoReflect.Descriptor instead.
func (*RiwayatGolonganPesertaData) Descriptor() ([]byte, []int) {
	return file_riwayatgolonganpesertabkn_message_proto_rawDescGZIP(), []int{0}
}

func (x *RiwayatGolonganPesertaData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RiwayatGolonganPesertaData) GetIdPns() string {
	if x != nil {
		return x.IdPns
	}
	return ""
}

func (x *RiwayatGolonganPesertaData) GetGolonganId() string {
	if x != nil {
		return x.GolonganId
	}
	return ""
}

func (x *RiwayatGolonganPesertaData) GetGolongan() string {
	if x != nil {
		return x.Golongan
	}
	return ""
}

func (x *RiwayatGolonganPesertaData) GetSkNomor() string {
	if x != nil {
		return x.SkNomor
	}
	return ""
}

func (x *RiwayatGolonganPesertaData) GetSkTanggal() string {
	if x != nil {
		return x.SkTanggal
	}
	return ""
}

func (x *RiwayatGolonganPesertaData) GetTmtGolongan() string {
	if x != nil {
		return x.TmtGolongan
	}
	return ""
}

func (x *RiwayatGolonganPesertaData) GetNoPertekBkn() string {
	if x != nil {
		return x.NoPertekBkn
	}
	return ""
}

func (x *RiwayatGolonganPesertaData) GetTglPertekBkn() string {
	if x != nil {
		return x.TglPertekBkn
	}
	return ""
}

func (x *RiwayatGolonganPesertaData) GetJumlahKreditUtama() string {
	if x != nil {
		return x.JumlahKreditUtama
	}
	return ""
}

func (x *RiwayatGolonganPesertaData) GetJumlahKreditTambahan() string {
	if x != nil {
		return x.JumlahKreditTambahan
	}
	return ""
}

func (x *RiwayatGolonganPesertaData) GetJenisKpId() string {
	if x != nil {
		return x.JenisKpId
	}
	return ""
}

func (x *RiwayatGolonganPesertaData) GetJenisKpNama() string {
	if x != nil {
		return x.JenisKpNama
	}
	return ""
}

func (x *RiwayatGolonganPesertaData) GetMasaKerjaGolonganTahun() string {
	if x != nil {
		return x.MasaKerjaGolonganTahun
	}
	return ""
}

func (x *RiwayatGolonganPesertaData) GetMasaKerjaGolonganBulan() string {
	if x != nil {
		return x.MasaKerjaGolonganBulan
	}
	return ""
}

var File_riwayatgolonganpesertabkn_message_proto protoreflect.FileDescriptor

var file_riwayatgolonganpesertabkn_message_proto_rawDesc = []byte{
	0x0a, 0x27, 0x72, 0x69, 0x77, 0x61, 0x79, 0x61, 0x74, 0x67, 0x6f, 0x6c, 0x6f, 0x6e, 0x67, 0x61,
	0x6e, 0x70, 0x65, 0x73, 0x65, 0x72, 0x74, 0x61, 0x62, 0x6b, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x72, 0x69, 0x77, 0x61, 0x79,
	0x61, 0x74, 0x67, 0x6f, 0x6c, 0x6f, 0x6e, 0x67, 0x61, 0x6e, 0x70, 0x65, 0x73, 0x65, 0x72, 0x74,
	0x61, 0x62, 0x6b, 0x6e, 0x22, 0xc7, 0x04, 0x0a, 0x1a, 0x52, 0x69, 0x77, 0x61, 0x79, 0x61, 0x74,
	0x47, 0x6f, 0x6c, 0x6f, 0x6e, 0x67, 0x61, 0x6e, 0x50, 0x65, 0x73, 0x65, 0x72, 0x74, 0x61, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x64, 0x5f, 0x70, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x64, 0x50, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x6f,
	0x6c, 0x6f, 0x6e, 0x67, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x67, 0x6f, 0x6c, 0x6f, 0x6e, 0x67, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x67,
	0x6f, 0x6c, 0x6f, 0x6e, 0x67, 0x61, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67,
	0x6f, 0x6c, 0x6f, 0x6e, 0x67, 0x61, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x5f, 0x6e, 0x6f,
	0x6d, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6b, 0x4e, 0x6f, 0x6d,
	0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6b, 0x5f, 0x74, 0x61, 0x6e, 0x67, 0x67, 0x61, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6b, 0x54, 0x61, 0x6e, 0x67, 0x67, 0x61,
	0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6d, 0x74, 0x5f, 0x67, 0x6f, 0x6c, 0x6f, 0x6e, 0x67, 0x61,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6d, 0x74, 0x47, 0x6f, 0x6c, 0x6f,
	0x6e, 0x67, 0x61, 0x6e, 0x12, 0x22, 0x0a, 0x0d, 0x6e, 0x6f, 0x5f, 0x70, 0x65, 0x72, 0x74, 0x65,
	0x6b, 0x5f, 0x62, 0x6b, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x6f, 0x50,
	0x65, 0x72, 0x74, 0x65, 0x6b, 0x42, 0x6b, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x67, 0x6c, 0x5f,
	0x70, 0x65, 0x72, 0x74, 0x65, 0x6b, 0x5f, 0x62, 0x6b, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x74, 0x67, 0x6c, 0x50, 0x65, 0x72, 0x74, 0x65, 0x6b, 0x42, 0x6b, 0x6e, 0x12, 0x2e,
	0x0a, 0x13, 0x6a, 0x75, 0x6d, 0x6c, 0x61, 0x68, 0x5f, 0x6b, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f,
	0x75, 0x74, 0x61, 0x6d, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6a, 0x75, 0x6d,
	0x6c, 0x61, 0x68, 0x4b, 0x72, 0x65, 0x64, 0x69, 0x74, 0x55, 0x74, 0x61, 0x6d, 0x61, 0x12, 0x34,
	0x0a, 0x16, 0x6a, 0x75, 0x6d, 0x6c, 0x61, 0x68, 0x5f, 0x6b, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f,
	0x74, 0x61, 0x6d, 0x62, 0x61, 0x68, 0x61, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x6a, 0x75, 0x6d, 0x6c, 0x61, 0x68, 0x4b, 0x72, 0x65, 0x64, 0x69, 0x74, 0x54, 0x61, 0x6d, 0x62,
	0x61, 0x68, 0x61, 0x6e, 0x12, 0x1e, 0x0a, 0x0b, 0x6a, 0x65, 0x6e, 0x69, 0x73, 0x5f, 0x6b, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x65, 0x6e, 0x69, 0x73,
	0x4b, 0x70, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x6a, 0x65, 0x6e, 0x69, 0x73, 0x5f, 0x6b, 0x70,
	0x5f, 0x6e, 0x61, 0x6d, 0x61, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6a, 0x65, 0x6e,
	0x69, 0x73, 0x4b, 0x70, 0x4e, 0x61, 0x6d, 0x61, 0x12, 0x39, 0x0a, 0x19, 0x6d, 0x61, 0x73, 0x61,
	0x5f, 0x6b, 0x65, 0x72, 0x6a, 0x61, 0x5f, 0x67, 0x6f, 0x6c, 0x6f, 0x6e, 0x67, 0x61, 0x6e, 0x5f,
	0x74, 0x61, 0x68, 0x75, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x6d, 0x61, 0x73,
	0x61, 0x4b, 0x65, 0x72, 0x6a, 0x61, 0x47, 0x6f, 0x6c, 0x6f, 0x6e, 0x67, 0x61, 0x6e, 0x54, 0x61,
	0x68, 0x75, 0x6e, 0x12, 0x39, 0x0a, 0x19, 0x6d, 0x61, 0x73, 0x61, 0x5f, 0x6b, 0x65, 0x72, 0x6a,
	0x61, 0x5f, 0x67, 0x6f, 0x6c, 0x6f, 0x6e, 0x67, 0x61, 0x6e, 0x5f, 0x62, 0x75, 0x6c, 0x61, 0x6e,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x6d, 0x61, 0x73, 0x61, 0x4b, 0x65, 0x72, 0x6a,
	0x61, 0x47, 0x6f, 0x6c, 0x6f, 0x6e, 0x67, 0x61, 0x6e, 0x42, 0x75, 0x6c, 0x61, 0x6e, 0x42, 0x04,
	0x5a, 0x02, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_riwayatgolonganpesertabkn_message_proto_rawDescOnce sync.Once
	file_riwayatgolonganpesertabkn_message_proto_rawDescData = file_riwayatgolonganpesertabkn_message_proto_rawDesc
)

func file_riwayatgolonganpesertabkn_message_proto_rawDescGZIP() []byte {
	file_riwayatgolonganpesertabkn_message_proto_rawDescOnce.Do(func() {
		file_riwayatgolonganpesertabkn_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_riwayatgolonganpesertabkn_message_proto_rawDescData)
	})
	return file_riwayatgolonganpesertabkn_message_proto_rawDescData
}

var file_riwayatgolonganpesertabkn_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_riwayatgolonganpesertabkn_message_proto_goTypes = []interface{}{
	(*RiwayatGolonganPesertaData)(nil), // 0: riwayatgolonganpesertabkn.RiwayatGolonganPesertaData
}
var file_riwayatgolonganpesertabkn_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_riwayatgolonganpesertabkn_message_proto_init() }
func file_riwayatgolonganpesertabkn_message_proto_init() {
	if File_riwayatgolonganpesertabkn_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_riwayatgolonganpesertabkn_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RiwayatGolonganPesertaData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_riwayatgolonganpesertabkn_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_riwayatgolonganpesertabkn_message_proto_goTypes,
		DependencyIndexes: file_riwayatgolonganpesertabkn_message_proto_depIdxs,
		MessageInfos:      file_riwayatgolonganpesertabkn_message_proto_msgTypes,
	}.Build()
	File_riwayatgolonganpesertabkn_message_proto = out.File
	file_riwayatgolonganpesertabkn_message_proto_rawDesc = nil
	file_riwayatgolonganpesertabkn_message_proto_goTypes = nil
	file_riwayatgolonganpesertabkn_message_proto_depIdxs = nil
}
