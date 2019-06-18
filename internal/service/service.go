package service

import (
	"SnowBrick-Backend/conf"
	"SnowBrick-Backend/internal/dao"
	"context"
)

// Service service.
type Service struct {
	c   *conf.Config
	dao *dao.Dao
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
