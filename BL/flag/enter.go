package flag

import (
	"blog-go/global"
	"errors"
	"os"

	"github.com/urfave/cli"
	"go.uber.org/zap"
)

var (
	sqlFlags = &cli.BoolFlag{
		Name:  "sql",
		Usage: "Initalizes the structure of the MySQL DB table",
	}
	sqlExportFlag = &cli.BoolFlag{
		Name:  "sql-export",
		Usage: "Exports SQL data to a specified file.",
	}
	sqlImportFlag = &cli.StringFlag{
		Name:  "sql-import",
		Usage: "Imports SQL data from a specified file.",
	}
	esFlag = &cli.BoolFlag{
		Name:  "es",
		Usage: "Initializes the Elasticsearch index.",
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
	case c.Bool(sqlExportFlag.Name):
		if err := SQLExport(); err != nil {
			global.Log.Error("Failed to export SQL data:", zap.Error(err))
		} else {
			global.Log.Info("Successfully exported SQL data")
		}
	case c.IsSet(sqlImportFlag.Name):
		if errs := SQLImport(c.String(sqlImportFlag.Name)); len(errs) > 0 {
			var combinedErrors string
			for _, err := range errs {
				combinedErrors += err.Error() + "\n"
			}
			err := errors.New(combinedErrors)
			global.Log.Error("Failed to import SQL data:", zap.Error(err))
		} else {
			global.Log.Info("Successfully imported SQL data")
		}
	case c.Bool(esFlag.Name):
		if err := Elasticsearch(); err != nil {
			global.Log.Error("Failed to create ES indices:", zap.Error(err))
		} else {
			global.Log.Info("Successfully created ES indices")
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
		sqlExportFlag,
		sqlImportFlag,
		esFlag,
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
