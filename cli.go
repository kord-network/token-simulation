package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kord-network/token-simulation/simulation"
)

func main() {
	nftDisputesCommand := flag.NewFlagSet("nft-disputes", flag.ExitOnError)
	agents := nftDisputesCommand.Int64("agents", 0, "Number of agents")

	if len(os.Args) < 2 {
		fmt.Println("simulation argument is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "nft-disputes":
		nftDisputesCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if nftDisputesCommand.Parsed() {
		simulation.Simulation(*agents)
	}
}
