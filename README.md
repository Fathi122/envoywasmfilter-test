# envoy wasm filter test

Build envoy-proxy with http wasm filter and http server backend images
 ```
./build-images.sh
```

Start  http backend server and envoy-proxy
```
./start-containers.sh
```

Test exposed endpoint
```
curl localhost:8090/test200
```

And Observe logs on logs/custom.log from envoy container
```
ENVOY_CONTAINER_ID=$(docker ps --filter "name=envoy" --format={{.ID}}) && docker exec -it ${ENVOY_CONTAINER_ID} sh -c 'tail -f logs/custom.log'
```
