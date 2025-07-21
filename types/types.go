package types
import(
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/ethereum/go-ethereum/core/types"
	"time"
)
/*
{
    "_id" : ObjectId("6790f5a10b84329ed8cc64f3"),
    "hash" : "0x786216b22f78a33d721c8347f878d78916c557181a8bcce3bc76708e1f3c1f2b",
    "chainId" : NumberInt(84532),
    "from" : "0x9b50d68e0d76ee781f693e35b545498fe0b5f6d7",
    "recipient" : "0x9b50d68e0d76ee781f693e35b545498fe0b5f6d7",
    "value" : NumberInt(0),
    "tokenAddress" : "0x1c38c477e8ab969400e3e29f4e389c7d7b9457f5",
    "amount" : NumberInt(1),
    "status" : "success",
    "type" : "mint",
    "createdAt" : ISODate("2025-01-22T13:41:53.395+0000"),
    "__v" : NumberInt(0)
}
*/
type Transaction struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Hash          string `json:"hash" bson:"hash"`
	ChainId       int    `json:"chain_id" bson:"chainId"`
	From          string `json:"from" bson:"from"`
	Recipient     string `json:"recipient" bson:"recipient"`
	Value         int    `json:"value" bson:"value"`
	TokenAddress  string `json:"token_address" bson:"tokenAddress"`
	Amount        int    `json:"amount" bson:"amount"`
	Status        string `json:"status" bson:"status"`
	Type          string `json:"type" bson:"type"`
	CreatedAt     time.Time `json:"createdAt" bson:"createdAt"`
	Source 	 	  string `json:"source" bson:"source"`
	StatusDescription string `json:"status_description" bson:"status_description"`
}
type TxWrap struct {
	From string
	To string 
	EthTx *types.Transaction
}
