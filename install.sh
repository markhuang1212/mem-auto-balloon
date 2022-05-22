#!/bin/sh

set -x

mkdir -p /opt/memdeflate

cp memdeflate /opt/memdeflate/

cp template/* /etc/systemd/system

systemctl daemon-reload
