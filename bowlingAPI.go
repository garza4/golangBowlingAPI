package main
import "fmt"

func main(){

}

func startGame(g Game){
	numPlayers := getNumPlayers(g)
	for i, num := range g.playerScore{
		g.playerScore[i] = 0
	} 
	
}



func getNumPlayers(g Game){
	return g.players
}

type Game struct{
	players int
	playerScore [players] int
	playerTurn [players] bool
	var m map[string]bool //map strings to boolean

}



