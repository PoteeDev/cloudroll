# cloudroll

## generate proto
```
protoc --go-grpc_out=. --go_out=.  --grpc-gateway_out=. proto/cloudroll.proto
```

## Test
### curl
```
curl -H "Authorization: Bearer yolo" localhost:8080/v1/ping
```
### gcurl
```
grpcurl -H "Authorization: Bearer yolo" -plaintext localhost:9090 CloudrollService.Ping
```