# golang Bowling Using gRPC
This code is an introduction to gRPC. This simple program uses golang and gRPC to simulate bowling using pseduorandom number generation and transmits data between the client and the server.
## 
To run the code you have to have grpc installed and protobuf. Here is a link to the grpc Golang page if curious about any grpc/golang details: https://grpc.io/docs/tutorials/basic/go/

Once installed run the commands
go run bowlingServer/main.go and go run bowlingClient/main.go
from the directory the repo was installed under and a game of bowling will commence!

Make sure $GOPATH is set and grpc is installed under the go/bin folder or the code will not work.
Used Go version 1.10.4 to write this code. 
