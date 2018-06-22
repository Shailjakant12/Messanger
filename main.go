package main

import (
	"messanger/publisher"
	"messanger/subscriber"
	"fmt"
)

func main(){
	fmt.Println("Calling Publisher... sending msg")
	publisher.SendData()
	fmt.Println("Calling Subscriber...")
	subscriber.ReceiveData()
}