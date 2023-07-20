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
grpcurl -H "Authorization: Bearer V1PHUChS6im6fjgz4CCIehEyKKJoteShuS6LxXAqJ3A7izZCHRhlZxJax6kkboMq-8GaySE" -plaintext localhost:9090 CloudrollService.Ping
```

access_token:"V1PHUChS6im6fjgz4CCIehEyKKJoteShuS6LxXAqJ3A7izZCHRhlZxJax6kkboMq-8GaySE"

## Actions

```
grpcurl -H "Authorization: Bearer <TOKEN>" -plaintext -d '{"name":"testTeam"}' localhost:9000 CloudrollService.CreateTeam
```

```
grpcurl -H "Authorization: Bearer <TOKEN>" -plaintext -d '{"id":"876422190021083137"}' localhost:9000 CloudrollService.JoinTeam
```

```
grpcurl -H "Authorization: Bearer <TOKEN>" -plaintext -d '{"name":"task1", "description":"task description", "points":100}' localhost:9000 CloudrollService.AddTask
```



# 
taskName

points
timeOfSolve

# Radar
solveTime
hards
taskAmount