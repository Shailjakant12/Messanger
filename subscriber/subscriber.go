package subscriber

import (
	"encoding/json"
	"fmt"
	"github.com/jjeffery/stomp"
)

func ReceiveData() {

	conn, err := stomp.Dial("tcp", "localhost:61613")
	if err != nil {
		fmt.Println(err)
	}
	sub, err := conn.Subscribe("/queue/myqueue1", stomp.AckClient)
	if err != nil {
		fmt.Println(err)
	}
	for {
		var actualmsg string
		m := <-sub.C
		err = json.Unmarshal(m.Body, &actualmsg)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(actualmsg)

		// acknowledge the message
		err = conn.Ack(m)
		if err != nil {
			fmt.Println(err)
		}
	}
}
