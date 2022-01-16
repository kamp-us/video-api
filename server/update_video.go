package server

import (
	"context"

	api "github.com/kamp-us/video-api/rpc/video-api"

	"github.com/twitchtv/twirp"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *VideoAPIServer) UpdateVideo(ctx context.Context, req *api.UpdateVideoRequest) (*emptypb.Empty, error) {

	if err := validateUpdateVideoRequest(req); err != nil {
		return nil, err
	}

	if err := s.backend.UpdateVideo(ctx,
		req.ID,
		convertToStringPtr(req.Name),
		convertToStringPtr(req.URI)); err != nil {
		return nil, twirp.InternalErrorWith(err)
	}
	return &emptypb.Empty{}, nil
}

func validateUpdateVideoRequest(req *api.UpdateVideoRequest) error {
	if req.ID == "" {
		return twirp.RequiredArgumentError("id")
	}
	return nil
}
