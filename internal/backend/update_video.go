package backend

import (
	"context"
	"github.com/gosimple/slug"
	"github.com/kamp-us/video-api/internal/models"
)

func (b *PostgreSQLBackend) UpdateVideo(ctx context.Context, id string, name *string, URI *string) error {
	video := models.Video{}
	result := b.DB.First(&video, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	updates := models.Video{Name: *name, Slug: slug.MakeLang(*name, "tr"), URI: *URI}
	result = b.DB.Model(&video).Updates(updates)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
