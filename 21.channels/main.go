package main

import (
	"fmt"
	// "math/rand"
	"time"
)

// func promesj(input chan int , output chan int) {

// 	for num := range input {
// 		fmt.Println("processing number is ", num)
// 		time.Sleep(1 * time.Second)
// 		output <- num
// 	}
// }

func sendEmail(emailchan chan string, done chan bool) {
	defer func() {
		done <- true
	}()

	for email := range emailchan {
		fmt.Println("sending email to ", email)
		time.Sleep(100 * time.Millisecond)
	}

}

func main() {

	// message := make(chan int) // buffered channel

	// result := make(chan int)

	// go  promesj(message, result)

	// for {

	// 	message <- rand.Intn(100)
	// 	resultValue := <-result
	// 	fmt.Println("result is ", resultValue)
	// }

	// emailchan := make(chan struct {
	// 	name      string
	// 	subject   string
	// 	recipient string
	// }, 100) // unbuffered channel

	// emailchan <- struct{ name, subject, recipient string }{name: "Alice", subject: "Hello", recipient: "Bob"}
	// emailchan <- struct{ name, subject, recipient string }{name: "Charlie", subject: "Hi", recipient: "David"}

	// for i := 0; i < 2; i++ {
	// 	email := <-emailchan
	// 	fmt.Println("name ", email.name)
	// 	fmt.Println("subject ", email.subject)
	// 	fmt.Println("recipient ", email.recipient)
	// 	fmt.Println("-------------")
	// }

	emailchan := make(chan string, 100) // unbuffered channel

	done := make(chan bool)

	go sendEmail(emailchan, done)

	for i:= 0 ; i > -1 ; i++ {

		emailchan <- fmt.Sprintf("%d@gmail.com", i)
	}

	fmt.Println("All emails sent")

	close(emailchan)
	<-done

}
