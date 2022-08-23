package config

import (
	"encoding/json"
	"github.com/GoAdminGroup/go-admin/modules/utils"
)

// FileUploadEngine is a file upload engine.
type FileUploadEngine struct {
	Name   string                 `json:"name,omitempty" yaml:"name,omitempty" ini:"name,omitempty"`
	Config map[string]interface{} `json:"config,omitempty" yaml:"config,omitempty" ini:"config,omitempty"`
}

func (f FileUploadEngine) JSON() string {
	if f.Name == "" {
		return ""
	}
	if len(f.Config) == 0 {
		f.Config = nil
	}
	return utils.JSON(f)
}

func GetFileUploadEngineFromJSON(m string) FileUploadEngine {
	var f FileUploadEngine
	if m == "" {
		return f
	}
	_ = json.Unmarshal([]byte(m), &f)
	return f
}
