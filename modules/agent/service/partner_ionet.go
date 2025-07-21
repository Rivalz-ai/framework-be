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
	url     = "https://api.intelligence.io.solutions/api/v1/workflows/run"
	llm_url = "https://api.intelligence.io.solutions/api/v1/chat/completions"
)

func ProcessIONetTask(ctx context.Context, agent_name, _type string, params map[string]interface{}, api_key string) (string, error) {
	if _type == "llm" {
		return LLM(ctx, params["text"].(string), params, api_key)
	}
	fmt.Println("agent_type: ", _type)
	if agent_name == "" {
		return "", errors.New("agent_name is nil")
	}
	text := params["text"].(string)
	if text == "" {
		return "", errors.New("text is nil")
	}
	model := params["model"].(string)
	if model == "" {
		model = "deepseek-ai/DeepSeek-R1"
	}
	header := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + api_key,
	}
	body := map[string]interface{}{
		"text":        text,
		"agent_names": []string{agent_name},
	}
	args := map[string]interface{}{
		"model": model,
		"type":  _type,
	}
	if _type == "classify" {
		args["classify_by"] = []string{
			"Narrative",
			"Lyric",
			"Drama",
			"Painting",
			"Sculpture",
			"Architecture",
			"Photography",
			"Music",
			"Dance",
			"Theater",
			"Film",
			"Folk Literature",
			"Contemporary Art",
			"Poetry",
			"Short Story",
			"Novel",
			"Epic",
			"Mythology",
			"Fairy Tale",
			"Legend",
			"Proverb",
			"Riddle",
			"Essay",
			"Memoir",
			"Biography",
			"Autobiography",
			"Satire",
			"Tragedy",
			"Comedy",
			"Opera",
			"Musical Theater",
			"Puppetry",
			"Ballet",
			"Folk Dance",
			"Contemporary Dance",
			"Drawing",
			"Printmaking",
			"Ceramics",
			"Textile Art",
			"Calligraphy",
			"Graphic Design",
			"Illustration",
			"Animation",
			"Documentary Film",
			"Experimental Film",
			"Street Art",
			"Installation Art",
			"Performance Art",
			"Digital Art",
			"Video Art",
			"Sound Art",
			"Comic",
			"Graphic Novel",
			"Folksong",
			"Classical Music",
			"Jazz",
			"Pop Music",
			"Traditional Music",
			"Choral Music",
			"Orchestral Music",
			"Electronic Music",
			"Fashion Design",
			"Interior Design",
			"Landscape Architecture",
			"Conceptual Art",
			"Mime",
			"Stand-up Comedy",
			"Spoken Word",
		}
	}
	body["args"] = args
	code, resp, err := http.SendRequest(url, "POST", header, body)
	log.Info("IONet response", "IONet", resp)
	if err != nil {
		log.Error("Error sending request to IONet: "+err.Error(), "IONet", resp)
		return "", err
	}
	if code != 200 {
		log.Error("IONet response code: "+strconv.Itoa(code), "IONet", resp)
		return "", errors.New("failed to process ionet task")
	}
	if _type == "classify" {
		return AgentClassificationParse(resp)
	}
	return AgentSummaryParse(resp)
}
func AgentSummaryParse(m map[string]interface{}) (string, error) {
	if m["result"] == nil {
		return "", errors.New("result is nil")
	}
	result := m["result"].(map[string]interface{})
	if result["summary"] == nil {
		return "", errors.New("summary is nil")
	}
	if result["key_points"] == nil {
		return "", errors.New("key_points is nil")
	}
	response := fmt.Sprintf("Summary: %s", result["summary"])
	return response, nil
}
func AgentClassificationParse(m map[string]interface{}) (string, error) {
	if m["result"] == nil {
		return "", errors.New("result is nil")
	}
	//convert m["result"] to string
	result := m["result"].(string)
	return result, nil
}

/*
1. request

	{
		"model": "meta-llama/Llama-3.3-70B-Instruct",
		"messages": [{"role": "user", "content": "how are you?"}],
		"reasoning_content": true,
		"temperature": 0.7
	}

2. response:

	{
	    "id": "019776d4-3ace-8e19-b058-bceb7568a28e",
	    "object": "chat.completion",
	    "created": 1750045308,
	    "model": "meta-llama/Llama-3.3-70B-Instruct",
	    "choices": [
	        {
	            "index": 0,
	            "message": {
	                "role": "assistant",
	                "reasoning_content": null,
	                "content": "I'm just a language model, so I don't have feelings or emotions like humans do, but I'm functioning properly and ready to help with any questions or tasks you may have. How about you? How's your day going so far?",
	                "tool_calls": []
	            },
	            "logprobs": null,
	            "finish_reason": "stop",
	            "stop_reason": null
	        }
	    ],
	    "usage": {
	        "prompt_tokens": 39,
	        "total_tokens": 89,
	        "completion_tokens": 50,
	        "prompt_tokens_details": null
	    },
	    "prompt_logprobs": null
	}
*/
func LLM(ctx context.Context, content string, params map[string]interface{}, api_key string) (string, error) {
	header := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + api_key,
	}
	model := params["model"].(string)
	if model == "" {
		model = "meta-llama/Llama-3.3-70B-Instruct"
	}
	body := map[string]interface{}{
		"model":             model,
		"messages":          []map[string]interface{}{{"role": "assistant", "content": content}},
		"reasoning_content": true,
		"temperature":       0.7,
	}
	code, resp, err := http.SendRequest(llm_url, "POST", header, body)
	log.Info("IONet response", "IONet", resp)
	if err != nil {
		log.Error("Error sending request to IONet: "+err.Error(), "IONet", resp)
		return "", err
	}
	if code != 200 {
		log.Error("IONet response code: "+strconv.Itoa(code), "IONet", resp)
		return "", errors.New("failed to process ionet task")
	}
	if resp["choices"] == nil {
		log.Error("IONet response choices is nil", "IONet", resp)
		return "", errors.New("choices is nil")
	}
	choices, err := utils.ItoSlice(resp["choices"])
	if err != nil {
		log.Error("IONet response choices is nil", "IONet", resp)
		return "", errors.New("choices is nil")
	}
	if len(choices) == 0 {
		log.Error("IONet response choices is empty", "IONet", resp)
		return "", errors.New("choices is empty")
	}
	choice, err := utils.ItoDictionary(choices[0])
	if err != nil {
		log.Error("IONet response choices is nil", "IONet", resp)
		return "", errors.New("choices is nil")
	}
	if choice["message"] == nil {
		log.Error("IONet response message is nil", "IONet", resp)
		return "", errors.New("message is nil")
	}
	message := choice["message"].(map[string]interface{})
	if message["content"] == nil {
		log.Error("IONet response content is nil", "IONet", resp)
		return "", errors.New("content is nil")
	}
	rep_content := message["content"].(string)
	return rep_content, nil
}
