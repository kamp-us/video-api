package backend

import (
	"context"
	"github.com/kamp-us/video-api/internal/models"
	"gorm.io/gorm"
)

type Backender interface {
	GetVideo(ctx context.Context, id string) (*models.Video, error)
	CreateVideo(ctx context.Context, userID string, name string, URI string) (*models.Video, error)
	UpdateVideo(ctx context.Context, id string, name *string, URI *string) error
	DeleteVideo(ctx context.Context, id string) error
}

type PostgreSQLBackend struct {
	DB *gorm.DB
}

func NewPostgreSQLBackend(db *gorm.DB) Backender {
	return &PostgreSQLBackend{
		DB: db,
	}
}
