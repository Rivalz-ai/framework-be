package config

import (
	"fmt"
	"strings"

	"github.com/Rivalz-ai/framework-be/framework/config/vault"
	"github.com/Rivalz-ai/framework-be/framework/utils"
)

func GetConfig(vault *vault.Vault, config_path string) map[string]string {
	config_path = strings.TrimSuffix(config_path, "/")
	m := utils.DictionaryString()
	m["HOST"] = vault.ReadVAR(fmt.Sprintf("%s/HOST", config_path))
	m["PORT"] = vault.ReadVAR(fmt.Sprintf("%s/PORT", config_path))
	m["DB"] = vault.ReadVAR(fmt.Sprintf("%s/DB", config_path))
	m["USERNAME"] = vault.ReadVAR(fmt.Sprintf("%s/USERNAME", config_path))
	m["PASSWORD"] = vault.ReadVAR(fmt.Sprintf("%s/PASSWORD", config_path))
	m["USE_SSL"] = vault.ReadVAR(fmt.Sprintf("%s/USE_SSL", config_path))
	m["TIME_ZONE"] = vault.ReadVAR(fmt.Sprintf("%s/TIME_ZONE", config_path))
	m["TABLE_PREFIX"] = vault.ReadVAR(fmt.Sprintf("%s/TABLE_PREFIX", config_path))
	m["ENABLE_APM"] = vault.ReadVAR(fmt.Sprintf("%s/ENABLE_APM", config_path))
	m["MAX_CONN"] = vault.ReadVAR(fmt.Sprintf("%s/MAX_CONN", config_path))
	m["MIN_CONN"] = vault.ReadVAR(fmt.Sprintf("%s/MIN_CONN", config_path))
	m["MAX_IDLE_CONN"] = vault.ReadVAR(fmt.Sprintf("%s/MAX_IDLE_CONN", config_path))
	m["MAX_LIFETIME_CON"] = vault.ReadVAR(fmt.Sprintf("%s/MAX_LIFETIME_CON", config_path))
	m["CONN_TIMEOUT"] = vault.ReadVAR(fmt.Sprintf("%s/CONN_TIMEOUT", config_path))
	m["GORM_LOG_LEVEL"] = vault.ReadVAR(fmt.Sprintf("%s/GORM_LOG_LEVEL", config_path))
	m["NUM_NODE"] = vault.ReadVAR(fmt.Sprintf("%s/NUM_NODE", config_path))
	return m
}
func MergeConfig(global, local map[string]string) map[string]string {
	m := utils.DictionaryString()
	if utils.Map_contains(global, "HOST") || utils.Map_contains(local, "HOST") {
		if utils.Map_contains(local, "HOST") && local["HOST"] != "" {
			m["HOST"] = local["HOST"]
		} else {
			m["HOST"] = global["HOST"]
		}
	}
	if utils.Map_contains(global, "PORT") || utils.Map_contains(local, "PORT") {
		if utils.Map_contains(local, "PORT") && local["PORT"] != "" {
			m["PORT"] = local["PORT"]
		} else {
			m["PORT"] = global["PORT"]
		}
	}
	if utils.Map_contains(global, "DB") || utils.Map_contains(local, "DB") {
		if utils.Map_contains(local, "DB") && local["DB"] != "" {
			m["DB"] = local["DB"]
		} else {
			m["DB"] = global["DB"]
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
	if utils.Map_contains(global, "TIME_ZONE") || utils.Map_contains(local, "TIME_ZONE") {
		if utils.Map_contains(local, "TIME_ZONE") && local["TIME_ZONE"] != "" {
			m["TIME_ZONE"] = local["TIME_ZONE"]
		} else {
			m["TIME_ZONE"] = global["TIME_ZONE"]
		}
	}
	if utils.Map_contains(global, "TABLE_PREFIX") || utils.Map_contains(local, "TABLE_PREFIX") {
		if utils.Map_contains(local, "TABLE_PREFIX") && local["TABLE_PREFIX"] != "" {
			m["TABLE_PREFIX"] = local["TABLE_PREFIX"]
		} else {
			m["TABLE_PREFIX"] = global["TABLE_PREFIX"]
		}
	}
	if utils.Map_contains(global, "ENABLE_APM") || utils.Map_contains(local, "ENABLE_APM") {
		if utils.Map_contains(local, "ENABLE_APM") && local["ENABLE_APM"] != "" {
			m["ENABLE_APM"] = local["ENABLE_APM"]
		} else {
			m["ENABLE_APM"] = global["ENABLE_APM"]
		}
	}
	if utils.Map_contains(global, "MAX_CONN") || utils.Map_contains(local, "MAX_CONN") {
		if utils.Map_contains(local, "MAX_CONN") && local["MAX_CONN"] != "" {
			m["MAX_CONN"] = local["MAX_CONN"]
		} else {
			m["MAX_CONN"] = global["MAX_CONN"]
		}
	}
	if utils.Map_contains(global, "MIN_CONN") || utils.Map_contains(local, "MIN_CONN") {
		if utils.Map_contains(local, "MIN_CONN") && local["MIN_CONN"] != "" {
			m["MIN_CONN"] = local["MIN_CONN"]
		} else {
			m["MIN_CONN"] = global["MIN_CONN"]
		}
	}
	if utils.Map_contains(global, "MAX_IDLE_CONN") || utils.Map_contains(local, "MAX_IDLE_CONN") {
		if utils.Map_contains(local, "MAX_IDLE_CONN") && local["MAX_IDLE_CONN"] != "" {
			m["MAX_IDLE_CONN"] = local["MAX_IDLE_CONN"]
		} else {
			m["MAX_IDLE_CONN"] = global["MAX_IDLE_CONN"]
		}
	}
	if utils.Map_contains(global, "MAX_LIFETIME_CON") || utils.Map_contains(local, "MAX_LIFETIME_CON") {
		if utils.Map_contains(local, "MAX_LIFETIME_CON") && local["MAX_LIFETIME_CON"] != "" {
			m["MAX_LIFETIME_CON"] = local["MAX_LIFETIME_CON"]
		} else {
			m["MAX_LIFETIME_CON"] = global["MAX_LIFETIME_CON"]
		}
	}
	if utils.Map_contains(global, "CONN_TIMEOUT") || utils.Map_contains(local, "CONN_TIMEOUT") {
		if utils.Map_contains(local, "CONN_TIMEOUT") && local["CONN_TIMEOUT"] != "" {
			m["CONN_TIMEOUT"] = local["CONN_TIMEOUT"]
		} else {
			m["CONN_TIMEOUT"] = global["CONN_TIMEOUT"]
		}
	}
	if utils.Map_contains(global, "GORM_LOG_LEVEL") || utils.Map_contains(local, "GORM_LOG_LEVEL") {
		if utils.Map_contains(local, "GORM_LOG_LEVEL") && local["GORM_LOG_LEVEL"] != "" {
			m["GORM_LOG_LEVEL"] = local["GORM_LOG_LEVEL"]
		} else {
			m["GORM_LOG_LEVEL"] = global["GORM_LOG_LEVEL"]
		}
	}
	if utils.Map_contains(global, "NUM_NODE") || utils.Map_contains(local, "NUM_NODE") {
		if utils.Map_contains(local, "NUM_NODE") && local["NUM_NODE"] != "" {
			m["NUM_NODE"] = local["NUM_NODE"]
		} else {
			m["NUM_NODE"] = global["NUM_NODE"]
		}
	}
	return m
}
