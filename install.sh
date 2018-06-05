#!/bin/bash

set -e

run=task-deploy
workdir=$HOME/task-deploy
echo Workspace: $workdir
if [ ! -d "$workdir" ]; then
  mkdir -p $workdir
fi

#编译
go build -o $workdir/$run

#生成配置文件夹 复制项目示例配置文件到实际位置
if [ ! -d "$workdir/conf" ]; then
  mkdir -p  $workdir/conf
fi
if [ ! -f "$workdir/conf/config.yaml" ]; then
  cp conf/config.yaml.example $workdir/conf/config.yaml
fi


#生成systemd 文件
serviceFilePath=$workdir/$run.service
sed "s|\$name|${run}|g; s|\$work|${workdir}|g" systemd/myregister.service > $serviceFilePath
cp $serviceFilePath /etc/systemd/system/$run.service
systemctl enable $run.service

echo Install finish! but some step continue:
echo 1. Please monidify $workdir/$run/conf/config.yaml
echo 2. start service: \'systemctl start $run.service\'