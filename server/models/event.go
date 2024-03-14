package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type EventType string

const (
	SPECIFIC_DATES EventType = "specific_dates"
	DOW            EventType = "dow"
)

// Representation of an Event in the mongoDB database
type Event struct {
	Id      primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	OwnerId primitive.ObjectID `json:"ownerId" bson:"ownerId,omitempty"`
	Name    string             `json:"name" bson:"name,omitempty"`

	Duration             *float32             `json:"duration" bson:"duration,omitempty"`
	Dates                []primitive.DateTime `json:"dates" bson:"dates,omitempty"`
	NotificationsEnabled bool                 `json:"notificationsEnabled" bson:"notificationsEnabled,omitempty"`

	Type EventType `json:"type" bson:"type,omitempty"`

	// Availability responses
	Responses map[string]*Response `json:"responses" bson:"responses"`

	// Scheduled event
	ScheduledEvent  *CalendarEvent `json:"scheduledEvent" bson:"scheduledEvent,omitempty"`
	CalendarEventId string         `json:"calendarEventId" bson:"calendarEventId,omitempty"`

	// Remindees
	Remindees []Remindee `json:"remindees" bson:"remindees,omitempty"`
}

// A response object containing an array of times that the given user is available
type Response struct {
	Name         string               `json:"name" bson:",omitempty"`
	UserId       primitive.ObjectID   `json:"userId" bson:",omitempty"`
	User         *User                `json:"user" bson:",omitempty"`
	Availability []primitive.DateTime `json:"availability" bson:"availability"`
}

// Object containing information associated with the remindee
type Remindee struct {
	Email     string   `json:"email" bson:"email,omitempty"`
	TaskIds   []string `json:"-" bson:"taskIds,omitempty"` // Task IDs of the scheduled emails
	Responded *bool    `json:"responded" bson:"responded,omitempty"`
}
