package main
import("net";"log";"fmt";"time";"strconv")

func main(){
	// listen to incoming udp packets
	pc, err := net.ListenPacket("udp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	//create timestamp when packet is received from client
	buffer := make([]byte, 1024)
	_,y,_ := pc.ReadFrom(buffer) //y is the address of the client
	timestamp_2:=time.Now().UnixNano()
	fmt.Println("message received from client")
	

	//create another timestamp before sending packet to client and send both timestamps to client
	timestamp_3:=time.Now().UnixNano()
	pc.WriteTo([]byte(strconv.FormatInt(timestamp_2,10)), y)
	pc.WriteTo([]byte(strconv.FormatInt(timestamp_3,10)), y)
}