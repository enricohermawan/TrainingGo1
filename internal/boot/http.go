package boot

import (
	"log"
	"net/http"

	"go-skeleton/internal/config"

	"github.com/jmoiron/sqlx"

	skeletonData "go-skeleton/internal/data/skeleton"
	skeletonServer "go-skeleton/internal/delivery/http"
	skeletonHandler "go-skeleton/internal/delivery/http/skeleton"
	skeletonService "go-skeleton/internal/service/skeleton"
)

// HTTP will load configuration, do dependency injection and then start the HTTP server
func HTTP() error {
	var (
		s   skeletonServer.Server    // HTTP Server Object
		sd  skeletonData.Data        // Domain data layer
		ss  skeletonService.Service  // Domain service layer
		sh  *skeletonHandler.Handler // Domain handler
		cfg *config.Config           // Configuration object
	)

	err := config.Init()
	if err != nil {
		log.Fatalf("[CONFIG] Failed to initialize config: %v", err)
	}
	cfg = config.Get()
	// Open MySQL DB Connection
	db, err := sqlx.Open("mysql", cfg.Database.Master)
	if err != nil {
		log.Fatalf("[DB] Failed to initialize database connection: %v", err)
	}

	// Diganti dengan domain yang anda buat
	sd = skeletonData.New(db)
	ss = skeletonService.New(sd)
	sh = skeletonHandler.New(ss)

	s = skeletonServer.Server{
		Skeleton: sh,
	}

	if err := s.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}
