#!/bin/bash

count=$(ps -ef | grep deepin-camera | grep -w -c $USER)
echo $count
if [ $count -ge 2 ]
then
	echo "close deepin-camera"
	killall deepin-camera
else
	echo "open deepin-camera"
	dbus-send --session --print-reply --dest=com.lingmo.SessionManager /com/lingmo/StartManager com.lingmo.StartManager.Launch string:/usr/share/applications/deepin-camera.desktop
fi

