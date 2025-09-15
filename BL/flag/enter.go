package flag

import (
	"blog-go/global"
	"os"

	"github.com/urfave/cli"
	"go.uber.org/zap"
)

var (
	sqlFlags = &cli.BoolFlag{
		Name:  "sql",
		Usage: "Initalizes the structure of the MySQL DB table",
	}
)

func Run(c *cli.Context) {
	if c.NumFlags() > 1 {
		err := cli.NewExitError("Only one command can be specified", -4)
		global.Log.Error("Invalid command usage: ", zap.Error(err))
		os.Exit(-4)
	}
	switch {
	case c.Bool(sqlFlags.Name):
		if err := SQL(); err != nil {
			global.Log.Error("Failed to craete table structure:", zap.Error(err))
		} else {
			global.Log.Info("Successfully created table structure")
		}
	default:
		err := cli.NewExitError("Unknown command", -5)
		global.Log.Error(err.Error(), zap.Error(err))
	}
}

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Go Blog"
	app.Flags = []cli.Flag{
		sqlFlags,
	}
	app.Action = Run
	return app
}

func InitFlag() {
	if len(os.Args) <= 1 {
		return
	}
	app := NewApp()
	err := app.Run(os.Args)
	if err != nil {
		global.Log.Error("Application execution encountered an error:", zap.Error(err))
		os.Exit(-3)
	}
	os.Exit(0)
}
