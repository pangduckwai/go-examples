package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	//////////////
	// Greetings
	//////////////
	// log.SetPrefix("greetings: ")
	// log.SetFlags(0)

	// message, err := greetings.Hellos(os.Args[1:])
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(message)

	//////////
	// Mongo
	//////////
	// tree := mongo.Build()
	// err := mongo.Write(tree)
	// // list, err := mongo.Split(tree, 3)

	// // tree, err := mongo.Read()
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }

	// // mongo.Change(tree, 3)
	// mongo.Fill(tree)

	// err = mongo.Write(tree)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(tree.Tree(10))
	// // fmt.Println("Is the same:", mongo.Same(tree, mongo.Build()))

	// // for _, frag := range list {
	// // 	fmt.Println(frag.Tree(3))
	// // }

	// Equation
	expands := []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	for _, v := range expands {
		x := int(v)
		fmt.Printf("%2v - %4v\n", x, (78/(x+1))-3) // 100000/2(3x + 1)
		// fmt.Printf("%2v - %4v\n", x, (x*x)/-5-x-x+76)
	}
}
