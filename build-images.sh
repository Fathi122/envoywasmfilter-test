#!/bin/bash

 docker build -f Dockerfile-envoy --rm --no-cache -t envoy-proxy-test:latest .

 cd http-server && docker build -f Dockerfile --rm --no-cache -t httpserv-svc:latest .
 