package server

import (
	"context"
	api "github.com/kamp-us/video-api/rpc/video-api"
	"github.com/twitchtv/twirp"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (s *VideoAPIServer) DeleteVideo(ctx context.Context, req *api.DeleteVideoRequest) (*emptypb.Empty, error) {
	if err := validateDeleteVideoRequest(req); err != nil {
		return nil, err
	}
	err := s.backend.DeleteVideo(ctx, req.ID)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}
	return &emptypb.Empty{}, nil
}

func convertToStringPtr(value *wrapperspb.StringValue) *string {
	val := value.GetValue()
	return &val
}

func validateDeleteVideoRequest(req *api.DeleteVideoRequest) error {
	if req.ID == "" {
		return twirp.RequiredArgumentError("id")
	}
	return nil
}
