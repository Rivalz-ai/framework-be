package service

import (
	"github.com/Rivalz-ai/framework-be/framework/pubsub/constant"
	"github.com/Rivalz-ai/framework-be/framework/utils"
	"gorm.io/gorm"
)

/*
APM
*/
type ApmConfig struct {
	Enable bool `json:"enable,omitempty"`
}

/*
MongoDB config
*/
type MongoConfig struct {
	Enable bool `json:"enable,omitempty"`
}

/*
Influx config
*/
type InfluxConfig struct {
	Model map[string]interface{} `json:"model,omitempty"`
}

/*
Redis config
*/
type RedisConfig struct {
	Enable bool `json:"enable,omitempty"`
}

/*
Redis WS config
*/
type WSRedisConfig struct {
	Enable bool `json:"enable,omitempty"`
}

/*
ENV
*/
type ENVConfig struct {
	RequireAuthen bool `json:"enable,omitempty"`
}

/*
Publisher
*/
type PubConfig struct {
	Enable             bool                        `json:"enable,omitempty"`
	Schema             interface{}                 `json:"pub_schema,omitempty"`
	PublisherName      string                      `json:"publisher_name,omitempty"`
	TopicPath          string                      `json:"topic_path,omitempty"`
	BatcherBuilderType constant.BatcherBuilderType `json:"batcher_builder_type,omitempty"`
	CompressionType    constant.CompressionType    `json:"compression_type,omitempty"`
	CompressionLevel   constant.CompressionLevel   `json:"compression_level,omitempty"`
	PubType            string                      `json:"pub_type,omitempty"`
}

/*
Consumer
*/
type SubConfig struct {
	Schema           interface{}               `json:"sub_schema,omitempty"`
	NoUserRetry      bool                      `json:"no_use_retry,omitempty"`
	Topic            string                    `json:"topic,omitempty"`
	ConsumerGroup    string                    `json:"consumer_group,omitempty"`
	NumConsumer      string                    `json:"num_consumer,omitempty"`
	ConsumerType     constant.ConsumerType     `json:"consumer_type,omitempty"`
	ConsumerPosition constant.ConsumerPosition `json:"consumer_position,omitempty"`
	SubType          string                    `json:"sub_type,omitempty"`
}

/*
ElasticSearch config
*/
type ElasticSearchConfig struct {
	Enable bool `json:"enable,omitempty"`
}

/*
Postgres config
*/
type PostgresConfig struct {
	Enable     bool         `json:"enable,omitempty"`
	GormConfig *gorm.Config `json:"gorm_config,omitempty"`
}

/*
GenericMicroClientConfig
*/
type GenericMicroClientConfig struct {
	Client map[string]string `json:"generic_micro_client,omitempty"`
}

/*
GRPC white_list_method
*/
type GRPCConfig struct {
	WhiteListMethod []string `json:"white_list_method,omitempty"`
}

func ArgsToMapConfig(args []interface{}) map[string]interface{} {
	if len(args) == 0 {
		return nil
	}
	m := utils.Dictionary()
	for _, arg := range args {
		switch utils.GetType(arg) { //for grpc, if change this need to change in micro
		case "ENVConfig":
			obj := utils.ValuePTR(arg).(*ENVConfig)
			m["require_authen"] = obj.RequireAuthen
		case "ApmConfig":
			obj := utils.ValuePTR(arg).(*ApmConfig)
			m["apm"] = obj.Enable
		case "MongoConfig":
			obj := utils.ValuePTR(arg).(*MongoConfig)
			m["mongo"] = obj.Enable
		case "InfluxConfig":
			obj := utils.ValuePTR(arg).(*InfluxConfig)
			m["model_influx"] = obj.Model
		case "RedisConfig":
			obj := utils.ValuePTR(arg).(*RedisConfig)
			m["redis"] = obj.Enable
		case "WSRedisConfig":
			obj := utils.ValuePTR(arg).(*WSRedisConfig)
			m["redis_ws_event"] = obj.Enable
		case "ElasticSearchConfig":
			obj := utils.ValuePTR(arg).(*ElasticSearchConfig)
			m["elasticsearch"] = obj.Enable
		case "PostgresConfig":
			obj := utils.ValuePTR(arg).(*PostgresConfig)
			m["postgres"] = obj.Enable
		case "PubConfig":
			obj := utils.ValuePTR(arg).(*PubConfig)
			m["pub"] = obj.Enable
			m["pub_schema"] = obj.Schema
			m["publisher_name"] = obj.PublisherName
			m["topic_path"] = obj.TopicPath
			m["batcher_builder_type"] = obj.BatcherBuilderType
			m["compression_type"] = obj.CompressionType
			m["compression_level"] = obj.CompressionLevel
			m["pub_type"] = obj.PubType
		case "SubConfig":
			obj := utils.ValuePTR(arg).(*SubConfig)
			m["sub_schema"] = obj.Schema
			m["no_use_retry"] = obj.NoUserRetry
			m["topic"] = obj.Topic
			m["consumer_group"] = obj.ConsumerGroup
			m["num_consumer"] = obj.NumConsumer
			m["consumer_type"] = obj.ConsumerType
			m["consumer_position"] = obj.ConsumerPosition
			m["sub_type"] = obj.SubType
		case "GenericMicroClientConfig":
			obj := utils.ValuePTR(arg).(*GenericMicroClientConfig)
			m["generic_micro_client"] = obj.Client
		case "GRPCConfig":
			obj := utils.ValuePTR(arg).(*GRPCConfig)
			m["white_list_method"] = obj.WhiteListMethod
		//case "":
		default:
			//ignore
		}
	}
	return m
}
func GetArgByType(args []interface{}, t string) (int, interface{}) {
	for index, arg := range args {
		argType := utils.GetType(arg)
		if argType == t {
			switch argType {
			case "ENVConfig":
				return index, utils.ValuePTR(arg).(*ENVConfig)
			case "ApmConfig":
				return index, utils.ValuePTR(arg).(*ApmConfig)
			case "MongoConfig":
				return index, utils.ValuePTR(arg).(*MongoConfig)
			case "InfluxConfig":
				return index, utils.ValuePTR(arg).(*InfluxConfig)
			case "RedisConfig":
				return index, utils.ValuePTR(arg).(*RedisConfig)
			case "WSRedisConfig":
				return index, utils.ValuePTR(arg).(*WSRedisConfig)
			case "ElasticSearchConfig":
				return index, utils.ValuePTR(arg).(*ElasticSearchConfig)
			case "PostgresConfig":
				return index, utils.ValuePTR(arg).(*PostgresConfig)
			case "PubConfig":
				return index, utils.ValuePTR(arg).(*PubConfig)
			case "SubConfig":
				return index, utils.ValuePTR(arg).(*SubConfig)
			case "GenericMicroClientConfig":
				return index, utils.ValuePTR(arg).(*GenericMicroClientConfig)
			case "GRPCConfig":
				return index, utils.ValuePTR(arg).(*GRPCConfig)
			default:
				//ignore
			}
		}
	}
	return 0, nil
}
func GetPgConfig(args []interface{}) *PostgresConfig {
	_, obj := GetArgByType(args, "PostgresConfig")
	if obj != nil {
		return obj.(*PostgresConfig)
	}
	return nil
}
