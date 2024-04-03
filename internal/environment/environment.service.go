package environment

import (
	environment_buffers "environment/buffers"
	"google.golang.org/protobuf/encoding/protojson"
	"os"
)

// Get - returns singleton instance of Environment.
func Get() *environment_buffers.Environment {
	/**
	 * Load the `configuration.json` from the working directory.
	 */
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Read the configuration file.
	config, err := os.ReadFile(cwd + "/configuration.json")
	if err != nil {
		panic(err)
	}

	var environment environment_buffers.Environment = environment_buffers.Environment{}
	err = protojson.Unmarshal(config, &environment)
	if err != nil {
		println("Failed to unmarshal the configuration file.")
		println(err.Error())
		return nil
	}

	return &environment
}
