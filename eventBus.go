package main

import "fmt"

type Listeners struct {
	id string
	ip string
}

type EventBus struct {
	registry map[string][]Listeners
}

func (eb *EventBus) _topicExist(topic string) bool {
	_, ok := eb.registry[topic]
	return ok
}

func (eb *EventBus) Publish(topic, message string) error {
	if !eb._topicExist(topic) {
		return fmt.Errorf("Topic has not been registerred")
	}

	
}
