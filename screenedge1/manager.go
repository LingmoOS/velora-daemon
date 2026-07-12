// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package screenedge

import (
	"github.com/LingmoOS/velora-daemon/common/dsync"
	wm "github.com/LingmoOS/golang-github-lingmo-go-dbus-factory/session/com.lingmo.wm"
	"github.com/LingmoOS/golang-github-lingmo-go-lib/dbusutil"
)

//go:generate dbusutil-gen em -type Manager

const (
	TopLeft     = "left-up"
	TopRight    = "right-up"
	BottomLeft  = "left-down"
	BottomRight = "right-down"

	dbusServiceName = "org.deepin.dde.Zone1"
	dbusPath        = "/org/deepin/dde/Zone1"
	dbusInterface   = "org.lingmo.Zone1"

	wmDBusServiceName = "com.deepin.wm"
)

type Manager struct {
	service        *dbusutil.Service
	settings       *Settings
	wm             wm.Wm
	sessionSigLoop *dbusutil.SignalLoop
	syncConfig     *dsync.Config
}

func newManager(service *dbusutil.Service) *Manager {
	var m = new(Manager)
	m.service = service
	m.settings = NewSettings()
	m.wm = wm.NewWm(service.Conn())
	m.sessionSigLoop = dbusutil.NewSignalLoop(service.Conn(), 10)
	m.sessionSigLoop.Start()
	m.syncConfig = dsync.NewConfig("screen_edge", &syncConfig{m: m},
		m.sessionSigLoop, dbusPath, logger)
	return m
}

func (m *Manager) destroy() {
	m.settings.Destroy()
	m.sessionSigLoop.Stop()
	m.syncConfig.Destroy()
}

func (*Manager) GetInterfaceName() string {
	return dbusInterface
}
