// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"github.com/LingmoOS/velora-daemon/langselector1"
	"github.com/LingmoOS/golang-github-lingmo-go-lib/gettext"
)

func main() {
	gettext.InitI18n()
	gettext.BindTextdomainCodeset("dde-daemon", "UTF-8")
	gettext.Textdomain("dde-daemon")
	langselector.Run()
}
