package kafka

import (
	"context"
	"encoding/json"

	vault_cfg "github.com/Rivalz-ai/framework-be/framework/pubsub/config"
	srv_cfg "github.com/Rivalz-ai/framework-be/framework/service"
	"github.com/Rivalz-ai/framework-be/framework/utils"
	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"

	//"github.com/Rivalz-ai/framework-be/framework/utils/transform"
	//"github.com/Rivalz-ai/framework-be/framework/health"
	"os"

	"github.com/Rivalz-ai/framework-be/framework/base/event"
	ev "github.com/Rivalz-ai/framework-be/framework/base/event"

	//"github.com/Rivalz-ai/framework-be/framework/cache/redis"
	"errors"
	"fmt"

	"github.com/Rivalz-ai/framework-be/framework/config/vault"
	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/Rivalz-ai/framework-be/framework/pubsub/constant"
	"github.com/ThreeDotsLabs/watermill"

	//"math"
	"os/signal"
	"sync"
	"syscall"
	"time"

	//"errors"
	"strings"
)

// type ConsumeFn = func(messages <-chan *message.Message)
type Subscriber struct {
	//uniqueue id of subscriber, it will be worker_name+hostname
	id string
	//subscriber           pulsar.Consumer
	//topic where consumer will consume message: if multi topic then use "," for separate
	topic string
	//callback func for process message
	ProcessFn event.ConsumeFn
	//number of consumer per pod, 2pod wil be 2xnum_consumer, each consumer is a goroutine
	num_consumer int
	mu           sync.Mutex
	//schedule count current number of consumer
	current_num_consumer int
	//inject consume_time, finish_time to message or not, default: wil inject
	no_inject bool
	//if item fail, will push item to Retry letter topic or not, default: will retry
	no_use_retry bool
	// if no_use_retry=true, how long fail item was process again(second)
	retry_delay int64 //default 60s
	num_retry   int
	//
	configMap map[string]string
	//option of go-pulsar client
	//clientCfg pulsar.ClientOptions
	//
	//consumerCfg pulsar.ConsumerOptions
	//after process item, if success log item to general topic sucess, default false
	//mark local variable
	is_local bool
	//
	dynamic_subscription_name bool
	//
	consume_type string           //defaut latest
	publisher    *kafka.Publisher //for retry message
}

// Initial Publisher
func (sub *Subscriber) Initial(vault *vault.Vault, sub_path string, worker_name string, callbackfn event.ConsumeFn, args ...interface{}) error {
	log.Info("Initialing Kakfa subscriber...", "KAFKA")
	if sub_path == "" {
		panic("Vault config path for Kafka is empty")
		return nil
	}
	//
	hostname, err_h := os.Hostname()
	if err_h != nil {
		log.Warn(fmt.Sprintf("Error get Hostname: %s", err_h.Error()), "SUBSCRIBER")
	}
	//prd, stg, dev default to send to retry queue if item fail
	//local Env resend message
	if os.Getenv("ENV") == "local" {
		sub.SetNoUseRetry(true)
		sub.is_local = true
	}
	sub.id = worker_name + "_" + hostname
	//global event bus config path
	var localConfigMap, configMap map[string]string
	globalConfigMap := vault_cfg.GetConfig(vault, "pubsub/sub/kafka")
	if sub_path != "" {
		localConfigMap = vault_cfg.GetConfig(vault, sub_path)
		configMap = vault_cfg.MergeConfig(globalConfigMap, localConfigMap)
	} else {
		configMap = globalConfigMap
	}

	//read local config from args
	cfg := srv_cfg.ArgsToMapConfig(args)
	if cfg != nil {
		if val, ok := cfg["consumer_group"]; ok && len(val.(string)) > 0 {
			configMap["CONSUMER_GROUP"] = val.(string)
		}
		if val, ok := cfg["topic"]; ok && len(val.(string)) > 0 {
			configMap["TOPIC"] = val.(string)
		}
		if val, ok := cfg["num_consumer"]; ok && len(val.(string)) > 0 {
			configMap["NUM_CONSUMER"] = val.(string)
		}
		//retry enable => repush to topic again for retry later, if still can not success => store to db
		if configMap["NO_RETRY"] == "true" {
			sub.SetNoUseRetry(true)
		} else if configMap["NO_RETRY"] == "false" {
			sub.SetNoUseRetry(false)
		}
	}
	if !vault_cfg.ValidateConsumerInfo(configMap) {
		panic("Configs for consumers were wrong. Please check again!")
		return nil
	}
	sub.configMap = configMap
	sub.num_consumer = utils.ItoInt(sub.configMap["NUM_CONSUMER"])
	sub.topic = configMap["TOPIC"]
	//default retry
	if !sub.no_use_retry {
		sub.num_retry = 5
		if configMap["NUM_RETRY"] != "" {
			i_retry_time := utils.ItoInt(configMap["NUM_RETRY"])
			sub.num_retry = int(i_retry_time)
			if sub.num_retry <= 0 {
				sub.num_retry = 5
			}
		}
	}
	//init publisher for retry
	conf := NewProducerConfigShopify(configMap, 0)
	brokers_str := configMap["BROKERS"]
	brokers := utils.Explode(brokers_str, ",")
	fmt.Println("Kafka Connecting to brokers: ", brokers_str, " | topic: ", sub.topic)
	var err error
	sub.publisher, err = kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:               brokers,
			Marshaler:             kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: conf,
		},
		//watermill.NewStdLogger(false, false),
		nil,
	)
	//not production
	if err != nil {
		panic("Error Kafka Publisher, please check kafka configuration: " + err.Error())
		return errors.New(err.Error())
	}
	//
	if cfg != nil && cfg["consumer_position"] != nil && sub.consume_type == "" { //not set by code, set by code higher priority
		sub.consume_type = strings.ToLower(utils.ItoString(cfg["consumer_position"]))
	}
	//set default if not set
	if sub.consume_type == "" {
		sub.consume_type = "latest"
	}
	sub.ProcessFn = callbackfn
	//
	log.Info(fmt.Sprintf("%s %s: %s", "-Kafka consumer brokers: ", sub.configMap["BROKERS"]+" | "+sub.topic, " connected"), "CONSUMER")
	return nil
}

// Set call back function
func (sub *Subscriber) WithCallBackFn(fn event.ConsumeFn) {
	sub.ProcessFn = fn
}

// consume message
func (sub *Subscriber) Consume() error {
	log.Info(fmt.Sprintf("Number of cosumer: %s", utils.ItoString(sub.num_consumer)))
	brokers_str := sub.configMap["BROKERS"]
	brokers := utils.Explode(brokers_str, ",")
	consumer_group := sub.configMap["CONSUMER_GROUP"]
	//number of goroutine
	for i := 0; i < sub.num_consumer; i++ {
		//config
		conf := NewConsumerConfig(sub.configMap)
		//sarama.OffsetOldest get from last offset not yet commit
		//sarama.OffsetNewest  ignore all mesage just get new message after consumer start
		if sub.consume_type == "latest" {
			conf.Consumer.Offsets.Initial = sarama.OffsetNewest
		} else {
			conf.Consumer.Offsets.Initial = sarama.OffsetOldest
		}
		//var err Error
		subscriber, err := kafka.NewSubscriber( //reconfig subscriber because want to each goroutine(subscriber has differerence client_ID)
			kafka.SubscriberConfig{
				Brokers:               brokers,
				Unmarshaler:           kafka.DefaultMarshaler{},
				OverwriteSaramaConfig: conf,
				ConsumerGroup:         consumer_group,
			},
			//watermill.NewStdLogger(false, false),
			nil,
		)
		if err != nil {
			panic("Error create NewSubscriber, please check subscriber configuration: " + err.Error())
			return nil
		}
		//
		messages, err := subscriber.Subscribe(context.Background(), sub.topic)
		if err != nil {
			panic("Error Subscribe to Topic, please check subscriber configuration: " + err.Error())
			return nil
		}
		//
		go sub.ProcessMesasge(i+1, messages)
	}
	c := make(chan os.Signal)
	//
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	//
	close(c)
	log.Info("Consumers stop", "Consumer")
	return nil
}

func (sub *Subscriber) ProcessMesasge(i int, messages <-chan *message.Message) {
	//defer *health=false
	log.Info(fmt.Sprintf("Consumer: %s-[%s] started", sub.id, utils.ItoString(i)))

	for msg := range messages {
		var event ev.Event
		var err error
		//process message
		event, err = InjectComsumeTime(msg) //insert consume_time to event_header, to trace later
		if err != nil {
			log.Error(err.Error(), "Consumer", "ProcessMesasge", event)
			//just log no need stop process message
		}
		event, err = ExtractEvent(msg)
		if err != nil { //message format is worng, can not parse to event => ignore message
			log.Error(err.Error(), "Consumer", "ProcessMesasge", event)
			msg.Ack()
			return
		}
		err_p := sub.ProcessFn(event, context.Background()) // => callback Fn
		if err_p != nil {
			log.Error(err_p.Error(), "ProcessMesasge", event)
			//push message to retry topic
			//default rule
			if !sub.no_use_retry && !sub.is_local { //server ENV, and set no_use_retry=false, send message to Retry Letter Topic for reconsume latter
				//add log retry to event, to trace reason why this message was retried
				ev_header, err := utils.ItoDictionary(event.EventHeader)
				counter := 0
				if err == nil {
					if ev_header["log"] != nil && ev_header["log"] != "" {
						old_error := utils.ItoString(ev_header["log"])
						ev_header["log"] = old_error + " | " + err_p.Error()
					} else {
						ev_header["log"] = err_p.Error()
					}
					//get couter of message
					if ev_header["counter"] != nil {
						counter = utils.ItoInt(ev_header["counter"])
						counter += 1
					}
					ev_header["counter"] = counter
					event.EventHeader = ev_header
				}
				//push message to topic again, for retry consume later, we can not control retry time on kafka
				if sub.publisher != nil && counter < sub.num_retry {
					err := sub.RePublish(event, "")
					if err != nil {
						log.Error(err.Error(), "Consumer", "ProcessMesasge", event)
					}
				}
				msg.Ack()
			} else if !sub.is_local { //server ENV,and set no_use_retry=true => ACK, and do not send to Retry Letter Topic
				msg.Ack()
			} else { //local ENV, send Non ACK => pulsar resend again message forever(because pulsar think message was loss)
				msg.Nack()
			}
		} else {
			msg.Ack()
		}
	}
	log.Info(fmt.Sprintf("Consumer: %s-[%s] shutdown", sub.id, utils.ItoString(i)))
}
func (sub *Subscriber) RePublish(event ev.Event, msgKey string) error {
	data, err := json.Marshal(event)
	if err != nil {
		//log.Error(err.Error(),"EVENT_DRIVEN_SERIALIZE")
		return errors.New(err.Error())
	}
	if sub.publisher == nil {
		return errors.New("Publisher has nil")
	}
	msg := message.NewMessage(watermill.NewUUID(), data)
	err_p := sub.publisher.Publish(sub.topic, msg)
	if err_p != nil {
		return errors.New(err_p.Error())
	}
	return nil
}
func (sub *Subscriber) SetNoAck(no_ack bool) {
	//sub.no_ack = no_ack
}

func (sub *Subscriber) SetNoInject(no_inject bool) {
	sub.no_inject = no_inject
}
func (sub *Subscriber) SetNoCache(no_cache bool) {
	//sub.no_cache = no_cache
}
func (sub *Subscriber) GetConsumerSchema() interface{} {
	return nil
}
func (sub *Subscriber) GetRetryEnable() bool {
	return sub.no_use_retry

}
func (sub *Subscriber) GetTopicName() string {
	return sub.topic
}
func (sub *Subscriber) SetNoUseRetry(v bool) {
	sub.no_use_retry = v
}
func (sub *Subscriber) SetConsumerSchema(v string) {
	//sub.consumerCfg.Schema = pulsar.NewJSONSchema(v, nil)
}
func (sub *Subscriber) SetRetryEnable(v bool) {
	//sub.consumerCfg.RetryEnable = v
}
func (sub *Subscriber) SetSubscriptionType(v constant.ConsumerType) {
	/*
		switch v {
		case constant.Exclusive:
			sub.consumerCfg.Type = pulsar.Exclusive
		case constant.KeyShared:
			sub.consumerCfg.Type = pulsar.KeyShared
			sub.consumerCfg.KeySharedPolicy=&pulsar.KeySharedPolicy {
				//Mode: pulsar.KeySharedPolicyModeAutoSplit,
				AllowOutOfOrderDelivery: true,
			}
		case constant.Failover:
		default:
			sub.consumerCfg.Type = pulsar.Shared
		}*/
}
func (sub *Subscriber) SetDQLConfig(retryLetterTopic, deadLetterTopic string, maxDeliveries uint32) {
	/*sub.consumerCfg.DLQ = &pulsar.DLQPolicy{
		RetryLetterTopic: retryLetterTopic,
		DeadLetterTopic:  deadLetterTopic,
		MaxDeliveries:    maxDeliveries,
	}*/
}
func (sub *Subscriber) SetLogFail(v bool) {
	//sub.log_item_fail = v
}
func (sub *Subscriber) SetLogGeneralSuccess(v bool) {
	//sub.log_item_general_success = v
}
func (sub *Subscriber) SetLogSuccess(v bool) {
	//sub.log_item_success = v
}
func (sub *Subscriber) SetReceiverQueueSize(v int) {
	//sub.consumerCfg.ReceiverQueueSize = v
}
func (sub *Subscriber) SetSubscriptionInitialPosition(v string) {
	//sub.consume_type = v
}
func (sub *Subscriber) SetSubscriptionName(v string) {
	//sub.consumerCfg.SubscriptionName = v
}
func (sub *Subscriber) Clean() {
	//sub.Redis.Close()
}
func ExtractEvent(messages *message.Message) (ev.Event, error) {
	//
	event := ev.Event{}
	err := json.Unmarshal([]byte(messages.Payload), &event)
	if err != nil {
		return event, errors.New(err.Error())
	}
	return event, nil
	//
}
func InjectComsumeTime(messages *message.Message) (ev.Event, error) {
	//
	event, err := ExtractEvent(messages)
	if err != nil {
		return event, err
	}
	event_header_map := utils.Dictionary()
	if event.EventHeader != nil {
		var err_c error
		event_header_map, err_c = utils.ItoDictionary(event.EventHeader)
		if err_c != nil {
			return event, errors.New(err_c.Error())
		}
	}
	event_header_map["consume_time"] = time.Now()
	event.EventHeader = event_header_map
	return event, nil
}

/*
func InjectWorkerName(messages *message.Message, worker_name string)  error {
	//
	event, err := ExtractEvent(messages)
	if err != nil {
		return err
	}
	//event.ProcessedFlow = event.ProcessedFlow + "->" + worker_name
	//
	data, err_m := json.Marshal(&event)
	if err_m != nil {
		//log.Error(err.Error(),"EVENT_DRIVEN_SERIALIZE")
		return e.New(err_m.Error(), "EVENT_DRIVEN", "MARSHAL_EVENT")
	}
	messages.Payload = []byte(data)
	return nil
	//msg := message.NewMessage(watermill.NewUUID(), data)
}
*/
/*
func InjectFinishTime(event *ev.Event) {
	if event != nil {
		event.FinishTime = time.Now()
		event.ProcessingTime = (time.Now().Sub(event.ConsumeTime).Seconds())
	}
}
*/
