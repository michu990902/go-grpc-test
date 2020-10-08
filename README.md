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

ssl/tls
    https://grpc.io/docs/guides/auth/

    cert error:
        http://www.inanzzz.com/index.php/post/jo4y/using-tls-ssl-certificates-for-grpc-client-and-server-communications-in-golang-updated


# https://www.youtube.com/watch?v=jmqLJMFS_yI&feature=emb_logo
# https://www.youtube.com/watch?v=kVpB-uH6X-s
https://www.sohamkamani.com/golang/2019-01-01-jwt-authentication/
https://medium.com/@pliutau/getting-started-with-oauth2-in-go-2c9fae55d187