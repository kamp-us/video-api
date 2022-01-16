package backend

import (
	"context"
	"github.com/gosimple/slug"
	"github.com/kamp-us/video-api/internal/models"
)

func (b *PostgreSQLBackend) CreateVideo(ctx context.Context, userID string, name string, URI string) (*models.Video, error) {
	video := models.Video{
		UserID: userID,
		Slug:   slug.MakeLang(name, "tr"),
		Name:   name,
		URI:    URI,
	}

	result := b.DB.Create(&video)
	if result.Error != nil {
		return nil, result.Error
	}

	return &video, nil
}
