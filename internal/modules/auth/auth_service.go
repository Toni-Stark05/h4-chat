package auth

import (
	"github.com/samber/do"
	"log/slog"
)

type Service struct {
	logger *slog.Logger
}

func NewService(i *do.Injector) (*Service, error) {
	logger := do.MustInvoke[*slog.Logger](i)

	return &Service{logger}, nil
}

func (s *Service) Login() (string, error) {
	return "Test", nil
}
