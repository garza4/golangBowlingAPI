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
	//"sync"
	//"time"

	"google.golang.org/grpc"
	//"github.com/golang/protobuf/proto"
	pb "golangBowlingAPI/golangBowlingAPI"


)

//var port = flag.Int("port",10000, "The Server Port")



type BowlingServiceServer struct{}





//throw the bowling ball and get the result from the throw...
//pb is the generated protoc file. 
func (s *BowlingServiceServer) Bowl(ctx context.Context, throw *pb.Throw) (*pb.Score,error){
	score := throw.GetScore()

	return pb.GetScore(),nil

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


	fmt.Printf("Hellcooper \n")
}
