package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"suzaku/pkg/constant"
)

var Config config

type config struct {
	ServerIP        string          `yaml:"server_ip"`
	Tsl             Tsl             `yaml:"tsl"`
	Etcd            Etcd            `yaml:"etcd"`
	RPCKeepalive    RPCKeepalive    `yaml:"rpc_keepalive"`
	RPCPort         RPCPort         `yaml:"rpc_port"`
	RPCRegisterName RPCRegisterName `yaml:"rpc_register_name"`
	Mysql           Mysql           `yaml:"mysql"`
	Mongo           Mongo           `yaml:"mongo"`
	Redis           Redis           `yaml:"redis"`
	Kafka           Kafka           `yaml:"kafka"`
	JwtAuth         JwtAuth         `yaml:"jwt_auth"`
	Casbin          Casbin          `yaml:"casbin"`
	Environment     Environment     `yaml:"environment"`
	Monlog          Monlog          `yaml:"monlog"`
	Endpoints       Endpoints       `yaml:"endpoints"`
	API             API             `yaml:"api"`
	Websocket       Websocket       `yaml:"websocket"`
	Credential      Credential      `yaml:"credential"`
	Log             Log             `yaml:"log"`
	Abnormal        Abnormal        `yaml:"abnormal"`
	Zap             Zap             `yaml:"zap"`
	Secret          string          `yaml:"secret"`
	Callback        Callback        `yaml:"callback"`
	IosPush         IosPush         `yaml:"ios_push"`
}
type Tsl struct {
	CertFile string `yaml:"cert_file"`
	KeyFile  string `yaml:"key_file"`
}
type Etcd struct {
	Address      []string `yaml:"address"`
	Schema       string   `yaml:"schema"`
	ReadTimeout  int      `yaml:"read_timeout"`
	WriteTimeout int      `yaml:"write_timeout"`
	DialTimeout  int      `yaml:"dial_timeout"`
}
type CallbackBeforeSendMsg struct {
	Switch          bool `yaml:"switch"`
	TimeoutStrategy int  `yaml:"timeout_strategy"`
}
type CallbackAfterSendMsg struct {
	Switch bool `yaml:"switch"`
}
type C2C struct {
	CallbackBeforeSendMsg CallbackBeforeSendMsg `yaml:"callback_before_send_msg"`
	CallbackAfterSendMsg  CallbackAfterSendMsg  `yaml:"callback_after_send_msg"`
}
type StateChange struct {
	Switch bool `yaml:"switch"`
}
type State struct {
	StateChange StateChange `yaml:"state_change"`
}
type RPCKeepalive struct {
	IdleTimeout       int `yaml:"idle_timeout"`
	ForceCloseWait    int `yaml:"force_close_wait"`
	KeepAliveInterval int `yaml:"keep_alive_interval"`
	KeepAliveTimeout  int `yaml:"keep_alive_timeout"`
	MaxLifeTime       int `yaml:"max_life_time"`
}
type RPCPort struct {
	UserPort           []int `yaml:"user_port"`
	FriendPort         []int `yaml:"friend_port"`
	OfflineMessagePort []int `yaml:"offline_message_port"`
	OnlineRelayPort    []int `yaml:"online_relay_port"`
	GroupPort          []int `yaml:"group_port"`
	AuthPort           []int `yaml:"auth_port"`
	PushPort           []int `yaml:"push_port"`
	StatisticsPort     []int `yaml:"statistics_port"`
	MessageCmsPort     []int `yaml:"message_cms_port"`
	AdminCmsPort       []int `yaml:"admin_cms_port"`
	C2C                C2C   `yaml:"c2c"`
	State              State `yaml:"state"`
}
type RPCRegisterName struct {
	UserName               string `yaml:"user_name"`
	FriendName             string `yaml:"friend_name"`
	OfflineMessageName     string `yaml:"offline_message_name"`
	PushName               string `yaml:"push_name"`
	OnlineMessageRelayName string `yaml:"online_message_relay_name"`
	GroupName              string `yaml:"group_name"`
	AuthName               string `yaml:"auth_name"`
	CacheName              string `yaml:"cache_name"`
	StatisticsName         string `yaml:"statistics_name"`
	MessageCmsName         string `yaml:"message_cms_name"`
	AdminCmsName           string `yaml:"admin_cms_name"`
}

type Mysql struct {
	Address      []string `yaml:"address"`
	Username     string   `yaml:"username"`
	Password     string   `yaml:"password"`
	Db           string   `yaml:"db"`
	MaxOpenConn  int      `yaml:"max_open_conn"`
	MaxIdleConn  int      `yaml:"max_idle_conn"`
	ConnLifetime int      `yaml:"conn_lifetime"`
	Charset      string   `yaml:"charset"`
}
type Mongo struct {
	Address           []string `yaml:"address"`
	Username          string   `yaml:"username"`
	Password          string   `yaml:"password"`
	Db                string   `yaml:"db"`
	Direct            bool     `yaml:"direct"`
	Timeout           int      `yaml:"timeout"`
	MaxPoolSize       int      `yaml:"max_pool_size"`
	RetainChatRecords int      `yaml:"retain_chat_records"`
}
type Redis struct {
	Address  []string `yaml:"address"`
	Db       int      `yaml:"db"`
	Password string   `yaml:"password"`
	Prefix   string   `yaml:"prefix"`
}
type Ws2Mschat struct {
	Addr  []string `yaml:"addr"`
	Topic string   `yaml:"topic"`
}
type Ms2Pschat struct {
	Addr  []string `yaml:"addr"`
	Topic string   `yaml:"topic"`
}
type ConsumerGroupID struct {
	MsgToMongo string `yaml:"msgToMongo"`
	MsgToMySQL string `yaml:"msgToMySql"`
	MsgToPush  string `yaml:"msgToPush"`
}
type Kafka struct {
	Ws2Mschat       Ws2Mschat       `yaml:"ws2mschat"`
	Ms2Pschat       Ms2Pschat       `yaml:"ms2pschat"`
	ConsumerGroupID ConsumerGroupID `yaml:"consumer_group_id"`
}
type JwtAuth struct {
	AuthMethod string `yaml:"auth_method"`
	IsDev      bool   `yaml:"is_dev"`
}
type Casbin struct {
	ModelPath string `yaml:"model-path"`
}
type Environment struct {
	RunModel string `yaml:"run_model"`
}
type Monlog struct {
	BatchSize     int   `yaml:"batch_size"`
	CommitTimeout int   `yaml:"commit_timeout"`
	Mongo         Mongo `yaml:"mongo"`
}
type Endpoints struct {
	API string `yaml:"api"`
}
type API struct {
	Port []int `yaml:"port"`
}
type Websocket struct {
	Port            []int `yaml:"port"`
	WriteWait       int   `yaml:"write_wait"`
	PongWait        int   `yaml:"pong_wait"`
	MaxMessageSize  int   `yaml:"max_message_size"`
	ReadBufferSize  int   `yaml:"read_buffer_size"`
	WriteBufferSize int   `yaml:"write_buffer_size"`
}
type Tencent struct {
	AppID     int    `yaml:"app_id"`
	Region    string `yaml:"region"`
	Bucket    string `yaml:"bucket"`
	SecretID  string `yaml:"secret_id"`
	SecretKey string `yaml:"secret_key"`
}
type Minio struct {
	Bucket    string `yaml:"bucket"`
	Location  string `yaml:"location"`
	Endpoint  string `yaml:"endpoint"`
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
}
type Credential struct {
	Tencent Tencent `yaml:"tencent"`
	Minio   Minio   `yaml:"minio"`
}
type Log struct {
	StorageLocation string   `yaml:"storage_location"`
	RotationTime    int      `yaml:"rotation_time"`
	RotationCount   int      `yaml:"rotation_count"`
	Level           int      `yaml:"level"`
	EsAddress       []string `yaml:"es_address"`
	EsUsername      string   `yaml:"es_username"`
	EsPassword      string   `yaml:"es_password"`
	EsSwitch        bool     `yaml:"es_switch"`
}
type Abnormal struct {
	File  string `yaml:"file"`
	Level int    `yaml:"level"`
}
type Zap struct {
	Encoder       string  `json:"encoder" yaml:"encoder"`               // 编码器 console Or json
	Directory     string  `json:"directory"  yaml:"directory"`          // 日志文件夹
	ShowLine      bool    `json:"show_line" yaml:"show_line"`           // 显示行
	EncodeLevel   string  `json:"encode_level" yaml:"encode_level"`     // 编码级
	StacktraceKey string  `json:"stacktrace_key" yaml:"stacktrace_key"` // 栈名
	LogStdout     bool    `json:"log_stdout" yaml:"log_stdout"`         // 输出控制台
	Segment       Segment `json:"segment" yaml:"segment"`               // 日志分割
}
type Segment struct {
	MaxSize    int  `json:"maxsize" yaml:"maxsize"`
	MaxAge     int  `json:"maxage" yaml:"maxage"`
	MaxBackups int  `json:"maxbackups" yaml:"maxbackups"`
	LocalTime  bool `json:"localtime" yaml:"localtime"`
	Compress   bool `json:"compress" yaml:"compress"`
}
type CallbackBeforeSendSingleMsg struct {
	Enable                 bool `yaml:"enable"`
	CallbackTimeOut        int  `yaml:"callback_time_out"`
	CallbackFailedContinue bool `yaml:"callback_failed_continue"`
}
type CallbackAfterSendSingleMsg struct {
	Enable          bool `yaml:"enable"`
	CallbackTimeOut int  `yaml:"callback_time_out"`
}
type CallbackBeforeSendGroupMsg struct {
	Enable                 bool `yaml:"enable"`
	CallbackTimeOut        int  `yaml:"callback_time_out"`
	CallbackFailedContinue bool `yaml:"callback_failed_continue"`
}
type CallbackAfterSendGroupMsg struct {
	Enable          bool `yaml:"enable"`
	CallbackTimeOut int  `yaml:"callback_time_out"`
}
type CallbackWordFilter struct {
	Enable                 bool `yaml:"enable"`
	CallbackTimeOut        int  `yaml:"callback_time_out"`
	CallbackFailedContinue bool `yaml:"callback_failed_continue"`
}
type Callback struct {
	CallbackURL                 string                      `yaml:"callback_url"`
	CallbackBeforeSendSingleMsg CallbackBeforeSendSingleMsg `yaml:"callback_before_send_single_msg"`
	CallbackAfterSendSingleMsg  CallbackAfterSendSingleMsg  `yaml:"callback_after_send_single_msg"`
	CallbackBeforeSendGroupMsg  CallbackBeforeSendGroupMsg  `yaml:"callbackBeforeSendGroupMsg"`
	CallbackAfterSendGroupMsg   CallbackAfterSendGroupMsg   `yaml:"callback_after_send_group_msg"`
	CallbackWordFilter          CallbackWordFilter          `yaml:"callback_word_filter"`
}

type IosPush struct {
	PushSound  string `yaml:"push_sound"`
	BadgeCount bool   `yaml:"badge_count"`
}

func init() {
	var (
		runMode string
		buf     []byte
		err     error
	)
	runMode = os.Getenv(constant.EnvironmentExportKey)
	if runMode == "" {
		runMode = constant.EnvironmentDev
	}
	path, _ := os.Getwd()
	path += "/configs/" + runMode + ".yaml"
	buf, err = ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(buf, &Config)
}
