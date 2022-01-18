package server

import (
	"context"
	api "github.com/kamp-us/video-api/rpc/video-api"
	"github.com/twitchtv/twirp"
)

func (s VideoAPIServer) GetBatchVideos(ctx context.Context, req *api.GetBatchVideosRequest) (*api.GetBatchVideosResponse, error) {
	if err := validateGetBatchVideosRequest(req); err != nil {
		return nil, err
	}
	videos, err := s.backend.GetBatchVideos(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	var batch []*api.Video
	for _, model := range videos {
		video := &api.Video{
			ID:     model.ID.String(),
			UserId: model.UserID,
			Name:   model.Name,
			URI:    model.URI,
			Slug:   model.Slug,
		}
		batch = append(batch, video)
	}
	return &api.GetBatchVideosResponse{Videos: batch}, nil

}

func validateGetBatchVideosRequest(req *api.GetBatchVideosRequest) error {
	if len(req.Ids) == 0 {
		return twirp.RequiredArgumentError("ids")
	}
	return nil
}
