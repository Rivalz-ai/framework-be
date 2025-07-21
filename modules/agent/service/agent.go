package service

import (
	//import node service
	"github.com/Rivalz-ai/framework-be/server"
	//ethereum type
	"errors"

	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/Rivalz-ai/framework-be/framework/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//"strings"
	"context"
	//"github.com/Rivalz-ai/framework-be/define"
	"github.com/Rivalz-ai/framework-be/modules/agent/dto"
	"github.com/Rivalz-ai/framework-be/modules/agent/service/http"
	userDto "github.com/Rivalz-ai/framework-be/modules/user/dto"
	xdto "github.com/Rivalz-ai/framework-be/modules/x/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"

	//userdto "github.com/Rivalz-ai/framework-be/modules/user/dto"
	"time"

	projectDto "github.com/Rivalz-ai/framework-be"
	xService "github.com/Rivalz-ai/framework-be/modules/x/service"
	"go.mongodb.org/mongo-driver/bson"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"fmt"
	"strings"

	//"time"
	//"golang.org/x/sync/errgroup"
	"math"
)

type AgentService struct {
	server                *server.Server
	colXuser              *mongo.Collection
	colUser               *mongo.Collection
	colJob                *mongo.Collection
	colJobInstance        *mongo.Collection
	colJobInstalceResult  *mongo.Collection
	colXusersProject      *mongo.Collection
	colProject            *mongo.Collection
	colRAgent             *mongo.Collection
	colRagentProject      *mongo.Collection
	colPartner            *mongo.Collection
	colPartnerAgent       *mongo.Collection
	colPartnerModel       *mongo.Collection
	colPartnerUserHistory *mongo.Collection
	colPartnerThread      *mongo.Collection
}

func NewAgentService(sv *server.Server) (*AgentService, error) {
	db, err := sv.Mgo.GetDB("rome")
	if err != nil {
		return nil, err
	}
	return &AgentService{
		server:               sv,
		colXuser:             db.Collection("xusers"),
		colUser:              db.Collection("users"),
		colJob:               db.Collection("jobs"),
		colJobInstance:       db.Collection("job_instances"),
		colJobInstalceResult: db.Collection("job_instance_results"),
		colXusersProject:     db.Collection("xusers_projects"),
		//colRagentProject:db.Collection("ragent_projects"),
		colProject:            db.Collection("projects"),
		colRAgent:             db.Collection("ragents"),
		colPartner:            db.Collection("partners"),
		colPartnerAgent:       db.Collection("partner_agents"),
		colPartnerModel:       db.Collection("partner_models"),
		colPartnerUserHistory: db.Collection("partner_user_histories"),
		colPartnerThread:      db.Collection("partner_threads"),
	}, nil
}
func (s *AgentService) GetSwarm(ctx context.Context, _type, project_id string) (map[string]interface{}, error) {
	//project_id:="67939f2c642d85777429a866" //rivalz => tương lai dựa vào authen_key tìm project id
	//query  users with xUserId not null get number user
	objectId, err := primitive.ObjectIDFromHex(project_id)
	if err != nil {
		log.Error("Error when convert project_id to objectID: "+err.Error(), "GetAgentDemo-GetAgent")
		return nil, err
	}
	var project projectDto.Project
	filter := bson.M{"_id": objectId, "status": "active"}
	err = s.colProject.FindOne(ctx, filter).Decode(&project)
	//check not found
	if err == mongo.ErrNoDocuments {
		log.Error("Project not found: "+err.Error(), "GetAgentDemo-GetAgent")
		return nil, err
	} else if err != nil {
		log.Error("Error when get project: "+err.Error(), "GetAgentDemo-GetAgent")
		return nil, err
	}
	filter = bson.M{"project_id": project_id}
	//find rd: count xuserprojects
	count, err := s.colXusersProject.CountDocuments(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			count = 0
		} else {
			log.Error("Error when count xuserprojects: "+err.Error(), "GetAgentDemo-GetAgent")
			return nil, err
		}
	}
	//find ragent: count from ragents filter type=*
	filter = bson.M{"type": "*"}
	count_ragent, err := s.colRAgent.CountDocuments(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			count_ragent = 0
		} else {
			log.Error("Error when count ragents: "+err.Error(), "GetAgentDemo-GetAgent")
			return nil, err
		}
	}

	/*
		cursor, err := s.colXusersProject.Find(ctx, filter)
		if err != nil {
			log.Error("Error when get xuser: "+err.Error(),"GetAgentDemo-GetAgent")
			return nil,err
		}
		var xuserprojects []*dto.XuserProject
		if err = cursor.All(ctx, &xuserprojects); err != nil {
			log.Error("Error when get xuser: "+err.Error(),"GetAgentDemo-GetAgent")
			return nil,err
		}
		fmt.Println("len xusersprojects: ",len(xuserprojects))
		//
		x_ids:=[]string{}
		for _,user:=range xuserprojects{
			x_ids=append(x_ids,user.XId)
		}
		fmt.Println("x_ids: ",x_ids)
		//
		xSV,err:=xService.NewXUserService(s.server)
		if err!=nil{
			log.Error("Error when get xservice: "+err.Error(),"GetAgentDemo-NewXService")
			return nil,err
		}

		xusers,err:=xSV.GetXUserByListXID(ctx,x_ids)
		if err!=nil{
			log.Error("error when get xuser by list id: "+err.Error(),"GetAgentDemo-GetXUserByListID",x_ids)
		}
		fmt.Println("len xusers: ",len(xusers))
		//
		var xusersToken []*xdto.XUserToken
		for _,xuser:=range xusers{
			xuserToken:=&xdto.XUserToken{
				XId:xuser.XId,
				FollowersCount: xuser.Followers,
				FollowingCount: xuser.Following,
				TweetCount: xuser.Tweets,
				LikeCount: xuser.Likes,
				MediaCount: xuser.Medias,
				ExamplePost: xuser.ExamplePost,
				StyleDescription: xuser.StyleDescription,
			}
			xusersToken=append(xusersToken,xuserToken)
		}
	*/
	response := map[string]interface{}{
		"info": project,
		"resources": map[string]interface{}{
			//"rx": xusersToken,
			"total": count + (count_ragent * 3),
			"rx":    count,
			"rc":    count_ragent,
			"rd":    count_ragent,
			"re":    count_ragent,
		},
	}
	return response, nil
}

func (s *AgentService) GetRXCount(ctx context.Context, project_id string) (int, int, error) {
	//count all
	filter := bson.M{
		"project_id": project_id,
	}
	count, err := s.colXusersProject.CountDocuments(ctx, filter)
	if err != nil {
		log.Error("Error when count xusers: "+err.Error(), "GetRXAvailable")
		return 0, 0, err
	}
	//count vailable => get item has updated_at before 1 week
	filter = bson.M{
		"project_id": project_id,
		"updated_at": bson.M{"$lt": time.Now().Add(-7 * 24 * time.Hour)},
	}
	available, err := s.colXusersProject.CountDocuments(ctx, filter)
	if err != nil {
		log.Error("Error when count xusers: "+err.Error(), "GetRXAvailable")
		return 0, 0, err
	}
	return int(count), int(available), nil
}

func (s *AgentService) GetRX(ctx context.Context, num int, project_id string) ([]*xdto.XUser, error) {
	//project_id:="67939f2c642d85777429a866" //rivalz => tương lai dựa vào authen_key tìm project id
	filter := bson.M{
		"project_id": project_id,
		"updated_at": bson.M{"$lt": time.Now().Add(-7 * 24 * time.Hour)},
	}
	opts := options.Find().
		SetSort(bson.D{
			{"updated_at", 1}, // Cũ nhất trước, để đảm bảo all X agent đều dc dùng, sau có thể cải tiến bằng PoH
			{"followers", -1}, // Sắp xếp giảm dần theo followers
			{"likes", -1},     // Tiếp theo là likes
			{"tweets", -1},    // Cuối cùng là tweets
		}).
		SetLimit(int64(num)) // Giới hạn số lượng kết quả

	cursor, err := s.colXusersProject.Find(ctx, filter, opts)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		log.Error("Error when get xusers: "+err.Error(), "GetRX")
		return nil, err
	}
	var xusersproject []*dto.XuserProject
	if err = cursor.All(ctx, &xusersproject); err != nil {
		log.Error("Error when get xusers: "+err.Error(), "GetRX")
		return nil, err
	}
	fmt.Println("len xusersproject: ", len(xusersproject))
	x_ids := []string{}
	for _, user := range xusersproject {
		x_ids = append(x_ids, user.XId)
	}
	if len(x_ids) == 0 {
		return nil, nil
	}
	fmt.Println("len x_ids: ", len(x_ids))
	//get xusers by list x_ids
	xSV, err := xService.NewXUserService(s.server)
	if err != nil {
		log.Error("Error when get xservice: "+err.Error(), "GetRX")
		return nil, err
	}
	fmt.Println("x_ids: ", x_ids)
	xusers, err := xSV.GetXUserByListXID(ctx, x_ids)
	if err != nil {
		log.Error("error when get xuser by list id: "+err.Error(), "GetRX", x_ids)
	}
	fmt.Println("len xusers: ", len(xusers))
	return xusers, nil

}

/*
	{
	    "_id" : ObjectId("67d286262d905e28337a01f5"),
	    "job_id" : ObjectId("67d286262d905e28337a01f4"),
	    "url" : "https://google.com",
	    "status" : "processing",
	    "type" : NumberInt(0),
	    "created_time" : ISODate("2025-03-13T07:15:50.978+0000"),
	    "pickup_node" : ""
		"updated_time": ISODate("2025-03-13T07:15:50.978+0000"),
	}
*/
func (s *AgentService) GetJob(ctx context.Context, node_id string) (map[string]interface{}, error) {
	filter := bson.M{
		"status": "processing",
		"type": bson.M{
			"$in": bson.A{0, 1, 2},
		},
		"$or": []bson.M{
			{"pickup_node": ""},
			{"pickup_node": bson.M{"$ne": ""}, "updated_time": bson.M{"$lt": time.Now().Add(-5 * time.Minute)}},
		},
	}
	update := bson.M{
		"$set": bson.M{
			"pickup_node":  node_id,
			"updated_time": time.Now().UTC(),
		},
	}
	var job bson.M
	err := s.colJobInstance.FindOneAndUpdate(ctx, filter, update).Decode(&job)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Error("Error when getting job: "+err.Error(), "GetJob", filter)
		return nil, err
	}
	return job, nil
}
func (s *AgentService) AddJob(ctx context.Context, _type, interval int, session_id string, data map[string]interface{}, project_id, agent_type string) (map[string]interface{}, error) {
	//project_id:="67939f2c642d85777429a866" //rivalz => tương lai dựa vào authen_key tìm project id
	job := bson.M{
		"type":       _type,
		"interval":   interval,
		"data":       data,
		"session_id": session_id,
	}
	result, err := s.colJob.InsertOne(ctx, job)
	if err != nil {
		log.Error("Error when adding job: "+err.Error(), "AddJob", job)
		return nil, err
	}

	switch _type {
	case 0: //get URL
		url := utils.ItoString(data["url"])
		if url == "" {
			log.Error("Invalid URL: "+url, "AddJob", job)
			return nil, errors.New("Invalid URL")
		}
		//insert many to collection job_instance
		jobInstance := bson.M{
			"job_id":       result.InsertedID,
			"data":         bson.M{"url": url},
			"status":       "processing",
			"type":         0,
			"created_time": time.Now().UTC(),
			"pickup_node":  "",
			"interval":     interval,
			"session_id":   session_id,
			"name":         "get-url",
			"agent_type":   "RD",
		}
		_, err = s.colJobInstance.InsertOne(ctx, jobInstance)
		if err != nil {
			log.Error("Error when adding job instance: "+err.Error(), "AddJob", job)
			return nil, err
		}
	case 1: //get node health
		jobInstance := bson.M{
			"job_id":       result.InsertedID,
			"status":       "processing",
			"type":         1,
			"created_time": time.Now().UTC(),
			"pickup_node":  "",
			"session_id":   session_id,
			"name":         "get-node-health",
			"agent_type":   "RC",
		}
		_, err = s.colJobInstance.InsertOne(ctx, jobInstance)
		if err != nil {
			log.Error("Error when adding job instance: "+err.Error(), "AddJob", job)
			return nil, err
		}
	case 2: //execute command on terminal
		//loại này hổ trợ cho  RC,RE,RD nên cần lấy agent type từ ngoài
		if agent_type == "" {
			agent_type = "RE"
		}
		cmd := utils.ItoString(data["cmd"])
		if cmd == "" {
			log.Error("Invalid cmd: "+cmd, "AddJob", job)
			return nil, errors.New("Invalid cmd")
		}
		jobInstance := bson.M{
			"job_id":       result.InsertedID,
			"status":       "processing",
			"type":         2,
			"data":         cmd,
			"created_time": time.Now().UTC(),
			"pickup_node":  "",
			"session_id":   session_id,
			"name":         "exec-command",
			"agent_type":   agent_type,
		}
		_, err = s.colJobInstance.InsertOne(ctx, jobInstance)
		if err != nil {
			log.Error("Error when adding job instance: "+err.Error(), "AddJob", job)
			return nil, err
		}
	case 3: //xagent
		post, err := utils.ItoDictionary(data)
		if err != nil {
			log.Error("Invalid Post for X ", "AddJob", job)
			return nil, errors.New("Invalid posts")
		}
		content := utils.ItoString(post["content"])
		x_id := utils.ItoString(post["x_id"])
		if content == "" || x_id == "" {
			log.Error("Invalid Post for X ", "AddJob", job)
			return nil, errors.New("Invalid posts")
		}
		jobInstance := bson.M{
			"job_id": result.InsertedID,
			"status": "processing",
			"type":   3,
			"data": map[string]interface{}{
				"post":       content,
				"x_id":       x_id,
				"project_id": project_id,
			},
			"created_time": time.Now().UTC(),
			"pickup_node":  "",
			"session_id":   session_id,
			"name":         "post-x",
			"agent_type":   "RX",
		}
		_, err = s.colJobInstance.InsertOne(ctx, jobInstance)
		if err != nil {
			log.Error("Error when adding job instance: "+err.Error(), "AddJob", job)
			return nil, err
		}
		//update update_at of xuserproject, this for set priority of this X lower, beacuse it was used
		filter := bson.M{"x_id": x_id, "project_id": project_id}
		update := bson.M{"$set": bson.M{"updated_at": time.Now().UTC()}}
		_, err = s.colXusersProject.UpdateOne(ctx, filter, update)
		if err != nil {
			log.Error("Error when updating xuserproject: "+err.Error(), "AddJob", job)
		}
	case 31, 32, 33: //X like post
		data, err := utils.ItoDictionary(data)
		if err != nil {
			log.Error("Invalid Like Post on X ", "AddJob", job)
			return nil, errors.New("Invalid posts")
		}
		post_url := utils.ItoString(data["post_url"])
		if post_url == "" {
			log.Error("Invalid Post URL: "+post_url, "AddJob", job)
			return nil, errors.New("Invalid post url")
		}
		arr := utils.Explode(post_url, "/")
		if len(arr) == 0 {
			log.Error("Invalid Post URL: "+post_url, "AddJob", job)
		}
		tweet_id := ""
		for i, v := range arr {
			if v == "status" {
				tweet_id = arr[i+1]
				break
			}
		}
		if tweet_id == "" {
			log.Error("Invalid Tweet ID: "+tweet_id, "AddJob", job)
			return nil, errors.New("Invalid tweet id")
		}
		x_id := utils.ItoString(data["x_id"])
		if x_id == "" {
			log.Error("Invalid Post for X ", "AddJob", job)
			return nil, errors.New("Invalid posts")
		}
		name := ""
		content := ""
		if _type == 31 {
			name = "like-post"
		} else if _type == 33 {
			name = "reply-post"
		} else if _type == 32 {
			name = "retweet-post"
		}
		if _type == 33 {
			content = utils.ItoString(data["content"])
			if content == "" {
				log.Error("Invalid Post for X ", "AddJob", job)
				return nil, errors.New("Invalid posts")
			}

		}
		jobInstance := bson.M{
			"job_id":       result.InsertedID,
			"status":       "processing",
			"type":         _type,
			"data":         map[string]interface{}{"tweet_id": tweet_id, "content": content, "x_id": x_id, "project_id": project_id},
			"created_time": time.Now().UTC(),
			"pickup_node":  "",
			"session_id":   session_id,
			"name":         name,
			"agent_type":   "RX",
		}
		_, err = s.colJobInstance.InsertOne(ctx, jobInstance)
		if err != nil {
			log.Error("Error when adding job instance: "+err.Error(), "AddJob", job)
			return nil, err
		}

	default:
		log.Error("Invalid job type: "+fmt.Sprint(_type), "AddJob", job)
		return nil, errors.New("Invalid job type")
	}
	data_resp := map[string]interface{}{
		"status": "processing",
		"job_id": result.InsertedID,
	}
	return data_resp, nil
}

func (s *AgentService) DeleteJob(ctx context.Context, jobID string) (*mongo.DeleteResult, error) {
	objectID, err := primitive.ObjectIDFromHex(jobID)
	if err != nil {
		log.Error("Invalid job ID: "+err.Error(), "DeleteJob", jobID)
		return nil, err
	}
	filter := bson.M{"_id": objectID}
	result, err := s.colJob.DeleteOne(ctx, filter)
	if err != nil {
		log.Error("Error when deleting job: "+err.Error(), "DeleteJob", jobID)
		return nil, err
	}
	return result, nil
}
func (s *AgentService) AddJobInstanceResult(ctx context.Context, request dto.JobInstanceResult, _type string) (*mongo.UpdateResult, error) {
	jobID := request.JobId
	jobInstanceID := request.JobInstanceId
	session_id := request.SessionId
	status := request.Status
	key := request.Key
	resultData := request.Result
	var obj map[string]interface{}
	if _type == "rc" || _type == "rd" || _type == "re" {
		obj = request.ResultObj
	}
	filter := bson.M{"job_instance_id": jobInstanceID}
	update := bson.M{
		"$set": bson.M{
			"session_id":      session_id,
			"job_instance_id": jobInstanceID,
			"job_id":          jobID,
			"key":             key,
			"status":          status,
			"result":          resultData,
			"result_obj":      obj,
			"updated_time":    time.Now().UTC(),
			"agent_type":      request.AgentType,
			"name":            request.Name,
		},
		"$setOnInsert": bson.M{ // Chỉ áp dụng khi insert mới
			"created_at": time.Now().UTC(),
		},
	}
	// Tùy chọn upsert
	opts := options.Update().SetUpsert(true)
	result, err := s.colJobInstalceResult.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Error("Error when adding job instance result: "+err.Error(), "AddJobInstanceResult", update)
		return nil, err
	}
	//update coljobinstance status ="success", updated_time current time, pickup_node="",counter=counter+1
	objectID, err := primitive.ObjectIDFromHex(jobInstanceID)
	if err != nil {
		log.Error("Invalid job instance ID: "+err.Error(), "AddJobInstanceResult", jobInstanceID)
		return nil, err
	}
	filter = bson.M{"_id": objectID}
	update = bson.M{
		"$set": bson.M{
			"status":       "success",
			"updated_time": time.Now().UTC(),
			//"pickup_node": "",
		},
		"$inc": bson.M{"counter": 1},
	}
	_, err = s.colJobInstance.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error("Error when updating job instance: "+err.Error(), "AddJobInstanceResult", filter)
		return nil, err
	}
	return result, nil
}

/*
"total_tasks": total_count,
"done": done_count,
"list_result_done": result_links,
"pending": pending_count,
"failed": failed_count,
"queue_size": queue_size,
"completion_percentage": completion_percentage
*/
func (s *AgentService) GetJobInstanceResult(ctx context.Context, session_id, _type string) (map[string]interface{}, error) {
	//count total from jobinstance: rx,rc,rd,re
	filter := bson.M{"session_id": session_id, "agent_type": strings.ToUpper(_type)}
	count, err := s.colJobInstance.CountDocuments(ctx, filter)
	if err != nil {
		log.Error("Error when get job instance: "+err.Error(), "GetJobInstanceResult", filter)
		return nil, err
	}

	//
	//query all by type from jobinstanceResult
	filter = bson.M{"session_id": session_id, "agent_type": strings.ToUpper(_type)}
	cursor, err := s.colJobInstalceResult.Find(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Error("Job instance result not found: "+err.Error(), "GetJobInstanceResult", filter)
			return nil, errors.New("Job instance result not found")
		}
		log.Error("Error when get job instance result: "+err.Error(), "GetJobInstanceResult", filter)
		return nil, err
	}
	defer cursor.Close(ctx)
	var jobResults []*dto.JobInstanceResult
	if err = cursor.All(ctx, &jobResults); err != nil {
		log.Error("Error when get job instance result: "+err.Error(), "GetJobInstanceResult", filter)
		return nil, err
	}

	//calculate the field on above comment
	total_count := count
	done_count := 0
	pending_count := 0
	failed_count := 0
	queue_size := 100
	var results_obj []map[string]interface{}
	var result_failed []map[string]interface{}
	for _, jobResult := range jobResults {
		//
		//get one job instance
		objectId, err := primitive.ObjectIDFromHex(jobResult.JobInstanceId)
		if err != nil {
			log.Error("Error when convert job instance id to objectID: "+err.Error(), "GetJobInstanceResult")
			return nil, err
		}
		num_loop := 0
		filter = bson.M{
			"_id": objectId,
		}
		var jobInstance dto.JobInstance
		err = s.colJobInstance.FindOne(ctx, filter).Decode(&jobInstance)
		if err != nil {
			log.Error("Error when get job instance: "+err.Error(), "GetJobInstanceResult", filter)
			return nil, err
		}
		num_loop = jobInstance.Counter
		//
		if jobResult.Status == "success" || jobResult.Status == "done" { //done
			done_count++
			if _type == "rx" {
				arr := utils.Explode(jobResult.Result, "||")
				if len(arr) != 3 {
					log.Error("Invalid result: "+jobResult.Result, "GetJobInstanceResult")
					return nil, errors.New("Invalid result")
				}
				x := map[string]interface{}{
					"id":         arr[0],
					"data":       arr[2],
					"task_id":    arr[1],
					"num_loop":   num_loop,
					"agent_type": "RX",
					"name":       jobResult.Name,
				}
				results_obj = append(results_obj, x)
			} else {
				jobResult.ResultObj["num_loop"] = num_loop
				jobResult.ResultObj["agent_type"] = jobResult.AgentType
				results_obj = append(results_obj, jobResult.ResultObj)
			}
		} else if jobResult.Status == "failed" { //failed
			failed_count++
			if _type == "rx" {
				arr := utils.Explode(jobResult.Result, "||")
				if len(arr) != 3 {
					log.Error("Invalid result: "+jobResult.Result, "GetJobInstanceResult")
					return nil, errors.New("Invalid result")
				}
				x := map[string]interface{}{
					"id":         arr[0],
					"error":      arr[2],
					"task_id":    arr[1],
					"num_loop":   num_loop,
					"agent_type": "RX",
					"name":       jobResult.Name,
				}
				result_failed = append(result_failed, x)
			} else {
				jobResult.ResultObj["num_loop"] = num_loop
				jobResult.ResultObj["agent_type"] = jobResult.AgentType
				result_failed = append(result_failed, jobResult.ResultObj)
			}
		}
	}

	completion_percentage := 0
	if total_count > 0 {
		completion_percentage = (done_count * 100) / int(total_count)
	}
	pending_count = int(count) - (done_count + failed_count)
	if pending_count < 0 {
		pending_count = 0
	}
	//remove null when response
	if len(results_obj) == 0 {
		results_obj = []map[string]interface{}{}
	}
	if len(result_failed) == 0 {
		result_failed = []map[string]interface{}{}
	}
	result := map[string]interface{}{
		"total_tasks":           total_count,
		"done":                  done_count,
		"list_result_done":      results_obj,
		"list_failed":           result_failed,
		"pending":               pending_count,
		"failed":                failed_count,
		"queue_size":            queue_size,
		"completion_percentage": completion_percentage,
	}
	return result, nil
}

func (s *AgentService) RegisterRAgent(ctx context.Context, request dto.RAgentRequest) (*dto.RAgent, error) {
	// Check if agent already exists with this wallet address and device ID
	filter := bson.M{
		"wallet_address": request.WalletAddress,
		"device_id":      request.DeviceID,
	}

	// Get current time for timestamps
	now := time.Now().UTC()

	// Find and update existing agent or create a new one
	update := bson.M{
		"$set": bson.M{
			"os":             request.OS,
			"version":        request.Version,
			"type":           request.Type,
			"last_ping_time": now,
			"updated_at":     now,
			"status":         "active",
		},
		"$setOnInsert": bson.M{
			"created_at": now,
		},
	}

	opts := options.FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After)

	var ragent dto.RAgent
	err := s.colRAgent.FindOneAndUpdate(ctx, filter, update, opts).Decode(&ragent)
	if err != nil {
		log.Error("Error when registering RAgent: "+err.Error(), "RegisterRAgent", request)
		return nil, err
	}

	return &ragent, nil
}

func (s *AgentService) PingRAgent(ctx context.Context, deviceID string) (*dto.RAgent, error) {
	// Find the agent by wallet address and device ID
	filter := bson.M{
		"device_id": deviceID,
	}

	// Update the last ping time
	now := time.Now().UTC()
	update := bson.M{
		"$set": bson.M{
			"last_ping_time": now,
			"updated_at":     now,
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var ragent dto.RAgent
	err := s.colRAgent.FindOneAndUpdate(ctx, filter, update, opts).Decode(&ragent)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Error("RAgent not found for ping: "+err.Error(), "PingRAgent", filter)
			return nil, errors.New("Agent not registered")
		}
		log.Error("Error when updating RAgent ping time: "+err.Error(), "PingRAgent", filter)
		return nil, err
	}

	return &ragent, nil
}

func (s *AgentService) UpdateRAgentWallet(ctx context.Context, oldWallet, newWallet, deviceID, code string) (*dto.RAgent, error) {
	// First, find the agent to check the code
	findFilter := bson.M{
		"device_id":      deviceID,
		"wallet_address": oldWallet,
	}

	var existingAgent dto.RAgent
	err := s.colRAgent.FindOne(ctx, findFilter).Decode(&existingAgent)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Error("RAgent not found for code verification: "+err.Error(), "UpdateRAgentWallet", findFilter)
			return nil, errors.New("Agent not found with the provided wallet address and device ID")
		}
		log.Error("Error when retrieving RAgent for code verification: "+err.Error(), "UpdateRAgentWallet", findFilter)
		return nil, err
	}

	// Verify the code
	if existingAgent.Code != code {
		log.Error("Invalid verification code provided: "+code, "UpdateRAgentWallet")
		return nil, errors.New("Invalid verification code")
	}

	// Check if the code has expired (5 minutes)
	now := time.Now().UTC()
	codeAge := now.Sub(existingAgent.CodeGeneratedAt)
	if codeAge > 5*time.Minute {
		log.Error("Verification code has expired", "UpdateRAgentWallet")
		return nil, errors.New("Verification code has expired. Please generate a new code")
	}

	// Code is valid and not expired, proceed with wallet update
	// Find the agent by old wallet address and device ID
	filter := bson.M{
		"wallet_address": oldWallet,
		"device_id":      deviceID,
	}

	// Update the wallet address and keep track of the old one
	update := bson.M{
		"$set": bson.M{
			"wallet_address":     newWallet,
			"old_wallet_address": oldWallet,
			"updated_at":         now,
			"last_ping_time":     now, // Also update ping time since it's an active operation
			"code":               "",  // Clear the code after successful use
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var ragent dto.RAgent
	err = s.colRAgent.FindOneAndUpdate(ctx, filter, update, opts).Decode(&ragent)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Error("RAgent not found for wallet update: "+err.Error(), "UpdateRAgentWallet", filter)
			return nil, errors.New("Agent not found with the provided wallet address and device ID")
		}
		log.Error("Error when updating RAgent wallet: "+err.Error(), "UpdateRAgentWallet", filter)
		return nil, err
	}

	return &ragent, nil
}

func (s *AgentService) GetRAgentVersion(ctx context.Context) (map[string]interface{}, error) {
	// Get the current RAgent version from server configuration
	ragentVersion := s.server.ExtendConfig.RAgentVersion

	// Return version information
	result := map[string]interface{}{
		"last_version": ragentVersion,
	}

	return result, nil
}

func (s *AgentService) GetRAgentByDeviceID(ctx context.Context, deviceID string) (*dto.RAgent, error) {
	// Find the agent by device ID
	filter := bson.M{
		"device_id": deviceID,
	}

	var ragent dto.RAgent
	err := s.colRAgent.FindOne(ctx, filter).Decode(&ragent)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Error("RAgent not found with device ID: "+deviceID, "GetRAgentByDeviceID")
			return nil, errors.New("Agent not found with the provided device ID")
		}
		log.Error("Error when retrieving RAgent by device ID: "+err.Error(), "GetRAgentByDeviceID", filter)
		return nil, err
	}

	return &ragent, nil
}

func (s *AgentService) GenerateRAgentCode(ctx context.Context, deviceID string) (map[string]interface{}, error) {
	// Generate a random 6-digit code
	code := fmt.Sprintf("%06d", utils.Random(0, 999999))

	// Find the agent by device ID
	filter := bson.M{
		"device_id": deviceID,
	}

	// Get current time for timestamps
	now := time.Now().UTC()

	// Update the agent with the new code
	update := bson.M{
		"$set": bson.M{
			"code":              code,
			"code_generated_at": now,
			"updated_at":        now,
			"last_ping_time":    now, // Also update ping time since it's an active operation
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var ragent dto.RAgent
	err := s.colRAgent.FindOneAndUpdate(ctx, filter, update, opts).Decode(&ragent)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Error("RAgent not found for code generation: "+deviceID, "GenerateRAgentCode", filter)
			return nil, errors.New("Agent not found with the provided device ID")
		}
		log.Error("Error when generating code for RAgent: "+err.Error(), "GenerateRAgentCode", filter)
		return nil, err
	}

	// Return code information
	result := map[string]interface{}{
		"code":              code,
		"code_generated_at": now,
		"device_id":         deviceID,
		"wallet_address":    ragent.WalletAddress,
	}

	return result, nil
}

func (s *AgentService) GetRAgent(ctx context.Context, num int, agentType string) ([]*dto.RAgent, error) {
	// Create filter based on agent type
	filter := bson.M{}

	// If a specific agent type is requested (not "*"), add it to the filter
	if agentType != "*" {
		filter["type"] = agentType
	}

	// Only include active agents
	filter["status"] = "active"

	// Set up query options
	opts := options.Find().
		SetLimit(int64(num)).
		SetSort(bson.D{
			{"last_ping_time", -1}, // Most recently active first
		})

	// Execute the query
	cursor, err := s.colRAgent.Find(ctx, filter, opts)
	if err != nil {
		log.Error("Error when finding RAgents: "+err.Error(), "GetRAgent", filter)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Decode results
	var ragents []*dto.RAgent
	if err = cursor.All(ctx, &ragents); err != nil {
		log.Error("Error when decoding RAgents: "+err.Error(), "GetRAgent", filter)
		return nil, err
	}

	// If no agents found, return empty slice (not nil)
	if len(ragents) == 0 {
		return []*dto.RAgent{}, nil
	}

	return ragents, nil
}

// RXSwarmItem represents a swarm item with project details
type RXSwarmItem struct {
	ProjectID   string `json:"project_id" bson:"project_id"`
	ProjectName string `json:"project_name" bson:"project_name"`
	Status      string `json:"status" bson:"status"`
}

func (s *AgentService) GetRXSwarmList(ctx context.Context, page, pageSize int) (map[string]interface{}, error) {
	log.Info("GetRXSwarmList", "page", page)
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	// Use aggregation to get unique project_ids first
	pipeline := []bson.M{
		{"$group": bson.M{
			"_id": "$project_id",
		}},
		{"$project": bson.M{
			"project_id": "$_id",
			"_id":        0,
		}},
	}

	// Count total unique projects for pagination
	countPipeline := append(pipeline, bson.M{"$count": "total"})
	countCursor, err := s.colXusersProject.Aggregate(ctx, countPipeline)
	if err != nil {
		log.Error("Error counting unique projects: "+err.Error(), "GetRXSwarmList")
		return nil, err
	}
	defer countCursor.Close(ctx)

	var countResult []bson.M
	if err = countCursor.All(ctx, &countResult); err != nil {
		log.Error("Error decoding count result: "+err.Error(), "GetRXSwarmList")
		return nil, err
	}
	// Get total count from aggregation result
	totalCount := int64(0)
	fmt.Printf("countResult: %+v\n", countResult)
	if len(countResult) > 0 && countResult[0]["total"] != nil {
		totalCount = int64(countResult[0]["total"].(int32))
	}
	// Add pagination to the pipeline
	paginatedPipeline := append(pipeline,
		bson.M{"$sort": bson.M{"project_id": 1}},
		bson.M{"$skip": (page - 1) * pageSize},
		bson.M{"$limit": pageSize},
	)

	// Run the aggregation
	cursor, err := s.colXusersProject.Aggregate(ctx, paginatedPipeline)
	if err != nil {
		log.Error("Error finding unique projects: "+err.Error(), "GetRXSwarmList")
		return nil, err
	}
	defer cursor.Close(ctx)

	var uniqueProjects []bson.M
	if err = cursor.All(ctx, &uniqueProjects); err != nil {
		log.Error("Error decoding unique projects: "+err.Error(), "GetRXSwarmList")
		return nil, err
	}
	// Extract project IDs
	projectIDs := make([]string, 0, len(uniqueProjects))
	for _, project := range uniqueProjects {
		projectID, ok := project["project_id"].(string)
		if ok {
			projectIDs = append(projectIDs, projectID)
		}
	}

	// Get project details for each project_id
	var projectObjectIDs []primitive.ObjectID
	for _, projectID := range projectIDs {
		objID, err := primitive.ObjectIDFromHex(projectID)
		if err != nil {
			log.Error("Invalid project ID format: "+projectID, "GetRXSwarmList")
			continue
		}
		projectObjectIDs = append(projectObjectIDs, objID)
	}

	var projects []projectDto.Project
	if len(projectObjectIDs) > 0 {
		projectCursor, err := s.colProject.Find(ctx, bson.M{"_id": bson.M{"$in": projectObjectIDs}})
		if err != nil {
			log.Error("Error finding projects: "+err.Error(), "GetRXSwarmList")
			return nil, err
		}
		defer projectCursor.Close(ctx)

		if err = projectCursor.All(ctx, &projects); err != nil {
			log.Error("Error decoding projects: "+err.Error(), "GetRXSwarmList")
			return nil, err
		}
	}

	// Create a map of project ID to project name for easy lookup
	projectNameMap := make(map[string]string)
	for _, project := range projects {
		projectID := project.ID.Hex()
		projectNameMap[projectID] = project.Name
	}

	// Get X users for each project and create ONLY ONE swarm item per project
	var swarmItems []RXSwarmItem
	projectAdded := make(map[string]bool) // Track which projects have been added

	for _, projectID := range projectIDs {
		// Skip if we've already added this project
		if projectAdded[projectID] {
			continue
		}

		// Get a sample X user for this project (we just need one to represent the project)
		userFilter := bson.M{"project_id": projectID}
		userOpts := options.FindOne()

		var xuser dto.XuserProject
		err := s.colXusersProject.FindOne(ctx, userFilter, userOpts).Decode(&xuser)
		if err != nil {
			log.Error("Error finding X user for project: "+err.Error(), "GetRXSwarmList", projectID)
			continue
		}

		projectName := projectNameMap[projectID]
		if projectName == "" {
			projectName = "Unknown Project"
		}

		// Create exactly one swarm item for this project
		swarmItem := RXSwarmItem{
			ProjectID:   projectID,
			ProjectName: projectName,
			Status:      "active",
		}
		swarmItems = append(swarmItems, swarmItem)

		// Mark this project as added
		projectAdded[projectID] = true
	}

	// Calculate pagination info
	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))
	if totalPages < 1 {
		totalPages = 1
	}

	// Return the response
	result := map[string]interface{}{
		"items":       swarmItems,
		"page":        page,
		"page_size":   pageSize,
		"total_pages": totalPages,
		"total_items": totalCount,
	}

	return result, nil
}

type SwarmItem struct {
	ProjectID   string `json:"project_id" bson:"project_id"`
	ProjectName string `json:"project_name" bson:"project_name"`
	UniqueName  string `json:"unique_name" bson:"unique_name"`
	Status      string `json:"status" bson:"status"`
	Logo        string `json:"logo" bson:"logo"`
}

func (s *AgentService) GetSwarmList(ctx context.Context, x_id string) ([]SwarmItem, error) {
	log.Info("GetSwarmList", "x_id", x_id)
	if x_id == "" {
		return nil, errors.New("Invalid X ID")
	}
	//find user by xUserId and decode to dto.Xuser
	filter := bson.M{"xUserId": x_id}
	var user userDto.User
	err := s.colUser.FindOne(ctx, filter).Decode(&user)
	found_user := true
	if err != nil {
		if err == mongo.ErrNoDocuments {
			//
			found_user = false
			//return nil, errors.New("NOT_FOUND")
		} else {
			log.Error("Error finding user: "+err.Error(), "GetSwarmList")
			return nil, err
		}
	}

	var swarmItems []SwarmItem

	//find all projects with creator=user.WalletAddress
	filter = bson.M{"creator": user.WalletAddress}
	if !found_user {
		filter = bson.M{"x_id": x_id}
	}
	cursor, err := s.colProject.Find(ctx, filter)
	if err != nil {
		log.Error("Error finding projects: "+err.Error(), "GetSwarmList")
		return nil, err
	}
	defer cursor.Close(ctx)
	var projects []projectDto.Project
	if err = cursor.All(ctx, &projects); err != nil {
		log.Error("Error decoding projects: "+err.Error(), "GetSwarmList")
		return nil, err
	}
	err = nil
	if !found_user {
		err = errors.New("NOT_FOUND")
	}
	if len(projects) == 0 { //get by creator wallet if not exist get by x_id
		swarmItems = make([]SwarmItem, 0)
		return swarmItems, err
	} else {
		found_user = true
		err = nil
	}
	for _, project := range projects {
		swarmItems = append(swarmItems, SwarmItem{
			ProjectID:   project.ID.Hex(),
			ProjectName: project.Name,
			UniqueName:  project.UniqueName,
			Logo:        project.Logo,
			Status:      "active",
		})
	}
	return swarmItems, err
}

// get tweet content from url
func (s *AgentService) GetTweetContent(ctx context.Context, url string) (string, error) {
	return http.GetTweetContent(url)
}
