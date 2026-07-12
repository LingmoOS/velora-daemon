// SPDX-FileCopyrightText: 2018 - 2026 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"fmt"
	"sync"

	_ "github.com/LingmoOS/velora-daemon/display1"
	"github.com/LingmoOS/velora-daemon/loader"

	_ "github.com/LingmoOS/velora-daemon/audio1"
	_ "github.com/LingmoOS/velora-daemon/bluetooth1"
	_ "github.com/LingmoOS/velora-daemon/screenedge1"

	// depends: network
	_ "github.com/LingmoOS/velora-daemon/calltrace"
	_ "github.com/LingmoOS/velora-daemon/clipboard1"
	_ "github.com/LingmoOS/velora-daemon/debug"

	_ "github.com/LingmoOS/velora-daemon/gesture1"
	_ "github.com/LingmoOS/velora-daemon/housekeeping"
	_ "github.com/LingmoOS/velora-daemon/inputdevices1"
	_ "github.com/LingmoOS/velora-daemon/keybinding1"
	_ "github.com/LingmoOS/velora-daemon/lastore1"

	_ "github.com/LingmoOS/velora-daemon/grub_gfx"
	_ "github.com/LingmoOS/velora-daemon/screensaver1"
	_ "github.com/LingmoOS/velora-daemon/service_trigger"
	_ "github.com/LingmoOS/velora-daemon/session/eventlog"
	_ "github.com/LingmoOS/velora-daemon/session/power1"
	_ "github.com/LingmoOS/velora-daemon/sessionwatcher1"
	_ "github.com/LingmoOS/velora-daemon/systeminfo1"
	_ "github.com/LingmoOS/velora-daemon/timedate1"
	_ "github.com/LingmoOS/velora-daemon/trayicon1"
	_ "github.com/LingmoOS/velora-daemon/x_event_monitor1"
)

var (
	moduleLocker sync.Mutex
)

func (s *SessionDaemon) checkDependencies(module loader.Module, enabled bool) error {
	if enabled {
		depends := module.GetDependencies()
		for _, n := range depends {
			if !s.getConfigValue(n) {
				return fmt.Errorf("Dependency lose: %v", n)
			}
		}
		return nil
	}

	for _, m := range loader.List() {
		if m == nil || m.Name() == module.Name() {
			continue
		}

		if m.IsEnable() && isStrInList(module.Name(), m.GetDependencies()) {
			return fmt.Errorf("Can not disable this module '%s', because of it was depended by'%s'",
				module.Name(), m.Name())
		}
	}
	return nil
}

func isStrInList(item string, list []string) bool {
	for _, v := range list {
		if item == v {
			return true
		}
	}
	return false
}
