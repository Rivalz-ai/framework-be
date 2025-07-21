package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/Rivalz-ai/framework-be/modules/agent/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *AgentService) GetPartnerList(ctx context.Context, x_id string) ([]dto.Partner, error) {
	//check if x_id is not empty
	if x_id == "" {
		return nil, errors.New("x_id is required")
	}
	//check x_id exist in colXuser
	filter := bson.M{
		"xId": x_id,
	}
	count, err := s.colXuser.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, errors.New("x_id not found")
	}
	//
	filter = bson.M{
		"active": true,
	}
	cursor, err := s.colPartner.Find(ctx, filter)
	if err != nil {
		log.Error("Error finding partners: "+err.Error(), "GetPartnerList")
		return nil, err
	}
	defer cursor.Close(ctx)
	var partners []dto.Partner
	if err = cursor.All(ctx, &partners); err != nil {
		log.Error("Error decoding partners: "+err.Error(), "GetPartnerList")
		return nil, err
	}
	for i := 0; i < len(partners); i++ {
		//query agent list
		agents, err := s.GetPartnerAgentList(ctx, partners[i].ID.Hex(), x_id)
		if err != nil {
			return nil, err
		}
		partners[i].Agents = agents
		partners[i].ApiKey = ""
	}
	return partners, nil
}
func (s *AgentService) GetPartnerAgentList(ctx context.Context, partner_id, x_id string) ([]dto.PartnerAgent, error) {
	//check if x_id is not empty
	if x_id == "" {
		return nil, errors.New("x_id is required")
	}
	//check x_id exist in colXuser
	filter := bson.M{
		"xId": x_id,
	}
	count, err := s.colXuser.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, errors.New("x_id not found")
	}
	//
	filter = bson.M{
		"partner_id": partner_id,
		"active":     true,
	}
	cursor, err := s.colPartnerAgent.Find(ctx, filter)
	if err != nil {
		log.Error("Error finding partner agents: "+err.Error(), "GetPartnerAgentList")
		return nil, err
	}
	defer cursor.Close(ctx)
	var partnerAgents []dto.PartnerAgent
	if err = cursor.All(ctx, &partnerAgents); err != nil {
		log.Error("Error decoding partner agents: "+err.Error(), "GetPartnerAgentList")
		return nil, err
	}
	if len(partnerAgents) == 0 {
		//create empty array
		partnerAgents = []dto.PartnerAgent{}
	}
	return partnerAgents, nil
}
func (s *AgentService) GetPartnerModelList(ctx context.Context, partner_id, x_id string) ([]dto.PartnerModel, error) {
	//check if x_id is not empty
	if x_id == "" {
		return nil, errors.New("x_id is required")
	}
	//check x_id exist in colXuser
	filter := bson.M{
		"xId": x_id,
	}
	count, err := s.colXuser.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, errors.New("x_id not found")
	}
	//
	filter = bson.M{
		"partner_id": partner_id,
		"active":     true,
	}
	cursor, err := s.colPartnerModel.Find(ctx, filter)
	if err != nil {
		log.Error("Error finding partner models: "+err.Error(), "GetPartnerModelList")
		return nil, err
	}
	defer cursor.Close(ctx)
	var partnerModels []dto.PartnerModel
	if err = cursor.All(ctx, &partnerModels); err != nil {
		log.Error("Error decoding partner models: "+err.Error(), "GetPartnerModelList")
		return nil, err
	}
	if len(partnerModels) == 0 {
		//create empty array
		partnerModels = []dto.PartnerModel{}
	}
	return partnerModels, nil
}

// method CreatePartnerAgenTask
func (s *AgentService) CreatePartnerAgentTask(ctx context.Context, x_id, thread_id string, params map[string]interface{}) (string, error) {
	//check if x_id is not empty
	if x_id == "" {
		return "", errors.New("x_id is required")
	}
	//check x_id exist in colXuser
	filter := bson.M{
		"xId": x_id,
	}
	count, err := s.colXuser.CountDocuments(ctx, filter)
	if err != nil {
		return "", err
	}
	if count == 0 {
		return "", errors.New("x_id not found")
	}
	//get agent name by thread_id
	objectId, err := primitive.ObjectIDFromHex(thread_id)
	if err != nil {
		return "", err
	}
	var thread dto.PartnerThread
	err = s.colPartnerThread.FindOne(ctx, bson.M{"_id": objectId}).Decode(&thread)
	if err != nil {
		return "", err
	}
	agent_name := thread.AgentName
	//get partner id from agent by agent_name
	filter = bson.M{
		"name": agent_name,
	}
	var agent dto.PartnerAgent
	err = s.colPartnerAgent.FindOne(ctx, filter).Decode(&agent)
	if err != nil {
		return "", err
	}
	partner_id := agent.PartnerID
	//check if project_id is not empty
	if partner_id == "" {
		return "", errors.New("partner_id is required")
	}
	//get partner from colPartner by project_id
	objectId, err = primitive.ObjectIDFromHex(partner_id)
	if err != nil {
		return "", err
	}
	var partner dto.Partner
	err = s.colPartner.FindOne(ctx, bson.M{
		"_id": objectId,
	}).Decode(&partner)
	if err != nil {
		return "", err
	}
	api_key := partner.ApiKey
	params["username"] = partner.UserName
	fmt.Printf("agent: %+v\n", agent)
	switch partner_id {
	case "682d7977dd2e0a0a3d2b88f7":
		result, err := ProcessIONetTask(ctx, agent_name, agent.Type, params, api_key)
		if err != nil {
			return "", err
		}
		err = s.StorePartnerUserHistory(ctx, partner_id, agent_name, x_id, params["text"].(string), result, thread_id)
		if err != nil {
			return "", err
		}
		return result, nil
	case "682d7988dd2e0a0a3d2b88f8": //holloword
		history, err := s.GetPartnerUserHistory(ctx, thread_id, x_id, 0, 20)
		if err != nil {
			return "", err
		}
		//
		result, err := ProcessHoloWorldTask(ctx, agent_name, params, api_key, history)
		if err != nil {
			return "", err
		}
		err = s.StorePartnerUserHistory(ctx, partner_id, agent_name, x_id, params["text"].(string), result, thread_id)
		if err != nil {
			return "", err
		}
		return result, nil
	case "6853d9df3e186849757cda25": //fetch.ai
		result, err := ProcessFETCH_AI_Task(ctx, agent_name, agent.Type, params, api_key)
		if err != nil {
			return "", err
		}
		err = s.StorePartnerUserHistory(ctx, partner_id, agent_name, x_id, params["text"].(string), result, thread_id)
		if err != nil {
			return "", err
		}
		return result, nil
	default:
		return "", errors.New("invalid project_id")
	}
}

func (s *AgentService) StorePartnerUserHistory(ctx context.Context, project_id, agent_name, x_id, text, result, thread_id string) error {
	data := bson.M{
		"partner_id": project_id,
		"agent_name": agent_name,
		"x_id":       x_id,
		"text":       text,
		"result":     result,
		"created_at": time.Now().UTC(),
		"thread_id":  thread_id,
	}
	_, err := s.colPartnerUserHistory.InsertOne(ctx, data)
	if err != nil {
		log.Error("Error inserting partner user history: "+err.Error(), "StorePartnerUserHistory")
		return err
	}
	//check if thread title is empty
	threadId, err := primitive.ObjectIDFromHex(thread_id)
	if err != nil {
		return err
	}
	var thread dto.PartnerThread
	err = s.colPartnerThread.FindOne(ctx, bson.M{"_id": threadId}).Decode(&thread)
	if err != nil {
		return err
	}
	if thread.Title == "" {
		s.UpdatePartnerThread(ctx, thread_id, s.GetFirst5Words(text))
	}
	return nil
}

// get history by agent_name sort by created_at desc and pagination
func (s *AgentService) GetPartnerUserHistory(ctx context.Context, thread_id string, x_id string, page, page_size int) ([]dto.PartnerUserHistory, error) {
	//check if x_id is not empty
	if x_id == "" {
		return nil, errors.New("x_id is required")
	}
	//check x_id exist in colXuser
	filter := bson.M{
		"xId": x_id,
	}
	count, err := s.colXuser.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, errors.New("x_id not found")
	}
	//get thread by thread_id
	objectId, err := primitive.ObjectIDFromHex(thread_id)
	if err != nil {
		return nil, err
	}
	var thread dto.PartnerThread
	err = s.colPartnerThread.FindOne(ctx, bson.M{"_id": objectId}).Decode(&thread)
	if err != nil {
		return nil, err
	}
	//get history by agent_name sort by created_at desc and pagination
	cursor, err := s.colPartnerUserHistory.Find(ctx, bson.M{
		"agent_name": thread.AgentName,
		"thread_id":  thread_id,
		"x_id":       x_id,
	}, options.Find().
		SetSort(bson.M{"created_at": -1}).
		SetSkip(int64(page*page_size)).
		SetLimit(int64(page_size)))
	if err != nil {
		log.Error("Error finding partner user history: "+err.Error(), "GetPartnerUserHistory")
		return nil, err
	}
	defer cursor.Close(ctx)
	var partnerUserHistories []dto.PartnerUserHistory
	if err = cursor.All(ctx, &partnerUserHistories); err != nil {
		log.Error("Error decoding partner user histories: "+err.Error(), "GetPartnerUserHistory")
		return nil, err
	}
	if len(partnerUserHistories) == 0 {
		//create empty array
		partnerUserHistories = []dto.PartnerUserHistory{}
	}
	return partnerUserHistories, nil
}

func (s *AgentService) CreatePartnerThread(ctx context.Context, x_id, agent_name string) (string, error) {
	//create thread and return thread_id by insertOne
	thread := dto.PartnerThread{
		ID:        primitive.NewObjectID(),
		XId:       x_id,
		AgentName: agent_name,
		CreatedAt: time.Now().UTC(),
	}
	_, err := s.colPartnerThread.InsertOne(ctx, thread)
	if err != nil {
		log.Error("Error creating partner thread: "+err.Error(), "CreatePartnerThread")
		return "", err
	}
	return thread.ID.Hex(), nil
}

func (s *AgentService) GetPartnerThread(ctx context.Context, x_id, agent_name string, page, page_size int) ([]dto.PartnerThread, error) {
	//add pagination, sort by created_at desc
	skip := int64(page * page_size)
	limit := int64(page_size)
	sort := bson.M{"created_at": -1}
	//find all threads by x_id and agent_name and title is not empty
	filter := bson.M{"x_id": x_id, "agent_name": agent_name, "title": bson.M{"$ne": ""}}
	cursor, err := s.colPartnerThread.Find(ctx, filter, options.Find().SetSkip(skip).SetLimit(limit).SetSort(sort))
	if err != nil {
		log.Error("Error finding partner threads: "+err.Error(), "GetPartnerThread")
		return nil, err
	}
	defer cursor.Close(ctx)
	var threads []dto.PartnerThread
	if err = cursor.All(ctx, &threads); err != nil {
		log.Error("Error decoding partner threads: "+err.Error(), "GetPartnerThread")
		return nil, err
	}
	if len(threads) == 0 {
		//create empty array
		threads = []dto.PartnerThread{}
	}
	return threads, nil
}

// update thread name by thread_id
func (s *AgentService) UpdatePartnerThread(ctx context.Context, thread_id, title string) error {
	objectId, err := primitive.ObjectIDFromHex(thread_id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"title": title}}
	_, err = s.colPartnerThread.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error("Error updating partner thread: "+err.Error(), "UpdatePartnerThread")
		return err
	}
	return nil
}

// func get first 5 words of text
func (s *AgentService) GetFirst5Words(text string) string {
	words := strings.Split(text, " ")
	if len(words) > 5 {
		return strings.Join(words[:5], " ")
	}
	return text
}
