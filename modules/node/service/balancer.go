package service

import (
	"errors"
	"math/rand"
	"sync"

	"github.com/Rivalz-ai/framework-be/define"
	"github.com/Rivalz-ai/framework-be/framework/utils"
	"github.com/Rivalz-ai/framework-be/server"
	"github.com/ethereum/go-ethereum/ethclient"
	//"fmt"
)

/*
type Node struct {
	NodeUrls string
	NodeSecrets string
	Client *ethclient.Client
	server *server.Server
}
*/

type NodeBalancer struct {
	nodes     []Node
	weights   []int
	sumWeight int
	mu        sync.Mutex
	server    *server.Server
}

// NewWeightedRPCBalancer khởi tạo balancer
func NewNodeBalancer(nodes []Node, sv *server.Server) *NodeBalancer {
	totalWeight := 0
	weights := make([]int, len(nodes))

	for i, node := range nodes {
		totalWeight += node.Weight
		weights[i] = totalWeight
	}

	return &NodeBalancer{
		nodes:     nodes,
		weights:   weights,
		sumWeight: totalWeight,
		server:    sv,
	}
}

func NewNodeService(server *server.Server, chainIds ...int) (*NodeBalancer, error) {
	var chainId int
	if len(chainIds) > 0 {
		chainId = chainIds[0]
	}

	var (
		nodeUrls    []string
		nodeSecrets []string
		nodeWeights []string
		nodeKeys    []string
	)
	switch chainId {
	case int(define.ETHEREUM):
		nodeUrls = utils.Explode(server.ExtendConfig.NodeEth.NodeUrls, ",")
		nodeSecrets = utils.Explode(server.ExtendConfig.NodeEth.NodeSecrets, ",")
		nodeWeights = utils.Explode(server.ExtendConfig.NodeEth.NodeWeights, ",")
		nodeKeys = utils.Explode(server.ExtendConfig.NodeEth.NodeKeys, ",")
	case int(define.BSC):
		nodeUrls = utils.Explode(server.ExtendConfig.NodeBsc.NodeUrls, ",")
		nodeSecrets = utils.Explode(server.ExtendConfig.NodeBsc.NodeSecrets, ",")
		nodeWeights = utils.Explode(server.ExtendConfig.NodeBsc.NodeWeights, ",")
		nodeKeys = utils.Explode(server.ExtendConfig.NodeBsc.NodeKeys, ",")
	case int(define.SOLANA):
		nodeUrls = utils.Explode(server.ExtendConfig.NodeSolana.NodeUrls, ",")
		nodeSecrets = utils.Explode(server.ExtendConfig.NodeSolana.NodeSecrets, ",")
		nodeWeights = utils.Explode(server.ExtendConfig.NodeSolana.NodeWeights, ",")
		nodeKeys = utils.Explode(server.ExtendConfig.NodeSolana.NodeKeys, ",")
	default:
		nodeUrls = utils.Explode(server.ExtendConfig.Node.NodeUrls, ",")
		nodeSecrets = utils.Explode(server.ExtendConfig.Node.NodeSecrets, ",")
		nodeWeights = utils.Explode(server.ExtendConfig.Node.NodeWeights, ",")
		nodeKeys = utils.Explode(server.ExtendConfig.Node.NodeKeys, ",")
	}

	if len(nodeUrls) != len(nodeSecrets) || len(nodeUrls) != len(nodeWeights) || len(nodeUrls) != len(nodeKeys) {
		panic("NodeUrls,NodeSecrets,NodeWeights,NodeKeys must have the same length")
	}
	var nodes []Node
	total_weight := 0
	for i := 0; i < len(nodeUrls); i++ {
		w := utils.ItoInt(nodeWeights[i])
		if w < 0 {
			panic("NodeWeights must be greater than 0")
		}
		n := Node{
			URL:    nodeUrls[i],
			Secret: nodeSecrets[i],
			Weight: w,
			Key:    nodeKeys[i],
		}
		client, err := ethclient.Dial(n.URL)
		if err != nil {
			panic(err)
		}
		n.Client = client
		n.ChainID = chainId
		nodes = append(nodes, n)
		total_weight += w
	}
	if total_weight != 100 {
		panic("Total weight of all nodes must be 100")
	}
	balancertmp := NewNodeBalancer(nodes, server)
	return balancertmp, nil
}

// GetNode chọn node theo trọng số, loại trừ node fail trước đó
func (b *NodeBalancer) GetNode(excluded map[string]bool) (Node, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	availableNodes := []Node{}
	availableWeights := []int{}
	sumWeight := 0

	// Lọc ra những node chưa bị fail
	for _, node := range b.nodes {
		if !excluded[node.Key] {
			sumWeight += node.Weight
			availableNodes = append(availableNodes, node)
			availableWeights = append(availableWeights, sumWeight)
		}
	}

	// Nếu không còn node nào khả dụng → return rỗng
	if len(availableNodes) == 0 {
		return Node{}, errors.New("No available node")
	}

	// Chọn node theo trọng số (Weighted Random)
	r := rand.Intn(sumWeight)
	for i, w := range availableWeights {
		if r < w {
			return availableNodes[i], nil
		}
	}

	return availableNodes[len(availableNodes)-1], nil // Fallback
}

// ko có loadbalancer nếu chọn fixed node
func (b *NodeBalancer) GetNodeByKey(key string) (Node, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, node := range b.nodes {
		if node.Key == key {
			return node, nil
		}
	}
	return Node{}, errors.New("Node not found")
}
