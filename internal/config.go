package internal

import (
	"errors"
	"flag"
)

var Cfg Config

// 代表程序输入参数的配置
type (
	//
	Config struct {
		Type         string
		CloneConfig  CloneConfig
		BranchConfig BranchConfig
	}
	// 切换分支配置
	BranchConfig struct {
		// 切换这个目录下的所有项目
		DirPath string
		// 切换分支的名字
		BranchName string
	}
	// 克隆组配置
	CloneConfig struct {
		// 组名
		GroupName string
		// Gitlab的地址. host:port
		Addr string
		// gitlab生成AccessToken,必须给予对应的权限. 最好把read权限都勾上生成
		AccessToken string
		// 克隆下来的项目保存的位置
		Output string
	}
)

//
func init() {
	flag.StringVar(&Cfg.Type, "t", "cg", "the business type that 'cg' or 'b'")
	initCloneConfig()
	initBranchConfig()
}

// Initialize Branch Config
func initBranchConfig() {
	flag.StringVar(&Cfg.BranchConfig.BranchName, "branch", "", "The branch name")
	flag.StringVar(&Cfg.BranchConfig.DirPath, "dir", "", "Need to switch the directory path of the branch")
}

// Initialize Clone Config flag
func initCloneConfig() {
	var addrUsage = "The ip and port such as 127.0.0.1:8080"
	flag.StringVar(&Cfg.CloneConfig.Addr, "addr", "", addrUsage)
	flag.StringVar(&Cfg.CloneConfig.GroupName, "group", "", "The group name")
	var tokenUsage = "Generate token in the gitlab. setting->Access Tokens. give only read authorization"
	flag.StringVar(&Cfg.CloneConfig.AccessToken, "token", "", tokenUsage)
	flag.StringVar(&Cfg.CloneConfig.Output, "out", "", "output path")
}

// 检测参数
func (c *Config) CheckParam() error {
	err := Cfg.CloneConfig.CheckParam()
	if err != nil {
		return err
	}
	return nil
}

// 检查克隆需要的参数
func (c *CloneConfig) CheckParam() error {
	if c.Addr == "" {
		return errors.New("the Addr(-addr) is empty")
	}
	if c.AccessToken == "" {
		return errors.New("the Access Token(-token) is empty")
	}
	if c.Output == "" {
		c.Output = "."
	}
	return nil
}

// 检查克隆需要的参数
func (c *BranchConfig) CheckParam() error {
	if c.DirPath == "" {
		return errors.New("the DirPath(-dir) is empty")
	}
	if c.BranchName == "" {
		return errors.New("the BranchName(-name) is empty")
	}
	return nil
}
