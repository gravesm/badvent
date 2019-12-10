package main

import (
	"fmt"
	"github.com/gravesm/badvent/pkg/exercises"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "badvent",
		Usage: "run some bad code",
		Action: func(c *cli.Context) error {
			switch day := c.Args().Get(0); day {
			case "1a":
				fmt.Println(exercises.Day1a(os.Stdin))
			case "1b":
				fmt.Println(exercises.Day1b(os.Stdin))
			case "2a":
				fmt.Println(exercises.Day2a(os.Stdin))
			case "2b":
				fmt.Println(exercises.Day2b(os.Stdin))
			case "3a":
				fmt.Println(exercises.Day3a(os.Stdin))
			case "3b":
				fmt.Println(exercises.Day3b(os.Stdin))
			case "4a":
				fmt.Println(exercises.Day4a(os.Stdin))
			case "4b":
				fmt.Println(exercises.Day4b(os.Stdin))
			case "5a":
				fmt.Println(exercises.Day5a(os.Stdin))
			default:
				cli.ShowAppHelp(c)
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
