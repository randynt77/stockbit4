package config

import (
	"fmt"
	"log"
	"stockbit4/pkg/util"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type (
	// AppConfig - hold data from config file and all conns
	AppConfig struct {
		dbConns  *DatabaseObjects
		DbConfig DBConfig
	}
	DatabaseObjects struct {
		Movie DatabaseDependency `cfg:"movie"`
	}
	DatabaseDependency struct {
		Master *sqlx.DB
		Slave  *sqlx.DB
	}
	DBConfig struct {
		Database map[string]*struct {
			Name         string `gcfg:"name"`
			Master       string `gcfg:"master"`
			Slave        string `gcfg:"slave"`
			MaxConn      string `gcfg:"maxConn"`
			IdleConn     string `gcfg:"idleConn"`
			Username     string `gcfg:"username"`
			MasterSecret string `gcfg:"masterSecret"`
			SlaveSecret  string `gcfg:"slaveSecret"`
		}
	}
)

func InitConfig(path ...string) *AppConfig {
	dbCfg := DBConfig{}
	isSuccess := util.ReadModuleConfig(&dbCfg, "/etc/stockbit4/", "db")
	if !isSuccess {
		return nil
	}
	return &AppConfig{
		DbConfig: dbCfg,
	}
}
func (g *AppConfig) GetDatabaseConns() *DatabaseObjects {
	if len(g.DbConfig.Database) == 0 {
		log.Fatal("No valid db host exist!")
	}

	if g.dbConns != nil {
		return g.dbConns
	}

	dbConns := make(map[string]DatabaseDependency)
	for k, v := range g.DbConfig.Database {
		var maxConn, idleConn int
		var err error
		maxConn, err = strconv.Atoi(v.MaxConn)
		if err != nil || maxConn == 0 {
			maxConn = 20
		}
		idleConn, err = strconv.Atoi(v.IdleConn)
		if err != nil || idleConn == 0 {
			idleConn = 5
		}

		dbMaster, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", v.Username, v.MasterSecret, v.Master, v.Name))
		if err != nil {
			log.Fatal(err, " failed connect db")
		}
		dbMaster.SetMaxOpenConns(maxConn)
		dbMaster.SetMaxIdleConns(idleConn)

		dbSlave, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", v.Username, v.SlaveSecret, v.Slave, v.Name))
		if err != nil {
			log.Fatal(err, " failed connect db")
		}
		dbSlave.SetMaxOpenConns(maxConn)
		dbSlave.SetMaxIdleConns(idleConn)

		dbConns[k] = DatabaseDependency{Master: dbMaster, Slave: dbSlave}
	}

	g.dbConns = &DatabaseObjects{}
	err := util.MapToStructGeneric(g.dbConns, &dbConns, "cfg", true)
	if err != nil {
		log.Fatal("Couldn't set Database Conn Object", err)
	}

	return g.dbConns
}
