package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func missingAction(c *cli.Context) {
	fmt.Fprintln(os.Stderr, "TODO :)")
}

func mainAction(c *cli.Context) {
	fmt.Fprintln(os.Stderr, "TODO :)")
}

func main() {

	// global level flags
	flagz := []cli.Flag{
		cli.StringFlag{
			Name:  "f",
			Usage: "Foo",
		},
		cli.StringFlag{
			Name:  "b",
			Usage: "Bar",
		},
	}

	// Commands
	cmdz := []cli.Command{
		{
			Name: "inferyes",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "f",
					Usage: "Foo.",
				},
			},
			Usage:  "Inferyes inference engine",
			Action: mainAction,
		},
		{
			Name: "separate",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "f",
					Usage: "Foo.",
				},
			},
			Usage:  "Separate splits and strips errors handling",
			Action: missingAction,
		},
		{
			Name: "merge",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "f",
					Usage: "Foo.",
				},
			},
			Usage:  "Merge handlers back",
			Action: missingAction,
		},
		{
			Name: "split",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "f",
					Usage: "Foo.",
				},
			},
			Usage:  "Split off the current error handling to an errfile",
			Action: missingAction,
		},
		{
			Name: "strip",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "f",
					Usage: "Foo.",
				},
			},
			Usage:  "Strip off the error handling from the client code",
			Action: missingAction,
		},
		{
			Name: "init",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "f",
					Usage: "Foo.",
				},
			},
			Usage:  "Create an empty errfile",
			Action: missingAction,
		},
		{
			Name: "ls",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "u",
					Usage: "Unhandled errors only.",
				},
			},
			Usage:  "List the handling of all errors",
			Action: missingAction,
		},
		{
			Name: "show",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "f",
					Usage: "Foo.",
				},
			},
			Usage:  "Show handling related to a function",
			Action: missingAction,
		},
		{
			Name: "status",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "f",
					Usage: "Foo.",
				},
			},
			Usage:  "Show cumulative statistics",
			Action: missingAction,
		},
	}

	app := cli.NewApp()
	app.Flags = flagz
	app.Commands = cmdz
	app.Usage = "strip / add error handling to a go file"
	app.Version = "0.0.0"
	app.Action = missingAction

	app.Run(os.Args)
}
