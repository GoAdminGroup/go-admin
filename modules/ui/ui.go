package ui

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
)

type Service struct {
	NavButtons *types.Buttons
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

func NewService(btns *types.Buttons) *Service {
	return &Service{
		NavButtons: btns,
	}
}

func (s *Service) UpdateButtons() {

}

func (s *Service) RemoveOrShowSiteNavButton(remove bool) {
	if remove {
		*s.NavButtons = (*s.NavButtons).RemoveSiteNavButton()
	} else {
		*s.NavButtons = (*s.NavButtons).AddNavButton(icon.Gear, types.NavBtnSiteName,
			action.JumpInNewTab(config.Url("/info/site/edit"),
				language.GetWithScope("site setting", "config")))
	}
}

func (s *Service) RemoveOrShowInfoNavButton(remove bool) {
	if remove {
		*s.NavButtons = (*s.NavButtons).RemoveInfoNavButton()
	} else {
		*s.NavButtons = (*s.NavButtons).AddNavButton(icon.Info, types.NavBtnInfoName,
			action.JumpInNewTab(config.Url("/application/info"),
				language.GetWithScope("system info", "system")))
	}

}

func (s *Service) RemoveOrShowToolNavButton(remove bool) {
	if remove {
		*s.NavButtons = (*s.NavButtons).RemoveToolNavButton()
	} else {
		*s.NavButtons = (*s.NavButtons).AddNavButton(icon.Wrench, types.NavBtnToolName,
			action.JumpInNewTab(config.Url("/info/generate/new"),
				language.GetWithScope("tool", "tool")))
	}

}

func (s *Service) RemoveOrShowPlugNavButton(remove bool) {
	if remove {
		*s.NavButtons = (*s.NavButtons).RemovePlugNavButton()
	} else {
		*s.NavButtons = (*s.NavButtons).AddNavButton(icon.Plug, types.NavBtnToolName,
			action.JumpInNewTab(config.Url("/plugin"),
				language.GetWithScope("plugin", "plugin")))
	}

}
