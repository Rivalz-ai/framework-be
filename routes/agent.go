package routes

import (
	"github.com/Rivalz-ai/framework-be/middleware"
	"github.com/Rivalz-ai/framework-be/modules/agent/handler"

	//log "github.com/sirupsen/logrus"
	"github.com/Rivalz-ai/framework-be/server"
)

func AgentRoutesInit(sv *server.Server, secretKey string) {
	//Init routes
	handler := handler.AgentHandler{}
	handler.Init(sv)
	agent := sv.SVC.Group("/agent")
	agentV2 := sv.SVC.Group("/api/v2/agent")
	agentV2.Get("/task/:type/stats", handler.GetJobStatusBySessionID)
	/////////public route
	agentV2.Get("/swarm/list", handler.GetSwarmList)
	agent.Get("/swarm", handler.GetSwarmInfo)
	agent.Get("/rx", handler.GetRX)
	agent.Get("/rx/swarm", handler.GetRXSwarmList)
	agent.Get("/ragent", handler.GetRAgent)
	agent.Get("/task", handler.GetJob)
	agent.Get("/tweet/content", handler.GetTweetContent)
	agent.Post("/task", handler.AddJob)
	agent.Delete("/task", handler.DeleteJob)
	agent.Post("/task/result", handler.AddJobInstanceResult)
	agent.Post("/ragent", handler.RegisterRAgent)
	agent.Get("/ragent/ping", handler.PingRAgent)
	agent.Put("/ragent/wallet/change", handler.UpdateRAgentWallet)
	agent.Get("/ragent/version", handler.GetRAgentVersion)
	agent.Get("/ragent/device", handler.GetRAgentByDeviceID)
	agent.Get("/ragent/wallet/code", handler.GenerateRAgentCode)
	//partner routes
	agentV2.Get("/partner/history", handler.GetPartnerUserHistory)
	agentV2.Get("/partner/list", handler.GetPartners)
	agentV2.Get("/partner/agents", handler.GetPartnerAgents)
	agentV2.Get("/partner/models", handler.GetPartnerModels)
	agentV2.Post("/partner/task", handler.CreatePartnerAgentTask)
	agentV2.Post("/partner/thread", handler.CreatePartnerAgentThread)
	agentV2.Get("/partner/thread", handler.GetPartnerAgentThread)
	agentV2.Get("/rx/count/:project_id", handler.GetRXCount)
	agentV2.Get("/user/type", handler.GetUserType)
	// Protected routes
	agent.Use(middleware.JWTAuthMiddleware(secretKey))
	//the routes after this will be protected
	//hello.Post("/", handler.<FunctionName>)

}
