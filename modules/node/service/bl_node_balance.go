package service
import (
	"context"
	"errors"
	//"fmt"
	
)


func (b *NodeBalancer)GetBalance(ctx context.Context,walletAddress,token_address,ABI string,decimals int) (int64,error){
	excluded := make(map[string]bool)
	for i:=0;i<len(b.nodes);i++{//number of retry time
		n,err:=b.GetNode(excluded)
		if err!=nil{
			return 0,err
		}
		result,err:=n.GetBalance(ctx,walletAddress,token_address,ABI,decimals)
		if err!=nil{
			excluded[n.Key]=true
			continue
		}
		return result,nil
	}
	return 0,errors.New("All nodes provider are failed")
}
