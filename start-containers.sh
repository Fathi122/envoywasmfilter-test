#!/bin/bash
docker run -d --name svc_httpheaderslogs httpserv-svc:latest

docker run --name envoy -d -p 8090:8090 --link svc_httpheaderslogs envoy-proxy-test:latest -c envoy.yaml --log-path logs/custom.log -l debug
