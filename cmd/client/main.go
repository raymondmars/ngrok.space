package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/inconshreveable/mousetrap"
	"github.com/ngrok-space/internal/app/client"
	"github.com/ngrok-space/internal/app/client/localserver"
	"github.com/ngrok-space/internal/pkg/log"
	"github.com/ngrok-space/internal/pkg/util"
)

func init() {
	if runtime.GOOS == "windows" {
		if mousetrap.StartedByExplorer() {
			fmt.Println("Don't double-click ngrok!")
			fmt.Println("You need to open cmd.exe and run it from the command line!")
			time.Sleep(5 * time.Second)
			os.Exit(1)
		}
	}
}

func main() {
	if WithLocalServer == "1" {
		go localserver.StartServer(fmt.Sprintf("https://%s.%s", DefaultSubDomain, DefaultDomain))
	}

	// parse options
	opts, err := ParseArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// set up logging
	log.LogTo(opts.Logto, opts.Loglevel)

	// read configuration file
	config, err := client.LoadConfiguration(opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// seed random number generator
	seed, err := util.RandomSeed()
	if err != nil {
		fmt.Printf("Couldn't securely seed the random number generator!")
		os.Exit(1)
	}
	rand.Seed(seed)

	client.NewController().Run(config)
}
