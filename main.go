package main

import (
	"permissions-connector-s2/internal"

	"github.com/untrustedmodders/go-plugify"
)

const pluginName = "permissions-connector-s2"

func init() {
	plugin := internal.NewPermsConnectorPlugin()
	plugin.Plugin = plugify.NewPlugin(pluginName, plugin.OnPluginStart, plugin.OnPluginUpdate, plugin.OnPluginEnd)
}

func main() {}
