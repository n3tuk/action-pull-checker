package main

import (
	"github.com/n3tuk/action-pull-requester/cmd"
)

var (
	version = "v0.0.0"
	date    = "2023-01-01 00:00:00"
	commit  = "0000000"
	branch  = "main"
)

func main() {
	cmd.Version = version
	cmd.BuildDate = date
	cmd.Commit = commit
	cmd.Branch = branch
	cmd.Execute()
}
