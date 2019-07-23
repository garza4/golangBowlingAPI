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
		scoreStack = append(scoreStack,play.Score)
		fmt.Println(play)
		if strike.Pop() == true{
			scoreStack[indexHistory] += play.Score
		}

		if spare.Pop() == true{
			fmt.Println("was a spare, update old score ", scoreStack[indexHistory])
			scoreStack[indexHistory] += play.Score
			fmt.Println("Good Spare! ", scoreStack[indexHistory])
		}
		//fmt.Println(play.History, " looking for spares and strikes")
		if(play.History == "STRIKE"){
			strike.Push(true)
			strike.Push(true)
			indexHistory = i

		}else if(play.History == "SPARE"){
			spare.Push(true)
			indexHistory = i
		}else{
			continue
		}

		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(play, " ", i)
	}

	for _,element := range scoreStack{
		score += element
	}
	
	fmt.Println("Game Over: \n", "Score was ", score)


}


