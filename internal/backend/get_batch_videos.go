package backend

import (
	"context"
	"github.com/kamp-us/video-api/internal/models"
)

func (b *PostgreSQLBackend) GetBatchVideos(ctx context.Context, ids []string) ([]*models.Video, error) {
	var videos []*models.Video
	result := b.DB.Find(&videos, ids)
	if result.Error != nil {
		return nil, result.Error
	}

	return videos, nil
}
