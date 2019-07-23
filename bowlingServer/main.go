package main
import( 
	//"encoding/json"
	//"flag"
	"log"
	//"io/ioutil"
	"fmt"
	"context"
	//"math"
	"net"
	"math/rand"

	"google.golang.org/grpc"
	//"github.com/golang/protobuf/proto"
	pb "golangBowlingAPI/golangBowlingAPI"


)

//var port = flag.Int("port",10000, "The Server Port")



type BowlingServiceServer struct{}

type PlayerScoreStack struct{
	history string
	score int32
}


var strike int32 = 10
var spare int32 = 10
//throw the bowling ball and get the result from the throw...
//pb is the generated protoc file. 
var player = new(PlayerScoreStack)

func (s *BowlingServiceServer) Bowl(ctx context.Context, throw *pb.Throw) (*pb.PlayerScoreStack,error){//also return the struct for history
	//player := new(PlayerScoreStack)
	//totalPins := 0
	throw.Pins = rand.Int31n(11) //number of pins hit
	fmt.Println(throw.GetPins(), " were hit")
	if throw.GetPins() == 0{
		fmt.Println("No pins hit")
	}
	
	if(throw.GetPins() == strike){
		fmt.Println("STRIKE")
		player.history = "STRIKE"//append(player.history, "STRIKE")
		player.score += strike
	//elif second turn and totalPins=10: player.history = append(player.history("SPARE"))
	}else if(throw.GetPins() + rand.Int31n(11 - throw.GetPins()) == spare){
		 fmt.Println("SPARE")
		 player.history = "SPARE"//append(player.history, "SPARE")
		 player.score += spare

	}else{
		//player.history = append(player.history, "")
		player.score += throw.GetPins()
	}
	player.history = ""
	fmt.Println(player.history)



	return &pb.PlayerScoreStack{History: player.history, Score: player.score},nil

}









func main(){
	address := "localhost:50051"
	lis,err := net.Listen("tcp",address)
	if err != nil {
		log.Fatalln("Failed to Listen: %v", err)
	}
	bServer := &BowlingServiceServer{}
	grpcServer := grpc.NewServer()
	pb.RegisterBowlingServiceServer(grpcServer,bServer)
	if err := grpcServer.Serve(lis);err != nil{
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Printf("Hellcooper \n")
}
