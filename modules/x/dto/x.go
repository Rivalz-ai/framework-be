package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
XUser
{
    "_id" : ObjectId("67a36a126db2afcf18bfa203"),
    "userId" : "67a045614875e0fee748ddb5",
    "xId" : "1563510242",
    "xUsername" : "suhoangson",
    "name" : "su hoang son",
    "description" : "",
    "avatar" : "https://pbs.twimg.com/profile_images/1834247183849414656/RmrYrn-x_normal.jpg",
    "bio" : "",
    "followers" : NumberInt(18),
    "following" : NumberInt(35),
    "tweets" : NumberInt(43),
    "likes" : NumberInt(66),
    "medias" : NumberInt(0),
    "avgImpression" : NumberInt(0),
    "accessToken" : "Y0N6SGFva042akJzaFhzM3J4S3hJdnFRZUExUWFvTjVHbThaVEVUcnhQWFloOjE3Mzg3NjI3Njk2Mjc6MToxOmF0OjE",
    "refreshToken" : "dy1qYkJFUzZadEYzRnJ1dlVyX1VhTTVCcndzZDdOQkZkMUZpTlNDSlJ1dm9TOjE3Mzg3NjI3Njk2Mjc6MToxOnJ0OjE",
    "expiresAt" : ISODate("2025-03-07T13:39:30.467+0000"),
    "createdAt" : ISODate("2013-07-02T15:56:47.000+0000"),
    "lastUpdate" : ISODate("2025-02-05T13:39:30.946+0000"),
    "__v" : NumberInt(0)
}
*/
/*

   "expires_in": 7200, //expire time of access_token
   "access_token": "xx",
   "refresh_token": "xx" //refresh token will be changed after each time renew access_token

*/

type XUser struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserId           string             `json:"user_id" bson:"userId"`
	XId              string             `json:"x_id" bson:"xId"`
	XUsername        string             `json:"x_username" bson:"xUsername"`
	Name             string             `json:"name" bson:"name"`
	Description      string             `json:"description" bson:"description"`
	Avatar           string             `json:"avatar" bson:"avatar"`
	Bio              string             `json:"bio" bson:"bio"`
	Followers        int                `json:"followers" bson:"followers"`
	Following        int                `json:"following" bson:"following"`
	Tweets           int                `json:"tweets" bson:"tweets"`
	Likes            int                `json:"likes" bson:"likes"`
	Medias           int                `json:"medias" bson:"medias"`
	AvgImpression    int                `json:"avg_impression" bson:"avgImpression"`
	AccessToken      string             `json:"access_token" bson:"accessToken"`
	RefreshToken     string             `json:"refresh_token " bson:"refreshToken"`
	ExpiresAt        *time.Time         `json:"expires_at" bson:"expiresAt"`
	SocialScore      int                `json:"social_score" bson:"socialScore"`
	CreatedAt        *time.Time         `json:"created_time" bson:"createdAt"`
	LastUpdate       *time.Time         `json:"updated_time" bson:"lastUpdate"`
	TokenValid       bool               `json:"token_valid" bson:"token_valid"`
	ClientId         string             `json:"client_id" bson:"client_id"`
	ClientSecret     string             `json:"client_secret" bson:"client_secret"`
	ExamplePost      string             `json:"example_post" bson:"example_post"`
	StyleDescription string             `json:"style_description" bson:"style_description"`
}

type XUserToken struct {
	XId          string `json:"x_id"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	//Reputation int `json:"reputation"`
	//
	FollowersCount   int    `json:"followers_count"`
	FollowingCount   int    `json:"following_count"`
	TweetCount       int    `json:"tweet_count"`
	ListedCount      int    `json:"listed_count"`
	LikeCount        int    `json:"like_count"`
	MediaCount       int    `json:"media_count"` //
	ExamplePost      string `json:"example_post"`
	StyleDescription string `json:"style_description"`
}

type ConnectX struct {
	XAccessToken  string     `json:"x_access_token"`
	XRefreshToken string     `json:"x_refresh_token"`
	ExpiresAt     *time.Time `json:"expires_at"`
	ClientId      string     `json:"client_id"`
	ClientSecret  string     `json:"-"`
}
type MKPConnect struct {
	XId       string `json:"x_id"`
	XUsername string `json:"x_username"`
}

/*
"followers_count": 20,
"following_count": 35,
"tweet_count": 44,
"listed_count": 0,
"like_count": 65,
"media_count": 0
*/
type XInfo struct {
	XId          string     `json:"x_id"`
	XName        string     `json:"x_name"`
	XUser        string     `json:"x_user"`
	XUserName    string     `json:"x_user_name"`
	XAvatar      string     `json:"x_avatar"`
	XCreatedAt   *time.Time `json:"x_created_at"`
	XDescription string     `json:"x_description"`
	//
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
	TweetCount     int `json:"tweet_count"`
	ListedCount    int `json:"listed_count"`
	LikeCount      int `json:"like_count"`
	MediaCount     int `json:"media_count"` //
}

type App struct {
	XClientId     string `json:"x_client_id" bson:"x_client_id"`
	XClientSecret string `json:"-"`
	NumUser       int64  `json:"num_user" bson:"num_user"`
	Status        bool   `json:"status" bson:"status"`
}

type LogUserDisconnect struct {
	UserId    string     `json:"user_id" bson:"user_id"`
	XId       string     `json:"x_id" bson:"x_id"`
	CreatedAt *time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" bson:"updated_at"`
}

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
