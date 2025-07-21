package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Agent struct {
	XEmail     string `json:"x_email"`
	XUserName  string `json:"x_username"`
	XPassword  string `json:"x_password"`
	X2FASecret string `json:"x_2fa_secret"`
}

// RAgent represents a Rivalz Agent node in the system
type RAgent struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	WalletAddress    string             `json:"wallet_address" bson:"wallet_address"`
	OldWalletAddress string             `json:"old_wallet_address" bson:"old_wallet_address"`
	DeviceID         string             `json:"device_id" bson:"device_id"`
	Type             string             `json:"type" bson:"type"`
	OS               string             `json:"os" bson:"os"`
	Version          string             `json:"version" bson:"version"`
	LastPingTime     time.Time          `json:"last_ping_time" bson:"last_ping_time"`
	CreatedAt        time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at" bson:"updated_at"`
	Status           string             `json:"status" bson:"status"`
	Code             string             `json:"code" bson:"code"`
	CodeGeneratedAt  time.Time          `json:"code_generated_at" bson:"code_generated_at"`
}

// RAgentRequest is used for agent registration
type RAgentRequest struct {
	WalletAddress string `json:"wallet_address" binding:"required"`
	DeviceID      string `json:"device_id" binding:"required"`
	Type          string `json:"type" binding:"required"`
	OS            string `json:"os" binding:"required"`
	Version       string `json:"version" binding:"required"`
}

// RAgentPingRequest is used for agent ping updates
type RAgentPingRequest struct {
	WalletAddress string `json:"wallet_address" binding:"required"`
	DeviceID      string `json:"device_id" binding:"required"`
}

// RAgentUpdateWalletRequest is used for updating agent wallet address
type RAgentUpdateWalletRequest struct {
	OldWalletAddress string `json:"old_wallet_address" binding:"required"`
	NewWalletAddress string `json:"new_wallet_address" binding:"required"`
	DeviceID         string `json:"device_id" binding:"required"`
	Code             string `json:"code" binding:"required"`
}

type Job struct {
	ID          primitive.ObjectID     `json:"_id" bson:"_id,omitempty"`
	Type        int                    `json:"type" bson:"type"`
	SessionId   string                 `json:"session_id" bson:"session_id"`
	Data        map[string]interface{} `json:"data" bson:"data"`
	Interval    int                    `json:"interval" bson:"interval"` //second
	CreatedTime *time.Time             `json:"created_time" bson:"created_time"`
	Status      int                    `json:"status" bson:"status"`
	ProjectId   string                 `json:"project_id" bson:"project_id"`
}

/*
1. call RestFull AP, get URL job

	Data: {
		"urls": "https://url1.com,https://url2.com",
	}
*/
type JobInstanceResult struct {
	JobInstanceId string                 `json:"job_instance_id" bson:"job_instance_id"`
	JobId         string                 `json:"job_id" bson:"job_id"`
	SessionId     string                 `json:"session_id" bson:"session_id"`
	Key           string                 `json:"key" bson:"key"`
	Status        string                 `json:"status" bson:"status"`
	Result        string                 `json:"result" bson:"result"`
	ResultObj     map[string]interface{} `json:"result_obj" bson:"result_obj"`
	AgentType     string                 `json:"agent_type" bson:"agent_type"`
	Name          string                 `json:"name" bson:"name"`
	UpdatedTime   *time.Time             `json:"updated_time" bson:"updated_time"`
}
type XuserProject struct {
	ProjectId    string `json:"project_id" bson:"project_id"`
	XId          string `json:"x_id" bson:"x_id"`
	ClientId     string `json:"client_id" bson:"client_id"`
	ClientSecret string `json:"client_secret" bson:"client_secret"`
	XuserId      string `json:"xuser_id" bson:"xuser_id"`
}

/*
	{
	    "_id" : ObjectId("67e67514b466eccab178ae5d"),
	    "data" : {
	        "url" : "https://github.com"
	    },
	    "status" : "success",
	    "type" : NumberInt(0),
	    "created_time" : ISODate("2025-03-28T10:08:20.496+0000"),
	    "pickup_node" : "son-pc",
	    "interval" : NumberInt(60),
	    "session_id" : "427759d6-4e2e-469d-8980-340d3436ffa8",
	    "job_id" : ObjectId("67e67514b466eccab178ae5c"),
	    "updated_time" : ISODate("2025-03-28T10:51:07.094+0000"),
	    "counter" : NumberInt(5)
	}
*/
type JobInstance struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	JobId     string             `json:"job_id" bson:"job_id"`
	SessionId string             `json:"session_id" bson:"session_id"`
	Counter   int                `json:"counter" bson:"counter"`
}
type SwarmListResponse struct {
	Code    int         `json:"code" example:"0"`
	Message string      `json:"message"`
	Data    []SwarmItem `json:"data"`
}
type SwarmItem struct {
	ProjectID   string `json:"project_id" bson:"project_id"`
	ProjectName string `json:"project_name" bson:"project_name"`
	UniqueName  string `json:"unique_name" bson:"unique_name"`
	Logo        string `json:"logo" bson:"logo"`
	Status      string `json:"status" bson:"status"`
}

type Partner struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	ApiKey   string             `json:"api_key" bson:"api_key"`
	UserName string             `json:"username" bson:"username"`
	Active   bool               `json:"active" bson:"active"`
	Logo     string             `json:"logo" bson:"logo"`
	Agents   []PartnerAgent     `json:"agents" bson:"agents"`
}
type PartnerResponse struct {
	Code    int       `json:"code" example:"0"`
	Message string    `json:"message"`
	Data    []Partner `json:"data"`
}
type PartnerAgent struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	PartnerID string             `json:"partner_id" bson:"partner_id"`
	Title     string             `json:"title" bson:"title"`
	Type      string             `json:"type" bson:"type"`
	Name      string             `json:"name" bson:"name"`
	Active    bool               `json:"active" bson:"active"`
}
type PartnerModel struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title  string             `json:"title" bson:"title"`
	Name   string             `json:"name" bson:"name"`
	Active bool               `json:"active" bson:"active"`
}
type PartnerUserHistory struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	PartnerID string             `json:"partner_id" bson:"partner_id"`
	ThreadID  string             `json:"thread_id" bson:"thread_id"`
	AgentName string             `json:"agent_name" bson:"agent_name"`
	XId       string             `json:"x_id" bson:"x_id"`
	Text      string             `json:"text" bson:"text"`
	Result    string             `json:"result" bson:"result"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}
type PartnerUserHistoryResponse struct {
	Code    int                  `json:"code" example:"0"`
	Message string               `json:"message"`
	Data    []PartnerUserHistory `json:"data"`
}
type PartnerListResponse struct {
	Code    int       `json:"code" example:"0"`
	Message string    `json:"message"`
	Data    []Partner `json:"data"`
}
type PartnerAgentListResponse struct {
	Code    int            `json:"code" example:"0"`
	Message string         `json:"message"`
	Data    []PartnerAgent `json:"data"`
}
type PartnerModelListResponse struct {
	Code    int            `json:"code" example:"0"`
	Message string         `json:"message"`
	Data    []PartnerModel `json:"data"`
}
type PartnerAgentTaskResponse struct {
	Code    int    `json:"code" example:"0"`
	Message string `json:"message"`
}
type AgentTaskRequest struct {
	XId      string `json:"x_id"`
	Text     string `json:"text"`
	ThreadID string `json:"thread_id"`
	Args     *Args  `json:"args"`
}
type Args struct {
	Model string `json:"model"`
}
type AgentTaskResp struct {
	Result string `json:"result"`
	Model  string `json:"model"`
}
type AgentTaskResponse struct {
	Code    int           `json:"code" example:"0"`
	Message string        `json:"message"`
	Data    AgentTaskResp `json:"data"`
}
type PartnerThread struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	XId       string             `json:"x_id" bson:"x_id"`
	AgentName string             `json:"agent_name" bson:"agent_name"`
	Title     string             `json:"title" bson:"title"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

type ThreadRequest struct {
	AgentName string `json:"agent_name"`
	XId       string `json:"x_id"`
}
type ThreadResponse struct {
	Code    int           `json:"code" example:"0"`
	Message string        `json:"message"`
	Data    PartnerThread `json:"data"`
}
type RXCountResponse struct {
	Code    int    `json:"code" example:"0"`
	Message string `json:"message"`
	Data    struct {
		Total     int `json:"total"`
		Available int `json:"available"`
	} `json:"data"`
}
type UserTypeResponse struct {
	Code    int    `json:"code" example:"0"`
	Message string `json:"message"`
	Data    struct {
		Type string `json:"type" example:"UNAUTHENTICATED | USER | SWARM_OWNER"`
	} `json:"data"`
}
