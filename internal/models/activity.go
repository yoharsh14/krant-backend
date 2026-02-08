package models

import (
	"time"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// UserActivity represents user interactions with news
type UserActivity struct {
	ID           bson.ObjectID    `json:"id" bson:"_id,omitempty"`
	UserID       bson.ObjectID    `json:"user_id" bson:"user_id"`
	NewsID       bson.ObjectID    `json:"news_id" bson:"news_id"`
	ActivityType string           `json:"activity_type" bson:"activity_type"` // view, bookmark, share, unbookmark
	Metadata     ActivityMetadata `json:"metadata" bson:"metadata"`
	CreatedAt    time.Time        `json:"created_at" bson:"created_at"`
}

// ActivityMetadata represents additional activity information
type ActivityMetadata struct {
	Device    string `json:"device" bson:"device"`         // mobile, tablet, desktop
	Platform  string `json:"platform" bson:"platform"`     // ios, android, web
	TimeSpent int    `json:"time_spent" bson:"time_spent"` // seconds
}

// CreateActivityInput represents input for creating user activity
type CreateActivityInput struct {
	NewsID       bson.ObjectID    `json:"news_id" binding:"required"`
	ActivityType string           `json:"activity_type" binding:"required"`
	Metadata     ActivityMetadata `json:"metadata"`
}

// ActivityType constants
const (
	ActivityTypeView       = "view"
	ActivityTypeBookmark   = "bookmark"
	ActivityTypeUnbookmark = "unbookmark"
	ActivityTypeShare      = "share"
)

// ValidActivityTypes returns all valid activity types
func ValidActivityTypes() []string {
	return []string{
		ActivityTypeView,
		ActivityTypeBookmark,
		ActivityTypeUnbookmark,
		ActivityTypeShare,
	}
}

// IsValidActivityType checks if the activity type is valid
func IsValidActivityType(activityType string) bool {
	for _, t := range ValidActivityTypes() {
		if t == activityType {
			return true
		}
	}
	return false
}