package apmelk

import (
	"fmt"
	"strings"

	"github.com/Rivalz-ai/framework-be/framework/config/vault"
	"github.com/Rivalz-ai/framework-be/framework/utils"
	//"os"
)

func GetConfig(vault *vault.Vault, config_path string) map[string]string {
	config_path = strings.TrimSuffix(config_path, "/")
	m := utils.DictionaryString()
	m["ELASTIC_APM_SERVICE_NAME"] = vault.GetServiceName()
	m["ELASTIC_APM_SECRET_TOKEN"] = vault.ReadVAR(fmt.Sprintf("%s/ELASTIC_APM_SECRET_TOKEN", config_path))
	//m["ELASTIC_APM_ENVIRONMENT"] = vault.ReadVAR(fmt.Sprintf("%s/ELASTIC_APM_ENVIRONMENT", config_path))
	m["ELASTIC_APM_SERVER_URL"] = vault.ReadVAR(fmt.Sprintf("%s/ELASTIC_APM_SERVER_URL", config_path))
	m["ENABLE"] = vault.ReadVAR(fmt.Sprintf("%s/ENABLE", config_path))
	m["ENABLE_AGENT_DEBUG"] = vault.ReadVAR(fmt.Sprintf("%s/ENABLE_AGENT_DEBUG", config_path))
	m["ENABLE_APM"] = vault.ReadVAR(fmt.Sprintf("%s/ENABLE_APM", config_path))
	return m
}
