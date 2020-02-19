// goWallet project main.go
package main

import (
	"github.com/judwhite/go-svc/svc"
	"gowallet/dao/database"
	"gowallet/global"
	"gowallet/wallet"

	_ "github.com/icattlecoder/godaemon"
	log "github.com/sirupsen/logrus"
)

type Program struct {
	htpServe GinServer
}

func (p *Program) Init(env svc.Environment) error {
	if env.IsWindowsService() {

	} else {

	}

	var err error
	if err = global.LoadGlobalConfig(); err != nil {
		log.Panic(err)
	}
	return nil
}

// Start is called after Init. This method must be non-blocking.
func (p *Program) Start() error {
	var err error
	// 连接数据库
	err = database.OpenMysql(global.ProjectConfig.MysqlCfgs...)
	if err != nil {
		log.Panic(err)
	}

	err = wallet.WalletService().LoadWallets()
	if err != nil {
		log.Panic(err)
	}

	// 启动http服务
	err = p.htpServe.Start(global.ProjectConfig.Servers.Debug, global.ProjectConfig.Servers.Server)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// Stop is called in response to syscall.SIGINT, syscall.SIGTERM, or when a
// Windows Service is stopped.
func (p *Program) Stop() error {
	var err error

	p.htpServe.Stop()

	// 存储钱包
	if err = wallet.WalletService().StoreWallets(); err != nil {
		log.Error(err)
	}

	// 关闭数据库连接
	database.CloseMysql()

	log.Error("WALLET·服务退出")
	return nil
}

func main() {
	pro := &Program{}
	if err := svc.Run(pro); err != nil {
		log.Error(err)
	}
}
