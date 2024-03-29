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
		play,err := client.Bowl(context.Background(),&pb.Throw{})
		scoreStack = append(scoreStack,play.Score)
		fmt.Println("Score ", play.Score, " at index ", i)
		if(strike.Pop() == true){
			fmt.Println("STRIKE!! update old score of" , scoreStack[indexHistory])
			scoreStack[indexHistory] += play.GetScore()
			fmt.Println("UPDATE ", scoreStack[indexHistory], " at index ", indexHistory)
		}else if(spare.Pop() == true){
			fmt.Println("was a spare, update old score of ", scoreStack[indexHistory], " at index ", indexHistory)
			scoreStack[indexHistory] += play.GetScore()
			fmt.Println("UPDATE for SPARE! ", scoreStack[indexHistory], " at index ", indexHistory)
		}else{

			scoreStack[i] = play.GetScore()
		
		}
		if(play.History == "STRIKE"){
			strike.Push(true)
			strike.Push(true)
			indexHistory = i
			//fmt.Println("save index ", indexHistory, " STRIKE")

		}else if(play.History == "SPARE"){
			spare.Push(true)
			indexHistory = i
		}else{
			continue
		}

		if err != nil{
			fmt.Println(err)
		}
		//fmt.Println(scoreStack)
	}
	fmt.Println(scoreStack)
	for _,element := range scoreStack{
		score += element
	}
	
	fmt.Println("Game Over: \n", "Score was ", score)


}


