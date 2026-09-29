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

	loadedUserTimers [64]s2sdk.TimerId

	log *logger.Logger
}

func NewPermsConnectorPlugin() *PermsConnectorPlugin {
	return &PermsConnectorPlugin{}
}

func (p *PermsConnectorPlugin) OnPluginStart() error {
	p.PluginID = int64(p.Plugin.ID())

	var err error

	p.log, err = logger.NewWithOptions(logger.Options{PluginName: p.Plugin.Name(), Folder: "permissions"})
	if err != nil {
		return err
	}

	p.config = config.NewConfig()
	p.config.Settings, err = config.LoadConfig("permissions", "connector.yml",
		config.Settings{
			NotifyMode: config.NotifyMode{
				Console: true,
				Chat:    true,
				Hint:    true,
			},
			HintDuration:        5,
			TimeMode:            0,
			AbbreviateTime:      false,
			FormatTimestamp:     "2006-01-02 15:04:05",
			DisplayActiveGroups: true,
			HiddenGroups:        []string{"Default"},
		},
	)

	const translationFile = "permissions-connector-s2.yaml"
	if err := t.LoadTranslation([]string{
		filepath.Join(p.Plugin.Location(), translationFile),
		filepath.Join(plugify.BaseDir(), "translations", translationFile),
		filepath.Join(plugify.ConfigsDir(), translationFile),
		filepath.Join(plugify.DataDir(), translationFile),
	}); err != "" {
		p.log.Errorf("Error loading translations: %s", err)
	}

	//permissions.SetConnectorID(p.PluginID)

	p.commands = commands.NewCommands(p.PluginID, p.config, p.log)
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
