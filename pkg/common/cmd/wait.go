package cmd

import (
	"log"
	"net"
	"time"
)

func WaitForService(host string) {
	log.Println("Waiting for %s", host);

	for {
		log.Printf("Testing connection to %s", host);
		conn, err := net.Dial("tcp", host);

		if err == nil{
			_ = conn.Close();
			log.Printf("%s is up!!", host);
			return
		}

		time.Sleep(time.Millisecond * 500);
	}
}