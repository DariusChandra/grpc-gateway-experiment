read https://github.com/grpc-ecosystem/grpc-gateway for more info

Generate stubs using
```
protoc -I . \
    --go_out ./gen/go/ --go_opt paths=source_relative \
    --go-grpc_out ./gen/go/ --go-grpc_opt paths=source_relative \
    service/v1/your_service.proto
```

Generate reverse-proxy using
```
protoc -I . --grpc-gateway_out ./gen/go \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    service/v1/your_service.proto


protoc -I . --grpc-gateway_out ./gen/go \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    service/v1/your_service.proto
```

Generate using buf
1. at service directory, do `buf mod update` `buf build`
2. at root directory, do `buf generate service`