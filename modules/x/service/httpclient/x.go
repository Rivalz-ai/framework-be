package httpclient

import (
	"errors"
	"time"

	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/Rivalz-ai/framework-be/framework/utils"
	"github.com/Rivalz-ai/framework-be/modules/x/dto"
	"github.com/go-resty/resty/v2"
	//"fmt"
)

/*
	{
	    "data": {
	        "profile_image_url": "https://pbs.twimg.com/profile_images/1834247183849414656/RmrYrn-x_normal.jpg",
	        "public_metrics": {
	            "followers_count": 20,
	            "following_count": 35,
	            "tweet_count": 44,
	            "listed_count": 0,
	            "like_count": 65,
	            "media_count": 0
	        },
	        "name": "su hoang son",
	        "id": "1563510242",
	        "created_at": "2013-07-02T15:56:47.000Z",
	        "description": "",
	        "username": "suhoangson"
	    }
	}
*/
func GetXInfo(base_url string, access_token string, logData map[string]interface{}) (*dto.XInfo, error) {
	url := base_url + "/users/me?user.fields=profile_image_url,public_metrics,description,created_at"
	logData["x_url"] = url
	client := resty.New()
	//
	var respData interface{}
	//response error
	var respErr interface{}
	//
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+access_token).
		SetResult(&respData).
		SetError(&respErr).
		Get(url)
	if err != nil {
		log.Error("Error when get X info: "+err.Error(), "GetXInfo", logData)
		return nil, err
	}
	logData["response_data"] = respData
	logData["response_error"] = respErr
	if res.StatusCode() != 200 {
		data, err := utils.ItoDictionary(respErr)
		if err != nil {
			log.Error("Error when get X info, decode Response Error from X: "+err.Error(), "GetXInfo", logData)
			return nil, err
		}
		if data["detail"] != nil {
			return nil, errors.New(data["detail"].(string))
		} else {
			return nil, errors.New("Something wrong from X")
		}

	}
	dataobj, err := utils.ItoDictionary(respData)
	if err != nil {
		log.Error("Error when get X info, decode Response Data from X: "+err.Error(), "GetXInfo", logData)
		return nil, err
	}
	if dataobj["data"] == nil {
		logData["data"] = respData
		log.Error("Error when get X info, data in json response is nil", "GetXInfo", logData)
		return nil, errors.New("data is invalid")
	}
	data, err := utils.ItoDictionary(dataobj["data"])
	if err != nil {
		log.Error("Error when get X info, decode Data from X: "+err.Error(), "GetXInfo", logData)
		return nil, err
	}

	/*
			XId string `json:"xId"`
		    XName string `json:"xName"`
		    XUser string `json:"xUser"`
		    XAvatar string `json:"xAvatar"`
		    XCreatedAt time.Time `json:"xCreatedAt"`
		    XDescription string `json:"xDescription"`
	*/
	// Parse chuỗi thành time.Time
	parsedTime, err := time.Parse(time.RFC3339, data["created_at"].(string))
	if err != nil {
		logData["x_created_at"] = data["created_at"]
		log.Error("Error when parse time created_at: "+err.Error(), "GetXInfo", logData)
		return nil, err
	}
	if data["public_metrics"] == nil {
		logData["data"] = respData
		log.Error("Error when get X info, public_metrics in json response is nil", "GetXInfo", logData)
		return nil, errors.New("public_metrics is invalid")
	}
	metrics, err := utils.ItoDictionary(data["public_metrics"])
	if err != nil {
		log.Error("Error when get X info, decode public_metrics from X: "+err.Error(), "GetXInfo", logData)
		return nil, err
	}
	xInfo := &dto.XInfo{
		XId:            data["id"].(string),
		XName:          data["name"].(string),
		XUserName:      data["username"].(string),
		XAvatar:        data["profile_image_url"].(string),
		XCreatedAt:     &parsedTime,
		XDescription:   data["description"].(string),
		FollowersCount: utils.ItoInt(metrics["followers_count"]),
		FollowingCount: utils.ItoInt(metrics["following_count"]),
		TweetCount:     utils.ItoInt(metrics["tweet_count"]),
		ListedCount:    utils.ItoInt(metrics["listed_count"]),
		LikeCount:      utils.ItoInt(metrics["like_count"]),
		MediaCount:     utils.ItoInt(metrics["media_count"]),
	}
	return xInfo, nil
}
