package server

import (
	"github.com/kamp-us/video-api/internal/backend"
)

type VideoAPIServer struct {
	backend backend.Backender
}

func NewVideoAPIServer(backend backend.Backender) *VideoAPIServer {
	return &VideoAPIServer{backend: backend}
}
