package main
import (
	"fmt"
	"google.golang.org/grpc"
	pb "golangBowlingAPI/golangBowlingAPI"
)

func main(){
	conn,err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err != nil{
		fmt.println(err)
	}
	defer conn.Close()
	client := pb.NewBowlingServiceClient(conn)
	play,err := client.Bowl(context.Background(),&pb.Throw{Score: 3}) //play bowling? 
	if err != nil {
	fmt.Println(err)
	}
	fmt.Println(play)


}


