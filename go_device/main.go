package main

import (
	"conormacpherson/iot/sisemographs/config"
	"conormacpherson/iot/sisemographs/messages"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
)

var (
	min                            float64       = -1.0
	max                            float64       = 1.0
	qos                            int           = 1
	cloudIOTCoreMQTTBridgeHostName string        = "mqtt.googleapis.com"
	cloudIOTCoreMQTTBridgePort     string        = "8883"
	configs                        config.Values = config.GetConfig()
)

func main() {
	ctx := context.Background()

	fmt.Printf("%+v", configs)

	client, err := pubsub.NewClient(ctx, configs.GCP.ProjectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	jobSize := configs.Simulation.JobSize

	for i := 0; i < jobSize; i++ {
		result := publish(ctx, client)
		id, err := result.Get(ctx)
		if err != nil {
			fmt.Fprintf(os.Stdout, "failed to publish: %v\n", err)
			return
		}
		fmt.Fprintf(os.Stdout, "Published message #%d. Message ID: %v\n", i, id)
		time.Sleep((time.Duration(configs.Simulation.Interval) * time.Second))
	}
}

func publish(ctx context.Context, client *pubsub.Client) *pubsub.PublishResult {
	rand.Seed(time.Now().UnixNano())
	variance := rand.Float64()*(max-min) + min

	reading := messages.SeismographReading{Name: configs.Device.Name, Location: configs.Device.Location, Scale: variance, Time: time.Now().Format("2006-01-02T15:04:05")}
	topic := client.Topic("seismograph_readings")
	json, err := json.Marshal(reading)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not marshal reading to json: %v\n", err)
		return nil
	}

	attributes := map[string]string{"device_name": configs.Device.Name}
	return topic.Publish(ctx, &pubsub.Message{
		Data:       json,
		Attributes: attributes,
	})
}
