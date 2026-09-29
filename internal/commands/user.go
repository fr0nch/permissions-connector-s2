package commands

import (
	"fmt"
	trs "permissions-connector-s2/internal/pkg/translations"
	"permissions-connector-s2/internal/pkg/utils"
	"strconv"
	"strings"
	"time"

	permissions "github.com/fr0nch/go-plugify-permissions"

	s2sdk "github.com/fr0nch/go-plugify-s2sdk/v2"
)

func registerUserCommands(c *Commands) {
	c.addCommand("perms_users", CommandInfo{
		Handler:     c.commandUserList,
		MinArgs:     0,
		Permission:  "permissions.user.list",
		Description: "command.description.user.list",
		Example:     "command.example.user.list",
	})
	c.addCommand("perms_user_info", CommandInfo{
		Handler:     c.commandUserInfo,
		MinArgs:     1,
		Permission:  "permissions.user.info",
		Description: "command.description.user.info",
		Example:     "command.example.user.info",
	})
	c.addCommand("perms_user_check_permission", CommandInfo{
		Handler:     c.commandUserPermissionCheck,
		MinArgs:     2,
		Permission:  "permissions.user.permission.check",
		Description: "command.description.user.permission.check",
		Example:     "command.example.user.permission.check",
	})
	c.addCommand("perms_user_check_group", CommandInfo{
		Handler:     c.commandUserGroupCheck,
		MinArgs:     2,
		Permission:  "permissions.user.group.check",
		Description: "command.description.user.group.check",
		Example:     "command.example.user.group.check",
	})

	c.addCommand("perms_user_permissions", CommandInfo{
		Handler:     c.commandUserPermissions,
		MinArgs:     1,
		Permission:  "permissions.user.permission.info",
		Description: "command.description.user.permission.info",
		Example:     "command.example.user.permission.info",
	})
	c.addCommand("perms_user_groups", CommandInfo{
		Handler:     c.commandUserGroups,
		MinArgs:     1,
		Permission:  "permissions.user.group.info",
		Description: "command.description.user.group.info",
		Example:     "command.example.user.group.info",
	})

	c.addCommand("perms_user_add_permission", CommandInfo{
		Handler:     c.commandUserPermissionAdd,
		MinArgs:     3,
		Permission:  "permissions.user.permission.add",
		Description: "command.description.user.permission.add",
		Example:     "command.example.user.permission.add",
	})
	c.addCommand("perms_user_set_permission", CommandInfo{
		Handler:     c.commandUserPermissionSet,
		MinArgs:     3,
		Permission:  "permissions.user.permission.set",
		Description: "command.description.user.permission.set",
		Example:     "command.example.user.permission.set",
	})
	c.addCommand("perms_user_remove_permission", CommandInfo{
		Handler:     c.commandUserPermissionRemove,
		MinArgs:     2,
		Permission:  "permissions.user.permission.remove",
		Description: "command.description.user.permission.remove",
		Example:     "command.example.user.permission.remove",
	})

	c.addCommand("perms_user_add_group", CommandInfo{
		Handler:     c.commandUserGroupAdd,
		MinArgs:     3,
		Permission:  "permissions.user.group.add",
		Description: "command.description.user.group.add",
		Example:     "command.example.user.group.add",
	})
	c.addCommand("perms_user_set_group", CommandInfo{
		Handler:     c.commandUserGroupSet,
		MinArgs:     3,
		Permission:  "permissions.user.group.set",
		Description: "command.description.user.group.set",
		Example:     "command.example.user.group.set",
	})
	c.addCommand("perms_user_remove_group", CommandInfo{
		Handler:     c.commandUserGroupRemove,
		MinArgs:     2,
		Permission:  "permissions.user.group.remove",
		Description: "command.description.user.group.remove",
		Example:     "command.example.user.group.remove",
	})
}

func (c *Commands) commandUserInfo(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		target, err := findPlayerTarget(ctx)
		if err != nil {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, err.Error())
			return
		}

		if target.Slot == -1 {
			s2sdk.ReplyToCommand(ctx.ConCommandContext, ctx.Caller.Slot, "[Permissions Info]\nServer has full permissions :)")
			return
		}

		var immunity int32
		permissions.GetImmunity(target.SteamID64, &immunity)

		var _perms, _groups []string
		permissions.DumpPermissions(target.SteamID64, &_perms)
		permissions.GetUserGroups(target.SteamID64, &_groups)

		var cookies []string
		var cookieValues []any
		permissions.GetAllCookies(target.SteamID64, &cookies, &cookieValues)

		perms := utils.ParseGroupsOrNodes(_perms, nil)
		groups := utils.ParseGroupsOrNodes(_groups, nil)

		//s2sdk.PrintToConsole(ctx.Caller.Slot, "[Permissions Dump]\n")

		s2sdk.PrintToConsole(ctx.Caller.Slot, "[Permissions Info]\n ")

		s2sdk.PrintToConsole(ctx.Caller.Slot, fmt.Sprintf("Name: %s[%d]", target.Name, target.Slot))
		s2sdk.PrintToConsole(ctx.Caller.Slot, fmt.Sprintf("SteamID64: %d", target.SteamID64))
		s2sdk.PrintToConsole(ctx.Caller.Slot, fmt.Sprintf("Immunity: %d\n ", immunity))

		var data [][]any
		for _, group := range groups {
			var expires string
			if group.Timestamp > 0 {
				expires = utils.FormatTimestamp(group.Timestamp, c.config.Settings.FormatTimestamp)
			} else {
				expires = "never"
			}

			data = append(data, []any{group.Name, expires, group.Timestamp})
		}

		utils.PrintConsoleTableToClient(ctx.Caller.Slot, "Groups",
			[]string{"name", "expires", "ts"},
			data,
		)

		clear(data)
		data = data[:0]
		for i := range perms {
			status := "allow"
			name := perms[i].Name
			if strings.HasPrefix(name, "-") {
				status = "disallow"
				name = name[1:]
			}

			var expires string
			if perms[i].Timestamp > 0 {
				expires = utils.FormatTimestamp(perms[i].Timestamp, c.config.Settings.FormatTimestamp)
			} else {
				expires = "never"
			}

			data = append(data, []any{name, status, expires, perms[i].Timestamp})
		}

		utils.PrintConsoleTableToClient(ctx.Caller.Slot, "Permissions",
			[]string{"name", "state", "expires", "ts"},
			data,
		)

		clear(data)
		data = data[:0]
		for i := 0; i < len(cookies); i++ {
			data = append(data, []any{cookies[i], cookieValues[i]})
		}

		utils.PrintConsoleTableToClient(ctx.Caller.Slot, "Cookies",
			[]string{"name", "value"},
			data,
		)
	}()

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandUserList(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		data := make([][]any, 0, 64)
		for i := int32(0); i < 64; i++ {
			if !utils.IsValidClient(i) {
				continue
			}

			var groups []string
			var steamID64 = s2sdk.GetClientSteamID64(i)
			permissions.GetUserGroups(steamID64, &groups)

			data = append(data, []any{
				i,
				s2sdk.GetClientName(i),
				steamID64,
				strings.Join(groups, ", "),
			})
		}

		utils.PrintConsoleTableToClient(ctx.Caller.Slot, "Players List",
			[]string{"#", "Name", "SteamID64", "Groups"},
			data,
		)
	}()

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandUserPermissions(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		target, err := findPlayerTarget(ctx)
		if err != nil {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, err.Error())
			return
		}

		var perms []string
		permissions.DumpPermissions(target.SteamID64, &perms)

		parsed := utils.ParseGroupsOrNodes(perms, nil)

		data := make([][]any, 0)
		for i := range parsed {
			data = append(data, []any{
				parsed[i].Name,
				func() string {
					if parsed[i].Timestamp == 0 {
						return "-"
					}
					return utils.FormatTimestamp(parsed[i].Timestamp, c.config.Settings.FormatTimestamp)
				}(),
				parsed[i].Timestamp,
			})
		}

		utils.PrintConsoleTableToClient(ctx.Caller.Slot, s2sdk.GetClientName(target.Slot)+"`s permissions",
			[]string{"permission", "expires at", "ts"},
			data,
		)
	}()

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandUserGroups(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		target, err := findPlayerTarget(ctx)
		if err != nil {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, err.Error())
			return
		}

		var groups []string
		permissions.GetUserGroups(target.SteamID64, &groups)

		parsed := utils.ParseGroupsOrNodes(groups, nil)

		if ctx.ConCommandContext == s2sdk.ConCommandContext_Console {
			data := make([][]any, 0)
			for i := range parsed {
				data = append(data, []any{
					parsed[i].Name,
					func() string {
						if parsed[i].Timestamp == 0 {
							return "-"
						}
						return utils.FormatTimestamp(parsed[i].Timestamp, c.config.Settings.FormatTimestamp)
					}(),
					parsed[i].Timestamp,
				})
			}

			utils.PrintConsoleTableToClient(ctx.Caller.Slot, s2sdk.GetClientName(target.Slot)+"`s groups",
				[]string{"group", "expires at", "ts"},
				data,
			)
		} else {
			s2sdk.PrintToChat(ctx.Caller.Slot, s2sdk.GetClientName(target.Slot)+"`s groups:")
			for _, group := range parsed {
				var msg string
				if group.Timestamp > 0 {
					msg = trs.GetTranslation("user.loaded.active_groups.temporary", ctx.Caller.Language, group.Name,
						utils.FormatTimestamp(group.Timestamp, c.config.Settings.FormatTimestamp),
					)
				} else {
					msg = trs.GetTranslation("user.loaded.active_groups.permanent", ctx.Caller.Language, group.Name)
				}

				s2sdk.PrintToChatColored(ctx.Caller.Slot, msg)
			}
		}
	}()

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandUserPermissionCheck(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		target, err := findPlayerTarget(ctx)
		if err != nil {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, err.Error())
			return
		}

		var data GroupOrNode
		data.Name = ctx.Args[1]

		if strings.TrimSpace(data.Name) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				trs.GetTranslation("command.error.permission.empty", ctx.Caller.Language,
					trs.GetTranslation(ctx.Info.Example, ctx.Caller.Language)),
			)
			return
		}

		data.Result = permissions.HasPermission(target.SteamID64, data.Name)

		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, translateUserCommandResult(ctx.Caller, target, data, CmdSuccessResult{}))

		c.log.Infof("Command: CheckPermission | Result: %v | Actor: %s[%d] | Target: %s[%d]",
			data.Result,
			ctx.Caller.Name, ctx.Caller.SteamID64,
			target.Name, target.SteamID64,
		)
	}()

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandUserGroupCheck(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		target, err := findPlayerTarget(ctx)
		if err != nil {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, err.Error())
			return
		}

		groupName := ctx.Args[1]

		if strings.TrimSpace(groupName) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				trs.GetTranslation("command.error.group.empty_name", ctx.Caller.Language,
					trs.GetTranslation(ctx.Info.Example, ctx.Caller.Language)),
			)
			return
		}

		var timestamp int64
		result := permissions.HasGroupExtended(target.SteamID64, groupName, &timestamp)

		var msg string
		switch result {
		case permissions.Status_PermanentGroup:
			msg = trs.GetTranslation("command.group.check.has", ctx.Caller.Language, target.Name, target.Slot, groupName)
		case permissions.Status_TemporalGroup:
			msg = trs.GetTranslation("command.group.check.has_temporary", ctx.Caller.Language, target.Name, target.Slot, groupName,
				utils.FormatTimestamp(timestamp, c.config.Settings.FormatTimestamp),
			)
		case permissions.Status_GroupNotDefined:
			msg = trs.GetTranslation("command.error.user.group.not_defined", ctx.Caller.Language, target.Name, target.Slot, groupName)
		case permissions.Status_GroupNotFound:
			msg = trs.GetTranslation("command.error.group.not_found", ctx.Caller.Language, groupName)
		default:
			msg = trs.GetTranslation("command.error.unknown", ctx.Caller.Language)
		}

		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, msg)

		c.log.Infof("Command: CheckGroup | Result: %v | Actor: %s[%d] | Target: %s[%d] | Group: %s",
			result,
			ctx.Caller.Name, ctx.Caller.SteamID64,
			target.Name, target.SteamID64,
			groupName,
		)
	}()

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandUserPermissionAdd(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		target, err := findPlayerTarget(ctx)
		if err != nil {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, err.Error())
			return
		}

		var data GroupOrNode
		data.Name = ctx.Args[1]

		if strings.TrimSpace(data.Name) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				trs.GetTranslation("command.error.permission.empty", ctx.Caller.Language,
					trs.GetTranslation(ctx.Info.Example, ctx.Caller.Language)),
			)
			return
		}

		arg2, err := strconv.ParseUint(ctx.Args[2], 10, 32)
		if err != nil {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, trs.GetTranslation("command.error.invalid_time", ctx.Caller.Language, ctx.Info.Example))
			return
		}

		data.Duration = utils.TimeToSeconds(arg2, c.config.Settings.TimeMode)

		permCleared := data.Name
		if strings.HasPrefix(data.Name, "-") {
			data.NewStatus = permissions.Status_Disallow
			permCleared = permCleared[1:]
		} else {
			data.NewStatus = permissions.Status_Allow
		}

		data.OldStatus = permissions.HasPermissionExtended(target.SteamID64, data.Name, false, &data.PermSource, &data.OldTimestamp)

		if data.Duration > 0 {
			switch data.OldTimestamp {
			case -1:
				data.NewTimestamp = time.Now().Unix() + int64(data.Duration)
			case 0:
				utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
					trs.GetTranslation("command.error.user.permission.already_granted", ctx.Caller.Language,
						target.Name, target.Slot, data.Name,
					),
				)
				return
			default:
				data.NewTimestamp = data.OldTimestamp + int64(data.Duration)
			}
		}

		data.Result = permissions.AddPermission(c.pluginID, target.SteamID64, data.Name, data.NewTimestamp, false)

		var action permissions.Action
		if data.OldStatus > 2 {
			action = permissions.Action_Add
		} else {
			action = permissions.Action_Replace
		}

		msg := translateUserCommandResult(ctx.Caller, target, data,
			CmdSuccessResult{
				Key: utils.BuildTranslationPermKey("command.permission", action, data.OldStatus, data.NewStatus, data.OldTimestamp, data.NewTimestamp),
				Args: func() []any {
					args := []any{target.Name, target.Slot, permCleared}
					if data.Duration > 0 {
						args = []any{target.Name, target.Slot, permCleared, utils.FormatTimestamp(data.NewTimestamp, c.config.Settings.FormatTimestamp)}
					}
					return args
				}(),
			},
		)

		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, msg)

		c.log.Infof("Command: AddPermission | Result: %v | Actor: %s[%d] | Target: %s[%d] | Perm: %s | Duration: %d | Ts: %d",
			data.Result,
			ctx.Caller.Name, ctx.Caller.SteamID64,
			target.Name, target.SteamID64,
			data.Name,
			data.Duration,
			data.NewTimestamp,
		)
	}()

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandUserPermissionSet(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		target, err := findPlayerTarget(ctx)
		if err != nil {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, err.Error())
			return
		}

		var data GroupOrNode
		data.Name = ctx.Args[1]

		if strings.TrimSpace(data.Name) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				trs.GetTranslation("command.error.permission.empty", ctx.Caller.Language,
					trs.GetTranslation(ctx.Info.Example, ctx.Caller.Language)),
			)
			return
		}

		permCleared := data.Name
		if strings.HasPrefix(data.Name, "-") {
			data.NewStatus = permissions.Status_Disallow
			permCleared = permCleared[1:]
		} else {
			data.NewStatus = permissions.Status_Allow
		}

		argCount := len(ctx.Args)
		if argCount > 2 {
			arg2, err := strconv.ParseUint(ctx.Args[2], 10, 32)
			if err != nil {
				utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, trs.GetTranslation("command.error.invalid_time", ctx.Caller.Language, ctx.Info.Example))
				return
			}

			data.Duration = utils.TimeToSeconds(arg2, c.config.Settings.TimeMode)
		}

		data.OldStatus = permissions.HasPermissionExtended(target.SteamID64, data.Name, false, &data.PermSource, &data.OldTimestamp)

		if data.Duration > 0 {
			data.NewTimestamp = time.Now().Unix() + int64(data.Duration)
		}

		if data.OldStatus == data.NewStatus && data.OldTimestamp == data.NewTimestamp && data.PermSource == permissions.PermSource_User {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				trs.GetTranslation("command.error.user.permission.already_granted", ctx.Caller.Language,
					target.Name, target.Slot, data.Name,
				),
			)
			return
		}

		data.Result = permissions.SetPermission(c.pluginID, target.SteamID64, data.Name, data.NewTimestamp, false)

		var action permissions.Action
		if data.OldStatus > 2 {
			action = permissions.Action_Add
		} else {
			action = permissions.Action_Replace
		}

		msg := translateUserCommandResult(ctx.Caller, target, data,
			CmdSuccessResult{
				Key: utils.BuildTranslationPermKey("command.permission", action, data.OldStatus, data.NewStatus, data.OldTimestamp, data.NewTimestamp),
				Args: func() []any {
					args := []any{target.Name, target.Slot, permCleared}
					if data.Duration > 0 {
						args = []any{target.Name, target.Slot, permCleared, utils.FormatTimestamp(data.NewTimestamp, c.config.Settings.FormatTimestamp)}
					}
					return args
				}(),
			},
		)

		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, msg)

		c.log.Infof("Command: SetPermission | Result: %v | Actor: %s[%d] | Target: %s[%d] | Perm: %s | Duration: %d | Ts: %d",
			data.Result,
			ctx.Caller.Name, ctx.Caller.SteamID64,
			target.Name, target.SteamID64,
			data.Name,
			data.Duration,
			data.NewTimestamp,
		)
	}()

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandUserPermissionRemove(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		target, err := findPlayerTarget(ctx)
		if err != nil {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, err.Error())
			return
		}

		var data GroupOrNode
		data.Name = ctx.Args[1]

		if strings.TrimSpace(data.Name) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				trs.GetTranslation("command.error.permission.empty", ctx.Caller.Language,
					trs.GetTranslation(ctx.Info.Example, ctx.Caller.Language)),
			)
			return
		}

		permCleared := data.Name
		if strings.HasPrefix(data.Name, "-") {
			permCleared = permCleared[1:]
		}

		data.OldStatus = permissions.HasPermission(target.SteamID64, data.Name)

		result := permissions.RemovePermission(c.pluginID, target.SteamID64, data.Name, false, false)

		msg := translateUserCommandResult(ctx.Caller, target, data,
			CmdSuccessResult{
				Key:  utils.BuildTranslationPermKey("command.permission", permissions.Action_Remove, data.OldStatus, permissions.Status_PermNotFound, 0, 0),
				Args: []any{target.Name, target.Slot, permCleared},
			},
		)

		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, msg)

		c.log.Infof("Command: RemovePermission | Result: %v | Actor: %s[%d] | Target: %s[%d] | Perm: %s",
			result,
			ctx.Caller.Name, ctx.Caller.SteamID64,
			target.Name, target.SteamID64,
			data.Name,
		)
	}()

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandUserGroupAdd(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		target, err := findPlayerTarget(ctx)
		if err != nil {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, err.Error())
			return
		}

		var data GroupOrNode
		data.Name = ctx.Args[1]

		if strings.TrimSpace(data.Name) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				trs.GetTranslation("command.error.group.empty_name", ctx.Caller.Language,
					trs.GetTranslation(ctx.Info.Example, ctx.Caller.Language)),
			)
			return
		}

		arg2, err := strconv.ParseUint(ctx.Args[2], 10, 32)
		if err != nil {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, trs.GetTranslation("command.error.invalid_time", ctx.Caller.Language, ctx.Info.Example))
			return
		}

		permissions.HasGroupExtended(target.SteamID64, data.Name, &data.OldTimestamp)

		data.Duration = utils.TimeToSeconds(arg2, c.config.Settings.TimeMode)
		if data.Duration > 0 {

			switch data.OldTimestamp {
			case -1:
				data.NewTimestamp = time.Now().Unix() + int64(data.Duration)
			case 0:
				utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, trs.GetTranslation("command.error.user.group.already_exist", ctx.Caller.Language, target.Name, target.Slot, data.Name))
				return
			default:
				data.NewTimestamp = data.OldTimestamp + int64(data.Duration)
			}
		}

		data.Result = permissions.AddGroup(c.pluginID, target.SteamID64, data.Name, data.NewTimestamp, false)

		msg := translateUserCommandResult(ctx.Caller, target, data,
			CmdSuccessResult{
				Key: utils.BuildTranslationGroupKey("command.group", permissions.Action_Add, data.OldTimestamp, data.NewTimestamp),
				Args: func() []any {
					args := []any{target.Name, target.Slot, data.Name}
					if data.Duration > 0 {
						args = []any{target.Name, target.Slot, data.Name, utils.FormatTimestamp(data.NewTimestamp, c.config.Settings.FormatTimestamp)}
					}
					return args
				}(),
			},
		)

		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, msg)

		c.log.Infof("Command: AddGroup | Result: %v | Actor: %s[%d] | Target: %s[%d] | Group: %s | Duration: %d | Ts: %d",
			data.Result,
			ctx.Caller.Name, ctx.Caller.SteamID64,
			target.Name, target.SteamID64,
			data.Name,
			data.Duration,
			data.NewTimestamp,
		)
	}()

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandUserGroupSet(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		target, err := findPlayerTarget(ctx)
		if err != nil {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, err.Error())
			return
		}

		var data GroupOrNode
		data.Name = ctx.Args[1]

		if strings.TrimSpace(data.Name) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				trs.GetTranslation("command.error.group.empty_name", ctx.Caller.Language,
					trs.GetTranslation(ctx.Info.Example, ctx.Caller.Language)),
			)
			return
		}

		arg2, err := strconv.ParseUint(ctx.Args[2], 10, 32)
		if err != nil {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, trs.GetTranslation("command.error.invalid_time", ctx.Caller.Language, ctx.Info.Example))
			return
		}

		permissions.HasGroupExtended(target.SteamID64, data.Name, &data.OldTimestamp)

		data.Duration = utils.TimeToSeconds(arg2, c.config.Settings.TimeMode)
		if data.Duration > 0 {
			data.NewTimestamp = time.Now().Unix() + int64(data.Duration)
		}

		data.Result = permissions.AddGroup(c.pluginID, target.SteamID64, data.Name, data.NewTimestamp, false)

		msg := translateUserCommandResult(ctx.Caller, target, data,
			CmdSuccessResult{
				Key: utils.BuildTranslationGroupKey("command.group", permissions.Action_Replace, data.OldTimestamp, data.NewTimestamp),
				Args: func() []any {
					args := []any{target.Name, target.Slot, data.Name}
					if data.Duration > 0 {
						args = []any{target.Name, target.Slot, data.Name, utils.FormatTimestamp(data.NewTimestamp, c.config.Settings.FormatTimestamp)}
					}
					return args
				}(),
			},
		)

		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, msg)

		c.log.Infof("Command: SetGroup | Result: %v | Actor: %s[%d] | Target: %s[%d] | Perm: %s | Duration: %d | Ts: %d",
			data.Result,
			ctx.Caller.Name, ctx.Caller.SteamID64,
			target.Name, target.SteamID64,
			data.Name,
			data.Duration,
			data.NewTimestamp,
		)
	}()

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandUserGroupRemove(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		target, err := findPlayerTarget(ctx)
		if err != nil {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, err.Error())
			return
		}

		var data GroupOrNode
		data.Name = ctx.Args[1]

		if strings.TrimSpace(data.Name) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				trs.GetTranslation("command.error.group.empty_name", ctx.Caller.Language,
					trs.GetTranslation(ctx.Info.Example, ctx.Caller.Language)),
			)
			return
		}

		data.Result = permissions.RemoveGroup(c.pluginID, target.SteamID64, data.Name, false)

		msg := translateUserCommandResult(ctx.Caller, target, data,
			CmdSuccessResult{
				Key:  utils.BuildTranslationGroupKey("command.group", permissions.Action_Remove, 0, 0),
				Args: []any{target.Name, target.Slot, data.Name},
			},
		)

		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, msg)

		c.log.Infof("Command: RemoveGroup | Result: %v | Actor: %s[%d] | Target: %s[%d] | Group: %s",
			data.Result,
			ctx.Caller.Name, ctx.Caller.SteamID64,
			target.Name, target.SteamID64,
			data.Name,
		)
	}()

	return s2sdk.ResultType_Continue
}

func translateUserCommandResult(caller, target *Player, data GroupOrNode, additional CmdSuccessResult) string {
	var msg string

	switch data.Result {
	case permissions.Status_Success:
		msg = trs.GetTranslation(additional.Key, caller.Language, additional.Args...)
	case permissions.Status_Allow:
		msg = trs.GetTranslation("command.error.user.permission.allow", caller.Language, target.Name, target.Slot, data.Name)
	case permissions.Status_Disallow:
		msg = trs.GetTranslation("command.error.user.permission.disallow", caller.Language, target.Name, target.Slot, data.Name)
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
		msg = trs.GetTranslation("command.error.user.group.already_exist", caller.Language, target.Name, target.Slot, data.Name)
	case permissions.Status_PermAlreadyGranted:
		msg = trs.GetTranslation("command.error.user.permission.already_granted", caller.Language, target.Name, target.Slot, data.Name)
	default:
		msg = trs.GetTranslation("command.error.unknown", caller.Language)
	}

	return msg
}
