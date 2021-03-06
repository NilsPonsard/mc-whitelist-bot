package main

import (
	"os"
	"path"

	cli "github.com/jawher/mow.cli"
	"github.com/nilsponsard/mc-whitelist-bot/internal/commands"
	"github.com/nilsponsard/mc-whitelist-bot/internal/config"
	"github.com/nilsponsard/mc-whitelist-bot/pkg/files"
	"github.com/nilsponsard/mc-whitelist-bot/pkg/verbosity"
)

// Version will be set by the script build.sh
var version string

func main() {

	app := cli.App("mc-whitelist-bot", "starter project")
	app.Version("v version", version)

	defaultPath := files.ParsePath("~/.mc-whitelist-bot/")

	// arguments

	var (
		verbose     = app.BoolOpt("d debug", false, "Debug mode, more verbose operations")
		appPath     = app.StringOpt("a app-folder", defaultPath, "Path to the app files")
		disableLogs = app.BoolOpt("disable-logs", false, "Disable the saving of logs")
	)

	// Executed befor the commands

	app.Before = func() {

		parsedConfigPath := *appPath
		files.EnsureFolder(parsedConfigPath)

		// create the folder for the logs

		files.EnsureFolder(path.Join(defaultPath, "test"))

		// Configure the logs

		verbosity.SetupLog(*verbose, path.Join(defaultPath, "logs.txt"))

		verbosity.SetLogging(!*disableLogs)

		_, err := config.LoadConfig(path.Join(*appPath, "config.json"))
		if err != nil {
			verbosity.Error("cannot load config : ", err)
		}

	}

	// set subcommands

	commands.SetupCommands(app)

	// parse the args

	app.Run(os.Args)
}
