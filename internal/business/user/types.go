package user

// CreateUserInput represents input for creating a new user
type CreateUserInput struct {
	GoogleID     string   `json:"google_id" binding:"required"`
	Email        string   `json:"email" binding:"required,email"`
	Name         string   `json:"name" binding:"required"`
	ProfileImage string   `json:"profile_image"`
	TraderType   string   `json:"trader_type" binding:"required"`
	Interests    []string `json:"interests" binding:"required,min=1"`
}
//Purpose: Validates data when creating a new user (during profile setup).
// Why separate from User struct?
// This struct is for input validation only
// It doesn't include auto-generated fields like ID, CreatedAt, etc.
// Uses binding tags for validation



type UpdateUserInput struct {
	Name         *string   `json:"name,omitempty"`
	ProfileImage *string   `json:"profile_image,omitempty"`
	TraderType   *string   `json:"trader_type,omitempty"`
	Interests    *[]string `json:"interests,omitempty"`
}
// if instead of pointer we just use string there will be aProblem:
// Can't find d/f b/w "user didn't send name"  and "user sent empty string"



// UpdatePreferencesInput represents input for updating user preferences
type UpdatePreferencesInput struct {
	NotificationEnabled *bool   `json:"notification_enabled,omitempty"`
	Theme               *string `json:"theme,omitempty"`
	Language            *string `json:"language,omitempty"`
}



