package models

import (
	"time"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Notification represents a user notification
type Notification struct {
	ID        bson.ObjectID  `json:"id" bson:"_id,omitempty"`
	UserID    bson.ObjectID  `json:"user_id" bson:"user_id"`
	Title     string         `json:"title" bson:"title"`
	Message   string         `json:"message" bson:"message"`
	Type      string         `json:"type" bson:"type"` // news_alert, system, promotion
	NewsID    *bson.ObjectID `json:"news_id,omitempty" bson:"news_id,omitempty"`
	IsRead    bool           `json:"is_read" bson:"is_read"`
	CreatedAt time.Time      `json:"created_at" bson:"created_at"`
}

// CreateNotificationInput represents input for creating a notification
type CreateNotificationInput struct {
	UserID  bson.ObjectID  `json:"user_id" binding:"required"`
	Title   string         `json:"title" binding:"required"`
	Message string         `json:"message" binding:"required"`
	Type    string         `json:"type" binding:"required"`
	NewsID  *bson.ObjectID `json:"news_id,omitempty"`
}

// NotificationResponse represents notification data returned to client
type NotificationResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Type      string    `json:"type"`
	NewsID    *string   `json:"news_id,omitempty"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

// ToResponse converts Notification to NotificationResponse
func (n *Notification) ToResponse() NotificationResponse {
	resp := NotificationResponse{
		ID:        n.ID.Hex(),
		Title:     n.Title,
		Message:   n.Message,
		Type:      n.Type,
		IsRead:    n.IsRead,
		CreatedAt: n.CreatedAt,
	}

	if n.NewsID != nil {
		newsID := n.NewsID.Hex()
		resp.NewsID = &newsID
	}

	return resp
}

// NotificationType constants
const (
	NotificationTypeNewsAlert = "news_alert"
	NotificationTypeSystem    = "system"
	NotificationTypePromotion = "promotion"
)

// ValidNotificationTypes returns all valid notification types
func ValidNotificationTypes() []string {
	return []string{
		NotificationTypeNewsAlert,
		NotificationTypeSystem,
		NotificationTypePromotion,
	}
}

// IsValidNotificationType checks if the notification type is valid
func IsValidNotificationType(notifType string) bool {
	for _, t := range ValidNotificationTypes() {
		if t == notifType {
			return true
		}
	}
	return false
}	