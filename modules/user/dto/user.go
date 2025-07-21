package dto
import (	
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)
type UserAccessKey struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserId       string             `json:"user_id" bson:"user_id"`
	AccessKey    string             `json:"access_key" bson:"access_key"`
	Title		string             `json:"title" bson:"title"`
	CreatedTime  time.Time          `json:"created_time" bson:"created_time"`
	UpdatedTime  time.Time          `json:"updated_time" bson:"updated_time"`
}
/*
JWT rome nodejs parse
{
  "walletAddress": "0xc0fcfad6a5db86f566102b8a7115823068a93c50",
  "id": "67a045614875e0fee748ddb5",
  "roles": [
    "user"
  ],
  "iat": 1738895059,
  "exp": 1741487059
}
*/
/*
{
    "_id" : ObjectId("67b4077fc6312c0d7f6f2769"),
    "sys_partition" : NumberInt(-1),
    "walletAddress" : "0xe141eb5dbfb3bd9308e1599d17003b772a02a8d9",
    "name" : "",
    "email" : "",
    "xUserId" : "",
    "discordUserId" : "",
    "country" : "",
    "isVerified" : false,
    "reputation" : NumberInt(10),
    "reward" : NumberInt(0),
    "roles" : [
        "user"
    ],
    "createdAt" : ISODate("2025-02-18T04:07:27.175+0000"),
    "updatedAt" : ISODate("2025-02-18T04:07:27.175+0000"),
    "__v" : NumberInt(0)
}
*/
type User struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	SysPartition   int                `json:"sys_partition" bson:"sys_partition"`
	WalletAddress  string             `json:"wallet_address" bson:"walletAddress"`
	Name           string             `json:"name" bson:"name"`
	Email          string             `json:"email" bson:"email"`
	XUserId        string             `json:"x_user_id" bson:"xUserId"`
	DiscordUserId  string             `json:"discord_user_id" bson:"discordUserId"`
	Country        string             `json:"country" bson:"country"`
	IsVerified     bool               `json:"is_verified" bson:"isVerified"`
	Reputation     int                `json:"reputation" bson:"reputation"`
	Reward         int                `json:"reward" bson:"reward"`
	Roles          []string           `json:"roles" bson:"roles"`
	CreatedAt      time.Time          `json:"created_at" bson:"createdAt"`
	UpdatedAt      time.Time          `json:"updated_at" bson:"updatedAt"`
}