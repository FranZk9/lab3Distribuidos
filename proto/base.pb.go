// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: proto/base.proto

package proto

import (
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

type Mensaje struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
}

func (x *Mensaje) Reset() {
	*x = Mensaje{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mensaje) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mensaje) ProtoMessage() {}

func (x *Mensaje) ProtoReflect() protoreflect.Message {
	mi := &file_proto_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mensaje.ProtoReflect.Descriptor instead.
func (*Mensaje) Descriptor() ([]byte, []int) {
	return file_proto_base_proto_rawDescGZIP(), []int{0}
}

func (x *Mensaje) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

type Crearmensaje struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mensaje *Mensaje `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
}

func (x *Crearmensaje) Reset() {
	*x = Crearmensaje{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Crearmensaje) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Crearmensaje) ProtoMessage() {}

func (x *Crearmensaje) ProtoReflect() protoreflect.Message {
	mi := &file_proto_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Crearmensaje.ProtoReflect.Descriptor instead.
func (*Crearmensaje) Descriptor() ([]byte, []int) {
	return file_proto_base_proto_rawDescGZIP(), []int{1}
}

func (x *Crearmensaje) GetMensaje() *Mensaje {
	if x != nil {
		return x.Mensaje
	}
	return nil
}

type Respuestamensaje struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mensajeid string `protobuf:"bytes,1,opt,name=mensajeid,proto3" json:"mensajeid,omitempty"`
}

func (x *Respuestamensaje) Reset() {
	*x = Respuestamensaje{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Respuestamensaje) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Respuestamensaje) ProtoMessage() {}

func (x *Respuestamensaje) ProtoReflect() protoreflect.Message {
	mi := &file_proto_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Respuestamensaje.ProtoReflect.Descriptor instead.
func (*Respuestamensaje) Descriptor() ([]byte, []int) {
	return file_proto_base_proto_rawDescGZIP(), []int{2}
}

func (x *Respuestamensaje) GetMensajeid() string {
	if x != nil {
		return x.Mensajeid
	}
	return ""
}

type Estado struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
}

func (x *Estado) Reset() {
	*x = Estado{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_base_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Estado) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Estado) ProtoMessage() {}

func (x *Estado) ProtoReflect() protoreflect.Message {
	mi := &file_proto_base_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Estado.ProtoReflect.Descriptor instead.
func (*Estado) Descriptor() ([]byte, []int) {
	return file_proto_base_proto_rawDescGZIP(), []int{3}
}

func (x *Estado) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

type ConsultarLista struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Estado *Estado `protobuf:"bytes,1,opt,name=estado,proto3" json:"estado,omitempty"`
}

func (x *ConsultarLista) Reset() {
	*x = ConsultarLista{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_base_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsultarLista) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsultarLista) ProtoMessage() {}

func (x *ConsultarLista) ProtoReflect() protoreflect.Message {
	mi := &file_proto_base_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsultarLista.ProtoReflect.Descriptor instead.
func (*ConsultarLista) Descriptor() ([]byte, []int) {
	return file_proto_base_proto_rawDescGZIP(), []int{4}
}

func (x *ConsultarLista) GetEstado() *Estado {
	if x != nil {
		return x.Estado
	}
	return nil
}

type RespuestaLista struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Estadoid []string `protobuf:"bytes,1,rep,name=estadoid,proto3" json:"estadoid,omitempty"`
}

func (x *RespuestaLista) Reset() {
	*x = RespuestaLista{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_base_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaLista) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaLista) ProtoMessage() {}

func (x *RespuestaLista) ProtoReflect() protoreflect.Message {
	mi := &file_proto_base_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaLista.ProtoReflect.Descriptor instead.
func (*RespuestaLista) Descriptor() ([]byte, []int) {
	return file_proto_base_proto_rawDescGZIP(), []int{5}
}

func (x *RespuestaLista) GetEstadoid() []string {
	if x != nil {
		return x.Estadoid
	}
	return nil
}

var File_proto_base_proto protoreflect.FileDescriptor

var file_proto_base_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x21, 0x0a, 0x07, 0x4d, 0x65, 0x6e, 0x73,
	0x61, 0x6a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x22, 0x37, 0x0a, 0x0c, 0x43,
	0x72, 0x65, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x6d,
	0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x6e,
	0x73, 0x61, 0x6a, 0x65, 0x22, 0x30, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x6e, 0x73,
	0x61, 0x6a, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x6e,
	0x73, 0x61, 0x6a, 0x65, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x06, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x22, 0x36, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x06, 0x65, 0x73,
	0x74, 0x61, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x52, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f,
	0x22, 0x2c, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x4c, 0x69, 0x73,
	0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x69, 0x64, 0x32, 0x81,
	0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x34, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x1a,
	0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x12, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x1a, 0x14, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x4c, 0x69, 0x73,
	0x74, 0x61, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x53, 0x69, 0x73, 0x74, 0x65, 0x6d, 0x61, 0x73, 0x2d, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x69, 0x64, 0x6f, 0x73, 0x2d, 0x32, 0x30, 0x32, 0x33, 0x2d, 0x30, 0x32, 0x2f, 0x47,
	0x72, 0x75, 0x70, 0x6f, 0x32, 0x32, 0x2d, 0x4c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x69, 0x6f, 0x2d, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_base_proto_rawDescOnce sync.Once
	file_proto_base_proto_rawDescData = file_proto_base_proto_rawDesc
)

func file_proto_base_proto_rawDescGZIP() []byte {
	file_proto_base_proto_rawDescOnce.Do(func() {
		file_proto_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_base_proto_rawDescData)
	})
	return file_proto_base_proto_rawDescData
}

var file_proto_base_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_base_proto_goTypes = []interface{}{
	(*Mensaje)(nil),          // 0: grpc.Mensaje
	(*Crearmensaje)(nil),     // 1: grpc.Crearmensaje
	(*Respuestamensaje)(nil), // 2: grpc.Respuestamensaje
	(*Estado)(nil),           // 3: grpc.Estado
	(*ConsultarLista)(nil),   // 4: grpc.ConsultarLista
	(*RespuestaLista)(nil),   // 5: grpc.RespuestaLista
}
var file_proto_base_proto_depIdxs = []int32{
	0, // 0: grpc.Crearmensaje.mensaje:type_name -> grpc.Mensaje
	3, // 1: grpc.ConsultarLista.estado:type_name -> grpc.Estado
	1, // 2: grpc.MensajeService.Create:input_type -> grpc.Crearmensaje
	4, // 3: grpc.MensajeService.CreateLista:input_type -> grpc.ConsultarLista
	2, // 4: grpc.MensajeService.Create:output_type -> grpc.Respuestamensaje
	5, // 5: grpc.MensajeService.CreateLista:output_type -> grpc.RespuestaLista
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_base_proto_init() }
func file_proto_base_proto_init() {
	if File_proto_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mensaje); i {
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
		file_proto_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Crearmensaje); i {
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
		file_proto_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Respuestamensaje); i {
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
		file_proto_base_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Estado); i {
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
		file_proto_base_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsultarLista); i {
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
		file_proto_base_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaLista); i {
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
			RawDescriptor: file_proto_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_base_proto_goTypes,
		DependencyIndexes: file_proto_base_proto_depIdxs,
		MessageInfos:      file_proto_base_proto_msgTypes,
	}.Build()
	File_proto_base_proto = out.File
	file_proto_base_proto_rawDesc = nil
	file_proto_base_proto_goTypes = nil
	file_proto_base_proto_depIdxs = nil
}
