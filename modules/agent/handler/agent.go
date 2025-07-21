package handler

import (
	"github.com/Rivalz-ai/framework-be/modules/agent/dto"
	agentService "github.com/Rivalz-ai/framework-be/modules/agent/service"
	"github.com/Rivalz-ai/framework-be/server"
	"github.com/gofiber/fiber/v2"

	//"github.com/Rivalz-ai/framework-be/framework/log"
	"fmt"
	//"time"
	//"github.com/Rivalz-ai/framework-be/framework/apm/apmelk"
	"strings"

	"github.com/Rivalz-ai/framework-be/framework/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AgentHandler struct {
	sv *server.Server
}

func (s *AgentHandler) Init(sv *server.Server) {
	s.sv = sv
}

type XRefreshKoken struct {
	ID    primitive.ObjectID `bson:"_id"`
	Token string             `json:"token" bson:"token"`
}

var secretKeyAuthen = ""

type XUserToken struct {
	XId          string `json:"x_id"`
	RefreshToken string `json:"refresh_token"`
	Reputation   int    `json:"reputation"`
}

// APIResponse định nghĩa response chuẩn
type APIResponse struct {
	Code    int          `json:"code" example:"0"`
	Data    []XUserToken `json:"data"`
	Message string       `json:"message"`
}
type APIResponseSuccess struct {
	Code    int    `json:"code" example:"0"`
	Message string `json:"message"`
}
type APIResponseError struct {
	Code    int    `json:"code" example:"1"`
	Message string `json:"message"`
}

/*
"total_tasks": total_count,
"done": done_count,
"list_result_done": result_links,
"pending": pending_count,
"failed": failed_count,
"queue_size": queue_size,
"completion_percentage": completion_percentage
list_failed
*/
type JobStatusResponse struct {
	TotalTasks           int           `json:"total_tasks"`
	Done                 int           `json:"done"`
	Pending              int           `json:"pending"`
	Failed               int           `json:"failed"`
	QueueSize            int           `json:"queue_size"`
	CompletionPercentage int           `json:"completion_percentage"`
	ListResultDone       []interface{} `json:"list_result_done"`
	ListFailed           []interface{} `json:"list_failed"`
}

func (s *AgentHandler) GetSwarmInfo(c *fiber.Ctx) error {
	key := c.Query("authen_key")
	if key != secretKeyAuthen {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Authen Key"})
	}
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	project_id := c.Query("project_id")
	if project_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Project ID"})
	}
	xTokens, err := agentSV.GetSwarm(ctx, "rx", project_id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 0, "data": xTokens})
}

// @Router			/agent/rx [get]
func (s *AgentHandler) GetRX(c *fiber.Ctx) error {
	key := c.Query("authen_key")
	if key != secretKeyAuthen {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Authen Key"})
	}
	num_str := c.Query("num")
	if num_str == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Num"})
	}
	num := utils.ItoInt(num_str)
	if num <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Num"})
	}
	project_id := c.Query("project_id")
	if project_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Project ID"})
	}
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	xusers, err := agentSV.GetRX(ctx, num, project_id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 0, "data": xusers})
}
func (s *AgentHandler) GetRAgent(c *fiber.Ctx) error {
	key := c.Query("authen_key")
	if key != secretKeyAuthen {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Authen Key"})
	}
	num_str := c.Query("num")
	if num_str == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Num"})
	}
	num := utils.ItoInt(num_str)
	if num <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Num"})
	}
	type_p := c.Query("type")
	if type_p == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Type"})
	}
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	xusers, err := agentSV.GetRAgent(ctx, num, type_p)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 0, "data": xusers})
}
func (s *AgentHandler) GetJob(c *fiber.Ctx) error {
	key := c.Query("authen_key")
	if key != secretKeyAuthen {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Authen Key"})
	}
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	node_id := c.Query("node_id")
	if node_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Node ID"})
	}
	result, err := agentSV.GetJob(ctx, node_id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 0, "data": result})
}
func (s *AgentHandler) AddJob(c *fiber.Ctx) error {
	key := c.Query("authen_key")
	if key != secretKeyAuthen {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Authen Key"})
	}
	agent_type := c.Query("agent_type")
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	var request dto.Job
	err = c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	//valida data
	if request.Type < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Type"})
	}
	//
	if request.SessionId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Session ID"})
	}
	if request.ProjectId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Project ID"})
	}
	if len(request.Data) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Data"})
	}
	//check session id

	result, err := agentSV.AddJob(ctx, request.Type, request.Interval, request.SessionId, request.Data, request.ProjectId, agent_type)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 0, "data": result})
}
func (s *AgentHandler) DeleteJob(c *fiber.Ctx) error {
	key := c.Query("authen_key")
	if key != secretKeyAuthen {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Authen Key"})
	}
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	jobID := c.Query("job_id")
	result, err := agentSV.DeleteJob(ctx, jobID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 0, "data": result})
}
func (s *AgentHandler) AddJobInstanceResult(c *fiber.Ctx) error {
	key := c.Query("authen_key")
	type_p := c.Query("type")
	if key != secretKeyAuthen {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Authen Key"})
	}
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	var request dto.JobInstanceResult
	err = c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	//valida data
	if request.JobInstanceId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Job Instance ID"})
	}
	if request.JobId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Job ID"})
	}
	if request.Status == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Status"})
	}
	_type := "rx"
	if type_p != "" {
		_type = type_p
	}
	if _type != "rx" && _type != "rc" && _type != "re" && _type != "rd" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Type"})
	}
	//
	result, err := agentSV.AddJobInstanceResult(ctx, request, _type)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 0, "data": result})
}

// @Summary	Get Task status by session id
// @Description Get Task status by session id
// @Tags			agent
// @Accept			json
// @Produce		json
// @Param			type		path		string				true	"Type"
// @Param			thread_id		query		string				true	"Session ID"
// @Success		200			{object}	JobStatusResponse			"Success response status of task"
// @Failure		400			{object}	APIResponseError			"Bad Request Error"
// @Router		/agent/task/:type/stats [get]
func (s *AgentHandler) GetJobStatusBySessionID(c *fiber.Ctx) error {
	session_id := c.Query("thread_id")
	if session_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid session_id"})
	}
	_type := c.Params("type")
	if _type == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid type"})
	}
	//_type only support "rx","rc","re","rd"
	if _type != "rx" && _type != "rc" && _type != "re" && _type != "rd" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid type"})
	}
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	result, err := agentSV.GetJobInstanceResult(ctx, session_id, _type)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 0, "data": result})
}

func (s *AgentHandler) RegisterRAgent(c *fiber.Ctx) error {
	// Validate authentication key
	key := c.Query("authen_key")
	if key != secretKeyAuthen {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Authen Key"})
	}

	// Parse request body
	var request dto.RAgentRequest
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "Invalid request: " + err.Error()})
	}
	//lower wallet address
	request.WalletAddress = strings.ToLower(request.WalletAddress)

	// Validate required fields
	if request.WalletAddress == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "Wallet address is required"})
	}

	if request.DeviceID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "Device ID is required"})
	}
	request.DeviceID = strings.ToLower(request.DeviceID)

	// Validate agent type (optional - add allowed types if needed)
	validTypes := []string{"RD", "RC", "RX", "RE"}
	typeValid := false

	if request.Type != "" {
		for _, t := range validTypes {
			if request.Type == t {
				typeValid = true
				break
			}
		}
	} else {
		request.Type = "*" //all type
		typeValid = true
	}

	if !typeValid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "Invalid agent type. Allowed types: RD, RC, RX, RE, * "})
	}

	// Get service and register agent
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	agent, err := agentSV.RegisterRAgent(ctx, request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"code": 0, "data": agent})
}

func (s *AgentHandler) PingRAgent(c *fiber.Ctx) error {
	// Validate authentication key
	key := c.Query("authen_key")
	if key != secretKeyAuthen {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Authen Key"})
	}

	device_id := c.Query("device_id")
	if device_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid device_id"})
	}
	device_id = strings.ToLower(device_id)
	// Get service and ping agent
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	agent, err := agentSV.PingRAgent(ctx, device_id)
	if err != nil {
		if err.Error() == "Agent not registered" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"code": 1, "message": "Agent not found, please register first"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"code": 0, "data": agent})
}

func (s *AgentHandler) UpdateRAgentWallet(c *fiber.Ctx) error {
	// Validate authentication key
	key := c.Query("authen_key")
	if key != secretKeyAuthen {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Authen Key"})
	}

	// Parse request body
	var request dto.RAgentUpdateWalletRequest
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "Invalid request: " + err.Error()})
	}
	fmt.Printf("request: %+v\n", request)
	// Validate required fields
	if request.OldWalletAddress == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "Old wallet address is required"})
	}

	if request.NewWalletAddress == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "New wallet address is required"})
	}

	if request.DeviceID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "Device ID is required"})
	}

	if request.Code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "Verification code is required"})
	}

	// Ensure old and new wallet addresses are different
	if request.OldWalletAddress == request.NewWalletAddress {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "New wallet address must be different from the old one"})
	}
	request.DeviceID = strings.ToLower(request.DeviceID)
	request.OldWalletAddress = strings.ToLower(request.OldWalletAddress)
	request.NewWalletAddress = strings.ToLower(request.NewWalletAddress)
	// Get service and update wallet
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	agent, err := agentSV.UpdateRAgentWallet(ctx, request.OldWalletAddress, request.NewWalletAddress, request.DeviceID, request.Code)
	if err != nil {
		if err.Error() == "Agent not found with the provided wallet address and device ID" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"code": 1, "message": err.Error()})
		} else if err.Error() == "Invalid verification code" || err.Error() == "Verification code has expired. Please generate a new code" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"code": 0, "data": agent})
}

func (s *AgentHandler) GetRAgentVersion(c *fiber.Ctx) error {
	// Validate authentication key
	key := c.Query("authen_key")
	if key != secretKeyAuthen {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Authen Key"})
	}

	// Get service and retrieve version
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	versionInfo, err := agentSV.GetRAgentVersion(ctx)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"code": 0, "data": versionInfo})
}

func (s *AgentHandler) GetRAgentByDeviceID(c *fiber.Ctx) error {
	// Validate authentication key
	key := c.Query("authen_key")
	if key != secretKeyAuthen {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Authen Key"})
	}

	// Get device ID from query parameter
	deviceID := c.Query("device_id")
	if deviceID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "Device ID is required"})
	}
	deviceID = strings.ToLower(deviceID)

	// Get service and retrieve agent info
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	agent, err := agentSV.GetRAgentByDeviceID(ctx, deviceID)
	if err != nil {
		if err.Error() == "Agent not found with the provided device ID" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"code": 1, "message": err.Error()})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"code": 0, "data": agent})
}

func (s *AgentHandler) GenerateRAgentCode(c *fiber.Ctx) error {
	// Validate authentication key
	key := c.Query("authen_key")
	if key != secretKeyAuthen {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Authen Key"})
	}

	// Get device ID from query parameter
	deviceID := c.Query("device_id")
	if deviceID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "Device ID is required"})
	}
	deviceID = strings.ToLower(deviceID)

	// Get service and generate code
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	codeInfo, err := agentSV.GenerateRAgentCode(ctx, deviceID)
	if err != nil {
		if err.Error() == "Agent not found with the provided device ID" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"code": 1, "message": err.Error()})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"code": 0, "data": codeInfo})
}

func (s *AgentHandler) GetRXSwarmList(c *fiber.Ctx) error {
	// Validate authentication key
	key := c.Query("authen_key")
	if key != secretKeyAuthen {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Authen Key"})
	}

	// Parse pagination parameters
	pageStr := c.Query("page", "1")
	pageSizeStr := c.Query("page_size", "10")

	page := utils.ItoInt(pageStr)
	if page < 1 {
		page = 1
	}

	pageSize := utils.ItoInt(pageSizeStr)
	if pageSize < 1 {
		pageSize = 10
	}

	// Get service and retrieve RX swarm list
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	result, err := agentSV.GetRXSwarmList(ctx, page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"code": 0, "data": result})
}

// @Summary	Get list of swarms by user authen x_id
// @Description	Retrieve a list of swarms
// @Tags			agent
// @Accept			json
// @Produce		json
// @Param			x_id		query		string				true	"X ID"
// @Success		200			{object}	dto.SwarmListResponse			"Success Response with swarm list"
// @Failure		400			{object}	APIResponseError			"Bad Request Error"
// @Router			/agent/swarm/list [get]
func (s *AgentHandler) GetSwarmList(c *fiber.Ctx) error {
	x_id := c.Query("x_id")
	if x_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid X ID"})
	}
	// Get service and retrieve RX swarm list
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	result, err := agentSV.GetSwarmList(ctx, x_id)
	if err != nil {
		if err.Error() == "NOT_FOUND" {
			//return 404
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"code": 1, "message": "AUTHEN_FAILED"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	if len(result) == 0 {
		return c.JSON(fiber.Map{"code": 0, "message": "YOU_ARE_NOT_SWARM_OWNER", "data": result})
	}
	return c.JSON(fiber.Map{"code": 0, "data": result, "message": "SUCCESS"})
}

// @Summary	Get type of user authen x_id
// @Description	Retrieve a type of user
// @Tags			agent
// @Accept			json
// @Produce		json
// @Param			x_id		query		string				true	"X ID"
// @Success		200			{object}	dto.UserTypeResponse			"Success Response with user type"
// @Failure		400			{object}	APIResponseError			"Bad Request Error"
// @Router			/agent/user/type [get]
func (s *AgentHandler) GetUserType(c *fiber.Ctx) error {
	x_id := c.Query("x_id")
	if x_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid X ID"})
	}
	// Get service and retrieve RX swarm list
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	result, err := agentSV.GetSwarmList(ctx, x_id)
	if err != nil {
		if err.Error() == "NOT_FOUND" {
			//return 404
			return c.JSON(fiber.Map{"code": 0, "message": "AUTHEN_FAILED", "data": fiber.Map{"type": "UNAUTHENTICATED"}})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	if len(result) == 0 {
		return c.JSON(fiber.Map{"code": 0, "message": "YOU_ARE_NOT_SWARM_OWNER", "data": fiber.Map{"type": "USER"}})
	}
	return c.JSON(fiber.Map{"code": 0, "message": "SUCCESS", "data": fiber.Map{"type": "SWARM_OWNER"}})
}

// @Summary	Get partner user history
// @Description	Retrieve a list of partner user history
// @Tags			agent
// @Accept			json
// @Produce		json
// @Param			thread_id		query		string				true	"Thread ID"
// @Param			x_id			query		string				true	"X ID"
// @Param			page			query		string				true	"Page"
// @Param			page_size		query		string				true	"Page size"
// @Success		200			{object}	dto.PartnerUserHistoryResponse			"Success Response with partner user history"
// @Failure		400			{object}	APIResponseError			"Bad Request Error"
// @Router			/agent/partner/history [get]
func (s *AgentHandler) GetPartnerUserHistory(c *fiber.Ctx) error {
	thread_id := c.Query("thread_id")
	page_str := c.Query("page", "0")
	page_size_str := c.Query("page_size", "10")
	if thread_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "Thread ID is required"})
	}
	page := utils.ItoInt(page_str)
	if page < 0 {
		page = 0
	}
	page_size := utils.ItoInt(page_size_str)
	if page_size < 1 {
		page_size = 10
	}
	x_id := c.Query("x_id")
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	result, err := agentSV.GetPartnerUserHistory(ctx, thread_id, x_id, page, page_size)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 0, "data": result})
}

// @Summary	Get list of partners
// @Description	Retrieve a list of partners
// @Tags			agent
// @Accept			json
// @Produce		json
// @Param			x_id		query		string				true	"X ID"
// @Success		200			{object}	dto.PartnerListResponse			"Success Response with partner list"
// @Failure		400			{object}	APIResponseError			"Bad Request Error"
// @Router			/agent/partner/list [get]
func (s *AgentHandler) GetPartners(c *fiber.Ctx) error {
	x_id := c.Query("x_id")
	if x_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "X ID is required"})
	}
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	result, err := agentSV.GetPartnerList(ctx, x_id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 0, "data": result})
}

// @Summary	Get list of partner agents
// @Description	Retrieve a list of partner agents
// @Tags			agent
// @Accept			json
// @Produce		json
// @Param			partner_id		query		string				true	"Partner ID"
// @Success		200			{object}	dto.PartnerAgentListResponse			"Success Response with partner agent list"
// @Failure		400			{object}	APIResponseError			"Bad Request Error"
// @Router			/agent/partner/agents [get]
func (s *AgentHandler) GetPartnerAgents(c *fiber.Ctx) error {
	partner_id := c.Query("partner_id")
	x_id := c.Query("x_id")
	if partner_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "Partner ID is required"})
	}
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	result, err := agentSV.GetPartnerAgentList(ctx, partner_id, x_id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 0, "data": result})
}

// @Summary	Get list of partner models
// @Description	Retrieve a list of partner models
// @Tags			agent
// @Accept			json
// @Produce		json
// @Param			partner_id		query		string				true	"Partner ID"
// @Param			x_id			query		string				true	"X ID"
// @Success		200			{object}	dto.PartnerModelListResponse			"Success Response with partner model list"
// @Failure		400			{object}	APIResponseError			"Bad Request Error"
// @Router			/agent/partner/models [get]
func (s *AgentHandler) GetPartnerModels(c *fiber.Ctx) error {
	partner_id := c.Query("partner_id")
	if partner_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "Partner ID is required"})
	}
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	x_id := c.Query("x_id")
	result, err := agentSV.GetPartnerModelList(ctx, partner_id, x_id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 0, "data": result})
}

// @Summary	Create partner agent task
// @Description	Create a partner agent task
// @Tags			agent
// @Accept			json
// @Produce		json
// @Param		request	body	dto.AgentTaskRequest	true	"Agent Task Request"
// @Success		200			{object}	dto.PartnerAgentTaskResponse			"Success Response with partner agent task"
// @Failure		400			{object}	APIResponseError			"Bad Request Error"
// @Router			/agent/partner/task [post]
func (s *AgentHandler) CreatePartnerAgentTask(c *fiber.Ctx) error {
	request := new(dto.AgentTaskRequest)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	if request.XId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "X ID is required"})
	}
	if request.Text == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "Text is required"})
	}
	if request.ThreadID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "Thread ID is required"})
	}
	params := utils.Dictionary()
	params["text"] = request.Text
	params["model"] = request.Args.Model
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	result, err := agentSV.CreatePartnerAgentTask(ctx, request.XId, request.ThreadID, params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	if request.Args.Model == "" {
		request.Args.Model = "deepseek-ai/DeepSeek-R1"
	}
	response := dto.AgentTaskResponse{
		Code:    0,
		Message: "Success",
		Data: dto.AgentTaskResp{
			Result: result,
			Model:  request.Args.Model,
		},
	}
	return c.JSON(response)
}

// create thread
// @Summary	Create partner agent thread
// @Description	Create a partner agent thread
// @Tags			agent
// @Accept			json
// @Produce			json
// @Param			request	body	dto.ThreadRequest	true	"Thread Request"
// @Success		200			{object}	dto.ThreadResponse			"Success Response with partner agent thread"
// @Failure		400			{object}	APIResponseError			"Bad Request Error"
// @Router			/agent/partner/thread [post]
func (s *AgentHandler) CreatePartnerAgentThread(c *fiber.Ctx) error {
	request := new(dto.ThreadRequest)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	if request.XId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "X ID is required"})
	}
	if request.AgentName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "Agent name is required"})
	}
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	result, err := agentSV.CreatePartnerThread(ctx, request.XId, request.AgentName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 0, "data": fiber.Map{"thread_id": result}})
}

// @Summary	Get partner agent thread
// @Description	Get a partner agent thread
// @Tags			agent
// @Accept			json
// @Produce		json
// @Param			x_id		query		string				true	"X ID"
// @Param			agent_name		query		string				true	"Agent Name"
// @Param			page			query		string				true	"Page"
// @Param			page_size		query		string				true	"Page size"
// @Success		200			{object}	dto.ThreadResponse			"Success Response with partner agent thread"
// @Failure		400			{object}	APIResponseError			"Bad Request Error"
// @Router			/agent/partner/thread [get]
func (s *AgentHandler) GetPartnerAgentThread(c *fiber.Ctx) error {
	x_id := c.Query("x_id")
	agent_name := c.Query("agent_name")
	if x_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "X ID is required"})
	}
	if agent_name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "Agent name is required"})
	}
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	page_str := c.Query("page", "0")
	page_size_str := c.Query("page_size", "10")
	page := utils.ItoInt(page_str)
	if page < 0 {
		page = 0
	}
	page_size := utils.ItoInt(page_size_str)
	if page_size < 1 {
		page_size = 10
	}
	result, err := agentSV.GetPartnerThread(ctx, x_id, agent_name, page, page_size)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 0, "data": result})
}

// @Summary	Get RX count
// @Description	Get RX count
// @Tags			agent
// @Accept			json
// @Produce		json
// @Param			project_id		path		string				true	"Project ID"
// @Success		200			{object}	dto.RXCountResponse			"Success Response with RX count"
// @Failure		400			{object}	APIResponseError			"Bad Request Error"
// @Router			/agent/rx/count/{project_id} [get]
func (s *AgentHandler) GetRXCount(c *fiber.Ctx) error {
	project_id := c.Params("project_id")
	if project_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "Project ID is required"})
	}
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	total, available, err := agentSV.GetRXCount(ctx, project_id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 0, "data": fiber.Map{"total": total, "available": available}})
}

func (s *AgentHandler) GetTweetContent(c *fiber.Ctx) error {
	url := c.Query("url")
	if url == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "URL is required"})
	}
	ctx := c.Context()
	agentSV, err := agentService.NewAgentService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	fmt.Println(url)
	content, err := agentSV.GetTweetContent(ctx, url)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 0, "data": fiber.Map{"content": content}})
}
