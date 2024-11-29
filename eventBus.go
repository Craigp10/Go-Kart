package main

import (
	"github.com/google/uuid"
)

type EventBus struct {
	registry map[uuid.UUID]*Consumer
}

// 1 : 1 event bus. Messages are posted with a subscriber in mind
func NewEventBus() *EventBus {
	reg := make(map[uuid.UUID]*Consumer)
	return &EventBus{
		registry: reg,
	}
}

func (eb *EventBus) RegisterConsumer(con *Consumer) {
	if eb.registry[con.ID] != nil {
		return
	}

	eb.registry[con.ID] = con
}

// If we want to do some sort of grouped routing topics like this would be useful
// func (eb *EventBus) _topicExist(topic string) bool {
// 	_, ok := eb.registry[topic]
// 	return ok
// }

// func (eb *EventBus) Publish(topic, message string) error {
// 	if !eb._topicExist(topic) {
// 		return fmt.Errorf("Topic has not been registerred")
// 	}

// }
