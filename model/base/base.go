package base

import (
	"fmt"
	"time"
)

// Config 服务器配置
type GinConfig struct {
	Name             string        `yaml:"name" json:"name"`       // 服务名称，必填
	Version          string        `yaml:"version" json:"version"` // 服务版本，必填
	Host             string        `yaml:"host" json:"host"`       // 域名主机
	Port             int64         `yaml:"port" json:"port"`
	BroadcastIP      string        `yaml:"broadcastIP" json:"broadcastIP"`               // 广播的运行地址，默认为：IP
	BroadcastPort    int           `yaml:"broadcastPort" json:"broadcastPort"`           // 广播的运行端口，默认为：Port
	Timeout          int           `yaml:"timeout" json:"timeout"`                       // 优雅退出时的超时机制
	Debug            bool          `yaml:"debug" json:"debug"`                           // 是否开启调试
	Pprof            bool          `yaml:"pprof" json:"pprof"`                           // 是否监控性能
	ReadTimeout      time.Duration `yaml:"readtimeout" json:"readtimeout"`               // 读超时
	WriteTimeout     time.Duration `yaml:"writetimeout" json:"writetimeout"`             // 写超时
	DisableAccessLog bool          `yaml:"disable_access_log" json:"disable_access_log"` // disable_access_log
	OpenLimit        bool          `yaml:"openLimit" json:"openLimit"`                   //打开限流器
	UseMicroRandom   bool          `yaml:"useMicroRandom" json:"useMicroRandom"`         //使用轮询选择器
	UseOldSelected   bool          `yaml:"useOldSelected" json:"useOldSelected"`         //使用旧的micro缓存器，可能会增加带宽
}

// Addr 运行地址
func (i *GinConfig) Addr() string {
	return fmt.Sprintf("%s:%d", i.Host, i.Port)
}

// BroadcastAddr 广播的运行地址
func (i *GinConfig) BroadcastAddr() string {
	return fmt.Sprintf("%s:%d", i.BroadcastIP, i.BroadcastPort)
}

type OpenTaoBao struct {
	AppKey          string `yaml:"appKey" json:"appKey"`
	AppSecret       string `yaml:"appSecret" json:"appSecret"`
	ServerUrl       string `yaml:"serverUrl" json:"serverUrl"`
	ConnectTimeount int64  `yaml:"connectTimeout" json:"connectTimeout"`
	ReadTimeout     int64  `yaml:"readTimeout" json:"readTimeout"`
}
