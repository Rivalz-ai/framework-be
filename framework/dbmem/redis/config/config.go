package config

import (
	"fmt"
	"strings"

	"github.com/Rivalz-ai/framework-be/framework/config/vault"
	"github.com/Rivalz-ai/framework-be/framework/utils"
)

func GetConfig(vault *vault.Vault, config_path string) map[string]string {
	m := utils.DictionaryString()
	if config_path == "" {
		return nil
	}
	config_path = strings.TrimSuffix(config_path, "/")
	m["HOST"] = vault.ReadVAR(fmt.Sprintf("%s/HOST", config_path))
	m["DB"] = vault.ReadVAR(fmt.Sprintf("%s/DB", config_path))
	m["PASSWORD"] = vault.ReadVAR(fmt.Sprintf("%s/PASSWORD", config_path))
	m["CLUSTER"] = vault.ReadVAR(fmt.Sprintf("%s/CLUSTER", config_path))
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
	if utils.Map_contains(global, "DB") || utils.Map_contains(local, "DB") {
		if utils.Map_contains(local, "DB") && local["DB"] != "" {
			m["DB"] = local["DB"]
		} else {
			m["DB"] = global["DB"]
		}
	}
	if utils.Map_contains(global, "PASSWORD") || utils.Map_contains(local, "PASSWORD") {
		if utils.Map_contains(local, "PASSWORD") && local["PASSWORD"] != "" {
			m["PASSWORD"] = local["PASSWORD"]
		} else {
			m["PASSWORD"] = global["PASSWORD"]
		}
	}
	if utils.Map_contains(global, "CLUSTER") || utils.Map_contains(local, "CLUSTER") {
		if utils.Map_contains(local, "CLUSTER") && local["CLUSTER"] != "" {
			m["CLUSTER"] = local["CLUSTER"]
		} else {
			m["CLUSTER"] = global["CLUSTER"]
		}
	}
	return m
}
