package main
import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	st "github.com/golang-collections/collections/stack"
	pb "golangBowlingAPI/golangBowlingAPI"
)

func main(){
	conn,err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err != nil{
		fmt.Println(err)
	}
	defer conn.Close()
	client := pb.NewBowlingServiceClient(conn)
	//play,err := client.Bowl(context.Background(),&pb.Throw{Pins: 0}) //play bowling
	var scoreStack []int32 //array of type struct PlayerScoreStack
	//fmt.Println(pb.PlayerScoreStack, " scoreStack")
	var indexHistory = 0
	var strike = st.New()
	var spare = st.New()
	var score int32 = 0
	for i := 0; i < 10; i++{
		play,err := client.Bowl(context.Background(),&pb.Throw{Pins:0})
		scoreStack = append(scoreStack,play.Result)
		//fmt.Println(play.player)
		if strike.Pop() != nil && strike.Pop() == true{
			scoreStack[indexHistory] += play.Result
		}

		if spare.Pop() != nil && spare.Pop() == true{
			scoreStack[indexHistory] += play.Result
		}

		if(pb.PlayerScoreStack{}.History == "STRIKE"){
			strike.Push(true)
			strike.Push(true)
			indexHistory = i
		}else if(pb.PlayerScoreStack{}.History == "SPARE"){
			spare.Push(true)
		}else{
			continue
		}

		if err != nil{
			fmt.Println(err)
		}
		//fmt.Println(i, " ", play)
	}

	for _,element := range scoreStack{
		score += element
	}
	
	fmt.Println("Game Over: \n", "Score was ", score)


}


