package publisher

import (
	"encoding/json"
	"github.com/jjeffery/stomp"
        "fmt"
        "bufio"
        "os"
)

func SendData() {
        //establishing the connection between activeMQ server & publisher
	conn, err := stomp.Dial("tcp", "localhost:61613")
	if err != nil {
		fmt.Println(err)
	}
      
        //getting input from user
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Enter text: ")
        input, _ := reader.ReadString('\n')
       
 
    
    
	b, _ := json.Marshal(input)
	err = conn.Send(
		"/queue/myqueue1", // destination
		"text/plain",      // content-type
		b)                 // body    
	if err != nil {
		fmt.Println(err)
	}
	
  conn.Disconnect()
}

