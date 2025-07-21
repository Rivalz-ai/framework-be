
package dto
import (	
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)
type XUser struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserID        string    `json:"userId" bson:"userId"`
	XID           string    `json:"xId" bson:"xId"`
	XUsername     string    `json:"xUsername" bson:"xUsername"`
	Name          string    `json:"name" bson:"name"`
	Description   string    `json:"description" bson:"description"`
	Avatar        string    `json:"avatar" bson:"avatar"`
	Bio           string    `json:"bio" bson:"bio"`
	Followers     int       `json:"followers" bson:"followers"`
	Following     int       `json:"following" bson:"following"`
	Tweets        int       `json:"tweets" bson:"tweets"`
	Likes         int       `json:"likes" bson:"likes"`
	Medias        int       `json:"medias" bson:"medias"`
	AvgImpression int       `json:"avgImpression" bson:"avgImpression"`
	AccessToken   string    `json:"accessToken" bson:"accessToken"`
	RefreshToken  string    `json:"refreshToken" bson:"refreshToken"`
	ExpiresAt     time.Time `json:"expiresAt" bson:"expiresAt"`
	CreatedAt     time.Time `json:"createdAt" bson:"createdAt"`
	LastUpdate    time.Time `json:"lastUpdate" bson:"lastUpdate"`
	Partition	  int    	`json:"partition" bson:"partition"`
	FailCount	 int       `json:"fail_count" bson:"fail_count"`
	IsOriginToken bool     `json:"is_origin_token" bson:"is_origin_token"`
}