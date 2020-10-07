# go-grpc-test


https://github.com/grpc/grpc-go

go get -u google.golang.org/grpc

Unary: small requests, small response eg. get user
Client Streaming: big request, small response eg. file upload
Server streaming: small request, big response eg. get products
BiDi streaming: big request, big response eg. chat