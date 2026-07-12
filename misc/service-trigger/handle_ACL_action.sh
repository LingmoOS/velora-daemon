#!/bin/bash

available=$(busctl call -- com.lingmo.daemon.ACL /org/lingmo/security/hierarchical/Control org.lingmo.security.hierarchical.Control Availabled|awk '{print $2}')
if [ "$available" == "true" ];then
    msg=$(gettext -d dde-daemon 'You cannot run the unverified "%s", but you can change the settings in Security Center.')
    msg=$(printf "$msg" $2)
    gdbus call --session --dest org.freedesktop.Notifications --object-path /org/freedesktop/Notifications --method org.freedesktop.Notifications.Notify dde-control-center 0 preferences-system '' "$msg" '["actions","'$(gettext -d dde-daemon 'Proceed')'"]'  '{"x-deepin-action-actions":<"dbus-send,--session,--print-reply,--dest=com.lingmo.defender.hmiscreen,/com/lingmo/defender/hmiscreen,com.lingmo.defender.hmiscreen.ShowPage,string:securitytools,string:application-safety">}' 5000
else
    msg=$(gettext -d dde-daemon '"%s" did not pass the system security verification, and cannot run now')
    msg=$(printf "$msg" $2)
    notify-send -i preferences-system -a dde-control-center "$msg"
fi