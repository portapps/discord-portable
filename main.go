//go:generate go install -v github.com/kevinburke/go-bindata/go-bindata
//go:generate go-bindata -prefix res/ -pkg assets -o assets/assets.go res/Discord.lnk
//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest
package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/portapps/discord-portable/assets"
	"github.com/portapps/portapps/v2"
	"github.com/portapps/portapps/v2/pkg/log"
	"github.com/portapps/portapps/v2/pkg/shortcut"
	"github.com/portapps/portapps/v2/pkg/utl"
)

type config struct {
	Cleanup bool `yaml:"cleanup" mapstructure:"cleanup"`
}

var (
	app *portapps.App
	cfg *config
)

func init() {
	var err error

	// Default config
	cfg = &config{
		Cleanup: false,
	}

	// Init app
	if app, err = portapps.NewWithCfg("discord-portable", "Discord", cfg); err != nil {
		log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	utl.CreateFolder(app.DataPath)
	electronBinPath := utl.PathJoin(app.AppPath, utl.FindElectronAppFolder("app-", app.AppPath))

	app.Process = utl.PathJoin(electronBinPath, "Discord.exe")
	app.WorkingDir = electronBinPath

	// Cleanup on exit
	if cfg.Cleanup {
		defer func() {
			utl.Cleanup([]string{
				path.Join(os.Getenv("APPDATA"), "discord"),
			})
		}()
	}

	// Update settings
	settingsPath := utl.PathJoin(app.DataPath, "settings.json")
	if _, err := os.Stat(settingsPath); err == nil {
		log.Info().Msg("Update settings...")
		rawSettings, err := ioutil.ReadFile(settingsPath)
		if err == nil {
			jsonMapSettings := make(map[string]interface{})
			if err = json.Unmarshal(rawSettings, &jsonMapSettings); err != nil {
				log.Error().Err(err).Msg("Settings unmarshal")
			}
			log.Info().Interface("settings", jsonMapSettings).Msg("Current settings")

			jsonMapSettings["SKIP_HOST_UPDATE"] = true
			log.Info().Interface("settings", jsonMapSettings).Msg("New settings")

			jsonSettings, err := json.Marshal(jsonMapSettings)
			if err != nil {
				log.Error().Err(err).Msg("Settings marshal")
			}
			err = ioutil.WriteFile(settingsPath, jsonSettings, 0644)
			if err != nil {
				log.Error().Err(err).Msg("Write settings")
			}
		}
	}

	// Workaround for tray.png not found issue (https://github.com/portapps/discord-ptb-portable/issues/2)
	if err := assets.RestoreAssets(app.RootPath, "data"); err != nil {
		log.Error().Err(err).Msg("Cannot restore data assets")
	}

	// Copy default shortcut
	shortcutPath := path.Join(utl.StartMenuPath(), "Discord Portable.lnk")
	defaultShortcut, err := assets.Asset("Discord.lnk")
	if err != nil {
		log.Error().Err(err).Msg("Cannot load asset Discord.lnk")
	}
	err = ioutil.WriteFile(shortcutPath, defaultShortcut, 0644)
	if err != nil {
		log.Error().Err(err).Msg("Cannot write default shortcut")
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
		log.Error().Err(err).Msg("Cannot create shortcut")
	}
	defer func() {
		if err := os.Remove(shortcutPath); err != nil {
			log.Error().Err(err).Msg("Cannot remove shortcut")
		}
	}()

	defer app.Close()
	app.Launch(os.Args[1:])
}
