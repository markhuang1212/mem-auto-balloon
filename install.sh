#!/bin/sh

set -x

# Be root if not already
[ "$UID" -eq 0 ] || exec sudo "$0" "$@"

mkdir -p /opt/memdeflate

cp memdeflate /opt/memdeflate/

cp template/* /etc/systemd/system

systemctl daemon-reload
