# envoy wasm filter test

Build envoy-proxy with http wasm filter and http server backend
 ```
 docker build -f Dockerfile-envoy --rm --no-cache -t envoy-proxy-test:latest .

 cd http-server && docker build -f Dockerfile --rm --no-cache -t httpserv-svc:latest .
```
Start http server backend server
```
docker run -it --rm --name svc_httpheaderslogs httpserv-svc:latest
```

Start envoy-proxy
```
docker run -it --name envoy --rm  -p 8090:8090 --link svc_httpheaderslogs envoy-proxy-test:latest -c envoy.yaml --log-path logs/custom.log -l debug
```

Test exposed endpoint
```
curl localhost:8090/test200
```

And Observe logs on logs/custom.log from envoy container
```
ENVOY_CONTAINER_ID=$(docker ps --filter "name=envoy" --format={{.ID}}) && docker exec -it ${ENVOY_CONTAINER_ID} sh -c 'tail -f logs/custom.log
```