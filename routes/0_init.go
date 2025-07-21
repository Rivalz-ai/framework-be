package routes

import (
	//""
	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/Rivalz-ai/framework-be/middleware"
	"github.com/Rivalz-ai/framework-be/server"

	//"github.com/swaggo/fiber-swagger"

	//"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	//"os"
	//"strings"
)

func RouteInit(sv *server.Server) {
	/*if sv==nil{
	    log.ErrorF("Server is nil")
	}*/
	//get secret ENV
	log.Info("Routes Initialization")
	secretKey := sv.Config.ReadVAR("http/config/JWT_SECRET")
	if secretKey == "" {
		log.ErrorF("Not found JWT_SECRET config")
	}
	middleware.AllowCORS(sv)
	//
	//sv.SVC.Get("/swagger/*",  swagger.HandlerDefault)
	sv.SVC.Get("/swagger/*", swagger.HandlerDefault)
	//hello routes
	HelloRoutesInit(sv, secretKey)
	//agent routes
	AgentRoutesInit(sv, secretKey)
}
