package main

import (  
	"bufio"
	"log"
	"os" 	
	"github.com/kimbbakar/Chat-App-With-RabbitMQ/send"
	"github.com/kimbbakar/Chat-App-With-RabbitMQ/receive"
)
 

func main() { 

	log.Print("What is your user Name: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()  
	userid := scanner.Text()


	for {
		log.Print("Want to 1. Receive or 2. Send ")

		scanner.Scan()  
		option := scanner.Text()	
	
		
		if option=="1"{
			receive.Receive(userid)
			break
		} else if option=="2" {
			send.Send(userid)		  
			break
		}
	}
 
}