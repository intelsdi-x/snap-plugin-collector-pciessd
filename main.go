package main

import (
	"github.com/intelsdi-x/snap-plugin-collector-pciessd/collector"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

func main() {
	plugin.StartCollector(collector.New(), collector.PLUGIN_NAME, collector.PLUGIN_VERSION)
}
