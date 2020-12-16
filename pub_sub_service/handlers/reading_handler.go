package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"pub_sub_service/config"

	"cloud.google.com/go/pubsub"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// ReadingHandler handles websocket connections by subscribing to a pub/sub topic on gcp, and relaying the messages to the front-end client.
func ReadingHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to set websocket upgrade: %+v", err)
		return
	}
	defer ws.Close()

	configValues := config.GetConfig()

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, configValues.GCP.ProjectID)
	subscription := client.Subscription(configValues.GCP.PubSub.Subscription)
	err = subscription.Receive(ctx, func(ctx context.Context, message *pubsub.Message) {
		ws.WriteMessage(websocket.TextMessage, message.Data)
	})

	fmt.Fprintf(os.Stdout, "%+v", err)
}
