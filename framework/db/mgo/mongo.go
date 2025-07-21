package mgo

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Rivalz-ai/framework-be/framework/config/vault"
	mgoCfg "github.com/Rivalz-ai/framework-be/framework/db/mgo/config"
	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/Rivalz-ai/framework-be/framework/utils"
	"go.elastic.co/apm/module/apmmongo/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MgoDB struct {
	Config  *vault.Vault
	DB      map[string]map[int]*mongo.Database
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

func (mgo *MgoDB) Initial(config *vault.Vault) error {
	if mgo.Config == nil {
		mgo.Config = config
	}
	serviceName := mgo.Config.GetServiceName()
	servicePath := strings.ReplaceAll(serviceName, ".", "/")
	check, err := mgo.Config.CheckPathExist(servicePath + "/db/mgo")
	if err != nil {
		panic("MongoDB Check Path Exist Error: " + err.Error())
	}
	dbConfigPath := fmt.Sprintf("%s/%s", servicePath, "db/mgo/general")
	dbConfigMap := mgoCfg.GetConfig(mgo.Config, dbConfigPath)
	//
	mgo.DB = make(map[string]map[int]*mongo.Database)
	mgo.NumNode = make(map[string]int)
	//
	if check { //multi db
		db_list := mgo.Config.ListItemByPath(servicePath + "/db/mgo")
		for _, dbname := range db_list {
			dbname = strings.TrimSuffix(dbname, "/")
			if dbname == "general" {
				continue
			}
			if !Map_contains(mgo.DB, dbname) {
				original_db_name := dbname
				if len(db_list) == 1 { //set default publish[main]
					dbname = "main"
				}
				//read config number of node on vault
				mgo.NumNode[dbname] = 1
				num_node_str := mgo.Config.ReadVAR(servicePath + "/db/mgo/" + dbname + "/config/NUM_NODE")
				if num_node_str != "" {
					mgo.NumNode[dbname] = utils.StringToInt(num_node_str)
					if mgo.NumNode[dbname] < 0 {
						panic("Invalid NUM_NODE value")
					}
				}
				//init db map
				mgo.DB[dbname] = make(map[int]*mongo.Database)
				//
				dbConfigMap["DB"] = original_db_name
				for i := 0; i < mgo.NumNode[dbname]; i++ {
					node_path := servicePath + "/db/mgo/" + dbname + "/node-" + utils.IntToS(i)
					mgo.DB[dbname][i] = &mongo.Database{}
					var err2 error
					mgo.DB[dbname][i], err2 = MongoInitByPath(mgo.Config, node_path, dbConfigMap)
					if err2 != nil {
						panic(err2)
					}
				}
				//
				if len(db_list) == 1 { //
					mgo.DB["main"] = mgo.DB[dbname]
				}
			}
		}
	} else { //single db
		mgo.DB["main"][0] = &mongo.Database{}
		var err error
		mgo.DB["main"][0], err = MongoInit(mgo.Config)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

// GetDB get db by dbname and node index
func (mgo *MgoDB) GetDB(dbname string, args ...int) (*mongo.Database, error) {
	//check dbname exist
	node_index := 0
	if len(args) > 0 {
		node_index = args[0]
	}
	if _, ok := mgo.DB[dbname]; ok {
		if _, ok := mgo.DB[dbname][node_index]; ok {
			return mgo.DB[dbname][node_index], nil
		} else {
			return nil, errors.New("Node index not found")
		}
	} else {
		return nil, errors.New("DB name not found")
	}
	return nil, nil
}
func MongoInitByPath(config *vault.Vault, path string, global_cfg map[string]string) (db *mongo.Database, err error) {
	log.Info("Initialing MongoDB...")
	if config == nil {
		return nil, errors.New("Vault Config is nil")
	}
	local_config := mgoCfg.GetConfig(config, path)
	cfg_map := mgoCfg.MergeConfig(global_cfg, local_config)
	//get db info
	host := cfg_map["HOST"]
	user := cfg_map["USERNAME"]
	pass := cfg_map["PASSWORD"]
	dbname := cfg_map["DB"]
	authdb := cfg_map["AUTHDB"]
	authenType := "SCRAM-SHA-256"
	//apm
	enableApm := false
	if cfg_map["ENABLE_APM"] == "true" {
		enableApm = true
	}
	//other info
	timezone := "UTC"
	_ = timezone
	ssl_mode := "disable"
	_ = ssl_mode
	port := "27017"
	_ = port
	if cfg_map["PORT"] != "" {
		port = cfg_map["PORT"]
	}
	if cfg_map["TIME_ZONE"] != "" {
		timezone = cfg_map["TIME_ZONE"]
	}
	if cfg_map["USE_SSL"] == "true" {
		ssl_mode = "enable"
	}
	// setup default config
	hostname, _ := os.Hostname()
	heartBeat := 15 * time.Second
	maxIdle := 180 * time.Second
	socketTimeout := 60 * time.Second
	connectTimeout := 60 * time.Second
	serverSelectTimeout := 60 * time.Second
	min := uint64(2)
	//retryWrites := false
	// setup options
	opt := &options.ClientOptions{
		AppName: &hostname,
		//RetryWrites: &retryWrites,
		Auth: &options.Credential{
			AuthMechanism: authenType, //c.Config.AuthMechanism,
			AuthSource:    authdb,
			Username:      user,
			Password:      pass,
		},
		ConnectTimeout:         &connectTimeout,
		HeartbeatInterval:      &heartBeat,
		MaxConnIdleTime:        &maxIdle,
		MinPoolSize:            &min,
		ServerSelectionTimeout: &serverSelectTimeout,
		SocketTimeout:          &socketTimeout,
	}

	if enableApm {
		opt.SetMonitor(apmmongo.CommandMonitor())
	}

	opt.ApplyURI(host)
	//default read secondary
	opt.ReadPreference = readpref.SecondaryPreferred()
	//
	//NewClient from mongo-driver not by sdk
	Client, err := mongo.NewClient(opt)
	if err != nil {
		panic("MongoDB Error: " + err.Error())
		return nil, err
	}
	err = Client.Connect(context.TODO())
	if err != nil {
		panic("MongoDB Error: " + err.Error())
		return nil, err
	}
	//
	err_db := Client.Ping(context.TODO(), nil)
	if err_db != nil {
		panic("MongoDB Connection Error: " + err_db.Error())
		return nil, err
	}
	database := Client.Database(dbname)
	fmt.Println("MongoDB Connected: ", host, dbname)
	return database, nil
}
func Map_contains(m map[string]map[int]*mongo.Database, item string) bool {
	if len(m) == 0 {
		return false
	}
	if _, ok := m[item]; ok {
		return true
	}
	return false
}

func MongoInit(config *vault.Vault) (db *mongo.Database, err error) {
	log.Info("Initialing MongoDB...")
	if config == nil {
		return nil, errors.New("Vault Config is nil")
	}
	service_name := config.GetServiceName()
	service_config_path := strings.ReplaceAll(service_name, ".", "/")
	//golbal config
	global_db_config_path := fmt.Sprintf("%s/%s", service_config_path, "db/mgo/general")
	global_config_map := mgoCfg.GetConfig(config, global_db_config_path)
	//get db info
	host := global_config_map["HOST"]
	user := global_config_map["USERNAME"]
	pass := global_config_map["PASSWORD"]
	dbname := global_config_map["DB"]
	authdb := global_config_map["AUTHDB"]
	authenType := "SCRAM-SHA-256"
	//apm
	enableApm := false
	if global_config_map["ENABLE_APM"] == "true" {
		enableApm = true
	}
	//
	//other info
	timezone := "UTC"
	_ = timezone
	ssl_mode := "disable"
	_ = ssl_mode
	port := "27017"
	_ = port
	if global_config_map["PORT"] != "" {
		port = global_config_map["PORT"]
	}
	if global_config_map["TIME_ZONE"] != "" {
		timezone = global_config_map["TIME_ZONE"]
	}
	if global_config_map["USE_SSL"] == "true" {
		ssl_mode = "enable"
	}
	// setup default config
	hostname, _ := os.Hostname()
	heartBeat := 15 * time.Second
	maxIdle := 180 * time.Second
	socketTimeout := 60 * time.Second
	connectTimeout := 60 * time.Second
	serverSelectTimeout := 60 * time.Second
	min := uint64(2)
	//retryWrites := false
	// setup options
	opt := &options.ClientOptions{
		AppName: &hostname,
		//RetryWrites: &retryWrites,
		Auth: &options.Credential{
			AuthMechanism: authenType, //c.Config.AuthMechanism,
			AuthSource:    authdb,
			Username:      user,
			Password:      pass,
		},
		ConnectTimeout:         &connectTimeout,
		HeartbeatInterval:      &heartBeat,
		MaxConnIdleTime:        &maxIdle,
		MinPoolSize:            &min,
		ServerSelectionTimeout: &serverSelectTimeout,
		SocketTimeout:          &socketTimeout,
	}

	if enableApm {
		opt.SetMonitor(apmmongo.CommandMonitor())
	}

	opt.ApplyURI(host)
	//default read secondary
	opt.ReadPreference = readpref.SecondaryPreferred()
	//
	//NewClient from mongo-driver not by sdk
	Client, err := mongo.NewClient(opt)
	if err != nil {
		panic("MongoDB Error: " + err.Error())
		return nil, err
	}
	err = Client.Connect(context.TODO())
	if err != nil {
		panic("MongoDB Error: " + err.Error())
		return nil, err
	}
	//
	err_db := Client.Ping(context.TODO(), nil)
	if err_db != nil {
		panic("MongoDB Connection Error: " + err_db.Error())
		return nil, err
	}
	database := Client.Database(dbname)
	fmt.Println("MongoDB Connected: ", host, dbname)
	return database, nil

}
