#!/bin/sh
case $1/$2 in
        pre/*)
        ;;
        post/*)
            gdbus call -y -d org.lingmo.Power1 -o /org/lingmo/Power1 -m org.lingmo.Power1.Refresh --timeout 2
        ;;
esac
