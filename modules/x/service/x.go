package x

import (
	"context"

	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/Rivalz-ai/framework-be/framework/utils"

	"github.com/Rivalz-ai/framework-be/modules/x/dto"
	"github.com/Rivalz-ai/framework-be/server"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"errors"
	"time"

	"github.com/Rivalz-ai/framework-be/define"

	"github.com/go-resty/resty/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"fmt"
)

type XUserService struct {
	server               *server.Server
	colXUser             *mongo.Collection
	colUser              *mongo.Collection
	colLog               *mongo.Collection
	colXusersProject     *mongo.Collection
	colApp               *mongo.Collection
	colLogUserDisconnect *mongo.Collection
	colXUserMkp          *mongo.Collection
}

func NewXUserService(sv *server.Server) (*XUserService, error) {
	//db:=sv.DB.GetDB("rome",1) => db index for scaling, default node-0
	//với những table có lượng data lớn cần tính toán index dựa trên data truyền vào cho từng method khi thực thi, ko tính toán tại hàm new service như các table nhỏ
	db, err := sv.Mgo.GetDB("rome")
	if err != nil {
		return nil, err
	}
	return &XUserService{
		server:               sv,
		colXUser:             db.Collection("xusers"),
		colUser:              db.Collection("users"),
		colLog:               db.Collection("logs"),
		colXusersProject:     db.Collection("xusersproject"),
		colApp:               db.Collection("x_app_id"),
		colLogUserDisconnect: db.Collection("log_user_disconnect"),
		colXUserMkp:          db.Collection("xusers_mkp"),
	}, nil
}

// func GetXUserByUserId
func (u *XUserService) GetXUserByUserId(ctx context.Context, user_id string) (*dto.XUser, error) {
	//find user
	var xuser dto.XUser
	filter := bson.M{"userId": user_id}
	// find user info
	err := u.colXUser.FindOne(ctx, filter).Decode(&xuser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(define.NOT_FOUND)
		}
		log.Error("error when find user: "+err.Error(), "GetXUserByUserId-FindUser", user_id)
		return nil, errors.New(define.INTERNAL_SERVER_ERROR)
	}
	return &xuser, nil
}
func (u *XUserService) GetXUserByListXID(ctx context.Context, xuser_id []string) ([]*dto.XUser, error) {
	//find user
	var xusers []*dto.XUser
	filter := bson.M{"xId": bson.M{"$in": xuser_id}}
	// find user info
	cursor, err := u.colXUser.Find(ctx, filter)
	if err != nil {
		log.Error("error when find user: "+err.Error(), "GetXUserByListID-FindUser", xuser_id)
		return nil, errors.New(define.INTERNAL_SERVER_ERROR)
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &xusers); err != nil {
		log.Error("Error read result:"+err.Error(), "GetXUserByListXID-scanCursor", xuser_id)
		return nil, err
	}
	return xusers, nil
}

/*
Response:

		{
			"token_type": "bearer",
	    	"expires_in": 7200,
	    	"access_token": "",
	    	"scope": "follows.read offline.access list.read tweet.write like.read users.read tweet.read follows.write",
	    	"refresh_token": ""
		}

Error Response:

		HTTP code: 400 Bad Request
		{
			"error": "invalid_request",
	    	"error_description": "Value passed for the token was invalid."
		}
*/
func RestClientReNewXToken(url, client_id, x_client_secret, refresh_token string) (*dto.XUserToken, error) {
	client := resty.New()
	//request  body
	requestBody := map[string]string{
		"refresh_token": refresh_token,
		"grant_type":    "refresh_token",
		"client_id":     client_id,
	}
	//fmt.Printf("Request Body: %v\n",requestBody)
	//response object
	var respData interface{}
	//response error
	var respErr interface{}
	//
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBasicAuth(client_id, x_client_secret).
		SetBody(requestBody).
		SetResult(&respData).
		SetError(&respErr).
		Post(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		data, err := utils.ItoDictionary(respErr)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(data["error_description"].(string))
	}
	data, err := utils.ItoDictionary(respData)
	if err != nil {
		return nil, err
	}
	expire := utils.ItoInt(data["expires_in"])
	if expire < 0 {
		return nil, errors.New("expires_in is invalid")
	}
	//fmt.Printf("Response Data: %v\n",data)
	token := &dto.XUserToken{
		AccessToken:  data["access_token"].(string),
		RefreshToken: data["refresh_token"].(string),
		ExpiresIn:    expire,
	}
	return token, nil
}

func (u *XUserService) Disconnect(ctx context.Context, user_id string) error {
	//set field xid trong user collection => empty => keep xid cho bước 3
	// get xid from user collection
	objectId, err2 := primitive.ObjectIDFromHex(user_id)
	if err2 != nil {
		log.Error("error when convert user_id to ObjectID: "+err2.Error(), "Disconnect-ConvertObjectID", user_id)
		return errors.New("user_id is invalid")
	}
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"xUserId": ""}}
	_, err := u.colUser.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	//xóa document trong xuser collection với userId = user(jwt)
	// get xid from xuser collection
	filter = bson.M{"userId": user_id}
	var xuser dto.XUser
	err = u.colXUser.FindOne(ctx, filter).Decode(&xuser)
	if err != nil {
		return err
	}
	_, err = u.colXUser.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	//xusers_project = xóa document có xid = xid đã xóa
	filter = bson.M{"x_id": xuser.XId}
	_, err = u.colXusersProject.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}

	// update num_user in app collection by decrement 1
	_, err = u.colApp.UpdateOne(ctx, bson.M{"x_client_id": xuser.ClientId}, bson.M{"$inc": bson.M{"num_user": -1}})
	if err != nil {
		return err
	}

	//add log
	current_time := time.Now().UTC()
	colLogUserDisconnect := &dto.LogUserDisconnect{
		UserId:    user_id,
		XId:       xuser.XId,
		CreatedAt: &current_time,
		UpdatedAt: &current_time,
	}
	_, err = u.colLogUserDisconnect.InsertOne(ctx, colLogUserDisconnect)
	if err != nil {
		return err
	}

	return nil
}

func (u *XUserService) GetApp(ctx context.Context) (*dto.App, error) {
	// get app with smallest num_user
	var app dto.App
	err := u.colApp.FindOne(ctx, bson.M{"status": true}, options.FindOne().SetSort(bson.M{"num_user": 1})).Decode(&app)
	if err != nil {
		return nil, err
	}
	return &app, nil
}

func (u *XUserService) MKPConnect(ctx context.Context, user_id, x_id, x_username string) error {
	//upsert x_id, user_id to xusers_mkp collection
	filter := bson.M{"x_id": x_id}
	update := bson.M{"$set": bson.M{"user_id": user_id, "x_username": x_username}}
	_, err := u.colXUserMkp.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}

func (u *XUserService) MKPDisconnect(ctx context.Context, user_id string) error {
	//delete x_id, user_id from xusers_mkp collection
	filter := bson.M{"user_id": user_id}
	_, err := u.colXUserMkp.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
