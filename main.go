//go:generate go get -v github.com/josephspurrier/goversioninfo/...
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	. "github.com/portapps/portapps"
)

func init() {
	App.ID = "discord-portable"
	App.Name = "Discord"
	Init()
}

func main() {
	App.MainPath = FindElectronMainFolder("app-")
	App.DataPath = CreateFolder(PathJoin(App.RootDataPath, "AppData", "Roaming", "discord"))
	App.Process = RootPathJoin("Update.exe")
	App.Args = []string{"--processStart", "Discord.exe"}
	App.WorkingDir = App.MainPath

	OverrideEnv("USERPROFILE", App.RootDataPath)
	Launch()
}
