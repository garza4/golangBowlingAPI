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
	history []string
	score int32
}


var strike int32 = 10
var spare int32 = 10
//throw the bowling ball and get the result from the throw...
//pb is the generated protoc file. 
func (s *BowlingServiceServer) Bowl(ctx context.Context, throw *pb.Throw) (*pb.Score,error){
	player := new(PlayerScoreStack)

	throw.Pins = rand.Int31n(10)
	if throw.GetPins() ==  strike{
		player.history = append(player.history, "STRIKE")
		player.score += strike
		//elif second turn and totalPins=10: player.history = append(player.history("SPARE"))
	}else{
		player.score += throw.GetPins()
	}
	fmt.Println(player.score)



	return &pb.Score{Result: player.score},nil

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
