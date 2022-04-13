package main

import (
  "fmt"
  "github.com/EntropiaTally/entropia-tally-cli/pkg/events"
  "github.com/EntropiaTally/entropia-tally-cli/pkg/util"
  "os"
)

type ParseCmd struct {
  All      bool   `short:"a" help:"Read the whole file including historical entries" default:"false"`
  Watch    bool   `short:"" help:"Keep reading file until closed." default:"false"`
  File     string `short:"f" help:"File to watch." type:"existingfile"`
  Location bool   `short:"l" help:"Attempt to get player location using OCR." default:"false"`
  Name     string `short:"n" help:"Player name, used to filter out globals and rare loots" default: ""`
}

func (p *ParseCmd) Run(ctx *Context) error {
  return util.ReadFile(p.File, p.All, p.Watch, func(row string) {

    // Parse row into *Event
    event, ok := events.Parse(row, p.Name)
    if !ok {
      return
    }

    // Serialize to JSON string
    serialized, ok := event.JSON()
    if !ok {
      return
    }

    // Print to Stdout
    fmt.Fprintln(os.Stdout, serialized)
  })
}
