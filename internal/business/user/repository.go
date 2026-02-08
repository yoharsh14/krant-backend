package user

import (
	"context"
	"errors"
	"time"
	"yoharsh14/krant-backend/internal/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// type Repository interface {
// 	 Create(ctx context.Context, user *models.User) (error)
// }


//// CREATE
// Create(ctx context.Context, user *models.User) error

// // READ
// FindByID(ctx context.Context, id bson.ObjectID) (*models.User, error)
// FindByGoogleID(ctx context.Context, googleID string) (*models.User, error)
// FindByEmail(ctx context.Context, email string) (*models.User, error)
// FindAll(ctx context.Context, pagination models.PaginationParams) ([]models.User, int64, error)
// FindByTraderType(ctx context.Context, traderType string, pagination models.PaginationParams) ([]models.User, int64, error)
// FindByInterest(ctx context.Context, interest string, pagination models.PaginationParams) ([]models.User, int64, error)

// // UPDATE
// Update(ctx context.Context, id bson.ObjectID, update bson.M) error
// UpdateProfile(ctx context.Context, id bson.ObjectID, input UpdateUserInput) error
// UpdatePreferences(ctx context.Context, id bson.ObjectID, input UpdatePreferencesInput) error
// UpdateLastLogin(ctx context.Context, id bson.ObjectID) error

// // BOOKMARKS
// AddBookmark(ctx context.Context, userID, newsID bson.ObjectID) error
// RemoveBookmark(ctx context.Context, userID, newsID bson.ObjectID) error
// HasBookmarked(ctx context.Context, userID, newsID bson.ObjectID) (bool, error)
// GetBookmarkedNews(ctx context.Context, userID bson.ObjectID) ([]bson.ObjectID, error)

// // READ HISTORY
// AddToReadHistory(ctx context.Context, userID, newsID bson.ObjectID) error
// GetReadHistory(ctx context.Context, userID bson.ObjectID, limit int) ([]models.ReadHistoryItem, error)

// // DELETE
// Delete(ctx context.Context, id bson.ObjectID) error

// // AGGREGATION / STATS
// GetUserStats(ctx context.Context, userID bson.ObjectID) (map[string]interface{}, error)
// CountByTraderType(ctx context.Context) (map[string]int64, error)

// // SEARCH
// Search(ctx context.Context, query string, pagination models.PaginationParams) ([]models.User, int64, error)

// // EXISTENCE CHECKS
// Exists(ctx context.Context, id bson.ObjectID) (bool, error)
// EmailExists(ctx context.Context, email string) (bool, error)
// GoogleIDExists(ctx context.Context, googleID string) (bool, error)

type Repository struct {
	coll *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		coll: db.Collection("users"),
	}
}


// ============================================================================
// CREATE OPERATIONS
// ============================================================================

// Create inserts a new user into the database
func (r *Repository) Create(ctx context.Context, user *models.User) error {
	// Generate new ObjectID if not set
	if user.ID.IsZero() {
		user.ID = bson.NewObjectID()
	}

	// Set timestamps
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	user.LastLogin = now

	// Initialize empty slices if nil
	if user.Interests == nil {
		user.Interests = []string{}
	}
	if user.BookmarkedNews == nil {
		user.BookmarkedNews = []bson.ObjectID{}
	}
	if user.ReadHistory == nil {
		user.ReadHistory = []models.ReadHistoryItem{}
	}

	// Set default preferences
	if user.Preferences.Theme == "" {
		user.Preferences.Theme = "light"
	}
	if user.Preferences.Language == "" {
		user.Preferences.Language = "en"
	}

	_, err := r.coll.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// ============================================================================
// READ OPERATIONS
// ============================================================================

// FindByID retrieves a user by their ObjectID
func (r *Repository) FindByID(ctx context.Context, id bson.ObjectID) (*models.User, error) {
	var user models.User
	err := r.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // User not found
		}
		return nil, err
	}
	return &user, nil
}

// FindByGoogleID retrieves a user by their Google OAuth ID
func (r *Repository) FindByGoogleID(ctx context.Context, googleID string) (*models.User, error) {
	var user models.User
	err := r.coll.FindOne(ctx, bson.M{"google_id": googleID}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // User not found
		}
		return nil, err
	}
	return &user, nil
}

// FindByEmail retrieves a user by their email address
func (r *Repository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.coll.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// FindAll retrieves all users with pagination
func (r *Repository) FindAll(ctx context.Context, pagination models.PaginationParams) ([]models.User, int64, error) {
	// Count total documents
	totalCount, err := r.coll.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	// Find with pagination
	findOptions := options.Find().
		SetSkip(int64(pagination.GetSkip())).
		SetLimit(int64(pagination.Limit)).
		SetSort(bson.D{{Key: "created_at", Value: -1}}) // Newest first

	cursor, err := r.coll.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}

// FindByTraderType retrieves users by trader type with pagination
func (r *Repository) FindByTraderType(ctx context.Context, traderType string, pagination models.PaginationParams) ([]models.User, int64, error) {
	filter := bson.M{"trader_type": traderType}

	totalCount, err := r.coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	findOptions := options.Find().
		SetSkip(int64(pagination.GetSkip())).
		SetLimit(int64(pagination.Limit)).
		SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := r.coll.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}

// FindByInterest retrieves users interested in a specific category
func (r *Repository) FindByInterest(ctx context.Context, interest string, pagination models.PaginationParams) ([]models.User, int64, error) {
	// $in operator checks if the array contains the value
	filter := bson.M{"interests": interest}

	totalCount, err := r.coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	findOptions := options.Find().
		SetSkip(int64(pagination.GetSkip())).
		SetLimit(int64(pagination.Limit)).
		SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := r.coll.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}

// ============================================================================
// UPDATE OPERATIONS
// ============================================================================

// Update updates a user's information
func (r *Repository) Update(ctx context.Context, id bson.ObjectID, update bson.M) error {
	filter := bson.M{"_id": id}

	// Always update the updated_at timestamp
	update["updated_at"] = time.Now()

	updateDoc := bson.M{"$set": update}

	result, err := r.coll.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

// UpdateProfile updates user's profile information
func (r *Repository) UpdateProfile(ctx context.Context, id bson.ObjectID, input UpdateUserInput) error {
	update := bson.M{}

	if input.Name != nil {
		update["name"] = *input.Name
	}
	if input.ProfileImage != nil {
		update["profile_image"] = *input.ProfileImage
	}
	if input.TraderType != nil {
		update["trader_type"] = *input.TraderType
	}
	if input.Interests != nil {
		update["interests"] = *input.Interests
	}

	if len(update) == 0 {
		return errors.New("no fields to update")
	}

	return r.Update(ctx, id, update)
}

// UpdatePreferences updates user's preferences
func (r *Repository) UpdatePreferences(ctx context.Context, id bson.ObjectID, input UpdatePreferencesInput) error {
	update := bson.M{}

	if input.NotificationEnabled != nil {
		update["preferences.notification_enabled"] = *input.NotificationEnabled
	}
	if input.Theme != nil {
		update["preferences.theme"] = *input.Theme
	}
	if input.Language != nil {
		update["preferences.language"] = *input.Language
	}

	if len(update) == 0 {
		return errors.New("no fields to update")
	}

	return r.Update(ctx, id, update)
}

// UpdateLastLogin updates the user's last login timestamp
// func (r *Repository) UpdateLastLogin(ctx context.Context, id bson.ObjectID) error {
// 	filter := bson.M{"_id": id}
// 	update := bson.M{
// 		"$set": bson.M{
// 			"last_login": time.Now(),
// 		},
// 	}

// 	_, err := r.coll.UpdateOne(ctx, filter, update)
// 	return err
// }

// // ============================================================================
// // BOOKMARK OPERATIONS
// // ============================================================================

// // AddBookmark adds a news article to user's bookmarks
// func (r *Repository) AddBookmark(ctx context.Context, userID, newsID bson.ObjectID) error {
// 	filter := bson.M{"_id": userID}
// 	update := bson.M{
// 		"$addToSet": bson.M{"bookmarked_news": newsID}, // $addToSet prevents duplicates
// 		"$set":      bson.M{"updated_at": time.Now()},
// 	}

// 	result, err := r.coll.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		return err
// 	}

// 	if result.MatchedCount == 0 {
// 		return mongo.ErrNoDocuments
// 	}

// 	return nil
// }

// // RemoveBookmark removes a news article from user's bookmarks
// func (r *Repository) RemoveBookmark(ctx context.Context, userID, newsID bson.ObjectID) error {
// 	filter := bson.M{"_id": userID}
// 	update := bson.M{
// 		"$pull": bson.M{"bookmarked_news": newsID}, // $pull removes the element
// 		"$set":  bson.M{"updated_at": time.Now()},
// 	}

// 	result, err := r.coll.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		return err
// 	}

// 	if result.MatchedCount == 0 {
// 		return mongo.ErrNoDocuments
// 	}

// 	return nil
// }

// // HasBookmarked checks if a user has bookmarked a specific news article
// func (r *Repository) HasBookmarked(ctx context.Context, userID, newsID bson.ObjectID) (bool, error) {
// 	filter := bson.M{
// 		"_id":             userID,
// 		"bookmarked_news": newsID,
// 	}

// 	count, err := r.coll.CountDocuments(ctx, filter)
// 	if err != nil {
// 		return false, err
// 	}

// 	return count > 0, nil
// }

// // GetBookmarkedNews retrieves all bookmarked news IDs for a user
// func (r *Repository) GetBookmarkedNews(ctx context.Context, userID bson.ObjectID) ([]bson.ObjectID, error) {
// 	// Use projection to only return the bookmarked_news field
// 	projection := bson.M{"bookmarked_news": 1}

// 	var result struct {
// 		BookmarkedNews []bson.ObjectID `bson:"bookmarked_news"`
// 	}

// 	err := r.coll.FindOne(
// 		ctx,
// 		bson.M{"_id": userID},
// 		options.FindOne().SetProjection(projection),
// 	).Decode(&result)

// 	if err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return []bson.ObjectID{}, nil
// 		}
// 		return nil, err
// 	}

// 	return result.BookmarkedNews, nil
// }

// // ============================================================================
// // READ HISTORY OPERATIONS
// // ============================================================================

// // AddToReadHistory adds a news article to user's read history
// func (r *Repository) AddToReadHistory(ctx context.Context, userID, newsID bson.ObjectID) error {
// 	filter := bson.M{"_id": userID}

// 	readItem := models.ReadHistoryItem{
// 		NewsID: newsID,
// 		ReadAt: time.Now(),
// 	}

// 	update := bson.M{
// 		"$push": bson.M{
// 			"read_history": bson.M{
// 				"$each":  []models.ReadHistoryItem{readItem},
// 				"$slice": -100, // Keep only last 100 items
// 			},
// 		},
// 		"$set": bson.M{"updated_at": time.Now()},
// 	}

// 	result, err := r.coll.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		return err
// 	}

// 	if result.MatchedCount == 0 {
// 		return mongo.ErrNoDocuments
// 	}

// 	return nil
// }

// // GetReadHistory retrieves user's read history
// func (r *Repository) GetReadHistory(ctx context.Context, userID bson.ObjectID, limit int) ([]models.ReadHistoryItem, error) {
// 	projection := bson.M{"read_history": 1}

// 	var result struct {
// 		ReadHistory []models.ReadHistoryItem `bson:"read_history"`
// 	}

// 	err := r.coll.FindOne(
// 		ctx,
// 		bson.M{"_id": userID},
// 		options.FindOne().SetProjection(projection),
// 	).Decode(&result)

// 	if err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return []models.ReadHistoryItem{}, nil
// 		}
// 		return nil, err
// 	}

// 	// Return only the last 'limit' items
// 	history := result.ReadHistory
// 	if len(history) > limit {
// 		history = history[len(history)-limit:]
// 	}

// 	return history, nil
// }

// // ============================================================================
// // DELETE OPERATIONS
// // ============================================================================

// // Delete removes a user from the database
// func (r *Repository) Delete(ctx context.Context, id bson.ObjectID) error {
// 	filter := bson.M{"_id": id}

// 	result, err := r.coll.DeleteOne(ctx, filter)
// 	if err != nil {
// 		return err
// 	}

// 	if result.DeletedCount == 0 {
// 		return mongo.ErrNoDocuments
// 	}

// 	return nil
// }

// // ============================================================================
// // AGGREGATION OPERATIONS
// // ============================================================================

// // GetUserStats retrieves statistics about a user
// func (r *Repository) GetUserStats(ctx context.Context, userID bson.ObjectID) (map[string]interface{}, error) {
// 	var user models.User
// 	err := r.coll.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	stats := map[string]interface{}{
// 		"total_bookmarks":    len(user.BookmarkedNews),
// 		"total_read":         len(user.ReadHistory),
// 		"interests_count":    len(user.Interests),
// 		"trader_type":        user.TraderType,
// 		"member_since":       user.CreatedAt,
// 		"last_active":        user.LastLogin,
// 		"notification_enabled": user.Preferences.NotificationEnabled,
// 	}

// 	return stats, nil
// }

// // CountByTraderType counts users by trader type
// func (r *Repository) CountByTraderType(ctx context.Context) (map[string]int64, error) {
// 	pipeline := []bson.M{
// 		{
// 			"$group": bson.M{
// 				"_id":   "$trader_type",
// 				"count": bson.M{"$sum": 1},
// 			},
// 		},
// 	}

// 	cursor, err := r.coll.Aggregate(ctx, pipeline)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	results := make(map[string]int64)
// 	for cursor.Next(ctx) {
// 		var result struct {
// 			ID    string `bson:"_id"`
// 			Count int64  `bson:"count"`
// 		}
// 		if err := cursor.Decode(&result); err != nil {
// 			return nil, err
// 		}
// 		results[result.ID] = result.Count
// 	}

// 	return results, nil
// }

// // ============================================================================
// // SEARCH OPERATIONS
// // ============================================================================

// // Search searches users by name or email
// func (r *Repository) Search(ctx context.Context, query string, pagination models.PaginationParams) ([]models.User, int64, error) {
// 	// Case-insensitive regex search
// 	filter := bson.M{
// 		"$or": []bson.M{
// 			{"name": bson.M{"$regex": query, "$options": "i"}},
// 			{"email": bson.M{"$regex": query, "$options": "i"}},
// 		},
// 	}

// 	totalCount, err := r.coll.CountDocuments(ctx, filter)
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	findOptions := options.Find().
// 		SetSkip(int64(pagination.GetSkip())).
// 		SetLimit(int64(pagination.Limit)).
// 		SetSort(bson.D{{Key: "created_at", Value: -1}})

// 	cursor, err := r.coll.Find(ctx, filter, findOptions)
// 	if err != nil {
// 		return nil, 0, err
// 	}
// 	defer cursor.Close(ctx)

// 	var users []models.User
// 	if err = cursor.All(ctx, &users); err != nil {
// 		return nil, 0, err
// 	}

// 	return users, totalCount, nil
// }

// // ============================================================================
// // EXISTENCE CHECKS
// // ============================================================================

// // Exists checks if a user exists by ID
// func (r *Repository) Exists(ctx context.Context, id bson.ObjectID) (bool, error) {
// 	count, err := r.coll.CountDocuments(ctx, bson.M{"_id": id})
// 	if err != nil {
// 		return false, err
// 	}
// 	return count > 0, nil
// }

// // EmailExists checks if an email is already registered
// func (r *Repository) EmailExists(ctx context.Context, email string) (bool, error) {
// 	count, err := r.coll.CountDocuments(ctx, bson.M{"email": email})
// 	if err != nil {
// 		return false, err
// 	}
// 	return count > 0, nil
// }

// // GoogleIDExists checks if a Google ID is already registered
// func (r *Repository) GoogleIDExists(ctx context.Context, googleID string) (bool, error) {
// 	count, err := r.coll.CountDocuments(ctx, bson.M{"google_id": googleID})
// 	if err != nil {
// 		return false, err
// 	}
// 	return count > 0, nil
// }