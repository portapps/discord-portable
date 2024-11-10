//go:generate go install -v github.com/kevinburke/go-bindata/v4/go-bindata
//go:generate go-bindata -prefix res/ -pkg assets -o assets/assets.go res/Discord.lnk res/pinned_update.json
//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest
package main

import (
	"encoding/json"
	"io"
	"os"
	"path"
	"strings"

	"github.com/portapps/discord-portable/assets"
	"github.com/portapps/portapps/v3"
	"github.com/portapps/portapps/v3/pkg/log"
	"github.com/portapps/portapps/v3/pkg/registry"
	"github.com/portapps/portapps/v3/pkg/shortcut"
	"github.com/portapps/portapps/v3/pkg/utl"
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
	electronAppPath := app.ElectronAppPath()

	app.Process = utl.PathJoin(electronAppPath, "Discord.exe")
	app.Args = []string{
		"--user-data-dir=" + app.DataPath,
	}
	app.WorkingDir = electronAppPath

	// Cleanup on exit
	if cfg.Cleanup {
		defer func() {
			regKey := registry.Key{
				Key:  `HKCU\SOFTWARE\Discord`,
				Arch: "32",
			}
			if regKey.Exists() {
				if err := regKey.Delete(true); err != nil {
					log.Error().Err(err).Msg("Cannot remove registry key")
				}
			}
			utl.Cleanup([]string{
				path.Join(os.Getenv("APPDATA"), "discord"),
				path.Join(os.Getenv("TEMP"), "Discord Crashes"),
			})
		}()
	}

	// Update settings
	settingsPath := utl.PathJoin(app.DataPath, "settings.json")
	if _, err := os.Stat(settingsPath); err == nil {
		log.Info().Msg("Update settings...")
		rawSettings, err := os.ReadFile(settingsPath)
		if err == nil {
			jsonMapSettings := make(map[string]interface{})
			if err = json.Unmarshal(rawSettings, &jsonMapSettings); err != nil {
				log.Error().Err(err).Msg("Settings unmarshal")
			}
			log.Info().Interface("settings", jsonMapSettings).Msg("Current settings")

			jsonMapSettings["SKIP_HOST_UPDATE"] = true
			jsonMapSettings["USE_PINNED_UPDATE_MANIFEST"] = true
			log.Info().Interface("settings", jsonMapSettings).Msg("New settings")

			jsonSettings, err := json.Marshal(jsonMapSettings)
			if err != nil {
				log.Error().Err(err).Msg("Settings marshal")
			}
			err = os.WriteFile(settingsPath, jsonSettings, 0644)
			if err != nil {
				log.Error().Err(err).Msg("Write settings")
			}
		}
	} else {
		fo, err := os.Create(settingsPath)
		if err != nil {
			log.Error().Err(err).Msg("Cannot create settings.json")
		}
		defer fo.Close()
		if _, err = io.Copy(fo, strings.NewReader(`{
  "SKIP_HOST_UPDATE": true,
  "USE_PINNED_UPDATE_MANIFEST": true
}`)); err != nil {
			log.Error().Err(err).Msg("Cannot write to settings.json")
		}
	}

	// Copy pinned_update.json
	pinnedUpdate, err := assets.Asset("pinned_update.json")
	if err != nil {
		log.Error().Err(err).Msg("Cannot load asset pinned_update.json")
	}
	err = os.WriteFile(utl.PathJoin(app.DataPath, "pinned_update.json"), pinnedUpdate, 0644)
	if err != nil {
		log.Error().Err(err).Msg("Cannot write pinned_update.json")
	}

	// Copy default shortcut
	shortcutPath := path.Join(utl.StartMenuPath(), "Discord Portable.lnk")
	defaultShortcut, err := assets.Asset("Discord.lnk")
	if err != nil {
		log.Error().Err(err).Msg("Cannot load asset Discord.lnk")
	}
	err = os.WriteFile(shortcutPath, defaultShortcut, 0644)
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
