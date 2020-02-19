package database

import (
	"github.com/lw000/gocommon/db/mysql"
	log "github.com/sirupsen/logrus"
)

const (
	DB_WALLET = "wallet"
)

var (
	// 数据库
	dbServer map[string]*tymysql.Mysql
)

func init() {
	dbServer = make(map[string]*tymysql.Mysql)
}

// 打开数据库连接
func OpenMysql(cfgs ...tymysql.JsonConfig) error {
	for _, cfg := range cfgs {
		srv := &tymysql.Mysql{}
		if err := srv.OpenWithJsonConfig(&cfg); err != nil {
			log.Error(err)
			return err
		}
		log.WithFields(log.Fields{"mysql": cfg.Database}).Info("数据库连接成功")
		dbServer[cfg.Database] = srv
	}
	return nil
}

func CloseMysql() {
	for _, srv := range dbServer {
		if err := srv.Close(); err != nil {
			log.Error(err)
		}
	}
}

// 获取数据库连接
func GetMysql(database string) *tymysql.Mysql {
	srv, exists := dbServer[database]
	if exists {
		return srv
	}
	return nil
}
