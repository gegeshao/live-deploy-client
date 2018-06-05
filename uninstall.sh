#!/bin/bash

set -e

run=task-deploy
workdir=$HOME/task-deploy

if [ ! -f "/etc/systemd/system/$run.service" ]; then
  echo "不存在服务"
  exit 0
fi

systemctl stop $run.service
systemctl disable $run.service
rm /etc/systemd/system/$run.service

if [ -d "$workdir" ]; then
  rm -rf $workdir
fi

echo Uninstall finish!