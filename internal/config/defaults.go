package config

import (
	"fmt"
	"os/user"
)

var (

	// server holds server configuration
	server = map[string]string{
		"name":    "<adaptor-name>-adaptor",
		"port":    "10007",
		"version": "v1.0.0",
	}

	// mesh holds mesh configuration
	mesh = map[string]string{
		"name":     "<adaptor-name>",
		"status":   "not installed",
		"traceurl": "http://localhost:14268/api/traces",
		"version":  "0.6.0",
	}

	// operations holds the supported operations inside mesh
	operations = map[string]interface{}{
		installMesh: map[string]interface{}{
			"type": "INSTALL",
			"properties": map[string]string{
				"description": "<adaptor-name> installation",
				"version":     "latest",
			},
		},
		installSample: map[string]interface{}{
			"type": "INSTALL",
			"properties": map[string]string{
				"description": "Sample application installation",
				"version":     "latest",
			},
		},
	}

	// Viper configuration
	filepath = fmt.Sprintf("%s/.<adaptor-name>", GetHome())
	filename = "config.yml"
	filetype = "yaml"
)

// GetHome returns the home path
func GetHome() string {
	usr, _ := user.Current()
	return usr.HomeDir
}
