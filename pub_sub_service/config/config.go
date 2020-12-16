package config

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Values contains configuration variables read from environment
// varialbes with the same name.
type Values struct {
	GCP        GCP        `yaml:"GCP"`
	Simulation Simulation `yaml:"SIMULATION"`
	Device     Device     `yaml:"DEVICE"`
	Server     Server     `yaml:"SERVER"`
}

// GCP contains configuration values related to the gcp project
type GCP struct {
	PubSub    PubSub `yaml:"PUB_SUB"`
	ProjectID string `yaml:"PROJECT_ID"`
}

// Device contains configuration values for a device gcp's iot core
type Device struct {
	Name     string `yaml:"NAME"`
	Path     string `yaml:"PATH"`
	Location string `yaml:"LOCATION"`
}

// PubSub contains configuration values related to pub sub on gcp.
type PubSub struct {
	TopicID      string `yaml:"TOPIC_ID"`
	TopicName    string `yaml:"TOPIC_NAME"`
	Subscription string `yaml:"SUBSCRIPTION"`
}

// Simulation contians configuration values related to running the simulation. i.e. how many "jobs"/"messages"
// are sent, at which interval, etc.
type Simulation struct {
	Interval int `yaml:"INTERVAL"`
	JobSize  int `yaml:"JOB_SIZE"`
}

// Server contains configuration values related to running the web server.
type Server struct {
	Port int `yaml:"PORT"`
}

// GetConfig reads and returns environment variables containing config values.
func GetConfig() Values {
	config := Values{}

	configPath := os.Getenv("PUB_SUB_SERVICE_LOGGING_PATH")
	if configPath == "" {
		var err error
		configPath, err = filepath.Abs("config/config.yaml")
		if err != nil {
			panic(errors.New("could not find config file in the provided location or in the default"))
		}
	}

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}
