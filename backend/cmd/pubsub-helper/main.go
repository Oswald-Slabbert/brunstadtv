package main

// This is just a simple helper, to set up a topic & push subscription
// on the pubsub emulator, and then send some message.
// Cleanup is done at the end so the program can be re-run w/o errors
import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/bcc-code/brunstadtv/backend/events"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/davecgh/go-spew/spew"
)

func create(projectID, topicID string) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Printf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	topic, err := client.CreateTopic(ctx, topicID)
	if err != nil {
		fmt.Printf("pubsub.NewClient: %v", err)
	}

	t, err := client.CreateSubscription(ctx, "bgjobs", pubsub.SubscriptionConfig{
		Topic: topic,
		PushConfig: pubsub.PushConfig{
			Endpoint: "http://host.docker.internal:8078/api/message",
		},
	})
	if err != nil {
		fmt.Printf("CreateTopic: %v", err)
	}
	fmt.Printf("Sub Created: %v\n", t)
}

func send(projectID, topicID string) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Printf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	e := cloudevents.NewEvent()
	e.SetSource(events.SourceMediaBanken)
	e.SetType(events.TypeAssetDelivered)
	e.SetData(cloudevents.ApplicationJSON, &events.AssetDelivered{
		//JSONMetaPath: "randomstring/sample.json",
		JSONMetaPath: "7233_TEMA2_Simen.json",
	})

	data, err := json.Marshal(e)
	spew.Dump(string(data))
	topic := client.Topic(topicID)
	msg := topic.Publish(ctx, &pubsub.Message{
		Data: data,
	})

	_, err = msg.Get(ctx)
	fmt.Printf("Sent: %v\n", err)
}

func refreshView(projectID, topicID string) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Printf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	e := cloudevents.NewEvent()
	e.SetSource(events.SourceCloudScheduler)
	e.SetType(events.TypeRefreshView)
	e.SetData(cloudevents.ApplicationJSON, &events.RefreshView{
		ViewName: "episodes_access",
		Force:    false,
	})

	data, err := json.Marshal(e)
	spew.Dump(string(data))
	topic := client.Topic(topicID)
	msg := topic.Publish(ctx, &pubsub.Message{
		Data: data,
	})

	_, err = msg.Get(ctx)
	fmt.Printf("Sent: %v\n", err)
}

func del(projectID, topicID string) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Printf("pubsub.NewClient: %v", err)
	}
	defer client.Close()
	sub := client.Subscription("bgjobs")
	err = sub.Delete(ctx)
	if err != nil {
		fmt.Printf("Error deleting sub: %v", err)
	}

	topic := client.Topic(topicID)
	err = topic.Delete(ctx)
	if err != nil {
		fmt.Printf("Error deleting sub: %v", err)
	}

	fmt.Printf("Deleted")
}

func translationSync(projectID string, topicID string) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Printf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	e := cloudevents.NewEvent()
	e.SetSource(events.SourceMediaBanken)
	e.SetType(events.TypeAssetDelivered)
	e.SetData(cloudevents.ApplicationJSON, &events.AssetDelivered{
		//JSONMetaPath: "randomstring/sample.json",
		JSONMetaPath: "7233_TEMA2_Simen.json",
	})

	data, err := json.Marshal(e)
	spew.Dump(string(data))
	topic := client.Topic(topicID)
	msg := topic.Publish(ctx, &pubsub.Message{
		Data: data,
	})

	_, err = msg.Get(ctx)
	fmt.Printf("Sent: %v\n", err)
}

func main() {
	task := flag.String("task", "", "")
	flag.Parse()
	switch *task {
	case "create":
		create("btv-local", "background-jobs")
	case "delete":
		del("btv-local", "background-jobs")
	case "refreshView":
		refreshView("btv-local", "background-jobs")
	default:
		create("btv-local", "background-jobs")
		/*
			send("btv-local", "background-jobs")
		*/

		refreshView("btv-local", "background-jobs")

		time.Sleep(1 * time.Second)
		del("btv-local", "background-jobs")
	}

}
