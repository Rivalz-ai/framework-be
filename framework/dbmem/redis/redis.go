package redis

import (
	"fmt"
	"strings"

	"github.com/Rivalz-ai/framework-be/framework/config/vault"
	rcf "github.com/Rivalz-ai/framework-be/framework/dbmem/redis/config"
	"github.com/Rivalz-ai/framework-be/framework/utils"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Config *vault.Vault
	Client *redis.Client
}

func NewClient(addr, password string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}
func (r *Redis) Initial(config *vault.Vault) {
	fmt.Println("Initial Redis")
	if r.Config == nil {
		r.Config = config
	}
	serviceName := r.Config.GetServiceName()
	servicePath := strings.ReplaceAll(serviceName, ".", "/")
	check, err := r.Config.CheckItemExist(servicePath + "/redis")
	if err != nil {
		panic("Redis Vaul Config not Exist : " + err.Error())
	}
	if !check {
		panic("Redis Vaul Config not Exist")
	}
	configPath := fmt.Sprintf("%s/%s", servicePath, "redis")
	configMap := rcf.GetConfig(r.Config, configPath)
	if configMap == nil {
		panic("Redis Config not Exist")
	}
	if configMap["HOST"] == "" {
		panic("Redis Host not Exist")
	}
	if configMap["DB"] == "" {
		panic("Redis DB not Exist")
	}
	r.Client = NewClient(configMap["HOST"], configMap["PASSWORD"], utils.ItoInt(configMap["DB"]))
	if r.Client == nil {
		panic("Redis Client Init fail")
	}
	fmt.Println("Initial Redis Success")
}
