#!/bin/bash

# clean up previous containers running
declare dockerimages=("envoy" "svc_httpheaderslogs")
for id in "${dockerimages[@]}";do docker rm -f $id;done

docker run -d  --name svc_httpheaderslogs httpserv-svc:latest

docker run -d --name envoy -p 8090:8090 --link svc_httpheaderslogs envoy-proxy-test:latest -c envoy.yaml --log-path logs/custom.log -l debug
