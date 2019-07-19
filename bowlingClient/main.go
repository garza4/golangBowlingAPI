package main
import (
	"fmt"
	"google.golang.org/grpc"
	pb "golangBowlingAPI/golangBowlingAPI"
)

func main(){
	conn,err := grpc.Dial("localhost:50051")
	if err != nil{
	
	}
	defer conn.Close()
	client := pb.NewBowlingAPIClient(conn)
	play,err := client.Bowl(context.Background(),&pb.Throw()) //play bowling? 
	fmt.printF("Hello Cooper")


}


