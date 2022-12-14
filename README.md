# step-by-step-go-grpc

## Step 1. Download and set PATH protoc

[ここ](https://developers.google.com/protocol-buffers/docs/downloads) から `protoc` を DL してインストール！

### My Recommand

use `asdf` and install protoc.

```console
asdf plugin-add protoc https://github.com/paxosglobal/asdf-protoc.git
asdf install protoc latest
```

> **Warning**
> インストールできるバージョンが最新じゃない。(適当に PR 出そうと思って忘れてた)

## Step 2. Install Go protoco buffer pliugin

```console
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## Step 3. Define proto file

##### helloworld.proto

```proto
syntax = "proto3";

option go_package = "helloworld";

package helloworld;

message HelloRequest { string name = 1; }

message HelloReply { string message = 1; }

service Greeter {
  rpc SayHello(HelloRequest) returns(HelloReply) {}
}
```

## Step 4. Generate code from proto file

```console
mkdir helloworld

protoc --go_out=helloworld --go_opt=paths=source_relative --go-grpc_out=helloworld --go-grpc_opt=paths=source_relative -I=./ helloworld.proto
```

## Step 5. Initialize Go projects.

```console
go mod init github.com/ryutah/step-by-step-go-grpc
go mod tidy
```

## Step 6. Write code!

```txt
s *server github.com/ryutah/step-by-step-go-grpc/helloworld.GreeterServer
```

## Step 7. Run Server

```console
go run .
```

## Step 8. Option implements reflection API

[see](https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md#enable-server-reflection)

```diff
diff --git a/main.go b/main.go
index c1c49e9..66928c9 100644
--- a/main.go
+++ b/main.go
@@ -7,6 +7,7 @@ import (

        "github.com/ryutah/step-by-step-go-grpc/helloworld"
        "google.golang.org/grpc"
+       "google.golang.org/grpc/reflection"
 )

 // リクエストを受け取るサーバの実装をする
@@ -36,6 +37,8 @@ func main() {
        s := grpc.NewServer()
        helloworld.RegisterGreeterServer(s, &server{})

+       reflection.Register(s)
+
        if err := s.Serve(lis); err != nil {
                panic(err)
        }
```

## Step 9. Exec API

```console
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

```console
grpcurl -plaintext -d '{"name": "ryutah"}' 127.0.0.1:50001 helloworld.Greeter/SayHello
```
