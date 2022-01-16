package server

import (
	"context"
	api "github.com/kamp-us/video-api/rpc/video-api"
	"github.com/twitchtv/twirp"
)

func (s *VideoAPIServer) CreateVideo(ctx context.Context, req *api.CreateVideoRequest) (*api.Video, error) {
	if err := validateCreateVideoRequest(req); err != nil {
		return nil, err
	}

	video, err := s.backend.CreateVideo(ctx, req.UserId, req.Name, req.URI)
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

func validateCreateVideoRequest(req *api.CreateVideoRequest) error {
	if req.UserId == "" {
		return twirp.RequiredArgumentError("user_id")
	}
	if req.Name == "" {
		return twirp.RequiredArgumentError("name")
	}
	if req.URI == "" {
		return twirp.RequiredArgumentError("URI")
	}
	return nil
}
