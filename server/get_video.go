package server

import (
	"context"
	api "github.com/kamp-us/video-api/rpc/video-api"
	"github.com/twitchtv/twirp"
)

func (s *VideoAPIServer) GetVideo(ctx context.Context, req *api.GetVideoRequest) (*api.Video, error) {

	if err := validateGetVideoRequest(req); err != nil {
		return nil, err
	}

	video, err := s.backend.GetVideo(ctx, req.ID)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &api.Video{
		ID:     video.ID.String(),
		UserId: video.UserID,
		Name:   video.Name,
		URI:    video.URI,
		Slug:   video.Slug,
	}, nil

}

func validateGetVideoRequest(req *api.GetVideoRequest) error {
	if req.ID == "" {
		return twirp.RequiredArgumentError("id")
	}
	return nil
}
