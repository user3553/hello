package main

import (
	"os"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "hello"
	app.Usage = "fight the loneliness!"

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "lang, l",
			Value: "english",
			Usage: "language for the greeting",
			EnvVar: "APP_LANG",
		},
	}
	app.Action = func(c *cli.Context) {
		name := "someone"
		if len(c.Args()) > 0 {
			name = c.Args()[0]
		}
		if c.String("lang") == "spanish" {
			println("Hola", name)
		} else {
			println("Hello", name)
		}
	}
	app.Commands = []cli.Command{
 		{
    			Name:      "add",
    			ShortName: "a",
    			Usage:     "add a task to the list",
    			Action: func(c *cli.Context) {
      				println("added task: ", c.Args().First())
    			},
  		},
  		{
    			Name:      "complete",
    			ShortName: "c",
    			Usage:     "complete a task on the list",
    			Action: func(c *cli.Context) {
      				println("completed task: ", c.Args().First())
    			},
  		},
  		{
    			Name:      "template",
    			ShortName: "r",
    			Usage:     "options for task templates",
    			Subcommands: []cli.Command{
      			{
        			Name:  "add",
  			      	Usage: "add a new template",
        			Action: func(c *cli.Context) {
            				println("new task template: ", c.Args().First())
        			},
      			},
      			{
        			Name:  "remove",
        			Usage: "remove an existing template",
        			Action: func(c *cli.Context) {
          				println("removed task template: ", c.Args().First())
        			},
      			},
    			},
  		},     
	}
	app.Run(os.Args)
}
