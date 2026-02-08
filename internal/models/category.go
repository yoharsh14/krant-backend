package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// Category represents a news category
type Category struct {
	ID             bson.ObjectID  `json:"id" bson:"_id,omitempty"`
	Name           string         `json:"name" bson:"name"`
	Slug           string         `json:"slug" bson:"slug"`
	Description    string         `json:"description" bson:"description"`
	Icon           string         `json:"icon" bson:"icon"`
	Color          string         `json:"color" bson:"color"`
	Order          int            `json:"order" bson:"order"`
	IsActive       bool           `json:"is_active" bson:"is_active"`
	ParentCategory *bson.ObjectID `json:"parent_category,omitempty" bson:"parent_category,omitempty"`
	CreatedAt      time.Time      `json:"created_at" bson:"created_at"`
}

// CreateCategoryInput represents input for creating a category
type CreateCategoryInput struct {
	Name           string         `json:"name" binding:"required"`
	Slug           string         `json:"slug" binding:"required"`
	Description    string         `json:"description"`
	Icon           string         `json:"icon"`
	Color          string         `json:"color"`
	Order          int            `json:"order"`
	ParentCategory *bson.ObjectID `json:"parent_category,omitempty"`
}

// UpdateCategoryInput represents input for updating a category
type UpdateCategoryInput struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Icon        *string `json:"icon,omitempty"`
	Color       *string `json:"color,omitempty"`
	Order       *int    `json:"order,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty"`
}

// CategoryResponse represents category data returned to client
type CategoryResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Color       string `json:"color"`
	Order       int    `json:"order"`
	IsActive    bool   `json:"is_active"`
}

// ToResponse converts Category to CategoryResponse
func (c *Category) ToResponse() CategoryResponse {
	return CategoryResponse{
		ID:          c.ID.Hex(),
		Name:        c.Name,
		Slug:        c.Slug,
		Description: c.Description,
		Icon:        c.Icon,
		Color:       c.Color,
		Order:       c.Order,
		IsActive:    c.IsActive,
	}
}