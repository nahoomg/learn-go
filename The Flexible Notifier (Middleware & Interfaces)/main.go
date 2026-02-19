package main

import (
	"fmt"
)

type Notifier interface{
	Send(message string)
}

type Email struct{}

type Sms struct{}

func (m Email) Send(message string){
	fmt.Printf("Sending Email: %s\n", message)
}

func (m Sms) Send(message string){
	fmt.Printf("Sending Email: %s\n", message)
}

func NewAlertCounter() func() int{
	sum := 0
	return func() int {
		sum++
		fmt.Printf("Total alerts sent so far: %d\n", sum)
		return sum
	}
}


func main(){
	notifiers := []Notifier{
    Email{},
    Sms{},
	}
	conunter:= NewAlertCounter()
	for _, n:= range notifiers{
		n.Send("System Overload!")
		conunter()
	}

}
