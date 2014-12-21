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

type Fooer struct{}
func (s *Fooer) Visit(node ast.Node) ast.Visitor {
	print("Oh such a beautiful node\n")
	return s
}

func mainAction(c *cli.Context) {

	codefile := c.GlobalString("f")

	fsetc := token.NewFileSet()
	fc, _ := parser.ParseFile(fsetc, codefile, nil, 0)

	fooer := Fooer{}

	for _, s := range fc.Decls {
		ast.Walk(&fooer, s)
	}

}

func main() {

	// global level flags
	flagz := []cli.Flag{
		cli.StringFlag{
			Name:  "f",
			Usage: "File",
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
	}

	app := cli.NewApp()
	app.Flags = flagz
	app.Commands = cmdz
	app.Usage = "infer types for autotype slices."
	app.Version = "0.0.0"
	app.Action = missingAction

	app.Run(os.Args)
}
