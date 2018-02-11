//go:generate go get -v github.com/josephspurrier/goversioninfo/...
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

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

	Papp.Process = PathJoin(electronBinPath, "Discord.exe")
	Papp.WorkingDir = electronBinPath

	// Update settings
	settingsPath := PathJoin(Papp.DataPath, "settings.json")
	if _, err := os.Stat(settingsPath); err == nil {
		Log.Info("Update settings...")
		rawSettings, err := ioutil.ReadFile(settingsPath)
		if err == nil {
			jsonMapSettings := make(map[string]interface{})
			json.Unmarshal(rawSettings, &jsonMapSettings)
			Log.Info("Current settings:", jsonMapSettings)

			jsonMapSettings["SKIP_HOST_UPDATE"] = true
			Log.Info("New settings:", jsonMapSettings)

			jsonSettings, err := json.Marshal(jsonMapSettings)
			if err != nil {
				Log.Error("Settings marshal:", err)
			}
			err = ioutil.WriteFile(settingsPath, jsonSettings, 0644)
			if err != nil {
				Log.Error("Write settings:", err)
			}
		}
	}

	Launch(os.Args[1:])
}
