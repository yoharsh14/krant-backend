package models

import (
	"time"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// News represents a news article
type News struct {
	ID              bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Title           string        `json:"title" bson:"title"`
	Description     string        `json:"description" bson:"description"`
	Content         string        `json:"content" bson:"content"`
	Author          string        `json:"author" bson:"author"`
	Source          NewsSource    `json:"source" bson:"source"`
	ImageURL        string        `json:"image_url" bson:"image_url"`
	Categories      []string      `json:"categories" bson:"categories"`
	Tags            []string      `json:"tags" bson:"tags"`
	TraderRelevance []string      `json:"trader_relevance" bson:"trader_relevance"`
	PublishedAt     time.Time     `json:"published_at" bson:"published_at"`
	Metrics         NewsMetrics   `json:"metrics" bson:"metrics"`
	Status          string        `json:"status" bson:"status"` // draft, published, archived
	CreatedAt       time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at" bson:"updated_at"`
}

// NewsSource represents the source of the news
type NewsSource struct {
	Name string `json:"name" bson:"name"`
	URL  string `json:"url" bson:"url"`
}

// NewsMetrics represents engagement metrics for a news article
type NewsMetrics struct {
	Views     int64 `json:"views" bson:"views"`
	Bookmarks int64 `json:"bookmarks" bson:"bookmarks"`
	Shares    int64 `json:"shares" bson:"shares"`
}

// CreateNewsInput represents input for creating news
type CreateNewsInput struct {
	Title           string     `json:"title" binding:"required"`
	Description     string     `json:"description" binding:"required"`
	Content         string     `json:"content" binding:"required"`
	Author          string     `json:"author"`
	Source          NewsSource `json:"source" binding:"required"`
	ImageURL        string     `json:"image_url"`
	Categories      []string   `json:"categories" binding:"required,min=1"`
	Tags            []string   `json:"tags"`
	TraderRelevance []string   `json:"trader_relevance" binding:"required,min=1"`
	PublishedAt     time.Time  `json:"published_at"`
}

// UpdateNewsInput represents input for updating news
type UpdateNewsInput struct {
	Title           *string     `json:"title,omitempty"`
	Description     *string     `json:"description,omitempty"`
	Content         *string     `json:"content,omitempty"`
	Author          *string     `json:"author,omitempty"`
	Source          *NewsSource `json:"source,omitempty"`
	ImageURL        *string     `json:"image_url,omitempty"`
	Categories      *[]string   `json:"categories,omitempty"`
	Tags            *[]string   `json:"tags,omitempty"`
	TraderRelevance *[]string   `json:"trader_relevance,omitempty"`
	Status          *string     `json:"status,omitempty"`
}

// NewsResponse represents news data returned to client
type NewsResponse struct {
	ID              string      `json:"id"`
	Title           string      `json:"title"`
	Description     string      `json:"description"`
	Content         string      `json:"content"`
	Author          string      `json:"author"`
	Source          NewsSource  `json:"source"`
	ImageURL        string      `json:"image_url"`
	Categories      []string    `json:"categories"`
	Tags            []string    `json:"tags"`
	TraderRelevance []string    `json:"trader_relevance"`
	PublishedAt     time.Time   `json:"published_at"`
	Metrics         NewsMetrics `json:"metrics"`
	IsBookmarked    bool        `json:"is_bookmarked"` // Set based on user context
	CreatedAt       time.Time   `json:"created_at"`
}

// ToResponse converts News to NewsResponse
func (n *News) ToResponse(isBookmarked bool) NewsResponse {
	return NewsResponse{
		ID:              n.ID.Hex(),
		Title:           n.Title,
		Description:     n.Description,
		Content:         n.Content,
		Author:          n.Author,
		Source:          n.Source,
		ImageURL:        n.ImageURL,
		Categories:      n.Categories,
		Tags:            n.Tags,
		TraderRelevance: n.TraderRelevance,
		PublishedAt:     n.PublishedAt,
		Metrics:         n.Metrics,
		IsBookmarked:    isBookmarked,
		CreatedAt:       n.CreatedAt,
	}
}

// NewsListResponse represents paginated news response
type NewsListResponse struct {
	News       []NewsResponse `json:"news"`
	Page       int            `json:"page"`
	Limit      int            `json:"limit"`
	TotalCount int64          `json:"total_count"`
	TotalPages int            `json:"total_pages"`
	HasMore    bool           `json:"has_more"`
}

// NewsStatus constants
const (
	NewsStatusDraft     = "draft"
	NewsStatusPublished = "published"
	NewsStatusArchived  = "archived"
)

// ValidNewsStatuses returns all valid news statuses
func ValidNewsStatuses() []string {
	return []string{
		NewsStatusDraft,
		NewsStatusPublished,
		NewsStatusArchived,
	}
}

// IsValidNewsStatus checks if the news status is valid
func IsValidNewsStatus(status string) bool {
	for _, s := range ValidNewsStatuses() {
		if s == status {
			return true
		}
	}
	return false
}