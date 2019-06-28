package service

import (
	"SnowBrick-Backend/conf"
	"SnowBrick-Backend/internal/brick/dao"
	"context"
	"golang.org/x/sync/errgroup"
)

// Service service.
type Service struct {
	c   *conf.Config
	dao *dao.Dao
	eg  errgroup.Group
}

// New new a service and return.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:   c,
		dao: dao.New(c),
	}
	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	s.dao.Close()
}

func (s *Service) Go(f func() error) {
	s.eg.Go(f)
}
