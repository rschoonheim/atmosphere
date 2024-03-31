package environment

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path"
)

type Environment struct {
	// Name - the name of the running environment. E.g. "development", "production", etc.
	Name string `json:"environment"`

	// Endpoints - map of endpoints that exposes services to the outside world.
	Endpoints map[string]Endpoint `json:"endpoints"`
}

var environment *Environment

func init() {

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Read the configuration.json
	configPath := path.Join(cwd, "configuration.json")

	// Read json from configPath
	jsonData, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Unmarshal the json data into a generic map
	var data map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Unmarshal the json data into a variable
	jsonData, err = json.Marshal(data)
	if err != nil {
		slog.Error("Failed to unmarshal environment configuration: %v", err)
		return
	}

	var env Environment = Environment{}

	// Set the environment name
	if val, ok := data["environment"]; ok {
		env.Name = val.(string)
	}

	// Set the endpoints
	if val, ok := data["endpoints"]; ok {
		endpoints := val.([]interface{})
		env.Endpoints = make(map[string]Endpoint, len(endpoints))

		for k, v := range endpoints {
			endpoint := v.(map[string]interface{})

			// Set the endpoint name
			endpointName := fmt.Sprintf("endpoint_%d", k)
			if val, ok := endpoint["name"]; ok {
				endpointName = val.(string)
			}

			// Set the endpoint port
			endpointPort := 0
			if val, ok := endpoint["port"]; ok {
				endpointPort = int(val.(float64))
			}

			// Set the endpoint TLS certificate
			endpointTlsCert := TlsCertificate{}
			if val, ok := endpoint["tls_cert"]; ok {
				tlsCert := val.(map[string]interface{})

				// Set the certificate path
				if val, ok := tlsCert["cert"]; ok {
					endpointTlsCert.Cert = val.(string)
				}

				// Set the key path
				if val, ok := tlsCert["key"]; ok {
					endpointTlsCert.Key = val.(string)
				}
			}

			env.Endpoints[endpointName] = Endpoint{
				Name:    endpointName,
				Port:    endpointPort,
				TlsCert: endpointTlsCert,
			}
		}
	}

	environment = &env
}

// Get - returns singleton instance of the Environment struct.
func Get() *Environment {
	return environment
}

// HasDuplicatePorts - returns true if there are duplicate ports in the environment.
func (env *Environment) HasDuplicatePorts() int {
	ports := make(map[int]bool)
	for _, endpoint := range env.Endpoints {
		if _, ok := ports[endpoint.Port]; ok {
			return endpoint.Port
		}
		ports[endpoint.Port] = true
	}
	return 0
}
