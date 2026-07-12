// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package daemon

//go:generate go build -o target/ github.com/LingmoOS/velora-daemon/bin/dde-session-daemon
//go:generate go build -o target/ github.com/LingmoOS/velora-daemon/bin/dde-system-daemon
//go:generate go build -o target/ github.com/LingmoOS/velora-daemon/bin/grub2
//go:generate go build -o target/ github.com/LingmoOS/velora-daemon/bin/search
//go:generate go build -o target/ github.com/LingmoOS/velora-daemon/bin/backlight_helper
//go:generate go build -o target/ github.com/LingmoOS/velora-daemon/bin/langselector
//go:generate go build -o target/ github.com/LingmoOS/velora-daemon/bin/soundeffect
//go:generate go build -o target/ github.com/LingmoOS/velora-daemon/bin/dde-lockservice
//go:generate go build -o target/ github.com/LingmoOS/velora-daemon/bin/dde-authority
//go:generate go build -o target/ github.com/LingmoOS/velora-daemon/bin/default-terminal
//go:generate go build -o target/ github.com/LingmoOS/velora-daemon/bin/dde-greeter-setter
