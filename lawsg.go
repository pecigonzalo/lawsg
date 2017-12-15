package main

import (
	"github.com/mmcquillan/lawsg/config"
	"github.com/mmcquillan/lawsg/fetch"
)

var version string

func main() {
	config.Aws()
	var options config.Options
	config.Defaults(&options)
	config.Commands(&options)
	config.Saved(&options)
	config.EnvVars(&options)
	config.Flags(&options)
	config.Validate(&options)
	switch options.Command {
	case "groups":
		fetch.Groups(options)
	case "streams":
		fetch.Streams(options)
	case "get":
		fetch.Logs(options)
	case "version":
		fetch.Version(version)
	default:
		fetch.Help(options, version)
	}
}
