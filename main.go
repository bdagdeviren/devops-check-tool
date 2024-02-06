package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"plugin"
)

type BasePlugin interface {
	Run()
}

func main() {

	pluginDir := "plugins"

	files, err := ioutil.ReadDir(pluginDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".so" {
			plug, err := plugin.Open(fmt.Sprintf(pluginDir+"/%s", file.Name()))
			if err != nil {
				log.Fatal(err)
			}

			pluginSymbol, err := plug.Lookup("BasePlugin")
			if err != nil {
				log.Fatal(err)
			}

			var basePlugin BasePlugin
			basePlugin, ok := pluginSymbol.(BasePlugin)
			if !ok {
				log.Fatal("Invalid Plugin Type")
			}

			basePlugin.Run()
		}
	}
}
