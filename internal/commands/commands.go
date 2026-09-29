package commands

import (
	"errors"
	"fmt"
	"permissions-connector-s2/internal/pkg/config"
	"permissions-connector-s2/internal/pkg/utils"

	permissions "github.com/fr0nch/go-plugify-permissions"
	s2sdk "github.com/fr0nch/go-plugify-s2sdk/v2"
	t "github.com/fr0nch/go-plugify-translations"
	"github.com/fr0nch/logger"
)

type Player struct {
	Slot      int32
	Name      string
	SteamID64 uint64
	Language  string
}

func NewPlayer(slot int32) *Player {
	if slot > -1 {
		if !utils.IsValidClient(slot) {
			return nil
		}

		return &Player{
			Slot:      slot,
			Name:      s2sdk.GetClientName(slot),
			SteamID64: s2sdk.GetClientSteamID64(slot),
			Language:  s2sdk.GetClientLanguageCode(slot),
		}
	}

	return &Player{
		Slot:      slot,
		Name:      "Console",
		SteamID64: 0,
		Language:  s2sdk.GetServerLanguage(),
	}
}

type GroupOrNode struct {
	Name         string
	Duration     uint64
	NewTimestamp int64
	OldTimestamp int64
	PermSource   permissions.PermSource
	NewStatus    permissions.Status
	OldStatus    permissions.Status
	Result       permissions.Status
}

type CommandContext struct {
	CommandName       string
	ConCommandContext s2sdk.ConCommandContext
	Args              []string
	Info              *CommandInfo

	Caller *Player
}

type CommandHandler func(ctx *CommandContext) s2sdk.ResultType

type CommandInfo struct {
	Handler     CommandHandler
	MinArgs     int
	Permission  string
	Description string
	Example     string
}

type Commands struct {
	pluginID int64

	config *config.Config
	log    *logger.Logger

	handlers map[string]CommandInfo
}

func (c *Commands) addCommand(name string, info CommandInfo) {
	if c.handlers == nil {
		c.handlers = make(map[string]CommandInfo)
	}

	c.handlers[name] = info
}

func NewCommands(pluginID int64, config *config.Config, log *logger.Logger) *Commands {
	c := &Commands{
		pluginID: pluginID,
		config:   config,
		log:      log,
	}

	c.addCommand("perms", CommandInfo{
		Handler:     c.commandPerms,
		MinArgs:     0,
		Permission:  "",
		Description: "command.description.info",
		Example:     "command.example.info",
	})
	c.addCommand("perms_reload_cfg", CommandInfo{
		Handler:     c.commandReloadConfig,
		MinArgs:     0,
		Permission:  "permissions.reload_config",
		Description: "command.description.reload_config",
		Example:     "command.example.reload_config",
	})
	c.addCommand("perms_reload_translations", CommandInfo{
		Handler:     c.commandReloadTranslations,
		MinArgs:     0,
		Permission:  "permissions.reload_translations",
		Description: "command.description.reload_config",
		Example:     "command.example.reload_config",
	})

	registerGroupCommands(c)
	registerUserCommands(c)

	return c
}

func (c *Commands) Register() {
	var flags = s2sdk.ConVarFlag_Release | s2sdk.ConVarFlag_ServerCanExecute | s2sdk.ConVarFlag_LinkedConcommand | s2sdk.ConVarFlag_ClientCanExecute
	for cmdName, info := range c.handlers {
		s2sdk.AddConsoleCommand(cmdName, t.Translate(s2sdk.GetServerLanguage(), info.Description), flags, c.conCommandCallback, s2sdk.HookMode_Post)
	}
}

func (c *Commands) conCommandCallback(caller int32, context s2sdk.ConCommandContext, arguments []string) s2sdk.ResultType {
	commandName := arguments[0]
	command, exists := c.handlers[commandName]
	if !exists {
		return s2sdk.ResultType_Continue
	}

	c.log.Debugf("ConCommandCallback :: Command: '%s' :: Args: '%v'", commandName, arguments)

	ctx := &CommandContext{
		CommandName:       commandName,
		ConCommandContext: context,
		Args:              arguments[1:],
		Caller:            NewPlayer(caller),
		Info:              &command,
	}

	if caller > -1 {
		if command.Permission != "" && permissions.HasPermission(ctx.Caller.SteamID64, command.Permission) != permissions.Status_Allow {
			utils.ReplyToCommandCleared(caller, context,
				t.Translate(ctx.Caller.Language, "command.error.no_permission"),
			)
			return s2sdk.ResultType_Handled
		}
	}

	argsCount := len(ctx.Args)
	if argsCount < command.MinArgs {
		utils.ReplyToCommandCleared(caller, context,
			t.TranslateFormat(ctx.Caller.Language, "command.error.invalid_arguments_count", []any{t.Translate(ctx.Caller.Language, command.Example)}),
		)
		return s2sdk.ResultType_Handled
	}

	return command.Handler(ctx)
}

func findPlayerTarget(ctx *CommandContext) (*Player, error) {
	clients := s2sdk.ProcessTargetString(ctx.Caller.Slot, ctx.Args[0])
	clientCounts := len(clients)

	if clientCounts == 0 {
		return nil, errors.New(t.TranslateFormat(ctx.Caller.Language, "command.error.user.not_found", []any{t.Translate(ctx.Caller.Language, ctx.Info.Example)}))
	}

	if clientCounts > 1 {
		return nil, errors.New(t.TranslateFormat(ctx.Caller.Language, "command.error.user.multiple_found", []any{t.Translate(ctx.Caller.Language, ctx.Info.Example)}))
	}

	target := NewPlayer(clients[0])
	if target == nil {
		return nil, errors.New(t.TranslateFormat(ctx.Caller.Language, "command.error.user.not_valid", []any{ctx.Info.Example}))
	}

	return target, canAffectUser(ctx.Caller, target)
}

func canAffectUser(caller, target *Player) error {
	if caller.Slot == -1 {
		return nil
	}

	resultAffect := permissions.CanAffectUser(caller.SteamID64, target.SteamID64)
	var err error

	switch resultAffect {
	case permissions.Status_Disallow:
		err = errors.New(t.TranslateFormat(caller.Language, "command.error.affect_user.disallow", []any{target.Name, target.SteamID64}))
	case permissions.Status_ActorUserNotFound:
		err = errors.New(t.Translate(caller.Language, "command.error.actor.not_found"))
	case permissions.Status_TargetUserNotFound:
		err = errors.New(t.TranslateFormat(caller.Language, "command.error.target.not_found", []any{target.Name, target.SteamID64}))
	}

	return err
}

func (c *Commands) commandPerms(ctx *CommandContext) s2sdk.ResultType {
	var data [][]any
	for name := range c.handlers {
		if ctx.Caller.Slot > -1 {
			if permissions.HasPermission(ctx.Caller.SteamID64, c.handlers[name].Permission) != permissions.Status_Allow {
				continue
			}
		}

		if name == "perms" {
			continue
		}

		data = append(data,
			[]any{
				name,
				t.Translate(ctx.Caller.Language, c.handlers[name].Description),
				t.Translate(ctx.Caller.Language, c.handlers[name].Example),
			},
		)
	}

	if len(data) == 0 {
		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
			t.Translate(ctx.Caller.Language, "command.error.no_available_commands"),
		)
		return s2sdk.ResultType_Continue
	}

	if ctx.ConCommandContext == s2sdk.ConCommandContext_Chat {
		s2sdk.ReplyToCommand(ctx.ConCommandContext, ctx.Caller.Slot, "[Permissions] Список команд:")

		for i := range data {
			s2sdk.ReplyToCommand(ctx.ConCommandContext, ctx.Caller.Slot,
				fmt.Sprintf("%s - %s", data[i][0].(string),
					t.Translate(ctx.Caller.Language, data[i][1].(string)),
				),
			)
		}

		return s2sdk.ResultType_Continue
	}

	if ctx.ConCommandContext == s2sdk.ConCommandContext_Console {
		utils.PrintConsoleTableToClient(ctx.Caller.Slot, "Список команд",
			[]string{"command", "description", "example"},
			data,
		)
	}

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandReloadConfig(ctx *CommandContext) s2sdk.ResultType {
	cfg, err := config.LoadConfig("permissions", "connector.yml",
		config.Settings{
			NotifyMode: config.NotifyMode{
				Console: true,
				Chat:    true,
				Hint:    true,
			},
			TimeMode:     0,
			HintDuration: 5,
		},
	)
	if err != nil {
		c.log.Errorf("Error reload config: %v", err)
		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, "[Permissions] Error reload config: "+err.Error())
	}

	c.config.Settings = cfg

	utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, "[Permissions] Config reload successfully.")
	c.log.Infof("Config reload successfully.")

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandReloadTranslations(ctx *CommandContext) s2sdk.ResultType {
	t.ReloadTransalations()

	utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, "[Permissions] Translations reload successfully.")
	c.log.Infof("Translations reload successfully.")

	return s2sdk.ResultType_Continue
}

type CmdSuccessResult struct {
	Key  string
	Args []any
}
