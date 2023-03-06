// +build tools

package main

import (
	- "github.com/smartystreets/goconvey"
	_ "github.com/git-chglog/git-chglog/cmd/git-chglog"
	_ "github.com/go-task/task/v3/cmd/task"
	_ "github.com/goreleaser/goreleaser"
	_ "github.com/kisielk/godepgraph"
	_ "github.com/siderolabs/conform/cmd/conform"
)
