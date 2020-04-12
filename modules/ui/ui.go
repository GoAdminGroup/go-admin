package ui

import (
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/template/types"
)

type Service struct {
	NavButtons types.Buttons
}

const ServiceKey = "ui"

func (s *Service) Name() string {
	return "ui"
}

func GetService(srv service.List) *Service {
	if v, ok := srv.Get("ui").(*Service); ok {
		return v
	}
	panic("wrong service")
}

func NewService(btns types.Buttons) *Service {
	return &Service{
		NavButtons: btns,
	}
}
