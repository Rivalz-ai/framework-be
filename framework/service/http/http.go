package http

import (
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/Rivalz-ai/framework-be/framework/apm/apmelk"
	"github.com/Rivalz-ai/framework-be/framework/base"
	"github.com/Rivalz-ai/framework-be/framework/config/vault"
	"github.com/Rivalz-ai/framework-be/framework/db/mgo"
	"github.com/Rivalz-ai/framework-be/framework/db/pg"
	"github.com/Rivalz-ai/framework-be/framework/dbmem/redis"
	"github.com/Rivalz-ai/framework-be/framework/log"
	vault_cfg "github.com/Rivalz-ai/framework-be/framework/pubsub/config"
	"github.com/Rivalz-ai/framework-be/framework/pubsub/kafka"
	srv_cfg "github.com/Rivalz-ai/framework-be/framework/service"
	"github.com/Rivalz-ai/framework-be/framework/utils"
	"github.com/gofiber/fiber/v2"
	"go.elastic.co/apm/module/apmfiber/v2"
	//"github.com/gofiber/contrib/swagger"
)

type Option struct {
	EnablePg    bool
	EnableES    bool
	EnableRedis bool
	EnablePub   bool
	EnableAPM   bool
	EnableMongo bool
}
type HTTP struct {
	Id string
	//
	env string
	//
	serviceName string
	//
	vaultServicePath string
	apmAgent         *apmelk.APMAgent
	//
	option Option
	//http/config
	generalConfiPath string
	//http general config
	host         string
	port         string
	jwtSecretKey string
	//key-value store management
	Config *vault.Vault
	//http server
	SVC *fiber.App
	//PostgresDB
	Pg *pg.PgDB
	//MongoDB
	Mgo *mgo.MgoDB
	//Publisher
	Pub map[string]*kafka.Publisher
	//Redis
	Redis *redis.Redis
}

func (h *HTTP) Init(service_name string, args ...interface{}) {
	//set service name
	h.serviceName = service_name
	//get ENV
	base.LoadENV()
	//init log
	log.Initial(service_name)
	//load base config
	h.host = os.Getenv("SERVICE_HOST")
	h.port = os.Getenv("SERVICE_PORT")
	h.env = os.Getenv("ENV")
	if h.env == "" {
		h.env = "local"
	}
	//key store config
	var config vault.Vault
	h.Config = &config
	h.Config.Initial(service_name)
	//set the service Id
	hostname, err := os.Hostname()
	//
	if err != nil {
		log.Warn("Can not get Hostname :"+err.Error(), "HOST_NAME")
		h.Id = h.Config.GetServiceName()
	} else {
		h.Id = fmt.Sprintf("%s-%s", h.Config.GetServiceName(), hostname)
	}
	//set the vault service path
	h.vaultServicePath = strings.ReplaceAll(h.Config.GetServiceName(), ".", "/")
	//overwrite htto host + port from Vault
	httpHostVault := h.Config.ReadVAR("http/config/SERVICE_HOST")
	httpPortVault := h.Config.ReadVAR("http/config/SERVICE_PORT")
	if httpHostVault != "" {
		h.host = httpHostVault
	}
	if httpPortVault != "" {
		h.port = httpPortVault
	}
	//owervrite logger
	logDestPathConfig := "logger/general/LOG_DEST"
	if log.LogMode() != 2 { // not in local, local just output log to std
		logDest := h.Config.ReadVAR(logDestPathConfig)
		if logDest == "kafka" {
			configMap := vault_cfg.GetConfig(h.Config, "logger/kafka")
			log.SetDestKafka(configMap)
		}
	}
	//
	cfg := srv_cfg.ArgsToMapConfig(args)
	if cfg != nil {
		//initial publisher
		if cfg["pub"] != nil {
			if utils.ItoBoolDefault(cfg["pub"]) {
				h.option.EnablePub = true
			}
		}
		if cfg["redis"] != nil {
			if utils.ItoBoolDefault(cfg["redis"]) {
				h.option.EnableRedis = true
			}
		}
		if cfg["elasticsearch"] != nil {
			if utils.ItoBoolDefault(cfg["elasticsearch"]) {
				h.option.EnableES = true
			}
		}
		if cfg["postgres"] != nil {
			if utils.ItoBoolDefault(cfg["postgres"]) {
				h.option.EnablePg = true
			}
		}
		if cfg["mongo"] != nil {
			if utils.ItoBoolDefault(cfg["mongo"]) {
				h.option.EnableMongo = true
			}
		}
		if cfg["redis"] != nil {
			if utils.ItoBoolDefault(cfg["redis"]) {
				h.option.EnableRedis = true
			}
		}
		if cfg["apm"] != nil {
			if utils.ItoBoolDefault(cfg["apm"]) {
				h.option.EnableAPM = true
			}
		}
	}
	//init apm client (not init for http server)
	if h.option.EnableAPM {
		h.apmAgent = &apmelk.APMAgent{}
		h.apmAgent.Initial(h.Config)
	}
	//init PostgresDB
	//buộc phải set enable trong code, ko thể chỉ dựa vào path config của postgres trên vault vì ko được tường mình
	if h.option.EnablePg {
		h.Pg = &pg.PgDB{}
		h.Pg.Initial(h.Config)
	}
	//init mongodb
	if h.option.EnableMongo {
		h.Mgo = &mgo.MgoDB{}
		h.Mgo.Initial(h.Config)
	}
	//init publisher
	if h.option.EnablePub {
		h.InitPublisher()
	}
	//init redis
	if h.option.EnableRedis {
		h.Redis = &redis.Redis{}
		h.Redis.Initial(h.Config)
	}
	//JETsecret
	jwtPath := "http/config/JWT_SECRET"
	jwtSecretKey := h.Config.ReadVAR(jwtPath)
	if jwtSecretKey == "" {
		panic("JWT_SECRET is required")
	}
	//
	enablePrintRoutes := false
	if h.env == "local" {
		enablePrintRoutes = true
	}
	//
	h.SVC = fiber.New(fiber.Config{
		//Prefork:       true,
		CaseSensitive: true,
		//StrictRouting: true,//When enabled, the router treats /foo and /foo/ as different. Otherwise, the router treats /foo and /foo/ as the same.
		ServerHeader:      "Rivalz-AI",
		AppName:           "Rivalz-AI v1.0.0",
		EnablePrintRoutes: enablePrintRoutes,
		BodyLimit:         10 * 1024 * 1024, // 10MB
		/*
			Prefork: Enables use of theSO_REUSEPORTsocket option. This will spawn multiple Go processes listening on the same port. learn more about socket sharding. NOTE: if enabled, the application will need to be ran through a shell because prefork mode sets environment variables. If you're using Docker, make sure the app is ran with CMD ./app or CMD ["sh", "-c", "/app"]. For more info, see this issue comment.
		*/
	})
	//APM for Fiber http init
	if h.option.EnableAPM {
		h.SVC.Use(apmfiber.Middleware())
	}
}
func (h *HTTP) Start() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println(h.serviceName + " HTTP Gracefully shutting down...")
		_ = h.SVC.Shutdown()
	}()

	if err := h.SVC.Listen(":" + h.port); err != nil {
		panic(err)
	}

	fmt.Println("Running cleanup tasks...")
	// Your cleanup tasks go here
}
func (h *HTTP) InitPublisher() {
	fmt.Println("Initialing Publisher")
	//find publisher list
	check, err := h.Config.CheckPathExist(h.vaultServicePath + "/pub/kafka")
	if err != nil {
		panic("Publisher Check Path Exist Error: " + err.Error())
	}
	//
	h.Pub = make(map[string]*kafka.Publisher)
	if check { //custom publisher, list event
		event_list := h.Config.ListItemByPath(h.vaultServicePath + "/pub/kafka")
		config_path := h.vaultServicePath + "/pub/kafka"
		for _, event := range event_list {
			if !Map_PublisherContains(h.Pub, event) && event != "general" {
				original_topic := event   //event=topic name
				if len(event_list) == 1 { //set default publish[main]
					event = "main"
				}
				h.Pub[event] = &kafka.Publisher{}
				//topic name pass to publisher init, instead of get from vault in case single publiser
				pubCfg := &srv_cfg.PubConfig{
					TopicPath: original_topic, //topic name
				}
				err := h.Pub[event].Initial(h.Config, config_path, pubCfg)
				if err != nil {
					panic(err)
				}
			}
		}
	} else { //
		path := h.vaultServicePath + "/pub/kafka"
		check, err := h.Config.CheckItemExist(path)
		if err != nil {
			panic("Publisher config Check Item Exist Error: " + err.Error())
		}
		if !check {
			panic("Publisher config not found")
		}
		h.Pub["main"] = &kafka.Publisher{}
		err_p := h.Pub["main"].Initial(h.Config, path)
		if err_p != nil {
			panic(err_p)
		}
	}
}
func Map_PublisherContains(m map[string]*kafka.Publisher, item string) bool {
	if len(m) == 0 {
		return false
	}
	if _, ok := m[item]; ok {
		return true
	}
	return false
}
