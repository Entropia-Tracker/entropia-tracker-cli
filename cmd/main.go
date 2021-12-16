package main

import (
  "github.com/alecthomas/kong"
)

var Version string
var absoluteLogPath string

type Context struct {
  Debug bool
}

var cli struct {
  Debug   bool       `help:"Enable debug logs."`
  Parse   ParseCmd   `cmd:"" help:"Parse log file."`
  Version VersionCmd `cmd:"" help:"Print the version."`
}

func main() {
  ctx := kong.Parse(&cli)
  err := ctx.Run(&Context{Debug: cli.Debug})
  ctx.FatalIfErrorf(err)
}
