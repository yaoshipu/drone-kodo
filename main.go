package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

var build = "0" // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "KODO plugin"
	app.Usage = "KODO plugin"
	app.Action = run
	app.Version = fmt.Sprintf("1.0.%s", build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "host",
			Usage:  "KODO API host",
			Value:  "https://upload.qbox.me",
			EnvVar: "PLUGIN_HOST",
		},
		cli.StringFlag{
			Name:   "access-key",
			Usage:  "KODO access key",
			EnvVar: "PLUGIN_ACCESS_KEY",
		},
		cli.StringFlag{
			Name:   "secret-key",
			Usage:  "KODO secret key",
			EnvVar: "PLUGIN_SECRET_KEY",
		},
		cli.StringFlag{
			Name:   "bucket",
			Usage:  "KODO bucket",
			EnvVar: "PLUGIN_BUCKET",
		},
		cli.StringFlag{
			Name:   "key",
			Usage:  "upload file path",
			Value:  ".",
			EnvVar: "PLUGIN_KEY",
		},
		cli.StringFlag{
			Name:   "source",
			Usage:  "source file path",
			Value:  "/",
			EnvVar: "PLUGIN_SOURCE",
		},
		cli.BoolFlag{
			Name:   "delete",
			Usage:  "delete existing file",
			EnvVar: "PLUGIN_DELETE",
		},
		cli.StringFlag{
			Name:  "env-file",
			Usage: "source env file",
		},
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	if c.String("env-file") != "" {
		_ = godotenv.Load(c.String("env-file"))
	}

	plugin := Plugin{
		Endpoint: c.String("endpoint"),
		AK:       c.String("access-key"),
		SK:       c.String("secret-key"),
		Bucket:   c.String("bucket"),
		Region:   c.String("region"),
		Key:      c.String("key"),
		Source:   c.String("source"),
		Delete:   c.Bool("delete"),
	}

	return plugin.Exec()
}
