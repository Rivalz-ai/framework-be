package routes

import (
	"github.com/Rivalz-ai/framework-be/middleware"

	//log "github.com/sirupsen/logrus"
	"github.com/Rivalz-ai/framework-be/modules/hello/handler"
	"github.com/Rivalz-ai/framework-be/server"
)

func HelloRoutesInit(sv *server.Server, secretKey string) {
	//Init routes
	handler := handler.HelloHandler{}
	handler.Init(sv)
	hello := sv.SVC.Group("/hello")
	/////////public route
	hello.Get("/", handler.GetHello)
	// Protected routes
	hello.Use(middleware.JWTAuthMiddleware(secretKey))
	//the routes after this will be protected
	//hello.Post("/", handler.<FunctionName>)
}
