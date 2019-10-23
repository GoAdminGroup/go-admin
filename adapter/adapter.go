// Copyright 2019 GoAdmin Core Team.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package adapter

import (
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/template/types"
)

// WebFrameWork is a interface which is used as an adapter of
// framework and goAdmin. It must implement two methods. Use registers
// the routes and the corresponding handlers. Content writes the
// response to the corresponding context of framework.
type WebFrameWork interface {
	Use(interface{}, []plugins.Plugin) error
	Content(interface{}, types.GetPanel)
}
