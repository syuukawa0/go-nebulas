// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

/*
Package nebletpb is a generated protocol buffer package.

It is generated from these files:
	config.proto

It has these top-level messages:
	Config
	NetworkConfig
	ChainConfig
	RPCConfig
	AppConfig
	MiscConfig
	StatsConfig
	InfluxdbConfig
*/
package nebletpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Reporting modules.
type StatsConfig_ReportingModule int32

const (
	StatsConfig_Influxdb StatsConfig_ReportingModule = 0
)

var StatsConfig_ReportingModule_name = map[int32]string{
	0: "Influxdb",
}
var StatsConfig_ReportingModule_value = map[string]int32{
	"Influxdb": 0,
}

func (x StatsConfig_ReportingModule) String() string {
	return proto.EnumName(StatsConfig_ReportingModule_name, int32(x))
}
func (StatsConfig_ReportingModule) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorConfig, []int{6, 0}
}

// Neblet global configurations.
type Config struct {
	// Network config.
	Network *NetworkConfig `protobuf:"bytes,1,opt,name=network" json:"network,omitempty"`
	// Chain config.
	Chain *ChainConfig `protobuf:"bytes,2,opt,name=chain" json:"chain,omitempty"`
	// RPC config.
	Rpc *RPCConfig `protobuf:"bytes,3,opt,name=rpc" json:"rpc,omitempty"`
	// Stats config.
	Stats *StatsConfig `protobuf:"bytes,100,opt,name=stats" json:"stats,omitempty"`
	// Misc config.
	Misc *MiscConfig `protobuf:"bytes,101,opt,name=misc" json:"misc,omitempty"`
	// App Config.
	App *AppConfig `protobuf:"bytes,102,opt,name=app" json:"app,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *Config) GetNetwork() *NetworkConfig {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *Config) GetChain() *ChainConfig {
	if m != nil {
		return m.Chain
	}
	return nil
}

func (m *Config) GetRpc() *RPCConfig {
	if m != nil {
		return m.Rpc
	}
	return nil
}

func (m *Config) GetStats() *StatsConfig {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *Config) GetMisc() *MiscConfig {
	if m != nil {
		return m.Misc
	}
	return nil
}

func (m *Config) GetApp() *AppConfig {
	if m != nil {
		return m.App
	}
	return nil
}

type NetworkConfig struct {
	// Neb seed node address.
	Seed []string `protobuf:"bytes,1,rep,name=seed" json:"seed,omitempty"`
	// Listen addresses.
	Listen []string `protobuf:"bytes,2,rep,name=listen" json:"listen,omitempty"`
	// Network node privateKey address. If nil, generate a new node.
	PrivateKey string `protobuf:"bytes,3,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// Network ID
	NetworkId uint32 `protobuf:"varint,4,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
}

func (m *NetworkConfig) Reset()                    { *m = NetworkConfig{} }
func (m *NetworkConfig) String() string            { return proto.CompactTextString(m) }
func (*NetworkConfig) ProtoMessage()               {}
func (*NetworkConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{1} }

func (m *NetworkConfig) GetSeed() []string {
	if m != nil {
		return m.Seed
	}
	return nil
}

func (m *NetworkConfig) GetListen() []string {
	if m != nil {
		return m.Listen
	}
	return nil
}

func (m *NetworkConfig) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *NetworkConfig) GetNetworkId() uint32 {
	if m != nil {
		return m.NetworkId
	}
	return 0
}

type ChainConfig struct {
	// ChainID.
	ChainId uint32 `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// genesis conf file path
	Genesis string `protobuf:"bytes,2,opt,name=genesis,proto3" json:"genesis,omitempty"`
	// Data dir.
	Datadir string `protobuf:"bytes,11,opt,name=datadir,proto3" json:"datadir,omitempty"`
	// Key dir.
	Keydir string `protobuf:"bytes,12,opt,name=keydir,proto3" json:"keydir,omitempty"`
	// Coinbase.
	Coinbase string `protobuf:"bytes,21,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
	// Miner.
	Miner string `protobuf:"bytes,22,opt,name=miner,proto3" json:"miner,omitempty"`
	// Passphrase.
	Passphrase string `protobuf:"bytes,23,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
	// Lowest GasPrice.
	GasPrice string `protobuf:"bytes,24,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	// Max GasLimit.
	GasLimit string `protobuf:"bytes,25,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	// Supported signature cipher list. ["ECC_SECP256K1"]
	SignatureCiphers []string `protobuf:"bytes,26,rep,name=signature_ciphers,json=signatureCiphers" json:"signature_ciphers,omitempty"`
}

func (m *ChainConfig) Reset()                    { *m = ChainConfig{} }
func (m *ChainConfig) String() string            { return proto.CompactTextString(m) }
func (*ChainConfig) ProtoMessage()               {}
func (*ChainConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{2} }

func (m *ChainConfig) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *ChainConfig) GetGenesis() string {
	if m != nil {
		return m.Genesis
	}
	return ""
}

func (m *ChainConfig) GetDatadir() string {
	if m != nil {
		return m.Datadir
	}
	return ""
}

func (m *ChainConfig) GetKeydir() string {
	if m != nil {
		return m.Keydir
	}
	return ""
}

func (m *ChainConfig) GetCoinbase() string {
	if m != nil {
		return m.Coinbase
	}
	return ""
}

func (m *ChainConfig) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

func (m *ChainConfig) GetPassphrase() string {
	if m != nil {
		return m.Passphrase
	}
	return ""
}

func (m *ChainConfig) GetGasPrice() string {
	if m != nil {
		return m.GasPrice
	}
	return ""
}

func (m *ChainConfig) GetGasLimit() string {
	if m != nil {
		return m.GasLimit
	}
	return ""
}

func (m *ChainConfig) GetSignatureCiphers() []string {
	if m != nil {
		return m.SignatureCiphers
	}
	return nil
}

type RPCConfig struct {
	// RPC listen addresses.
	RpcListen []string `protobuf:"bytes,1,rep,name=rpc_listen,json=rpcListen" json:"rpc_listen,omitempty"`
	// HTTP listen addresses.
	HttpListen []string `protobuf:"bytes,2,rep,name=http_listen,json=httpListen" json:"http_listen,omitempty"`
	// Enabled HTTP modules.["api", "admin"]
	HttpModule []string `protobuf:"bytes,3,rep,name=http_module,json=httpModule" json:"http_module,omitempty"`
}

func (m *RPCConfig) Reset()                    { *m = RPCConfig{} }
func (m *RPCConfig) String() string            { return proto.CompactTextString(m) }
func (*RPCConfig) ProtoMessage()               {}
func (*RPCConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{3} }

func (m *RPCConfig) GetRpcListen() []string {
	if m != nil {
		return m.RpcListen
	}
	return nil
}

func (m *RPCConfig) GetHttpListen() []string {
	if m != nil {
		return m.HttpListen
	}
	return nil
}

func (m *RPCConfig) GetHttpModule() []string {
	if m != nil {
		return m.HttpModule
	}
	return nil
}

type AppConfig struct {
	LogLevel          string `protobuf:"bytes,1,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	LogFile           string `protobuf:"bytes,2,opt,name=log_file,json=logFile,proto3" json:"log_file,omitempty"`
	EnableCrashReport bool   `protobuf:"varint,3,opt,name=enable_crash_report,json=enableCrashReport,proto3" json:"enable_crash_report,omitempty"`
}

func (m *AppConfig) Reset()                    { *m = AppConfig{} }
func (m *AppConfig) String() string            { return proto.CompactTextString(m) }
func (*AppConfig) ProtoMessage()               {}
func (*AppConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{4} }

func (m *AppConfig) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

func (m *AppConfig) GetLogFile() string {
	if m != nil {
		return m.LogFile
	}
	return ""
}

func (m *AppConfig) GetEnableCrashReport() bool {
	if m != nil {
		return m.EnableCrashReport
	}
	return false
}

type MiscConfig struct {
	// Default encryption ciper when create new keystore file.
	DefaultKeystoreFileCiper string `protobuf:"bytes,1,opt,name=default_keystore_file_ciper,json=defaultKeystoreFileCiper,proto3" json:"default_keystore_file_ciper,omitempty"`
}

func (m *MiscConfig) Reset()                    { *m = MiscConfig{} }
func (m *MiscConfig) String() string            { return proto.CompactTextString(m) }
func (*MiscConfig) ProtoMessage()               {}
func (*MiscConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{5} }

func (m *MiscConfig) GetDefaultKeystoreFileCiper() string {
	if m != nil {
		return m.DefaultKeystoreFileCiper
	}
	return ""
}

type StatsConfig struct {
	// Enable metrics or not.
	EnableMetrics   bool                          `protobuf:"varint,1,opt,name=enable_metrics,json=enableMetrics,proto3" json:"enable_metrics,omitempty"`
	ReportingModule []StatsConfig_ReportingModule `protobuf:"varint,2,rep,packed,name=reporting_module,json=reportingModule,enum=nebletpb.StatsConfig_ReportingModule" json:"reporting_module,omitempty"`
	// Influxdb config.`
	Influxdb *InfluxdbConfig `protobuf:"bytes,11,opt,name=influxdb" json:"influxdb,omitempty"`
}

func (m *StatsConfig) Reset()                    { *m = StatsConfig{} }
func (m *StatsConfig) String() string            { return proto.CompactTextString(m) }
func (*StatsConfig) ProtoMessage()               {}
func (*StatsConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{6} }

func (m *StatsConfig) GetEnableMetrics() bool {
	if m != nil {
		return m.EnableMetrics
	}
	return false
}

func (m *StatsConfig) GetReportingModule() []StatsConfig_ReportingModule {
	if m != nil {
		return m.ReportingModule
	}
	return nil
}

func (m *StatsConfig) GetInfluxdb() *InfluxdbConfig {
	if m != nil {
		return m.Influxdb
	}
	return nil
}

type InfluxdbConfig struct {
	// Host.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Port.
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// Database name.
	Db string `protobuf:"bytes,3,opt,name=db,proto3" json:"db,omitempty"`
	// Auth user.
	User string `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	// Auth password.
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
}

func (m *InfluxdbConfig) Reset()                    { *m = InfluxdbConfig{} }
func (m *InfluxdbConfig) String() string            { return proto.CompactTextString(m) }
func (*InfluxdbConfig) ProtoMessage()               {}
func (*InfluxdbConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{7} }

func (m *InfluxdbConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *InfluxdbConfig) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *InfluxdbConfig) GetDb() string {
	if m != nil {
		return m.Db
	}
	return ""
}

func (m *InfluxdbConfig) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *InfluxdbConfig) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "nebletpb.Config")
	proto.RegisterType((*NetworkConfig)(nil), "nebletpb.NetworkConfig")
	proto.RegisterType((*ChainConfig)(nil), "nebletpb.ChainConfig")
	proto.RegisterType((*RPCConfig)(nil), "nebletpb.RPCConfig")
	proto.RegisterType((*AppConfig)(nil), "nebletpb.AppConfig")
	proto.RegisterType((*MiscConfig)(nil), "nebletpb.MiscConfig")
	proto.RegisterType((*StatsConfig)(nil), "nebletpb.StatsConfig")
	proto.RegisterType((*InfluxdbConfig)(nil), "nebletpb.InfluxdbConfig")
	proto.RegisterEnum("nebletpb.StatsConfig_ReportingModule", StatsConfig_ReportingModule_name, StatsConfig_ReportingModule_value)
}

func init() { proto.RegisterFile("config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 700 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xeb, 0x6a, 0xf3, 0x46,
	0x10, 0xad, 0xed, 0x5c, 0xa4, 0x71, 0xe2, 0x24, 0x9b, 0xdb, 0x26, 0xa1, 0x4d, 0x10, 0x04, 0x0c,
	0x01, 0x43, 0xd3, 0xfe, 0xed, 0x8f, 0x62, 0x28, 0x84, 0x24, 0x25, 0xa8, 0x0f, 0x20, 0x74, 0x19,
	0xcb, 0x4b, 0xd6, 0xd2, 0xb2, 0xbb, 0xce, 0x85, 0x3e, 0x43, 0x1f, 0xa5, 0x8f, 0xf4, 0xbd, 0xcb,
	0xc7, 0xac, 0x56, 0x72, 0x6c, 0xbe, 0x7f, 0x9a, 0x73, 0xce, 0xce, 0xee, 0x9c, 0x99, 0x11, 0xec,
	0xe5, 0x75, 0x35, 0x13, 0xe5, 0x44, 0xe9, 0xda, 0xd6, 0x2c, 0xa8, 0x30, 0x93, 0x68, 0x55, 0x16,
	0xfd, 0xd7, 0x87, 0x9d, 0xa9, 0xa3, 0xd8, 0xaf, 0xb0, 0x5b, 0xa1, 0x7d, 0xaf, 0xf5, 0x2b, 0xef,
	0xdd, 0xf4, 0xc6, 0xc3, 0xfb, 0xf3, 0x49, 0x2b, 0x9b, 0xfc, 0xdd, 0x10, 0x8d, 0x32, 0x6e, 0x75,
	0xec, 0x0e, 0xb6, 0xf3, 0x79, 0x2a, 0x2a, 0xde, 0x77, 0x07, 0x4e, 0x57, 0x07, 0xa6, 0x04, 0x7b,
	0x79, 0xa3, 0x61, 0xb7, 0x30, 0xd0, 0x2a, 0xe7, 0x03, 0x27, 0x3d, 0x5e, 0x49, 0xe3, 0x97, 0xa9,
	0x17, 0x12, 0x4f, 0x39, 0x8d, 0x4d, 0xad, 0xe1, 0xc5, 0x66, 0xce, 0x7f, 0x08, 0x6e, 0x73, 0x3a,
	0x0d, 0x1b, 0xc3, 0xd6, 0x42, 0x98, 0x9c, 0xa3, 0xd3, 0x9e, 0xac, 0xb4, 0xcf, 0xc2, 0xe4, 0x5e,
	0xea, 0x14, 0x74, 0x7b, 0xaa, 0x14, 0x9f, 0x6d, 0xde, 0xfe, 0xa7, 0x52, 0xed, 0xed, 0xa9, 0x52,
	0xd1, 0xbf, 0xb0, 0xbf, 0x56, 0x2b, 0x63, 0xb0, 0x65, 0x10, 0x0b, 0xde, 0xbb, 0x19, 0x8c, 0xc3,
	0xd8, 0x7d, 0xb3, 0x33, 0xd8, 0x91, 0xc2, 0x58, 0xa4, 0xba, 0x09, 0xf5, 0x11, 0xbb, 0x86, 0xa1,
	0xd2, 0xe2, 0x2d, 0xb5, 0x98, 0xbc, 0xe2, 0xa7, 0xab, 0x34, 0x8c, 0xc1, 0x43, 0x8f, 0xf8, 0xc9,
	0x7e, 0x06, 0xf0, 0xd6, 0x25, 0xa2, 0xe0, 0x5b, 0x37, 0xbd, 0xf1, 0x7e, 0x1c, 0x7a, 0xe4, 0xa1,
	0x88, 0xfe, 0xef, 0xc3, 0xf0, 0x8b, 0x71, 0xec, 0x02, 0x02, 0x67, 0x1d, 0x89, 0x7b, 0x4e, 0xbc,
	0xeb, 0xe2, 0x87, 0x82, 0x71, 0xd8, 0x2d, 0xb1, 0x42, 0x23, 0x8c, 0xf3, 0x3e, 0x8c, 0xdb, 0x90,
	0x98, 0x22, 0xb5, 0x69, 0x21, 0x34, 0x1f, 0x36, 0x8c, 0x0f, 0xe9, 0xd9, 0xaf, 0xf8, 0x49, 0xc4,
	0x9e, 0x23, 0x7c, 0xc4, 0x2e, 0x21, 0xc8, 0x6b, 0x51, 0x65, 0xa9, 0x41, 0x7e, 0xea, 0x98, 0x2e,
	0x66, 0x27, 0xb0, 0xbd, 0x10, 0x15, 0x6a, 0x7e, 0xe6, 0x88, 0x26, 0x60, 0xbf, 0x00, 0xa8, 0xd4,
	0x18, 0x35, 0xd7, 0x74, 0xe6, 0xdc, 0xd7, 0xd9, 0x21, 0xec, 0x0a, 0xc2, 0x32, 0x35, 0x89, 0xd2,
	0x22, 0x47, 0xce, 0x9b, 0x94, 0x65, 0x6a, 0x5e, 0x28, 0x6e, 0x49, 0x29, 0x16, 0xc2, 0xf2, 0x8b,
	0x8e, 0x7c, 0xa2, 0x98, 0xdd, 0xc1, 0x91, 0x11, 0x65, 0x95, 0xda, 0xa5, 0xc6, 0x24, 0x17, 0x6a,
	0x8e, 0xda, 0xf0, 0x4b, 0xe7, 0xf2, 0x61, 0x47, 0x4c, 0x1b, 0x3c, 0x92, 0x10, 0x76, 0xc3, 0x43,
	0xde, 0x6a, 0x95, 0x27, 0xbe, 0x31, 0x4d, 0xbb, 0x42, 0xad, 0xf2, 0xa7, 0xae, 0x37, 0x73, 0x6b,
	0x55, 0xb2, 0xd6, 0x38, 0x20, 0x68, 0x43, 0xb0, 0xa8, 0x8b, 0xa5, 0x44, 0x3e, 0x58, 0x09, 0x9e,
	0x1d, 0x12, 0x19, 0x08, 0xbb, 0x61, 0xa1, 0x22, 0x64, 0x5d, 0x26, 0x12, 0xdf, 0x50, 0xba, 0xde,
	0x84, 0x71, 0x20, 0xeb, 0xf2, 0x89, 0x62, 0xea, 0x1b, 0x91, 0x33, 0x21, 0xb1, 0xed, 0x8e, 0xac,
	0xcb, 0xbf, 0x84, 0x44, 0x36, 0x81, 0x63, 0xac, 0xd2, 0x4c, 0x62, 0x92, 0xeb, 0xd4, 0xcc, 0x13,
	0x8d, 0xaa, 0xd6, 0xd6, 0x8d, 0x4a, 0x10, 0x1f, 0x35, 0xd4, 0x94, 0x98, 0xd8, 0x11, 0xd1, 0x23,
	0xc0, 0x6a, 0x94, 0xd9, 0x1f, 0x70, 0x55, 0xe0, 0x2c, 0x5d, 0x4a, 0x4b, 0x03, 0x66, 0x6c, 0xad,
	0xd1, 0xdd, 0x42, 0x4e, 0xa1, 0xf6, 0xef, 0xe0, 0x5e, 0xf2, 0xe8, 0x15, 0x74, 0xef, 0x94, 0xf8,
	0xe8, 0x5b, 0x0f, 0x86, 0x5f, 0x96, 0x88, 0xdd, 0xc2, 0xc8, 0x3f, 0x66, 0x81, 0x56, 0x8b, 0xdc,
	0xb8, 0x0c, 0x41, 0xbc, 0xdf, 0xa0, 0xcf, 0x0d, 0xc8, 0x5e, 0xe0, 0xb0, 0x79, 0xa6, 0xa8, 0xca,
	0xd6, 0x1e, 0xf2, 0x6f, 0x74, 0x7f, 0xfb, 0xc3, 0xe5, 0x9c, 0xc4, 0xad, 0xba, 0x71, 0x2e, 0x3e,
	0xd0, 0xeb, 0x00, 0xfb, 0x1d, 0x02, 0x51, 0xcd, 0xe4, 0xf2, 0xa3, 0xc8, 0xdc, 0x90, 0x0e, 0xef,
	0xf9, 0x2a, 0xd3, 0x83, 0x67, 0xfc, 0x5a, 0x76, 0xca, 0xe8, 0x1a, 0x0e, 0x36, 0x32, 0xb3, 0x3d,
	0x08, 0x5a, 0xf9, 0xe1, 0x4f, 0xd1, 0x07, 0x8c, 0xd6, 0x0f, 0xd3, 0xf6, 0xce, 0x6b, 0x63, 0xbd,
	0x33, 0xee, 0x9b, 0x30, 0xe7, 0x79, 0xdf, 0x6d, 0x94, 0xfb, 0x66, 0x23, 0xe8, 0x17, 0x99, 0x5f,
	0xd8, 0x7e, 0x91, 0x91, 0x66, 0x69, 0x50, 0xbb, 0x15, 0x0d, 0x63, 0xf7, 0x4d, 0x6b, 0x42, 0x23,
	0xfe, 0x5e, 0xeb, 0x82, 0x6f, 0x37, 0x1d, 0x6f, 0xe3, 0x6c, 0xc7, 0xfd, 0x57, 0x7f, 0xfb, 0x1e,
	0x00, 0x00, 0xff, 0xff, 0xc9, 0xd3, 0x1c, 0x31, 0x67, 0x05, 0x00, 0x00,
}
