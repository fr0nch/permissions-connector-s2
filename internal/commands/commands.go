package commands

import (
	"errors"
	"fmt"
	"permissions-connector-s2/internal/pkg/config"
	"permissions-connector-s2/internal/pkg/permissions"
	trs "permissions-connector-s2/internal/pkg/translations"
	"permissions-connector-s2/internal/pkg/utils"

	s2sdk "github.com/fr0nch/go-plugify-s2sdk/v2"
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
			Language:  s2sdk.GetClientLanguage(slot),
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
		s2sdk.AddConsoleCommand(cmdName, trs.GetTranslation(info.Description, s2sdk.GetServerLanguage()), flags, c.conCommandCallback, s2sdk.HookMode_Post)
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
				trs.GetTranslation("command.error.no_permission", ctx.Caller.Language),
			)
			return s2sdk.ResultType_Handled
		}
	}

	argsCount := len(ctx.Args)
	if argsCount < command.MinArgs {
		utils.ReplyToCommandCleared(caller, context,
			trs.GetTranslation("command.error.invalid_arguments_count", ctx.Caller.Language,
				trs.GetTranslation(command.Example, ctx.Caller.Language)),
		)
		return s2sdk.ResultType_Handled
	}

	return command.Handler(ctx)
}

func findPlayerTarget(ctx *CommandContext) (*Player, error) {
	clients := s2sdk.ProcessTargetString(ctx.Caller.Slot, ctx.Args[0])
	clientCounts := len(clients)

	if clientCounts == 0 {
		return nil, errors.New(trs.GetTranslation("command.error.user.not_found", ctx.Caller.Language, ctx.Info))
	}

	if clientCounts > 1 {
		return nil, errors.New(trs.GetTranslation("command.error.user.multiple_found", ctx.Caller.Language, ctx.Info.Example))
	}

	target := NewPlayer(clients[0])
	if target == nil {
		return nil, errors.New(trs.GetTranslation("command.error.user.not_valid", ctx.Caller.Language, ctx.Info.Example))
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
		err = errors.New(trs.GetTranslation("command.error.affect_user.disallow", caller.Language, target.Name, target.SteamID64))
	case permissions.Status_ActorUserNotFound:
		err = errors.New(trs.GetTranslation("command.error.actor.not_found", caller.Language))
	case permissions.Status_TargetUserNotFound:
		err = errors.New(trs.GetTranslation("command.error.target.not_found", caller.Language, target.Name, target.SteamID64))
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
				trs.GetTranslation(c.handlers[name].Description, ctx.Caller.Language),
				trs.GetTranslation(c.handlers[name].Example, ctx.Caller.Language),
			},
		)
	}

	if len(data) == 0 {
		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
			trs.GetTranslation("command.error.no_available_commands", ctx.Caller.Language),
		)
		return s2sdk.ResultType_Continue
	}

	if ctx.ConCommandContext == s2sdk.ConCommandContext_Chat {
		s2sdk.ReplyToCommand(ctx.ConCommandContext, ctx.Caller.Slot, "[Permissions] Список команд:")

		for i := range data {
			s2sdk.ReplyToCommand(ctx.ConCommandContext, ctx.Caller.Slot,
				fmt.Sprintf("%s - %s", data[i][0].(string),
					trs.GetTranslation(data[i][1].(string), ctx.Caller.Language),
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
	//p.log.Infof("Config reload successfully.")

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandReloadTranslations(ctx *CommandContext) s2sdk.ResultType {
	err := trs.LoadTranslation("permissions-connector-s2")
	if err != nil {
		c.log.Errorf("Error reload t: %v", err)
		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, "[Permissions] Error reload t: "+err.Error())
	}

	utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, "[Permissions] Translations reload successfully.")
	//p.log.Infof("Translations reload successfully.")

	return s2sdk.ResultType_Continue
}

type CmdSuccessResult struct {
	Key  string
	Args []any
}

func translateCommandResult(caller, target *Player, data GroupOrNode, additional CmdSuccessResult) string {
	var msg string

	switch data.Result {
	case permissions.Status_Success:
		msg = trs.GetTranslation(additional.Key, caller.Language, additional.Args...)
	case permissions.Status_Allow:
		msg = trs.GetTranslation("command.permission.allow", caller.Language, target.Name, target.Slot, data.Name)
	case permissions.Status_Disallow:
		msg = trs.GetTranslation("command.permission.disallow", caller.Language, target.Name, target.Slot, data.Name)
	case permissions.Status_PermNotFound:
		msg = trs.GetTranslation("command.error.permission.not_found", caller.Language)
	case permissions.Status_GroupNotFound:
		msg = trs.GetTranslation("command.error.group.not_found", caller.Language, data.Name)
	case permissions.Status_ChildGroupNotFound:
		msg = trs.GetTranslation("command.error.group.child_not_found", caller.Language, data.Name)
	case permissions.Status_ParentGroupNotFound:
		msg = trs.GetTranslation("command.error.group.parent_not_found", caller.Language, data.Name)
	case permissions.Status_ActorUserNotFound:
		msg = trs.GetTranslation("command.error.actor.not_found", caller.Language, target.Name, target.Slot)
	case permissions.Status_TargetUserNotFound:
		msg = trs.GetTranslation("command.error.target.not_found", caller.Language, target.Name, target.Slot)
	case permissions.Status_GroupAlreadyExist:
		msg = trs.GetTranslation("command.error.group.already_exist", caller.Language, target.Name, target.Slot, data.Name)
	case permissions.Status_PermAlreadyGranted:
		msg = trs.GetTranslation("command.error.permission.already_granted", caller.Language, target.Name, target.Slot, data.Name)
	default:
		msg = trs.GetTranslation("command.error.unknown", caller.Language)
	}

	return msg
}
