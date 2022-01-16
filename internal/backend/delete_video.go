package backend

import (
	"context"
	"github.com/kamp-us/video-api/internal/models"
)

func (b *PostgreSQLBackend) DeleteVideo(ctx context.Context, id string) error {
	video := models.Video{}
	result := b.DB.First(&video, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	result = b.DB.Delete(&video)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
