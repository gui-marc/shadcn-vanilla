package main

import (
	"github.com/alecthomas/kong"
	"github.com/gui-marc/shadcn-vanilla/internal/cmd"
)

func main() {
	cli := struct {
		Add     cmd.AddCmd     `cmd:"" help:"Add a component"`
		Install cmd.InstallCmd `cmd:"" help:"Initialize project configuration"`
	}{}

	ctx := kong.Parse(&cli)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
