package kafka

import (
	"crypto/tls"
	"crypto/x509"
	"strings"

	//"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/Rivalz-ai/framework-be/framework/log"
	//"github.com/Rivalz-ai/framework-be/framework/utils"
	//"github.com/Rivalz-ai/framework-be/framework/config/vault"
	"github.com/IBM/sarama"
	//for consumer
	sarama2 "github.com/Shopify/sarama"
	//"fmt"
	"time"

	"github.com/google/uuid"
)

/*
type=0: sync producer
type=1: async producer
*/

func NewProducerConfig(storeCfg map[string]string, pubType int) *sarama.Config {
	cfg := sarama.NewConfig()
	cfg.ClientID = "GK-1.0.1-Producer"
	cfg.Version = sarama.V3_3_2_0
	cfg.Metadata.AllowAutoTopicCreation = false
	cfg.Metadata.Full = true
	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true
	cfg.Producer.Retry.Max = 5
	cfg.Producer.Retry.Backoff = time.Millisecond * 10
	cfg.Producer.Compression = sarama.CompressionZSTD
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Flush.Frequency = 500 * time.Millisecond
	cfg.Net.MaxOpenRequests = 1
	if pubType == 1 {
		cfg.Producer.Idempotent = true
	}
	if storeCfg["USERNAME"] != "" && storeCfg["PASSWORD"] != "" {
		cfg.Net.SASL.Enable = true
		cfg.Net.SASL.User = storeCfg["USERNAME"]
		cfg.Net.SASL.Password = storeCfg["PASSWORD"]
		cfg.Net.SASL.Handshake = true
		/*if strings.ToLower(storeCfg["SHA_ALGORITHM"]) == "sha512" {
			conf.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA512} }
			conf.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512
		} else {
			conf.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA256} }
			conf.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA256

		}*/
	}
	if strings.ToLower(storeCfg["USE_SSL"]) == "yes" || strings.ToLower(storeCfg["USE_SSL"]) == "true" {
		cfg.Net.TLS.Enable = true
		cfg.Net.TLS.Config = createTLSConfiguration(storeCfg)
	}
	return cfg
}
func NewConsumerConfig(storeCfg map[string]string) *sarama2.Config {
	conf := sarama2.NewConfig()
	conf.Metadata.Full = true
	//conf.Version = sarama.V0_10_0_0
	conf.ClientID = "Platform-1.0.1-Consumer" + "-" + uuid.New().String()
	conf.Metadata.Full = true
	conf.Consumer.Group.Session.Timeout = 30 * time.Second
	conf.Consumer.MaxProcessingTime = 60 * 60 * time.Second
	conf.Consumer.Group.Rebalance.Timeout = 60 * time.Second
	conf.Consumer.Group.Heartbeat.Interval = 3 * time.Second
	if storeCfg["USERNAME"] != "" && storeCfg["PASSWORD"] != "" {
		conf.Net.SASL.Enable = true
		conf.Net.SASL.User = storeCfg["USERNAME"]
		conf.Net.SASL.Password = storeCfg["PASSWORD"]
		conf.Net.SASL.Handshake = true
		/*if strings.ToLower(storeCfg["SHA_ALGORITHM"]) == "sha512" {
			conf.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA512} }
			conf.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512
		} else {
			conf.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA256} }
			conf.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA256

		}*/
	}
	if strings.ToLower(storeCfg["USE_SSL"]) == "yes" || strings.ToLower(storeCfg["USE_SSL"]) == "true" {
		conf.Net.TLS.Enable = true
		conf.Net.TLS.Config = createTLSConfiguration(storeCfg)
	}
	return conf
}
func createTLSConfiguration(storeCfg map[string]string) (t *tls.Config) {
	verify_ssl := false
	if strings.ToLower(storeCfg["VERIFY_SSL"]) == "true" || strings.ToLower(storeCfg["VERIFY_SSL"]) == "yes" {
		verify_ssl = true
	}
	t = &tls.Config{
		InsecureSkipVerify: verify_ssl,
	}
	if storeCfg["SSL_CERT"] != "" && storeCfg["SSL_KEY"] != "" && storeCfg["SSL_CA"] != "" {
		cert, err := tls.X509KeyPair([]byte(storeCfg["SSL_CERT"]), []byte(storeCfg["SSL_KEY"]))
		if err != nil {
			log.Error(err.Error(), "KAFKA_SSL_CONFIG")
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM([]byte(storeCfg["SSL_CA"]))
		verify_ssl := false
		if storeCfg["VERIFY_SSL"] != "" {
			if strings.ToLower(storeCfg["VERIFY_SSL"]) == "yes" || strings.ToLower(storeCfg["VERIFY_SSL"]) == "true" {
				verify_ssl = true
			}
		}
		t = &tls.Config{
			Certificates:       []tls.Certificate{cert},
			RootCAs:            caCertPool,
			InsecureSkipVerify: verify_ssl,
		}
	}
	return t
}

func NewProducerConfigShopify(storeCfg map[string]string, pubType int) *sarama2.Config {
	cfg := sarama2.NewConfig()
	cfg.ClientID = "GK-1.0.1-Producer" + "-" + uuid.New().String()
	cfg.Version = sarama2.V2_1_0_0
	cfg.Metadata.AllowAutoTopicCreation = false
	cfg.Metadata.Full = true
	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true
	cfg.Producer.Retry.Max = 5
	cfg.Producer.Retry.Backoff = time.Millisecond * 10
	cfg.Producer.Compression = sarama2.CompressionZSTD
	cfg.Producer.RequiredAcks = sarama2.WaitForAll
	cfg.Producer.Flush.Frequency = 500 * time.Millisecond
	cfg.Net.MaxOpenRequests = 1
	if pubType == 1 {
		cfg.Producer.Idempotent = true
	}
	if storeCfg["USERNAME"] != "" && storeCfg["PASSWORD"] != "" {
		cfg.Net.SASL.Enable = true
		cfg.Net.SASL.User = storeCfg["USERNAME"]
		cfg.Net.SASL.Password = storeCfg["PASSWORD"]
		cfg.Net.SASL.Handshake = true
		/*if strings.ToLower(storeCfg["SHA_ALGORITHM"]) == "sha512" {
			conf.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA512} }
			conf.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512
		} else {
			conf.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA256} }
			conf.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA256

		}*/
	}
	if strings.ToLower(storeCfg["USE_SSL"]) == "yes" || strings.ToLower(storeCfg["USE_SSL"]) == "true" {
		cfg.Net.TLS.Enable = true
		cfg.Net.TLS.Config = createTLSConfiguration(storeCfg)
	}
	return cfg
}
