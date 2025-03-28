package main

import (
	"flag"
	"fmt"
	"mygo/app/console/commands"
)

func main() {
	runUserSeed := flag.Bool("userseed", false, "Run User Seed script")
	runPermit := flag.Bool("permit", false, "Run Permit script")
	flag.Parse()

	if *runUserSeed {
		commands.RunUserSeed()
	} else if *runPermit {
		commands.RunPermit()
	} else {
		fmt.Println("Please provide a valid flag: -userseed or -permit")
	}
}
