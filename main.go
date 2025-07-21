package main

import (
	"github.com/Rivalz-ai/framework-be/framework/service"
	"github.com/Rivalz-ai/framework-be/routes"
	"github.com/Rivalz-ai/framework-be/server"
	//"github.com/Rivalz-ai/framework-be/framework/base/event"
	//"fmt"
)

//		@title			Rome API
//		@version		1.0
//		@description	Rome API
//		@termsOfService	http://swagger.io/terms/
//		@contact.name	API Support
//		@contact.url	http://www.swagger.io/support
//		@contact.email	support@swagger.io
//		@license.name	Apache 2.0
//		@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//		@BasePath	/api/v2
//	 	@schemes https http
func main() {
	var server server.Server
	mgo_cfg := &service.MongoConfig{
		Enable: true,
	}
	apm_cfg := &service.ApmConfig{
		Enable: true,
	}

	pg_cfg := &service.PostgresConfig{
		Enable: true,
	}
	pubCfg := &service.PubConfig{
		Enable: true,
	}
	redisCfg := &service.RedisConfig{
		Enable: true,
	}
	//
	//server.Init("http.rivalz-rome",apm_cfg,mgo_cfg,pubCfg)
	server.Init("http.rivalz-rome", apm_cfg, mgo_cfg, pg_cfg, pubCfg, redisCfg)
	server.LoadExtendConfig()
	routes.RouteInit(&server)
	/*
		ev:=event.Event{
			EventName:"test",
			EventHeader:map[string]string{
				"header_data":"header_data",
			},
			EventData:map[string]string{
				"body_data":"body_data",
			},
		}
		err:=server.Pub["task"].Publish(ev)
		fmt.Println(err)
		err=server.Pub["fe_log"].Publish(ev)
		fmt.Println(err)
	*/
	server.Start()
}
