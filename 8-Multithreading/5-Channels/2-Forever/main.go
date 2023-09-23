package main

func main() {
	forever := make(chan bool)

	//// Solution
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(i)
	// 	}
	// 	forever <- true
	// }()

	<-forever
}
