// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/service/oidb/OidbSvcTrpcTcp0xE37_800.proto

package oidb

// Get Offline File Download
type OidbSvcTrpcTcp0XE37_800 struct {
	SubCommand uint32                       `protobuf:"varint,1,opt"`
	Field2     int32                        `protobuf:"varint,2,opt"`
	Body       *OidbSvcTrpcTcp0XE37_800Body `protobuf:"bytes,10,opt"`
	Field101   int32                        `protobuf:"varint,101,opt"`
	Field102   int32                        `protobuf:"varint,102,opt"`
	Field200   int32                        `protobuf:"varint,200,opt"`
	_          [0]func()
}

type OidbSvcTrpcTcp0XE37_800Body struct {
	SenderUid   string `protobuf:"bytes,10,opt"`
	ReceiverUid string `protobuf:"bytes,20,opt"`
	FileUuid    string `protobuf:"bytes,30,opt"`
	FileHash    string `protobuf:"bytes,40,opt"`
	_           [0]func()
}

type OidbSvcTrpcTcp0XE37Response struct {
	Command                  uint32             `protobuf:"varint,1,opt"`
	Seq                      int32              `protobuf:"varint,2,opt"`
	Upload                   *ApplyUploadRespV3 `protobuf:"bytes,19,opt"`
	BusinessId               int32              `protobuf:"varint,101,opt"`
	ClientType               int32              `protobuf:"varint,102,opt"`
	FlagSupportMediaPlatform int32              `protobuf:"varint,200,opt"`
	_                        [0]func()
}

type ApplyUploadRespV3 struct {
	RetCode                int32    `protobuf:"varint,10,opt"`
	RetMsg                 string   `protobuf:"bytes,20,opt"`
	TotalSpace             int64    `protobuf:"varint,30,opt"`
	UsedSpace              int64    `protobuf:"varint,40,opt"`
	UploadedSize           int64    `protobuf:"varint,50,opt"`
	UploadIp               string   `protobuf:"bytes,60,opt"`
	UploadDomain           string   `protobuf:"bytes,70,opt"`
	UploadPort             uint32   `protobuf:"varint,80,opt"`
	Uuid                   string   `protobuf:"bytes,90,opt"`
	UploadKey              []byte   `protobuf:"bytes,100,opt"`
	BoolFileExist          bool     `protobuf:"varint,110,opt"`
	PackSize               int32    `protobuf:"varint,120,opt"`
	UploadIpList           []string `protobuf:"bytes,130,rep"`
	UploadHttpsPort        int32    `protobuf:"varint,140,opt"`
	UploadHttpsDomain      string   `protobuf:"bytes,150,opt"`
	UploadDns              string   `protobuf:"bytes,160,opt"`
	UploadLanip            string   `protobuf:"bytes,170,opt"`
	FileAddon              string   `protobuf:"bytes,200,opt"`
	MediaPlatformUploadKey []byte   `protobuf:"bytes,220,opt"`
}
