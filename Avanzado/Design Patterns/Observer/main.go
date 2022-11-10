package main

import "fmt"

type Topic struct {
	observers []Observer
	name      string
	available bool
}

func NewTopic(name string) *Topic {
	return &Topic{
		name: name,
	}
}

func (i *Topic) UpdateAvailable() {
	fmt.Printf("Item %s is available\n", i.name)
	i.available = true
	i.broadcast()
}

func (i *Topic) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

func (i *Topic) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

type Observer interface {
	getId() string
	updateValue(value string)
}

type EmailClient struct {
	id string
}

func (eC *EmailClient) getId() string {
	return eC.id
}

func (eC *EmailClient) updateValue(value string) {
	fmt.Printf("Sending email - %s available from client %s\n", value, eC.id)
}

type SMSClient struct {
	id string
}

func (sC *SMSClient) getId() string {
	return sC.id
}

func (sC *SMSClient) updateValue(value string) {
	fmt.Printf("Sending SMS - %s available from client %s\n", value, sC.id)
}

func main() {
	nvidiaTopic := NewTopic("RTX 3080")
	firstObserver := &EmailClient{
		id: "12ab",
	}
	secondObserver := &SMSClient{
		id: "34dc",
	}
	nvidiaTopic.register(firstObserver)
	nvidiaTopic.register(secondObserver)
	nvidiaTopic.UpdateAvailable()
}
