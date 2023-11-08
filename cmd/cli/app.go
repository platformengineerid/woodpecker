// Copyright 2021 Woodpecker Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"go.woodpecker-ci.org/woodpecker/cli/common"
	"go.woodpecker-ci.org/woodpecker/cli/cron"
	"go.woodpecker-ci.org/woodpecker/cli/deploy"
	"go.woodpecker-ci.org/woodpecker/cli/exec"
	"go.woodpecker-ci.org/woodpecker/cli/info"
	"go.woodpecker-ci.org/woodpecker/cli/lint"
	"go.woodpecker-ci.org/woodpecker/cli/log"
	"go.woodpecker-ci.org/woodpecker/cli/loglevel"
	"go.woodpecker-ci.org/woodpecker/cli/pipeline"
	"go.woodpecker-ci.org/woodpecker/cli/registry"
	"go.woodpecker-ci.org/woodpecker/cli/repo"
	"go.woodpecker-ci.org/woodpecker/cli/secret"
	"go.woodpecker-ci.org/woodpecker/cli/user"
	"go.woodpecker-ci.org/woodpecker/version"
)

//go:generate go run docs.go app.go
func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "woodpecker-cli"
	app.Version = version.String()
	app.Usage = "command line utility"
	app.EnableBashCompletion = true
	app.Flags = common.GlobalFlags
	app.Commands = []*cli.Command{
		pipeline.Command,
		log.Command,
		deploy.Command,
		exec.Command,
		info.Command,
		registry.Command,
		secret.Command,
		repo.Command,
		user.Command,
		lint.Command,
		loglevel.Command,
		cron.Command,
	}

	zlog.Logger = zlog.Output(
		zerolog.ConsoleWriter{
			Out: os.Stderr,
		},
	)
	for _, command := range app.Commands {
		command.Before = common.SetupConsoleLogger
	}

	return app
}
