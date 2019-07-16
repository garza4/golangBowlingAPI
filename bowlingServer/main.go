package main
import "fmt"


type bowlingAPIServer struct{

}

type Game struct{
	numPlayers int
	playerScore []int
	playerTurn []bool

}

func startGame(g Game){
	numPlayers := getNumPlayers(g)
	for i, num := range g.playerScore{
		g.playerScore[i] = 0
	}

}


//throw the bowling ball
func (s *bowlingAPIServer) Throw(ctx context.Context, throw *pb.Throw) (pb.Feature,error){


}


//get the score after a throw
func (s *bowlingAPIServer) Score(ctx context.Context, score *pb.Score) (pb.Feature,error){


}





func main(){
	fmt.Printf("Hellcooper \n")
}
