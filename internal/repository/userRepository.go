package repository

import (
	"context"
	"errors"
	"yoharsh14/krant-backend/internal/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	user.ID = bson.NewObjectID()
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepository) FindByGoogleID(ctx context.Context, googleID string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"google_id": googleID}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // User not found
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id bson.ObjectID) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, id bson.ObjectID, update bson.M) error {
	filter := bson.M{"_id": id}
	updateDoc := bson.M{"$set": update}
	
	result, err := r.collection.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return err
	}
	
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	
	return nil
}

func (r *UserRepository) AddBookmark(ctx context.Context, userID, newsID bson.ObjectID) error {
	filter := bson.M{"_id": userID}
	update := bson.M{
		"$addToSet": bson.M{"bookmarked_news": newsID},
	}
	
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *UserRepository) RemoveBookmark(ctx context.Context, userID, newsID bson.ObjectID) error {
	filter := bson.M{"_id": userID}
	update := bson.M{
		"$pull": bson.M{"bookmarked_news": newsID},
	}
	
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}