package kafka

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/Rivalz-ai/framework-be/framework/config/vault"
	"github.com/Rivalz-ai/framework-be/framework/log"
	vault_cfg "github.com/Rivalz-ai/framework-be/framework/pubsub/config"
	"github.com/Rivalz-ai/framework-be/framework/pubsub/constant"
	srv_cfg "github.com/Rivalz-ai/framework-be/framework/service"
	"github.com/Rivalz-ai/framework-be/framework/utils"
)

type Publisher struct {
	syncPublisher sarama.SyncProducer
	asynPublisher sarama.AsyncProducer
	topic         string
}

// Initial Publisher
func (pub *Publisher) Initial(vault *vault.Vault, pub_path string, args ...interface{}) error {
	log.Info("Initialing Kafka publisher...", "KAFKA")
	//
	if pub_path == "" {
		panic("Vault config path for Kafka is empty")
		return errors.New("Vault config path for Kafka is empty")
	}
	//global  config path
	globalConfigMap := vault_cfg.GetConfig(vault, "pubsub/pub/kafka")
	var configMap map[string]string
	var publisherName string
	//read local config from args
	cfg := srv_cfg.ArgsToMapConfig(args)
	//
	topic_path := ""
	if cfg != nil {
		//set Publisher name
		if val, ok := cfg["publisher_name"]; ok && len(val.(string)) > 0 {
			publisherName = utils.ItoString(cfg["publisher_name"])
		}
		//topic truyền từ ngoài vào trong trường hợp publish nhiều topic nên ko lấy topic từ vault
		if val, ok := cfg["topic_path"]; ok && len(val.(string)) > 0 {
			topic_path = utils.ItoString(val)
		}
	}
	// TODO: check name
	if topic_path != "" { //=>multi topic, config = .../pub/kafka/genral + .../pub/kafka/event(topic)
		//if event_path exist then load pub config from event path instead service path
		localConfigPath := pub_path + "/general"
		topicConfigPath := pub_path + "/" + topic_path
		check, err := vault.CheckItemExist(localConfigPath)
		if err != nil {
			return errors.New(err.Error())
		}
		localConfigMap := utils.DictionaryString()
		if check {
			//local broker <service>/pub/kafka, user, password config, if not exist then use global config data/pubsub/pub/kafka
			localConfigMap = vault_cfg.GetConfig(vault, localConfigPath)
		}
		//only topic config
		topicConfigMap := vault_cfg.GetConfig(vault, topicConfigPath)
		//
		localConfigMap = vault_cfg.MergeConfig(localConfigMap, topicConfigMap)
		//merge global + local config
		configMap = vault_cfg.MergeConfig(globalConfigMap, localConfigMap)
	} else { //=> single topic all config in .../pub/kafka
		localConfigMap := vault_cfg.GetConfig(vault, pub_path)
		//merge global + local config
		configMap = vault_cfg.MergeConfig(globalConfigMap, localConfigMap)
	}
	if !vault_cfg.ValidateConnectionInfo(configMap) {
		return errors.New("Brokers or Topics not found. Please help to check again !")
	}
	//
	confSync := NewProducerConfig(configMap, 0)
	confAsyncSync := NewProducerConfig(configMap, 1)
	brokers_str := configMap["BROKERS"]
	pub.topic = configMap["TOPIC"]
	//
	var err error
	brokers := utils.Explode(brokers_str, ",")
	fmt.Println("Kafka Connecting to brokers: ", brokers_str, " | topic: ", pub.topic)
	//init publisher
	pub.syncPublisher, err = sarama.NewSyncProducer(brokers, confSync)
	if err != nil {
		panic(err)
		return errors.New(err.Error())
	}
	_ = confAsyncSync
	//init asyn publisher
	/*pub.asynPublisher,err=sarama.NewAsyncProducer(brokers, confAsyncSync)
	if err!=nil{
		panic(err)
		return errors.New(err.Error())
	}*/
	log.Info(fmt.Sprintf("%s %s %s %s: %s", "Kafka publisher ", publisherName, " brokers: ", brokers_str+" | "+pub.topic, " connected"), "KAFKA", "PUBLISHER")
	return nil
}

// publish message
func (pub *Publisher) Publish(data interface{}, args ...string) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{Topic: pub.topic, Value: sarama.StringEncoder(json)}
	if len(args) > 0 {
		key := args[0]
		msg.Key = sarama.StringEncoder(key)
	}
	_, _, errp := pub.syncPublisher.SendMessage(msg)
	if errp != nil {
		return errp
	}
	return nil
}

// publish message
type AsyncPubError struct {
	Id    string
	Error error
}

func (pub *Publisher) PublishAsync(data interface{}, id string, errorCh chan<- AsyncPubError, successChan chan<- string, msgKey string, args ...string) {
	json, err := json.Marshal(data)
	if err != nil {
		log.Error("Error when marshal data", "PublishAsync")
		return
	}
	msg := sarama.ProducerMessage{
		Topic:    pub.topic,
		Value:    sarama.StringEncoder(string(json)),
		Metadata: id,
	}
	if len(args) > 0 {
		key := args[0]
		msg.Key = sarama.StringEncoder(key)
	}
	pub.asynPublisher.Input() <- &msg
	select {
	case suc := <-pub.asynPublisher.Successes():
		successChan <- suc.Metadata.(string)
	case err := <-pub.asynPublisher.Errors():
		errobj := AsyncPubError{
			Id:    err.Msg.Metadata.(string),
			Error: err.Err,
		}
		errorCh <- errobj
	}
}

func (pub *Publisher) SetDisableBlockIfQueueFull(v bool) {
}

func (pub *Publisher) SetMaxPendingMessages(v int) {
}

func (pub *Publisher) SetDisableBatching(v bool) {
}

func (pub *Publisher) SetBatchingMaxPublishDelay(v time.Duration) {
}

func (pub *Publisher) SetBatchingMaxMessages(v uint) {
}

func (pub *Publisher) SetBatchingMaxSize(v uint) {
}

func (pub *Publisher) SetBatcherBuilderType(v constant.BatcherBuilderType) {
}

func (pub *Publisher) SetPublisherSchema(v string) {
}

func (pub *Publisher) SetSendTimeout(v time.Duration) {
}

func (pub *Publisher) SetConnectionTimeout(v time.Duration) {
}

func (pub *Publisher) SetOperationTimeout(v time.Duration) {
}

func (pub *Publisher) SetKeepAliveInterval(v time.Duration) {
}

func (pub *Publisher) SetMaxConnectionPerBrokers(v int) {
}

func (pub *Publisher) SetCompressionType(v constant.CompressionType) {
}
func (pub *Publisher) SetCompressionLevel(v constant.CompressionLevel) {
}

func (pub *Publisher) SetHashScheme(v constant.HashingScheme) {
}

func (pub *Publisher) SetDisableMultiSchema(v bool) {
}

func (pub *Publisher) SetIsOnlyPushUid(v bool) {
}

func (pub *Publisher) GetPublisherSchema() interface{} {
	return nil
}

func (pub *Publisher) GetTopicName() string {
	return ""
}
