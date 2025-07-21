package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/Rivalz-ai/framework-be/framework/utils"
	"github.com/Rivalz-ai/framework-be/modules/agent/service/http"
)

var (
	llm_url_fetch_ai = "https://api.asi1.ai/v1/chat/completions"
)

func ProcessFETCH_AI_Task(ctx context.Context, agent_name, _type string, params map[string]interface{}, api_key string) (string, error) {
	return FETCH_AI_LLM(ctx, params["text"].(string), params, api_key)
}
func FETCH_AI_LLM(ctx context.Context, content string, params map[string]interface{}, api_key string) (string, error) {
	fmt.Println("api_key:", api_key)
	header := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + api_key,
		"Accept":        "application/json",
	}
	model := params["model"].(string)
	if model == "" {
		model = "asi1-mini"
	}
	body := map[string]interface{}{
		"model":       model,
		"messages":    []map[string]interface{}{{"role": "assistant", "content": content}},
		"temperature": 0.7,
	}
	code, resp, err := http.SendRequest(llm_url_fetch_ai, "POST", header, body)
	log.Info("FETCH.AI response", "FETCH.AI", resp)
	if err != nil {
		log.Error("Error sending request to FETCH.AI: "+err.Error(), "FETCH.AI", resp)
		return "", err
	}
	if code != 200 {
		log.Error("FETCH.AI response code: "+strconv.Itoa(code), "FETCH.AI", resp)
		return "", errors.New("failed to process FETCH.AI task")
	}
	if resp["choices"] == nil {
		log.Error("FETCH.AI response choices is nil", "FETCH.AI", resp)
		return "", errors.New("choices is nil")
	}
	choices, err := utils.ItoSlice(resp["choices"])
	if err != nil {
		log.Error("FETCH.AI response choices is nil", "FETCH.AI", resp)
		return "", errors.New("choices is nil")
	}
	if len(choices) == 0 {
		log.Error("FETCH.AI response choices is empty", "FETCH.AI", resp)
		return "", errors.New("choices is empty")
	}
	choice, err := utils.ItoDictionary(choices[0])
	if err != nil {
		log.Error("FETCH.AI response choices is nil", "FETCH.AI", resp)
		return "", errors.New("choices is nil")
	}
	if choice["message"] == nil {
		log.Error("FETCH.AI response message is nil", "FETCH.AI", resp)
		return "", errors.New("message is nil")
	}
	message := choice["message"].(map[string]interface{})
	if message["content"] == nil {
		log.Error("FETCH.AI response content is nil", "FETCH.AI", resp)
		return "", errors.New("content is nil")
	}
	rep_content := message["content"].(string)
	return rep_content, nil
}
