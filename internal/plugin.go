package internal

import (
	"path/filepath"
	"permissions-connector-s2/internal/commands"
	"permissions-connector-s2/internal/pkg/config"

	s2sdk "github.com/fr0nch/go-plugify-s2sdk/v2"
	t "github.com/fr0nch/go-plugify-translations"
	"github.com/fr0nch/logger"
	"github.com/untrustedmodders/go-plugify"
)

type PermsConnectorPlugin struct {
	Plugin   plugify.Plugin
	PluginID int64

	config   *config.Config
	commands *commands.Commands

	loadedUserTimers map[uint64]s2sdk.TimerId

	log *logger.Logger
}

func NewPermsConnectorPlugin() *PermsConnectorPlugin {
	return &PermsConnectorPlugin{
		loadedUserTimers: make(map[uint64]s2sdk.TimerId),
	}
}

func (p *PermsConnectorPlugin) OnPluginStart() error {
	p.PluginID = int64(p.Plugin.ID())

	var err error

	p.log, err = logger.NewWithOptions(logger.Options{PluginName: p.Plugin.Name(), Folder: "permissions"})
	if err != nil {
		return err
	}

	p.config = config.NewConfig()
	p.config.Settings, err = config.LoadConfig(p.Plugin.Location(), "permissions", "permissions_connector_s2.settings.yaml", config.DefaultSettings())

	const translationFile = "permissions_connector_s2.translations.yaml"
	if err := t.LoadTranslation([]string{
		filepath.Join(p.Plugin.Location(), translationFile),
		filepath.Join(plugify.ConfigsDir(), "permissions", translationFile),
		filepath.Join(plugify.DataDir(), "permissions", translationFile),
	}); err != "" {
		p.log.Errorf("Error loading translations: %s", err)
	}

	//permissions.SetConnectorID(p.PluginID)

	p.commands = commands.NewCommands(p.PluginID, p.Plugin.Location(), p.config, p.log)
	p.commands.Register()

	p.RegisterEvents()

	return nil
}

func (p *PermsConnectorPlugin) OnPluginUpdate(dt float32) error {
	return nil
}

func (p *PermsConnectorPlugin) OnPluginEnd() error {
	//permissions.ResetConnectorID()
	return nil
}
