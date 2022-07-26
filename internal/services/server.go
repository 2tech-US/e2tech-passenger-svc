package services

import (
	"github.com/lntvan166/e2tech-passenger-svc/internal/config"
	"github.com/lntvan166/e2tech-passenger-svc/internal/db"
	"github.com/lntvan166/e2tech-passenger-svc/internal/pb"
)

type Server struct {
	DB     *db.Queries
	Config *config.Config
	pb.UnimplementedPassengerServiceServer
}
