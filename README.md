# go-grpc-test


https://github.com/grpc/grpc-go

go get -u google.golang.org/grpc

Unary: small requests, small response eg. get user
Client Streaming: big request, small response eg. file upload
Server streaming: small request, big response eg. get products
BiDi streaming: big request, big response eg. chat

errors:
    http://avi.im/grpc-errors/
    https://grpc.io/docs/guides/error/

deadlines:
    always use!!!
    https://grpc.io/blog/deadlines/
    good timings: 
        100ms-1s for small api
        for long api call: 5 min


!!!SSL:
    https://www.youtube.com/watch?v=jmqLJMFS_yI&feature=emb_logo
    https://github.com/techschool/pcbook-go
    cd ssl; ./gen.sh

TODO:
    makefile
    https://tutorialedge.net/golang/makefiles-for-go-developers/
    
    evanscli