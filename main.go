package main

import (
	"fmt"
	"github.com/Liang-jonas/jnote/Cmd"
)

func main() {
	cmd := Cmd.NewCmd()
	cmd.RegisterCmds(
		Cmd.VersionCmd(),
		Cmd.ServerCmd(),
	)
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
