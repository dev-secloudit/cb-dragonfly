package eventhandler

import (
	"fmt"

	"github.com/cloud-barista/cb-dragonfly/pkg/core/alert/eventhandler/event/slack"
	"github.com/cloud-barista/cb-dragonfly/pkg/core/alert/eventhandler/event/smtp"
	"github.com/cloud-barista/cb-dragonfly/pkg/core/alert/types"
)

const (
	SlackType = "slack"
	SMTPType  = "smtp"
)

var (
	EventTypes = make(map[string]EventHandler)
)

type EventHandler interface {
	ListEventHandlers() ([]types.AlertEventHandler, error)
	GetEventHandler(name string) (types.AlertEventHandler, error)
	CreateEventHandler(createOpts types.AlertEventHandlerReq) (types.AlertEventHandler, error)
	UpdateEventHandler(createOpts types.AlertEventHandlerReq) (types.AlertEventHandler, error)
	DeleteEventHandler(name string) error
}

func InitializeEventTypes() {
	EventTypes[SlackType] = slack.SlackHandler{}
	EventTypes[SMTPType] = smtp.SmtpHandler{}
}

func ListEventHandlers(eventType string) ([]types.AlertEventHandler, error) {
	// get specific event type handlers
	if eventType != "" {
		if _, ok := EventTypes[eventType]; !ok {
			return nil, fmt.Errorf("not found eventType with Name %s", eventType)
		}
		return EventTypes[eventType].ListEventHandlers()
	}
	// get all event type handlers
	var eventHandlerList []types.AlertEventHandler
	for _, handlers := range EventTypes {
		eventHandlers, err := handlers.ListEventHandlers()
		if err != nil {
			return nil, err
		}
		eventHandlerList = append(eventHandlerList, eventHandlers...)
	}
	return eventHandlerList, nil
}

func GetEventHandler(eventType string, eventHandlerName string) (types.AlertEventHandler, error) {
	if _, ok := EventTypes[eventType]; !ok {
		return types.AlertEventHandler{}, fmt.Errorf("not found eventType with Name %s", eventType)
	}
	return EventTypes[eventType].GetEventHandler(eventHandlerName)
}

func CreateEventHandler(eventType string, eventHandlerReq types.AlertEventHandlerReq) (types.AlertEventHandler, error) {
	if _, ok := EventTypes[eventType]; !ok {
		return types.AlertEventHandler{}, fmt.Errorf("not found eventType with Name %s", eventType)
	}
	return EventTypes[eventType].CreateEventHandler(eventHandlerReq)
}

func UpdateEventHandler(eventType string, eventHandlerReq types.AlertEventHandlerReq) (types.AlertEventHandler, error) {
	if _, ok := EventTypes[eventType]; !ok {
		return types.AlertEventHandler{}, fmt.Errorf("not found eventType with Name %s", eventType)
	}
	return EventTypes[eventType].UpdateEventHandler(eventHandlerReq)
}

func DeleteEventHandler(eventType string, eventHandlerName string) error {
	if _, ok := EventTypes[eventType]; !ok {
		return fmt.Errorf("not found eventType with Name %s", eventType)
	}
	return EventTypes[eventType].DeleteEventHandler(eventHandlerName)
}
