package service

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/Rivalz-ai/framework-be/modules/agent/dto"
)

// Định nghĩa struct để parse JSON từng dòng
type StreamChunk struct {
	Type      string `json:"type"`
	Content   string `json:"content"`
	MessageID string `json:"messageId"`
}

var url_holoworld = "https://holoworld-llm.com/api/chat"

func ProcessHoloWorldTask(ctx context.Context, agent_name string, params map[string]interface{}, api_key string, history []dto.PartnerUserHistory) (string, error) {
	params["name"] = agent_name
	params["api_key"] = api_key
	return AgentCharacteristic(params, history)
}
func AgentCharacteristic(params map[string]interface{}, history []dto.PartnerUserHistory) (string, error) {

	text := params["text"].(string)
	agent_id := params["name"].(string) //agent id
	username := params["username"].(string)
	api_key := params["api_key"].(string)
	if text == "" {
		return "", errors.New("text is nil")
	}
	/*
		payload := []byte(`{
			"characterId": "LsDn77VBTSzj6f4y8lhn",
			"agentId": "LsDn77VBTSzj6f4y8lhn",
			"input": "` + text + `",
			"username": "` + username + `",
			"chatHistory": [
				{
					"speaker": "Human",
					"content": "Hello!",
					"sentAt": "2023-04-03T10:15:30Z"
				},
				{
					"speaker": "AI",
					"content": "Hi there! How can I help you today?",
					"sentAt": "2023-04-03T10:15:35Z"
				}
			]
		}`)*/
	//create map for payload
	var historyMap []map[string]interface{}
	for _, h := range history {
		historyMapItem := map[string]interface{}{
			"speaker": "Human",
			"content": h.Text,
			"sentAt":  h.CreatedAt,
		}
		historyMap = append(historyMap, historyMapItem)
		historyMapItem = map[string]interface{}{
			"speaker": "AI",
			"content": h.Result,
			"sentAt":  h.CreatedAt,
		}
		historyMap = append(historyMap, historyMapItem)
	}
	payloadMap := make(map[string]interface{})
	payloadMap["characterId"] = agent_id
	payloadMap["agentId"] = agent_id
	payloadMap["input"] = text
	payloadMap["username"] = username
	payloadMap["chatHistory"] = historyMap
	payload, err := json.Marshal(payloadMap)
	if err != nil {
		return "", err
	}
	fmt.Printf("HoLo Agents: payload: %v\n", string(payload))
	req, err := http.NewRequest("POST", url_holoworld, bytes.NewBuffer(payload))
	if err != nil {
		log.Error("HoLo Agents: error new request: "+err.Error(), "HoloWorld", payload)
		return "", err
	}

	req.Header.Set("x-api-key", api_key)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("HoLo Agents: error do request: "+err.Error(), "HoloWorld", payload)
		return "", err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	var result string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		var chunk StreamChunk
		if err := json.Unmarshal([]byte(line), &chunk); err != nil {
			continue
		}

		if chunk.Type == "end" {
			break
		}

		result += chunk.Content
	}

	if err := scanner.Err(); err != nil {
		log.Error("HoLo Agents: error scanner: "+err.Error(), "HoloWorld", payload)
		return "", err
	}
	return result, nil
}
