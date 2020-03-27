package http

import (
	"net/http"

	"github.com/rs/cors"
	"go-skeleton/pkg/grace"
)

// SkeletonHandler ...
type SkeletonHandler interface {
	// Masukkan fungsi handler di sini
	SkeletonHandler(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	server   *http.Server
	Skeleton SkeletonHandler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	handler := cors.AllowAll().Handler(s.Handler())
	return grace.Serve(port, handler)
}
