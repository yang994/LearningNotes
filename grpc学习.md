# Hello gRPC!

## RPC
  * 远程过程调用，通俗理解：远程计算机提供一个服务，本地可直接调用其中的方法，而不需要知道其中网络的细节。

## 源码分析：

  `proto文件`
  ```
  syntax = "proto3";  //proto版本
  package helloworld; //包名

  service Greeter {
    // 定义一个Greeter服务,其中API为SayHello
    //过程是传入请求信息，返回相应信息
    rpc SayHello (HelloRequest) returns (HelloResponse) {}
  }

  // 请求信息的结构体
  message HelloRequest {
    string name = 1;
  }
  // 响应信息的结构体
  message HelloResponse {
    string message = 1;
  }
  ```
  `客户端文件`
  ```
  package main    //属于main包
  import (  
    "log"
    "os"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    pb "grpc/helloworld"  //这是通过之前的proto文件生成的包
  )

  const (
    address     = "localhost:50051"   //本地客户端的端口
    defaultName = "world"
  )

  func main() {
    // 本地节点与服务器建立连接的过程
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    //传入地址,拨号操作，得到一个grpc连接实例conn
    if err != nil {
      log.Fatalf("连接错误: %v", err)
    }
    defer conn.Close()
    c := pb.NewGreeterClient(conn)
    //根据连接实例，得到服务实例c

    // 请求服务器并打印响应
    name := defaultName
    if len(os.Args) > 1 {
      name = os.Args[1]
    }
    //os.Args 是取输入命令行中的参数，
    //如： run go abc.go 123  os.Args[0] = abc.go  os.Args[1] = 123

    r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
    //调用服务中的SayHello方法，传入ctx和请求信息结构体
    if err != nil {
      log.Fatalf("请求错误: %v", err)
    }
    log.Printf("服务器响应: %s", r.Message)

  }
  ```
  `服务器文件`
  ```
  package main
  import (
    "log"
    "net"
    "fmt"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    pb "grpc/helloworld"
  )
  const (
    port = ":50051"
  )

  // 实现 helloworld.GreeterServer 接口
  type server struct{}

  func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
    //这是proto文件 Greeter服务中要求提供的调用函数（接口） 收到请求信息，返回响应信息
    fmt.Println("客户端请求: "+in.Name)
    return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
  }


  func main() {
    lis, err := net.Listen("tcp", port)
    if err != nil {
      log.Fatalf("监听错误: %v", err)
    }//建立一个监听prot端口的tcp信息的协议，与客户端dial（拨号操作对应）

    s := grpc.NewServer() //建立一个grpc服务器
    pb.RegisterGreeterServer(s, &server{})  //将之前定义的server结构，在grpc服务器上登录，以便调用

    // 在grpc服务器上注册反射服务
    reflection.Register(s)
    if err := s.Serve(lis); err != nil {
      log.Fatalf("服务错误: %v", err)
    }
    //s.Server（s）  传入建立的监听器，队监听器传来的信息进行服务
  }
  ```
