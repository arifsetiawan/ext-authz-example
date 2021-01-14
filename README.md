
# Ext Authz Example

Example implementation of auth with ext-authz

```
go run main.go
```

## Build docker

```
CGO_ENABLED=0 GOOS=linux go build -o ./bin/ext-authz
```

```
docker build -t ext-authz:0.2.0 .
```

```
docker run -p 8088:8088 --name ext-authz ext-authz:0.2.0
```

## Test with Envoy

Run backend service
```
docker run -p 8080:80 kennethreitz/httpbin
```

Run envoy
```
envoy --config-path ./envoy/base-auth.json
```

curl
```
curl http://localhost:10000/get -v
```

ext-authz should show log of authz request
