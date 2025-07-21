package pg

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Rivalz-ai/framework-be/framework/config/vault"
	pgcf "github.com/Rivalz-ai/framework-be/framework/db/pg/config"
	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/Rivalz-ai/framework-be/framework/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"moul.io/zapgorm2"

	//"gorm.io/driver/postgres"
	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
)

type PgDB struct {
	Config  *vault.Vault
	DB      map[string]map[int]*gorm.DB
	NumNode map[string]int
	Key     int
	/*
		DB["user"][node-0]
		DB["user"][node-1]
		DB["user"][node-N]
		DB["parter"][node-0]
		DB["parter"][node-1]
		DB["parter"][node-n]
	*/
}

func (pg *PgDB) Initial(config *vault.Vault) error {
	if pg.Config == nil {
		pg.Config = config
	}
	serviceName := pg.Config.GetServiceName()
	servicePath := strings.ReplaceAll(serviceName, ".", "/")
	check, err := pg.Config.CheckPathExist(servicePath + "/db/postgres")
	if err != nil {
		panic("Postgres Check Path Exist Error: " + err.Error())
	}
	dbConfigPath := fmt.Sprintf("%s/%s", servicePath, "db/postgres/general")
	dbConfigMap := pgcf.GetConfig(pg.Config, dbConfigPath)
	//
	pg.DB = make(map[string]map[int]*gorm.DB)
	pg.NumNode = make(map[string]int)
	//
	if check { //multi db
		db_list := pg.Config.ListItemByPath(servicePath + "/db/postgres")
		for _, dbname := range db_list {
			dbname = strings.TrimSuffix(dbname, "/")
			if dbname == "general" {
				continue
			}
			if !Map_contains(pg.DB, dbname) {
				original_db_name := dbname
				if len(db_list) == 1 { //set default publish[main]
					dbname = "main"
				}
				//read config number of node on vault
				pg.NumNode[dbname] = 1
				num_node_str := pg.Config.ReadVAR(servicePath + "/db/postgres/" + dbname + "/config/NUM_NODE")
				if num_node_str != "" {
					pg.NumNode[dbname] = utils.StringToInt(num_node_str)
					if pg.NumNode[dbname] < 0 {
						panic("Invalid NUM_NODE value")
					}
				}
				//init db map
				pg.DB[dbname] = make(map[int]*gorm.DB)
				//
				dbConfigMap["DB"] = original_db_name
				for i := 0; i < pg.NumNode[dbname]; i++ {
					node_path := servicePath + "/db/postgres/" + dbname + "/node-" + utils.IntToS(i)
					pg.DB[dbname][i] = &gorm.DB{}
					var err2 error
					pg.DB[dbname][i], err2 = GormInitByPath(pg.Config, nil, node_path, dbConfigMap)
					if err2 != nil {
						panic(err2)
					}
				}
				//
				if len(db_list) == 1 { //
					pg.DB["main"] = pg.DB[dbname]
				}
			}
		}
	} else { //single db
		pg.DB["main"][0] = &gorm.DB{}
		var err error
		pg.DB["main"][0], err = GormInit(pg.Config, nil)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func GormInitByPath(config *vault.Vault, gcf *gorm.Config, path string, global_cfg map[string]string) (db *gorm.DB, err error) {
	log.Info("Initialing Postgres Gorm DB...")
	if config == nil {
		return nil, errors.New("Vault Config is nil")
	}
	local_config := pgcf.GetConfig(config, path)
	cfg_map := pgcf.MergeConfig(global_cfg, local_config)
	//get db info
	host := cfg_map["HOST"]
	user := cfg_map["USERNAME"]
	pass := cfg_map["PASSWORD"]
	dbname := cfg_map["DB"]
	//enableAPM:=false
	//if cfg_map["ENABLE_APM"]=="true"{enableAPM=true}
	//other info
	timezone := "UTC"
	ssl_mode := "disable"
	port := "5432"
	if cfg_map["PORT"] != "" {
		port = cfg_map["PORT"]
	}
	if cfg_map["TIME_ZONE"] != "" {
		timezone = cfg_map["TIME_ZONE"]
	}
	if cfg_map["USE_SSL"] == "true" {
		ssl_mode = "enable"
	}
	//log level default based on ENV
	db_log_level := "error"
	env := os.Getenv("ENV")
	if env == "local" {
		db_log_level = "info"
	}

	logLevel := logger.Silent
	if db_log_level == "info" {
		logLevel = logger.Info
	} else if db_log_level == "warn" {
		logLevel = logger.Warn
	} else if db_log_level == "error" {
		logLevel = logger.Error
	}
	//dns
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, pass, dbname, port, ssl_mode, timezone)
	//gorm log
	logger := zapgorm2.New(zap.L())
	logger.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks
	logger.LogLevel = logLevel
	//default gorm config
	gorm_cfg := &gorm.Config{}
	//overwrite gorm config from external
	if gcf != nil {
		gorm_cfg = gcf
	}
	gorm_cfg.Logger = logger
	gorm_cfg.NamingStrategy = schema.NamingStrategy{
		SingularTable: true,
	}
	gormDB, err := gorm.Open(postgres.Open(dsn), gorm_cfg)
	if err != nil {
		panic("failed to connect database")
		return nil, err
	}
	fmt.Println("Pg DB Connected: ", host, dbname)
	return gormDB, nil

}
func Map_contains(m map[string]map[int]*gorm.DB, item string) bool {
	if len(m) == 0 {
		return false
	}
	if _, ok := m[item]; ok {
		return true
	}
	return false
}

func GormInit(config *vault.Vault, gcf *gorm.Config) (db *gorm.DB, err error) {
	log.Info("Initialing Postgres Gorm DB...")
	if config == nil {
		return nil, errors.New("Vault Config is nil")
	}
	service_name := config.GetServiceName()
	service_config_path := strings.ReplaceAll(service_name, ".", "/")
	//golbal config
	global_db_config_path := fmt.Sprintf("%s/%s", service_config_path, "db/postgres")
	global_config_map := pgcf.GetConfig(config, global_db_config_path)
	//get db info
	host := global_config_map["HOST"]
	user := global_config_map["USERNAME"]
	pass := global_config_map["PASSWORD"]
	dbname := global_config_map["DB"]
	//other info
	timezone := "UTC"
	ssl_mode := "disable"
	port := "5432"
	if global_config_map["PORT"] != "" {
		port = global_config_map["PORT"]
	}
	if global_config_map["TIME_ZONE"] != "" {
		timezone = global_config_map["TIME_ZONE"]
	}
	if global_config_map["USE_SSL"] == "true" {
		ssl_mode = "enable"
	}
	//log level default based on ENV
	db_log_level := "error"
	env := os.Getenv("ENV")
	if env == "local" {
		db_log_level = "info"
	}
	//orverite log level from Vault config
	if global_config_map["GORM_LOG_LEVEL"] != "" {
		db_log_level = global_config_map["GORM_LOG_LEVEL"]
	}
	logLevel := logger.Silent
	if db_log_level == "info" {
		logLevel = logger.Info
	} else if db_log_level == "warn" {
		logLevel = logger.Warn
	} else if db_log_level == "error" {
		logLevel = logger.Error
	}

	//dns
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, pass, dbname, port, ssl_mode, timezone)
	//gorm log
	logger := zapgorm2.New(zap.L())
	logger.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks
	logger.LogLevel = logLevel
	//default gorm config
	gorm_cfg := &gorm.Config{}
	//overwrite gorm config from external
	if gcf != nil {
		gorm_cfg = gcf
	}
	gorm_cfg.Logger = logger
	gorm_cfg.NamingStrategy = schema.NamingStrategy{
		SingularTable: true,
	}
	gormDB, err := gorm.Open(postgres.Open(dsn), gorm_cfg)
	if err != nil {
		panic("failed to connect database")
		return nil, err
	}
	return gormDB, nil

}

// GetDB get db by dbname and node index
func (pg *PgDB) GetDB(dbname string, args ...int) (*gorm.DB, error) {
	//check dbname exist
	node_index := 0
	if len(args) > 0 {
		node_index = args[0]
	}
	if _, ok := pg.DB[dbname]; ok {
		if _, ok := pg.DB[dbname][node_index]; ok {
			return pg.DB[dbname][node_index], nil
		} else {
			return nil, errors.New("Node index not found")
		}
	} else {
		return nil, errors.New("DB name not found")
	}
	return nil, nil
}
