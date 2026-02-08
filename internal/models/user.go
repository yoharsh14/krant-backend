package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// User represents a user in the system
type User struct {
	ID            bson.ObjectID   `json:"id" bson:"_id,omitempty"`
	GoogleID      string               `json:"google_id" bson:"google_id"`
	Email         string               `json:"email" bson:"email"`
	Name          string               `json:"name" bson:"name"`
	ProfileImage  string               `json:"profile_image" bson:"profile_image"`
	TraderType    string               `json:"trader_type" bson:"trader_type"` // day_trader, swing_trader, long_term_investor, beginner
	Interests     []string             `json:"interests" bson:"interests"`
	BookmarkedNews []bson.ObjectID `json:"bookmarked_news" bson:"bookmarked_news"`
	ReadHistory   []ReadHistoryItem    `json:"read_history" bson:"read_history"`
	Preferences   UserPreferences      `json:"preferences" bson:"preferences"`
	CreatedAt     time.Time            `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at" bson:"updated_at"`
	LastLogin     time.Time            `json:"last_login" bson:"last_login"`
}

// ReadHistoryItem represents a news article read by the user
type ReadHistoryItem struct {
	NewsID bson.ObjectID `json:"news_id" bson:"news_id"`
	ReadAt time.Time          `json:"read_at" bson:"read_at"`
}

// UserPreferences represents user settings
type UserPreferences struct {
	NotificationEnabled bool   `json:"notification_enabled" bson:"notification_enabled"`
	Theme               string `json:"theme" bson:"theme"` // dark, light
	Language            string `json:"language" bson:"language"`
}


// UserResponse represents user data returned to client (without sensitive info)
type UserResponse struct {
	ID            string             `json:"id"`
	Email         string             `json:"email"`
	Name          string             `json:"name"`
	ProfileImage  string             `json:"profile_image"`
	TraderType    string             `json:"trader_type"`
	Interests     []string           `json:"interests"`
	Preferences   UserPreferences    `json:"preferences"`
	CreatedAt     time.Time          `json:"created_at"`
}

// ToResponse converts User to UserResponse
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:           u.ID.Hex(),
		Email:        u.Email,
		Name:         u.Name,
		ProfileImage: u.ProfileImage,
		TraderType:   u.TraderType,
		Interests:    u.Interests,
		Preferences:  u.Preferences,
		CreatedAt:    u.CreatedAt,
	}
}

// TraderTypes constants
const (
	TraderTypeDayTrader         = "day_trader"
	TraderTypeSwingTrader       = "swing_trader"
	TraderTypeLongTermInvestor  = "long_term_investor"
	TraderTypeBeginner          = "beginner"
)

// ValidTraderTypes returns all valid trader types
func ValidTraderTypes() []string {
	return []string{
		TraderTypeDayTrader,
		TraderTypeSwingTrader,
		TraderTypeLongTermInvestor,
		TraderTypeBeginner,
	}
}


// IsValidTraderType checks if the trader type is valid
func IsValidTraderType(traderType string) bool {
	for _, t := range ValidTraderTypes() {
		if t == traderType {
			return true
		}
	}
	return false
}