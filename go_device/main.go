package main

import (
	"conormacpherson/iot/sisemographs/config"
	"conormacpherson/iot/sisemographs/messages"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
)

var (
	minDelta                       int           = -1
	maxDelta                       int           = 1
	qos                            int           = 1
	cloudIOTCoreMQTTBridgeHostName string        = "mqtt.googleapis.com"
	cloudIOTCoreMQTTBridgePort     string        = "8883"
	configs                        config.Values = config.GetConfig()
)

func main() {
	ctx := context.Background()

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
	seed := rand.NewSource(time.Now().UnixNano())
	variance := minDelta + (rand.New(seed).Int() * maxDelta)
	reading := messages.SeismographReading{ID: configs.GCP.IOTCore.Device.Name, Location: configs.GCP.IOTCore.Device.Location, Scale: variance, Time: time.Now().Format("2006-01-02T15:04:05")}
	fmt.Println(reading)
	topic := client.Topic("seismograph_readings")
	return topic.Publish(ctx, &pubsub.Message{
		Data: []byte(fmt.Sprintf("%v", reading)),
	})
}
