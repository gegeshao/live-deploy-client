#!/bin/sh
gox -osarch="linux/386" -output="./bin/nginx-panel-x86"
gox -osarch="linux/amd64" -output="./bin/nginx-panel-x64"
