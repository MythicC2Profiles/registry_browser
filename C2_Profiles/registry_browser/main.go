package main

import (
	"MyContainer/registry_browser"

	"github.com/MythicMeta/MythicContainer"
)

func main() {
	// load up the agent functions directory so all the init() functions execute
	registry_browser.Initialize()
	// sync over definitions and listen
	MythicContainer.StartAndRunForever([]MythicContainer.MythicServices{
		MythicContainer.MythicServiceCustomBrowser,
	})
}
