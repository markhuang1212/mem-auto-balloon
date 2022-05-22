#!/bin/sh

set -x

# Be root if not already
[ "$UID" -eq 0 ] || exec sudo "$0" "$@"

mkdir -p /opt/memdeflate

cp memdeflate /opt/memdeflate/

cp template/memdeflate@.service /etc/systemd/system
cp template/memdeflate@.timer /etc/systemd/system

systemctl daemon-reload

if [[ -n $1 ]]
then
    VM_NAME=$1
    REPLACE_CMD=s/VM_NAME=/VM_NAME=$VM_NAME/g
    cat template/hooks.sh | sed $REPLACE_CMD | tee /etc/libvirt/hooks/qemu.d/$VM_NAME.sh
    systemctl restart virtqemud
fi
