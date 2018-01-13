//go:generate go get -v github.com/josephspurrier/goversioninfo/...
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	. "github.com/portapps/portapps"
)

func init() {
	Papp.ID = "discord-portable"
	Papp.Name = "Discord"
	Init()
}

func main() {
	Papp.AppPath = AppPathJoin("app")
	Papp.DataPath = AppPathJoin("data")

	electronBinPath := PathJoin(Papp.AppPath, FindElectronAppFolder("app-", Papp.AppPath))

	Papp.Process = PathJoin(Papp.AppPath, "Update.exe")
	Papp.Args = []string{"--processStart", "Discord.exe"}
	Papp.WorkingDir = electronBinPath

	Launch()
}
