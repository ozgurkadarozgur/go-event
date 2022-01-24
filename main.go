package main

import (
	"encoding/json"
	"log"
	"time"

	eventmanager "github.com/ozgurkadarozgur/go-event/event-manager"
)

const (
	eventChannel = "test-channel"
)

func main() {

	em := eventmanager.GetEventManager().GetClient()
	em.Subscribe(eventChannel, func(payload string) {
		log.Println("here is payload: ", payload)
	})

	eventData, _ := json.Marshal("selam")
	em.Publish(eventChannel, eventData)

	time.Sleep(time.Second * 2)
}
