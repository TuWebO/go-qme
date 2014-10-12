package main

import (
	"fmt"	
	"github.com/TuWebO/go-qme/qs"
)

func main() {
	fmt.Printf("=======================================================\n")
	fmt.Printf("Welcome to Qme, best enqueue app. \nTake your turn now!\n")
	fmt.Printf("=======================================================\n\n\n\n")

	// Creates and draw a queue
	// q := qs.NewQueue(1,"\"La cola de Agustín\"")
	// q.DrawQueue()
	
	// Creates a slice of queues and draw them
	queues := []qs.Queue{qs.NewQueue(1,"Cola 1"), qs.NewQueue(2, "Cola 2")}
	fmt.Printf("Draw a list of current queues\n")
	fmt.Printf("=============================\n")
	for i := 0; i < len(queues); i++ {
		queues[i].DrawQueue()
		fmt.Printf("------------\n")
	}
	
	fmt.Printf("\n\n")
	
	// Implements Stringer
	fmt.Printf("Stringer implementation of current queues\n")
	fmt.Printf("=========================================\n")
	for i := 0; i < len(queues); i++ {
		fmt.Println(queues[i])
		fmt.Printf("------------\n")
	}
	
	fmt.Printf("\n\n")
	
	// Implements Set method
	fmt.Printf("Implements Set method\n")
	fmt.Printf("=====================\n")
	queues[0].Setid(5)
	queues[1].Setid(6)
	fmt.Println(queues)
}