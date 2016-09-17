package main

import (
	"flag"
	"log"
	"time"
)

type MSG struct {
	Q chan interface{}
	P interface{}
}

//Client
func Ping(IN chan interface{}, PONG chan MSG) {
	var count uint64

	for {

		PONG <- MSG{IN, "Ping"}
		select {
		case <-IN:
			count++
		}
	}
}

//Server
func Pong(IN chan MSG) {

	var count uint64

	for {
		select {
		case msg := <-IN:
			if val, found := msg.P.(string); found == true {
				switch val {

				case "Ping":
					msg.Q <- MSG{nil, "Pong"}
					count++

				case "finished":
					log.Printf("Pong Count %d", count)
					return

				}
			}

		}
	}

}

func start_pings(PongChan chan MSG, Nping int) {

	for i := 0; i < Nping; i++ {
		PingsChan := make(chan interface{})
		go Ping(PingsChan, PongChan)
	}
}


func main() {

	Nsec := flag.Int("NSec", 5, "Number of seconds this program to run")
	Npings := flag.Int("Pings", 1, "Number of ping go-routine to run")
	
	flag.Parse()

	Pongchan := make(chan MSG, *Npings)

	go Pong(Pongchan)

	start_pings(Pongchan, *Npings)

	<-time.After(time.Second * time.Duration(*Nsec))

	Pongchan <- MSG{nil, "finished"}

	time.Sleep(1 * time.Second)

}
