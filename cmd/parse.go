package main

import (
  "fmt"
  "github.com/EntropiaTally/entropia-tally-cli/internal/misc"
  "github.com/EntropiaTally/entropia-tally-cli/internal/watcher"
  "github.com/EntropiaTally/entropia-tally-cli/parser"
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
  events := make(chan *misc.Event)

  go p.watchChatlog(events)

  for {
    event := <-events

    if event.Event == "EOF" {
      break
    }

    // JSON serialize parsed result
    serialized, err := event.JSON()
    if err != nil {
      continue
    }

    // Output to Stdout
    fmt.Fprintln(os.Stdout, serialized)
  }

  return nil
}

// watchChatlog for new entries
func (p *ParseCmd) watchChatlog(events chan *misc.Event) {
  msg := make(chan string)
  defer close(msg)

  go watcher.Parse(p.File, msg, p.All, p.Watch)

  for {
    row := <-msg

    if row == "EOF" {
      events <- &misc.Event{Event: "EOF"}
      break
    }

    result, ok := parser.Parse(row, p.Name)
    if !ok {
      continue
    }

    events <- result
  }
}
