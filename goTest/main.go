package main

import (
	"fmt"
	// Which path to import.
	pb "goTest/proto/consignment/importMe.go"
)

func main() {
	pb.sayHello()
	fmt.Println("Yes")
}
