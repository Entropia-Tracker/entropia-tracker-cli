package main

import (
  "fmt"
)

type VersionCmd struct {
}

func (v *VersionCmd) Run(ctx *Context) error {
  fmt.Println(Version)
  return nil
}
