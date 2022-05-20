package main

import (
	"fmt"
	"log"
	"flag"
)

func main() {
	//create a flag to get config file name
	config := flag.String("config","", "configuration file depending on stages , refer README file")
	flag.Parse()
	if(flag.NFlag !=1 ) {
		log.println("error with config file")
	}
	fmt.Println(config)
}