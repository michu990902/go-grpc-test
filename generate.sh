protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.

protoc blog/blogpb/blog.proto --go_out=plugins=grpc:.



protoc blog/blogpb/blog.proto --js_out=import_style=commonjs,binary:web_client --grpc-web_out=import_style=commonjs,mode=grpcwebtext:web_client