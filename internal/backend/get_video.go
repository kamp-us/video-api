package backend

import (
	"context"
	"github.com/kamp-us/video-api/internal/models"
)

func (b *PostgreSQLBackend) GetVideo(ctx context.Context, id string) (*models.Video, error) {
	video := models.Video{}
	result := b.DB.First(&video, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &video, nil
}
