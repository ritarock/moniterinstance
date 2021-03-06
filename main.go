package main

import (
	"fmt"
	"log"
	"moniterinstance/lib/action"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var profile string
	var term int

	app := cli.App{
		Name:  "moniterinstance",
		Usage: "moniter ec2 instance",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "profile",
				Usage:       "set aws profile",
				Value:       "default",
				Destination: &profile,
			},
			&cli.IntFlag{
				Name:        "term",
				Usage:       "set StartTime",
				Value:       60,
				Destination: &term,
			},
		},
		Action: func(c *cli.Context) error {
			var instanceName string
			if c.NArg() > 0 {
				instanceName = c.Args().Get(0)
				action.Run(profile, instanceName, term)
			} else {
				fmt.Println("Instance is not selected")
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
