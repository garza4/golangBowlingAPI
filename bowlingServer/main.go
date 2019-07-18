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
func (s *BowlingServiceServer) Bowl(ctx context.Context, throw *pb.Throw) (*pb.Throw,error){
	
	return throw,nil

}

//func (s *BowlingServiceServer) DisplayScore(ctx context.Context,throw *pb.Throw) (*pb.Score,error){
//	scoreArray := throw.Pins
//	*pb.Score score = 0
//	for i, boolean := range scoreArray{
//		if boolean{
//			score += i
//
//		}
//	
//	}
//	return throw,nil

//}







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
