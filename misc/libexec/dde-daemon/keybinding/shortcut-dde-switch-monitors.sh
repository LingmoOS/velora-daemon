#!/bin/bash

if ! outputs=$(dbus-send --print-reply --dest=org.lingmo.Display1 /org/lingmo/Display1 org.lingmo.Display1.ListOutputNames | grep -oP 'string "\K[^"]+'); then
    exit 1
fi

if [ $(echo "$outputs" | wc -l) -gt 1 ]; then
    dbus-send --print-reply --dest=org.lingmo.Osd1 / org.lingmo.Osd1.ShowOSD string:SwitchMonitors
fi
