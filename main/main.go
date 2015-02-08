package main

import (
    "fmt"
    "TwitterShell/handler" //Sterilizer
    "TwitterShell/twilio" //Twilio
    "github.com/op/go-logging"
)

var log = logging.MustGetLogger("TShell")
var format = logging.MustStringFormatter(
	"%{color}%{time:15:04:05.000} %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}",
)


func main() {


    //initializes the twilio send/receive function
    //initializes the handler sterlizer function
    //initializes the process function
    fmt.Printf("Hello world\n")
    log.Info("Hello world")

    fakeData := Twilio.TwilData{PhoneNum: "555-555-5555", InMessage: "I like trains!"}

    fmt.Println(fakeData)
    
    hand := make(chan Twilio.TwilData, 5)
    demo := make(chan Twilio.TwilData, 5)

    demo <- fakeData

    go Sterilizer.Sterlhand(demo, hand)
   
    fmt.Println(<-hand)

    _, twil := Twilio.Initialize()
   twil.GetTexts();

}

