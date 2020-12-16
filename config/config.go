package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Cfg 设置全局变量
var Cfg Config

// Config 配置参数
type Config struct {
	Cred Cred
	Host Host
}

// Cred 认证
type Cred struct {
	PublicKey  string
	PrivateKey string
}

// Host 配置
type Host struct {
	ImageID            string `toml: "ImageId"`
	LoginMode          string
	ChargeType         string
	CPU                int
	Memory             int
	NetCapability      string
	MachineType        string
	MinimalCPUPlatform string `toml: "MinimalCpuPlatform"`
	NetworkInterface   NetworkInterface
	Disks              Disk
}

// NetworkInterface 设置
type NetworkInterface struct {
	Bandwidth    int
	PayMode      string
	OperatorName string
}

// Disk 配置
type Disk struct {
	IsBoot string
	Size   int
	Type   string
}

// Init 解析配置文件，初始化变量
func Init(cfgFile string) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in current directory with name "ucloud" with no extension.
		viper.AddConfigPath(".")
		viper.SetConfigName("ucloud")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		err := viper.Unmarshal(&Cfg)
		if err != nil {
			fmt.Printf("解析配置文件出错, %v", err)
			// 配置文件解析错误直接退出程序
			os.Exit(1)
		}

	}
}
