#!/bin/sh
gox -osarch="linux/386" -output="./bin/x86/nginx-panel-client"
gox -osarch="linux/amd64" -output="./bin/x64/nginx-panel-client"
#gox -osarch="darwin/amd64" -output="./bin/macOS/nginx-panel-client"
#gox -osarch="linux/arm" -output="./bin/linux-arm/nginx-panel-client"