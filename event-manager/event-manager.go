package eventmanager

import "github.com/ozgurkadarozgur/go-event/client"

type EventManager interface {
	GetClient() client.Client
}

type eventManager struct {
	eventManagerClient client.Client
}

func GetEventManager() EventManager {
	return &eventManager{
		eventManagerClient: client.NewRedisClient(),
	}
}

func (em *eventManager) GetClient() client.Client {
	return em.eventManagerClient
}
