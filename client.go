package main
import("net";"fmt";"time";"strconv")

func main(){
	//Connect udp
	conn, err := net.Dial("udp", "127.0.0.1:8888")
	
	if err != nil {
		fmt.Println("Error")
	}
	
	defer conn.Close()

	//make storage arrays
	buffer_1 := make([]byte, 1024)
	buffer_2 := make([]byte, 1024)
	

	//record one timestamp before sending packet to server and one after receiving corresponding packet from server
	timestamp_1:=time.Now().UnixNano()
	conn.Write([]byte("message sent to server"))
	x,_:= conn.Read(buffer_1) // x is the number of bytes received by buffer_1
	timestamp_4:=time.Now().UnixNano()
	
	
	//convert received timestamps to int64 integers and print all 4 timestamps
	buffer_1=buffer_1[0:x]
	y,_:= conn.Read(buffer_2) // y is the number of bytes received by buffer_2
	buffer_2=buffer_2[0:y]
	timestamp_2,_:=strconv.Atoi(string(buffer_1))
	timestamp_3,_:=strconv.Atoi(string(buffer_2))
	fmt.Println(timestamp_1)
	fmt.Println(timestamp_2)
	fmt.Println(timestamp_3)
	fmt.Println(timestamp_4)
}