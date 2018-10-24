//go:generate go install -v github.com/kevinburke/go-bindata/go-bindata
//go:generate go-bindata -prefix res/ -pkg assets -o assets/assets.go res/Discord.lnk
//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/portapps/discord-portable/assets"
	. "github.com/portapps/portapps"
)

func init() {
	Papp.ID = "discord-portable"
	Papp.Name = "Discord"
	Init()
}

func main() {
	Papp.AppPath = AppPathJoin("app")
	Papp.DataPath = CreateFolder(AppPathJoin("data"))

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

	// Copy default shortcut
	shortcutPath := path.Join(os.Getenv("APPDATA"), "Microsoft", "Windows", "Start Menu", "Programs", "Discord Portable.lnk")
	defaultShortcut, err := assets.Asset("Discord.lnk")
	if err != nil {
		Log.Error("Cannot load asset Discord.lnk:", err)
	}
	err = ioutil.WriteFile(shortcutPath, defaultShortcut, 0644)
	if err != nil {
		Log.Error("Cannot write default shortcut:", err)
	}

	// Update default shortcut
	err = CreateShortcut(WindowsShortcut{
		ShortcutPath:     shortcutPath,
		TargetPath:       Papp.Process,
		Arguments:        WindowsShortcutProperty{Clear: true},
		Description:      WindowsShortcutProperty{Value: "Discord Portable by Portapps"},
		IconLocation:     WindowsShortcutProperty{Value: Papp.Process},
		WorkingDirectory: WindowsShortcutProperty{Value: Papp.AppPath},
	})
	if err != nil {
		Log.Error("Cannot create shortcut:", err)
	}

	Launch(os.Args[1:])
}
