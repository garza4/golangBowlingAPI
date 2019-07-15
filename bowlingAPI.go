package main
import "fmt"

func main(){
	fmt.Printf("Hellcooper \n")
}

func startGame(g Game){
	numPlayers := getNumPlayers(g)
	for i, num := range g.playerScore{
		g.playerScore[i] = 0
	}

}


func setNumPlayers(g Game, players int){
	g.numPlayers = players
}

func getNumPlayers(g Game) int {
	return g.numPlayers
}

var m map[string]bool
type Game struct{
	numPlayers int
	playerScore []int
	playerTurn []bool

}



