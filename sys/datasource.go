package sys

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/jwei2006/gweb/config"
	"os"
)

var engine *xorm.Engine

func InitDb() {
	if !config.Conf.Db {
		return
	}
	var e error
	engine, e = xorm.NewEngine(config.Conf.DbType, config.Conf.DbUrl)
	if e != nil {
		Error().Fatalln(e)
		return
	}
	engine.SetMaxIdleConns(config.Conf.MaxIdle)
	engine.SetMaxOpenConns(config.Conf.MaxActive)
	engine.ShowSQL(config.Conf.ShowSql)
	if !config.Conf.IsDebug && config.Conf.ShowSql {
		file, e := os.Create(config.Conf.LogInfoPath)
		if e != nil {
			Error().Fatalln(e)
			return
		}
		_log := xorm.NewSimpleLogger(file)
		engine.SetLogger(_log)
	}
}

func GetEngine() *xorm.Engine {
	return engine
}

func Close() {
	if engine != nil {
		engine.Close()
	}
}
