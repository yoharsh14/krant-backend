package user

import (
	"context"
	"log"
	"time"
	"yoharsh14/krant-backend/internal/models"

)

type Service interface {
	CreateUser(ctx context.Context,input CreateUserInput) (error)
	FetchByUserNameAndEmail(ctx context.Context) (models.UserResponse,error)
	UpdateUser(ctx context.Context) (models.UserResponse,error)
	ListAllUser(ctx context.Context) ([]models.UserResponse,error) 
	// GetUserByID(ctx context.Context, id bson.ObjectID) (*models.User, error)
	// GetUserByGoogleID(ctx context.Context, googleID string) (*models.User, error)
	// GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	// GetOrCreateUser(ctx context.Context, googleID, email, name, profileImage string) (*models.User, bool, error)
	// UpdateProfile(ctx context.Context, userID bson.ObjectID, input UpdateUserInput) (*models.User, error)
	// UpdatePreferences(ctx context.Context, userID bson.ObjectID, input UpdatePreferencesInput) (*models.User, error)
	// DeleteUser(ctx context.Context, userID bson.ObjectID) error
	// BookmarkNews(ctx context.Context, userID, newsID bson.ObjectID) error 
	// UnbookmarkNews(ctx context.Context, userID, newsID bson.ObjectID) error
	// IsNewsBookmarked(ctx context.Context, userID, newsID bson.ObjectID) (bool, error)
	// GetUserBookmarks(ctx context.Context, userID bson.ObjectID) ([]bson.ObjectID, error) 
	// MarkAsRead(ctx context.Context, userID, newsID bson.ObjectID) error
	// GetReadHistory(ctx context.Context, userID bson.ObjectID, limit int) ([]models.ReadHistoryItem, error)
	// GetUserStats(ctx context.Context, userID bson.ObjectID) (map[string]interface{}, error) 
	// SearchUsers(ctx context.Context, query string, page, limit int) ([]models.User, int64, error)
	// GetUsersByTraderType(ctx context.Context, traderType string, page, limit int) ([]models.User, int64, error) 
	// GetUsersByInterest(ctx context.Context, interest string, page, limit int) ([]models.User, int64, error)
	// IsProfileComplete(ctx context.Context, userID bson.ObjectID) (bool, error)
	// CompleteProfile(ctx context.Context, userID bson.ObjectID, traderType string, interests []string) (*models.User, error)
}

type svc struct{
	r *Repository
}

func NewService(repo *Repository) Service{
	return &svc{
		r:repo,
	}
}

func (s *svc) CreateUser(ctx context.Context,input CreateUserInput) (error){
	user := &models.User{
		GoogleID: input.GoogleID,
		Email: input.Email,
		Name: input.Name,
		ProfileImage: input.TraderType,
		TraderType:input.TraderType,
		Interests: input.Interests,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		LastLogin:time.Now(),
	}
	err :=s.r.Create(context.TODO(),user)
	if err !=nil{
		return err
	}else{
		log.Println("User create with documentId")
	}
	return nil
}
func (s*svc) FetchByUserNameAndEmail(ctx context.Context) (models.UserResponse,error){
	
	return models.UserResponse{},nil
}
func (s*svc) UpdateUser(ctx context.Context) (models.UserResponse,error){
	return models.UserResponse{},nil
	
}
func (s*svc) ListAllUser(ctx context.Context) ([]models.UserResponse,error){
	return []models.UserResponse{},nil
} 

// // GetUserByID retrieves a user by their ID
// func (s *svc) GetUserByID(ctx context.Context, id bson.ObjectID) (*models.User, error) {
// 	user, err := s.r.FindByID(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if user == nil {
// 		return nil, errors.New("user not found")
// 	}
// 	return user, nil
// }

// // GetUserByGoogleID retrieves a user by their Google OAuth ID
// func (s *svc) GetUserByGoogleID(ctx context.Context, googleID string) (*models.User, error) {
// 	user, err := s.r.FindByGoogleID(ctx, googleID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if user == nil {
// 		return nil, errors.New("user not found")
// 	}
// 	return user, nil
// }

// // GetUserByEmail retrieves a user by their email
// func (s *svc) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
// 	user, err := s.r.FindByEmail(ctx, email)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if user == nil {
// 		return nil, errors.New("user not found")
// 	}
// 	return user, nil
// }

// // GetOrCreateUser gets existing user or creates new one (for OAuth login)
// func (s *svc) GetOrCreateUser(ctx context.Context, googleID, email, name, profileImage string) (*models.User, bool, error) {
// 	// Try to find existing user
// 	user, err := s.r.FindByGoogleID(ctx, googleID)
// 	if err != nil {
// 		return nil, false, err
// 	}

// 	// If user exists, update last login and return
// 	if user != nil {
// 		err = s.r.UpdateLastLogin(ctx, user.ID)
// 		if err != nil {
// 			return nil, false, err
// 		}
// 		return user, false, nil // false = not a new user
// 	}

// 	// User doesn't exist, create minimal profile
// 	// They'll need to complete profile setup (trader type, interests)
// 	user = &models.User{
// 		GoogleID:       googleID,
// 		Email:          email,
// 		Name:           name,
// 		ProfileImage:   profileImage,
// 		TraderType:     "", // Will be set during profile setup
// 		Interests:      []string{},
// 		BookmarkedNews: []bson.ObjectID{},
// 		ReadHistory:    []models.ReadHistoryItem{},
// 		Preferences: models.UserPreferences{
// 			NotificationEnabled: true,
// 			Theme:               "light",
// 			Language:            "en",
// 		},
// 	}

// 	err = s.r.Create(ctx, user)
// 	if err != nil {
// 		return nil, false, err
// 	}

// 	return user, true, nil // true = new user created
// }

// // UpdateProfile updates user profile information
// func (s *svc) UpdateProfile(ctx context.Context, userID bson.ObjectID, input UpdateUserInput) (*models.User, error) {
// 	// Validate trader type if provided
// 	if input.TraderType != nil && !models.IsValidTraderType(*input.TraderType) {
// 		return nil, errors.New("invalid trader type")
// 	}

// 	// Update in database
// 	err := s.r.UpdateProfile(ctx, userID, input)
// 	if err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return nil, errors.New("user not found")
// 		}
// 		return nil, err
// 	}

// 	// Return updated user
// 	return s.r.FindByID(ctx, userID)
// }

// // UpdatePreferences updates user preferences
// func (s *svc) UpdatePreferences(ctx context.Context, userID bson.ObjectID, input UpdatePreferencesInput) (*models.User, error) {
// 	// Validate theme if provided
// 	if input.Theme != nil && *input.Theme != "light" && *input.Theme != "dark" {
// 		return nil, errors.New("invalid theme, must be 'light' or 'dark'")
// 	}

// 	err := s.r.UpdatePreferences(ctx, userID, input)
// 	if err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return nil, errors.New("user not found")
// 		}
// 		return nil, err
// 	}

// 	return s.r.FindByID(ctx, userID)
// }

// // DeleteUser deletes a user account
// func (s *svc) DeleteUser(ctx context.Context, userID bson.ObjectID) error {
// 	err := s.r.Delete(ctx, userID)
// 	if err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return errors.New("user not found")
// 		}
// 		return err
// 	}
// 	return nil
// }

// // ============================================================================
// // BOOKMARK OPERATIONS
// // ============================================================================

// // BookmarkNews adds a news article to user's bookmarks
// func (s *svc) BookmarkNews(ctx context.Context, userID, newsID bson.ObjectID) error {
// 	// Check if user exists
// 	exists, err := s.r.Exists(ctx, userID)
// 	if err != nil {
// 		return err
// 	}
// 	if !exists {
// 		return errors.New("user not found")
// 	}

// 	// Add bookmark
// 	err = s.r.AddBookmark(ctx, userID, newsID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // UnbookmarkNews removes a news article from user's bookmarks
// func (s *svc) UnbookmarkNews(ctx context.Context, userID, newsID bson.ObjectID) error {
// 	err := s.r.RemoveBookmark(ctx, userID, newsID)
// 	if err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return errors.New("user not found")
// 		}
// 		return err
// 	}
// 	return nil
// }

// // IsNewsBookmarked checks if a user has bookmarked a news article
// func (s *svc) IsNewsBookmarked(ctx context.Context, userID, newsID bson.ObjectID) (bool, error) {
// 	return s.r.HasBookmarked(ctx, userID, newsID)
// }

// // GetUserBookmarks retrieves all bookmarked news IDs for a user
// func (s *svc) GetUserBookmarks(ctx context.Context, userID bson.ObjectID) ([]bson.ObjectID, error) {
// 	return s.r.GetBookmarkedNews(ctx, userID)
// }

// // ============================================================================
// // READ HISTORY OPERATIONS
// // ============================================================================

// // MarkAsRead adds a news article to user's read history
// func (s *svc) MarkAsRead(ctx context.Context, userID, newsID bson.ObjectID) error {
// 	return s.r.AddToReadHistory(ctx, userID, newsID)
// }

// // GetReadHistory retrieves user's read history
// func (s *svc) GetReadHistory(ctx context.Context, userID bson.ObjectID, limit int) ([]models.ReadHistoryItem, error) {
// 	if limit <= 0 {
// 		limit = 50 // Default limit
// 	}
// 	if limit > 100 {
// 		limit = 100 // Max limit
// 	}

// 	return s.r.GetReadHistory(ctx, userID, limit)
// }

// // ============================================================================
// // USER STATISTICS
// // ============================================================================

// // GetUserStats retrieves user statistics
// func (s *svc) GetUserStats(ctx context.Context, userID bson.ObjectID) (map[string]interface{}, error) {
// 	stats, err := s.r.GetUserStats(ctx, userID)
// 	if err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return nil, errors.New("user not found")
// 		}
// 		return nil, err
// 	}
// 	return stats, nil
// }

// // ============================================================================
// // SEARCH & LISTING
// // ============================================================================

// // SearchUsers searches for users by name or email
// func (s *svc) SearchUsers(ctx context.Context, query string, page, limit int) ([]models.User, int64, error) {
// 	pagination := models.GetPaginationParams(page, limit)
// 	return s.r.Search(ctx, query, pagination)
// }

// // GetUsersByTraderType retrieves users by trader type
// func (s *svc) GetUsersByTraderType(ctx context.Context, traderType string, page, limit int) ([]models.User, int64, error) {
// 	if !models.IsValidTraderType(traderType) {
// 		return nil, 0, errors.New("invalid trader type")
// 	}

// 	pagination := models.GetPaginationParams(page, limit)
// 	return s.r.FindByTraderType(ctx, traderType, pagination)
// }

// // GetUsersByInterest retrieves users interested in a specific category
// func (s *svc) GetUsersByInterest(ctx context.Context, interest string, page, limit int) ([]models.User, int64, error) {
// 	pagination := models.GetPaginationParams(page, limit)
// 	return s.r.FindByInterest(ctx, interest, pagination)
// }

// // ============================================================================
// // PROFILE COMPLETION CHECK
// // ============================================================================

// // IsProfileComplete checks if user has completed their profile
// func (s *svc) IsProfileComplete(ctx context.Context, userID bson.ObjectID) (bool, error) {
// 	user, err := s.r.FindByID(ctx, userID)
// 	if err != nil {
// 		return false, err
// 	}
// 	if user == nil {
// 		return false, errors.New("user not found")
// 	}

// 	// Profile is complete if they have trader type and at least one interest
// 	return user.TraderType != "" && len(user.Interests) > 0, nil
// }

// // CompleteProfile completes user profile setup (after OAuth login)
// func (s *svc) CompleteProfile(ctx context.Context, userID bson.ObjectID, traderType string, interests []string) (*models.User, error) {
// 	// Validate inputs
// 	if !models.IsValidTraderType(traderType) {
// 		return nil, errors.New("invalid trader type")
// 	}
// 	if len(interests) == 0 {
// 		return nil, errors.New("at least one interest is required")
// 	}

// 	// Update user
// 	update := bson.M{
// 		"trader_type": traderType,
// 		"interests":   interests,
// 	}

// 	err := s.r.Update(ctx, userID, update)
// 	if err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return nil, errors.New("user not found")
// 		}
// 		return nil, err
// 	}

// 	return s.r.FindByID(ctx, userID)
// }