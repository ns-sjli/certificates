package main

import (
	"github.com/mvmaasakkers/certificates/commands"
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Certificates"
	app.Version = "v0.0.1-alpha1"
	app.Commands = []cli.Command{
		commands.CertificateCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}