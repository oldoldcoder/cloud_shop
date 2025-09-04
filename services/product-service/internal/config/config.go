package config

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
)

// Config 服务配置
type Config struct {
	Server struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`

	Database struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Database string `mapstructure:"database"`
	} `mapstructure:"database"`

	Redis struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
	} `mapstructure:"redis"`

	JWT struct {
		Secret string `mapstructure:"secret"`
	} `mapstructure:"jwt"`

	Migrate struct {
		Auto bool `mapstructure:"auto"`
	} `mapstructure:"migrate"`

	Nacos struct {
		Enabled    bool   `mapstructure:"enabled"`
		ServerAddr string `mapstructure:"server_addr"`
		Username   string `mapstructure:"username"`
		Password   string `mapstructure:"password"`
		Namespace  string `mapstructure:"namespace"`
		Group      string `mapstructure:"group"`
		DataId     string `mapstructure:"data_id"`
	} `mapstructure:"nacos"`

	Service struct {
		Name     string `mapstructure:"name"`
		Register bool   `mapstructure:"register"`
		IP       string `mapstructure:"ip"`
	} `mapstructure:"service"`
}

func LoadConfigFromNacos(serverAddr string, username, password, namespace, group, dataId string) (*Config, error) {
	// 解析 serverAddr，支持 "127.0.0.1:8848" 格式
	ip := serverAddr
	port := uint64(8848)
	if strings.Contains(serverAddr, ":") {
		parts := strings.Split(serverAddr, ":")
		ip = parts[0]
		if p, err := strconv.ParseUint(parts[1], 10, 64); err == nil {
			port = p
		}
	}

	sc := []constant.ServerConfig{{IpAddr: ip, Port: port}}

	// 对于 public namespace，使用空字符串作为 namespaceId
	namespaceId := namespace
	if namespace == "public" {
		namespaceId = ""
	}

	cc := constant.ClientConfig{
		NamespaceId: namespaceId, Username: username, Password: password,
		TimeoutMs: 5000, BeatInterval: 5000, LogLevel: "info",
	}
	cfgClient, err := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  &cc,
		ServerConfigs: sc,
	})
	if err != nil {
		return nil, err
	}

	content, err := cfgClient.GetConfig(vo.ConfigParam{DataId: dataId, Group: group})
	if err != nil {
		return nil, err
	}

	v := viper.New()
	v.SetConfigType("yaml")
	if err := v.ReadConfig(bytes.NewBufferString(content)); err != nil {
		return nil, err
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

// LoadConfig 从本地文件加载配置
func LoadConfig(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}
	return &c, nil
}
