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

	_ "github.com/kevinburke/go-bindata"
	"github.com/portapps/discord-portable/assets"
	. "github.com/portapps/portapps"
	"github.com/portapps/portapps/pkg/shortcut"
	"github.com/portapps/portapps/pkg/utl"
)

var (
	app *App
)

func init() {
	var err error

	// Init app
	if app, err = New("discord-portable", "Discord"); err != nil {
		Log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	utl.CreateFolder(app.DataPath)
	electronBinPath := utl.PathJoin(app.AppPath, utl.FindElectronAppFolder("app-", app.AppPath))

	app.Process = utl.PathJoin(electronBinPath, "Discord.exe")
	app.WorkingDir = electronBinPath

	// Update settings
	settingsPath := utl.PathJoin(app.DataPath, "settings.json")
	if _, err := os.Stat(settingsPath); err == nil {
		Log.Info().Msg("Update settings...")
		rawSettings, err := ioutil.ReadFile(settingsPath)
		if err == nil {
			jsonMapSettings := make(map[string]interface{})
			if err = json.Unmarshal(rawSettings, &jsonMapSettings); err != nil {
				Log.Error().Err(err).Msg("Settings unmarshal")
			}
			Log.Info().Interface("settings", jsonMapSettings).Msg("Current settings")

			jsonMapSettings["SKIP_HOST_UPDATE"] = true
			Log.Info().Interface("settings", jsonMapSettings).Msg("New settings")

			jsonSettings, err := json.Marshal(jsonMapSettings)
			if err != nil {
				Log.Error().Err(err).Msg("Settings marshal")
			}
			err = ioutil.WriteFile(settingsPath, jsonSettings, 0644)
			if err != nil {
				Log.Error().Err(err).Msg("Write settings")
			}
		}
	}

	// Workaround for tray.png not found issue (https://github.com/portapps/discord-ptb-portable/issues/2)
	if err := assets.RestoreAssets(app.RootPath, "data"); err != nil {
		Log.Error().Err(err).Msg("Cannot restore data assets")
	}

	// Copy default shortcut
	shortcutPath := path.Join(utl.StartMenuPath(), "Discord Portable.lnk")
	defaultShortcut, err := assets.Asset("Discord.lnk")
	if err != nil {
		Log.Error().Err(err).Msg("Cannot load asset Discord.lnk")
	}
	err = ioutil.WriteFile(shortcutPath, defaultShortcut, 0644)
	if err != nil {
		Log.Error().Err(err).Msg("Cannot write default shortcut")
	}

	// Update default shortcut
	err = shortcut.Create(shortcut.Shortcut{
		ShortcutPath:     shortcutPath,
		TargetPath:       app.Process,
		Arguments:        shortcut.Property{Clear: true},
		Description:      shortcut.Property{Value: "Discord Portable by Portapps"},
		IconLocation:     shortcut.Property{Value: app.Process},
		WorkingDirectory: shortcut.Property{Value: app.AppPath},
	})
	if err != nil {
		Log.Error().Err(err).Msg("Cannot create shortcut")
	}
	defer func() {
		if err := os.Remove(shortcutPath); err != nil {
			Log.Error().Err(err).Msg("Cannot remove shortcut")
		}
	}()

	app.Launch(os.Args[1:])
}
