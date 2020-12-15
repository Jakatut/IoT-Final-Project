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
}

// GCP contains configuration values related to the gcp project
type GCP struct {
	IOTCore   IOTCore `yaml:"IOT_CORE"`
	PubSub    PubSub  `yaml:"PUB_SUB"`
	ProjectID string  `yaml:"PROJECT_ID"`
}

// IOTCore contains configuration values related to gcp's iot core.
type IOTCore struct {
	Region     string `yaml:"GCP_CLOUD_IOT_CORE_REGION"`
	RegistryID string `yaml:"REGISTRY_ID"`
	Device     Device `yaml:"DEVICE"`
}

// Device contains configuration values for a device gcp's iot core
type Device struct {
	Name     string `yaml:"NAME"`
	Path     string `yaml:"PATH"`
	Location string `yaml:"LOCATION"`
}

// PubSub contains configuration values related to pub sub on gcp.
type PubSub struct {
	TopicID   string `yaml:"TOPIC_ID"`
	TopicName string `yaml:"TOPIC_NAME"`
}

// Simulation contians configuration values related to running the simulation. i.e. how many "jobs"/"messages"
// are sent, at which interval, etc.
type Simulation struct {
	Interval int `yaml:"INTERVAL"`
	JobSize  int `yaml:"JOB_SIZE"`
}

// GetConfig reads and returns environment variables containing config values.
func GetConfig() Values {
	config := Values{}

	configPath := os.Getenv("CONFIG_PATH")
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
