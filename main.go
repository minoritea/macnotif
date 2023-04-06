package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{}
	app.Usage = "MacOS notification generator"
	app.HideHelpCommand = true
	app.ArgsUsage = "message"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "title",
			Usage: "set notification title",
		},
	}

	app.Action = func(cc *cli.Context) error {
		message := cc.Args().First()
		if message == "" {
			return fmt.Errorf("message should not be blank")
		}
		script := fmt.Sprintf("display notification %q", message)
		title := cc.String("title")
		if title != "" {
			script += fmt.Sprintf(" with title %q", title)
		}
		return exec.Command("osascript", "-e", script).Run()
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
