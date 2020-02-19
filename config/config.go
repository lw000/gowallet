package config

import (
	"encoding/json"
	"github.com/lw000/gocommon/db/mysql"
	"github.com/lw000/gocommon/network/ws/cli"
	"io/ioutil"
)

type Server struct {
	Listen      int64
	Servername  []string
	Blacklist   []string
	Whitelist   []string
	Ssl         string
	SslCertfile string
	SslKeyfile  string
}

type Servers struct {
	Debug  int64
	Server []Server
}

// 工程配置结构体
type JSONConfig struct {
	MysqlCfgs []tymysql.JsonConfig
	WsCfg     tywscfg.JsonConfig
	Servers   Servers
}

// NewJSONConfig ...
func NewJSONConfig() *JSONConfig {
	return &JSONConfig{}
}

// 加载配置文件
func LoadJSONConfig(file string) (*JSONConfig, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var ccf CfgStruct
	if err = json.Unmarshal(data, &ccf); err != nil {
		return nil, err
	}

	cfg := NewJSONConfig()
	cfg.Servers.Debug = ccf.Servers.Debug

	for _, srv := range ccf.Servers.Server {
		cfg.Servers.Server = append(cfg.Servers.Server, Server{
			srv.Listen,
			srv.Servername,
			srv.Blacklist,
			srv.Whitelist,
			srv.Ssl,
			srv.SslCertfile,
			srv.SslKeyfile,
		})
	}

	// 数据接口配置
	for _, mycfg := range ccf.Mysqls {
		var c tymysql.JsonConfig
		c.Database = mycfg.Database
		c.Host = mycfg.Host
		c.MaxOdleConns = mycfg.MaxOdleConns
		c.MaxOpenConns = mycfg.MaxOpenConns
		c.Password = mycfg.Password
		c.Username = mycfg.Username
		cfg.MysqlCfgs = append(cfg.MysqlCfgs, c)
	}

	// ws配置
	cfg.WsCfg.Host = ccf.Ws.Host
	cfg.WsCfg.Path = ccf.Ws.Path

	// // Redis配置
	// cfg.RdsCfg.Host = ccf.Redis.Host
	// cfg.RdsCfg.Db = ccf.Redis.DB
	// cfg.RdsCfg.Psd = ccf.Redis.Psd
	// cfg.RdsCfg.PoolSize = ccf.Redis.PoolSize
	// cfg.RdsCfg.MinIdleConns = ccf.Redis.MinIdleConns

	// // 邮件配置
	// cfg.MailCfg.From = ccf.Mail.From
	// cfg.MailCfg.To = ccf.Mail.To
	// cfg.MailCfg.Pass = ccf.Mail.Pass
	// cfg.MailCfg.Host = ccf.Mail.Host
	// cfg.MailCfg.Port = ccf.Mail.Port

	return cfg, nil
}
