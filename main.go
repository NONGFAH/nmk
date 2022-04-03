package main

import (
	"context"
	"fmt"
	grpc2 "github.com/nongfah/nmk/grpc"
	"github.com/nongfah/nmk/system"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	if len(os.Args) > 1 {
		client()
	} else {
		server()
	}
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
}
func client() {

	host := os.Args[1]
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := grpc2.NewNMkGrpcClient(conn)
	control, _ := client.Control(context.Background())
	fmt.Println("连接成功")
	go func() {
		system.KeyboardMonitor(func(event system.Event) {
			control.Send(&grpc2.EventStreamRequest{
				Type:  grpc2.Type(event.Type),
				Code:  int32(event.Code),
				Value: int32(event.Value),
			})
		})
	}()
	go func() {
		system.MouseMonitor(func(event system.Event) {
			control.Send(&grpc2.EventStreamRequest{
				Type:  grpc2.Type(event.Type),
				Code:  int32(event.Code),
				Value: int32(event.Value),
			})
		})
	}()
}

func server() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Printf("Cannot Listen : %v\n", err)
		return
	}
	s := grpc.NewServer()
	serv := NMkGrpcServer{}
	grpc2.RegisterNMkGrpcServer(s, &serv)
	err = s.Serve(listener)
	fmt.Println("启动成功")
	if err != nil {
		fmt.Printf("Cannot Serve : %v\n", err)
		return
	}
}

type NMkGrpcServer struct {
}

func (N NMkGrpcServer) Control(controlServer grpc2.NMkGrpc_ControlServer) error {
	var x, y int32 = 0, 0
	for {
		select {
		case <-controlServer.Context().Done():
			fmt.Print("grpc 客户端掉线")
			return nil
		default:
			request, err := controlServer.Recv()
			if err != nil {
				return err
			}
			code, value := request.Code, request.Value
			if request.Type == grpc2.Type_EvMouseMove {
				code, value = request.Code-x, request.Value-y
				x, y = request.Code, request.Value
				if code > 100 || code < -100 {
					code = 0
				}
				if value > 100 || value < -100 {
					value = 0
				}
			}
			event := system.Event{
				Type:  system.Type(request.Type),
				Code:  system.Code(code),
				Value: system.Value(value),
			}
			err = system.Input(event)
			if err != nil {
				return err
			}
		}
	}
}
