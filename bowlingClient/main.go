package main
import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "golangBowlingAPI/golangBowlingAPI"
)

func main(){
	//score := 0
	conn,err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err != nil{
		fmt.Println(err)
	}
	defer conn.Close()
	client := pb.NewBowlingServiceClient(conn)
	//play,err := client.Bowl(context.Background(),&pb.Throw{Pins: 0}) //play bowling
	for i := 0; i < 10; i++{
		play,err := client.Bowl(context.Background(),&pb.Throw{Pins:0})
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(i, " ", play)
	}
	//if err != nil {
	//fmt.Println(err)
	//}
	//fmt.Println(play)
	fmt.Println("Game Over: \n", "Score was ", "____")


}


