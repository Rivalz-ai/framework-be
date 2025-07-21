package config

import (
	"fmt"
	"math"
	"strings"

	"github.com/Rivalz-ai/framework-be/framework/config/vault"
	"github.com/Rivalz-ai/framework-be/framework/utils"
)

func GetConfig(vault *vault.Vault, config_path string) map[string]string {
	config_path = strings.TrimSuffix(config_path, "/")
	m := utils.DictionaryString()
	m["BROKERS"] = vault.ReadVAR(fmt.Sprintf("%s/BROKERS", config_path))
	m["TOPIC"] = vault.ReadVAR(fmt.Sprintf("%s/TOPIC", config_path))
	m["USERNAME"] = vault.ReadVAR(fmt.Sprintf("%s/USERNAME", config_path))
	m["PASSWORD"] = vault.ReadVAR(fmt.Sprintf("%s/PASSWORD", config_path))
	m["USE_SSL"] = vault.ReadVAR(fmt.Sprintf("%s/USE_SSL", config_path))
	m["SSL_CA"] = vault.ReadVAR(fmt.Sprintf("%s/SSL_CA", config_path))
	m["SSL_CERT"] = vault.ReadVAR(fmt.Sprintf("%s/SSL_CERT", config_path))
	m["SSL_KEY"] = vault.ReadVAR(fmt.Sprintf("%s/SSL_KEY", config_path))
	m["VERIFY_SSL"] = vault.ReadVAR(fmt.Sprintf("%s/VERIFY_SSL", config_path))
	m["NUM_CONSUMER"] = vault.ReadVAR(fmt.Sprintf("%s/NUM_CONSUMER", config_path))
	m["CONSUMER_GROUP"] = vault.ReadVAR(fmt.Sprintf("%s/CONSUMER_GROUP", config_path))
	m["CONSUMER_TYPE"] = vault.ReadVAR(fmt.Sprintf("%s/CONSUMER_TYPE", config_path))
	m["TOKEN"] = vault.ReadVAR(fmt.Sprintf("%s/TOKEN", config_path))
	m["DEAD_LETTER_TOPIC"] = vault.ReadVAR(fmt.Sprintf("%s/DEAD_LETTER_TOPIC", config_path))
	m["RETRY_LETTER_TOPIC"] = vault.ReadVAR(fmt.Sprintf("%s/RETRY_LETTER_TOPIC", config_path))
	m["NO_RETRY"] = vault.ReadVAR(fmt.Sprintf("%s/NO_RETRY", config_path))
	m["RETRY_INTERVAL"] = vault.ReadVAR(fmt.Sprintf("%s/RETRY_INTERVAL", config_path))
	m["COMPRESS_TYPE"] = vault.ReadVAR(fmt.Sprintf("%s/COMPRESS_TYPE", config_path))
	m["NUM_RETRY"] = vault.ReadVAR(fmt.Sprintf("%s/NUM_RETRY", config_path))
	m["ALWAYS_LAST_CONSUME"] = vault.ReadVAR(fmt.Sprintf("%s/ALWAYS_LAST_CONSUME", config_path))
	m["SEND_TIMEOUT"] = vault.ReadVAR(fmt.Sprintf("%s/SEND_TIMEOUT", config_path))
	return m
}
func GetPubSubType(vault *vault.Vault, args ...interface{}) string {
	t := vault.ReadVAR("pubsub/general/TYPE")
	if t == "" {
		t = "pulsar"
	}
	return t
}
func MergeConfig(global, local map[string]string) map[string]string {
	m := utils.DictionaryString()
	if utils.Map_contains(global, "BROKERS") || utils.Map_contains(local, "BROKERS") {
		if utils.Map_contains(local, "BROKERS") && local["BROKERS"] != "" {
			m["BROKERS"] = local["BROKERS"]
		} else {
			m["BROKERS"] = global["BROKERS"]
		}
	}
	if utils.Map_contains(global, "TOPIC") || utils.Map_contains(local, "TOPIC") {
		if utils.Map_contains(local, "TOPIC") && local["TOPIC"] != "" {
			m["TOPIC"] = local["TOPIC"]
		} else {
			m["TOPIC"] = global["TOPIC"]
		}
	}
	if utils.Map_contains(global, "NUM_CONSUMER") || utils.Map_contains(local, "NUM_CONSUMER") {
		if utils.Map_contains(local, "NUM_CONSUMER") && local["NUM_CONSUMER"] != "" {
			m["NUM_CONSUMER"] = local["NUM_CONSUMER"]
		} else {
			m["NUM_CONSUMER"] = global["NUM_CONSUMER"]
		}
	}
	if utils.Map_contains(global, "USERNAME") || utils.Map_contains(local, "USERNAME") {
		if utils.Map_contains(local, "USERNAME") && local["USERNAME"] != "" {
			m["USERNAME"] = local["USERNAME"]
		} else {
			m["USERNAME"] = global["USERNAME"]
		}
	}
	if utils.Map_contains(global, "PASSWORD") || utils.Map_contains(local, "PASSWORD") {
		if utils.Map_contains(local, "PASSWORD") && local["PASSWORD"] != "" {
			m["PASSWORD"] = local["PASSWORD"]
		} else {
			m["PASSWORD"] = global["PASSWORD"]
		}
	}
	if utils.Map_contains(global, "USE_SSL") || utils.Map_contains(local, "USE_SSL") {
		if utils.Map_contains(local, "USE_SSL") && local["USE_SSL"] != "" {
			m["USE_SSL"] = local["USE_SSL"]
		} else {
			m["USE_SSL"] = global["USE_SSL"]
		}
	}
	if utils.Map_contains(global, "CONSUMER_GROUP") || utils.Map_contains(local, "CONSUMER_GROUP") {
		if utils.Map_contains(local, "CONSUMER_GROUP") && local["CONSUMER_GROUP"] != "" {
			m["CONSUMER_GROUP"] = local["CONSUMER_GROUP"]
		} else {
			m["CONSUMER_GROUP"] = global["CONSUMER_GROUP"]
		}
	}
	if utils.Map_contains(global, "SSL_CA") || utils.Map_contains(local, "SSL_CA") {
		if utils.Map_contains(local, "SSL_CA") && local["SSL_CA"] != "" {
			m["SSL_CA"] = local["SSL_CA"]
		} else {
			m["SSL_CA"] = global["SSL_CA"]
		}
	}
	if utils.Map_contains(global, "SSL_CERT") || utils.Map_contains(local, "SSL_CERT") {
		if utils.Map_contains(local, "SSL_CERT") && local["SSL_CERT"] != "" {
			m["SSL_CERT"] = local["SSL_CERT"]
		} else {
			m["SSL_CERT"] = global["SSL_CERT"]
		}
	}
	if utils.Map_contains(global, "SSL_KEY") || utils.Map_contains(local, "SSL_KEY") {
		if utils.Map_contains(local, "SSL_KEY") && local["SSL_KEY"] != "" {
			m["SSL_KEY"] = local["SSL_KEY"]
		} else {
			m["SSL_KEY"] = global["SSL_KEY"]
		}
	}
	if utils.Map_contains(global, "VERIFY_SSL") || utils.Map_contains(local, "VERIFY_SSL") {
		if utils.Map_contains(local, "VERIFY_SSL") && local["VERIFY_SSL"] != "" {
			m["VERIFY_SSL"] = local["VERIFY_SSL"]
		} else {
			m["VERIFY_SSL"] = global["VERIFY_SSL"]
		}
	}
	if utils.Map_contains(global, "TOKEN") || utils.Map_contains(local, "TOKEN") {
		if utils.Map_contains(local, "TOKEN") && local["TOKEN"] != "" {
			m["TOKEN"] = local["TOKEN"]
		} else {
			m["TOKEN"] = global["TOKEN"]
		}
	}
	if utils.Map_contains(global, "DEAD_LETTER_TOPIC") || utils.Map_contains(local, "DEAD_LETTER_TOPIC") {
		if utils.Map_contains(local, "DEAD_LETTER_TOPIC") && local["DEAD_LETTER_TOPIC"] != "" {
			m["DEAD_LETTER_TOPIC"] = local["DEAD_LETTER_TOPIC"]
		} else {
			m["DEAD_LETTER_TOPIC"] = global["DEAD_LETTER_TOPIC"]
		}
	}
	if utils.Map_contains(global, "RETRY_LETTER_TOPIC") || utils.Map_contains(local, "RETRY_LETTER_TOPIC") {
		if utils.Map_contains(local, "RETRY_LETTER_TOPIC") && local["RETRY_LETTER_TOPIC"] != "" {
			m["RETRY_LETTER_TOPIC"] = local["RETRY_LETTER_TOPIC"]
		} else {
			m["RETRY_LETTER_TOPIC"] = global["RETRY_LETTER_TOPIC"]
		}
	}
	if utils.Map_contains(global, "NO_RETRY") || utils.Map_contains(local, "NO_RETRY") {
		if utils.Map_contains(local, "NO_RETRY") && local["NO_RETRY"] != "" {
			m["NO_RETRY"] = local["NO_RETRY"]
		} else {
			m["NO_RETRY"] = global["NO_RETRY"]
		}
	}
	if utils.Map_contains(global, "RETRY_INTERVAL") || utils.Map_contains(local, "RETRY_INTERVAL") {
		if utils.Map_contains(local, "RETRY_INTERVAL") && local["RETRY_INTERVAL"] != "" {
			m["RETRY_INTERVAL"] = local["RETRY_INTERVAL"]
		} else {
			m["RETRY_INTERVAL"] = global["RETRY_INTERVAL"]
		}
	}
	if utils.Map_contains(global, "COMPRESS_TYPE") || utils.Map_contains(local, "COMPRESS_TYPE") {
		if utils.Map_contains(local, "COMPRESS_TYPE") && local["COMPRESS_TYPE"] != "" {
			m["COMPRESS_TYPE"] = local["COMPRESS_TYPE"]
		} else {
			m["COMPRESS_TYPE"] = global["COMPRESS_TYPE"]
		}
	}
	if utils.Map_contains(global, "NUM_RETRY") || utils.Map_contains(local, "NUM_RETRY") {
		if utils.Map_contains(local, "NUM_RETRY") && local["NUM_RETRY"] != "" {
			m["NUM_RETRY"] = local["NUM_RETRY"]
		} else {
			m["NUM_RETRY"] = global["NUM_RETRY"]
		}
	}
	if utils.Map_contains(global, "CONSUMER_TYPE") || utils.Map_contains(local, "CONSUMER_TYPE") {
		if utils.Map_contains(local, "CONSUMER_TYPE") && local["CONSUMER_TYPE"] != "" {
			m["CONSUMER_TYPE"] = local["CONSUMER_TYPE"]
		} else {
			m["CONSUMER_TYPE"] = global["CONSUMER_TYPE"]
		}
	}
	if utils.Map_contains(global, "ALWAYS_LAST_CONSUME") || utils.Map_contains(local, "ALWAYS_LAST_CONSUME") {
		if utils.Map_contains(local, "ALWAYS_LAST_CONSUME") && local["ALWAYS_LAST_CONSUME"] != "" {
			m["ALWAYS_LAST_CONSUME"] = local["ALWAYS_LAST_CONSUME"]
		} else {
			m["ALWAYS_LAST_CONSUME"] = global["ALWAYS_LAST_CONSUME"]
		}
	}
	if utils.Map_contains(global, "SEND_TIMEOUT") || utils.Map_contains(local, "SEND_TIMEOUT") {
		if utils.Map_contains(local, "SEND_TIMEOUT") && local["SEND_TIMEOUT"] != "" {
			m["SEND_TIMEOUT"] = local["SEND_TIMEOUT"]
		} else {
			m["SEND_TIMEOUT"] = global["SEND_TIMEOUT"]
		}
	}
	return m
}

// Check is config map correct
func ValidateConnectionInfo(configMap map[string]string) bool {
	if configMap["BROKERS"] == "" {
		return false
	}

	if configMap["TOPIC"] == "" {
		return false
	}

	return true
}

func ValidateConsumerInfo(configMap map[string]string) bool {
	isConnectionValidated := ValidateConnectionInfo(configMap)

	if utils.Map_contains(configMap, "CONSUMER_GROUP") {
		if configMap["CONSUMER_GROUP"] == "" {
			return false
		}
	}

	if utils.Map_contains(configMap, "NUM_CONSUMER") {
		if configMap["NUM_CONSUMER"] == "" {
			return false
		}
		num_consumer := utils.ItoInt(configMap["NUM_CONSUMER"])
		if num_consumer == math.MinInt32 {
			return false
		}
	}

	return isConnectionValidated
}
