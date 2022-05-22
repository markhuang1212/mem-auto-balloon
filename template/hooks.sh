#!/bin/sh

# Usage: 
# 1. Fill appropiate VM_NAME
# 2. Copy this file to /etc/libvirt/hooks/qemu.d

VM_NAME=

if [[ $1 == $VM_NAME ]] && [[ $2 == started ]]
then
    systemctl start memdeflate@$VM_NAME.timer
fi

if [[ $1 == $VM_NAME ]] && [[ $2 == stopped ]]
then
    systemctl stop memdeflate@$VM_NAME.timer
fi

