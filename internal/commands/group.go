package commands

import (
	"permissions-connector-s2/internal/pkg/permissions"
	trs "permissions-connector-s2/internal/pkg/translations"
	"permissions-connector-s2/internal/pkg/utils"
	"strconv"
	"strings"

	s2sdk "github.com/fr0nch/go-plugify-s2sdk/v2"
)

func registerGroupCommands(c *Commands) {
	c.addCommand("perms_groups", CommandInfo{
		Handler:     c.commandListGroups,
		MinArgs:     0,
		Permission:  "permissions.listgroups",
		Description: "command.description.listgroups",
		Example:     "command.example.listgroups",
	})
	c.addCommand("perms_group_permissions", CommandInfo{
		Handler:     c.commandGroupPermissions,
		MinArgs:     1,
		Permission:  "permissions.group.permissions",
		Description: "command.description.group.permissions",
		Example:     "command.example.group.permissions",
	})
	c.addCommand("perms_creategroup", CommandInfo{
		Handler:     c.commandCreateGroup,
		MinArgs:     4,
		Permission:  "permissions.creategroup",
		Description: "command.description.creategroup",
		Example:     "command.example.creategroup",
	})
	c.addCommand("perms_deletegroup", CommandInfo{
		Handler:     c.commandDeleteGroup,
		MinArgs:     1,
		Permission:  "permissions.deletegroup",
		Description: "command.description.deletegroup",
		Example:     "command.example.deletegroup",
	})
	c.addCommand("perms_group_parent_set", CommandInfo{
		Handler:     c.commandGroupParentSet,
		MinArgs:     2,
		Permission:  "permissions.group.parent.set",
		Description: "command.description.group.parent.set",
		Example:     "command.example.group.parent.set",
	})
}

func (c *Commands) commandCreateGroup(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		groupName := ctx.Args[0]

		priority, err := strconv.ParseInt(ctx.Args[1], 10, 32)
		if err != nil {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				trs.GetTranslation("command.error.invalid_number", ctx.Caller.Language, ctx.Info.Example),
			)
			return
		}

		parentName := ctx.Args[2]
		perms := strings.Split(ctx.Args[3], " ")

		result := permissions.CreateGroup(c.pluginID, groupName, perms, int32(priority), parentName, false)

		var msg string
		switch result {
		case permissions.Status_Success:
			msg = trs.GetTranslation("command.creategroup.success", ctx.Caller.Language, groupName)
		case permissions.Status_GroupAlreadyExist:
			msg = trs.GetTranslation("command.error.creategroup.already_exist", ctx.Caller.Language, groupName)
		case permissions.Status_ParentGroupNotFound:
			msg = trs.GetTranslation("command.error.group.parent_not_found", ctx.Caller.Language, parentName)
		default:
			msg = trs.GetTranslation("command.error.unknown", ctx.Caller.Language)
		}

		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, msg)

		c.log.Infof("Command: CreateGroup | Result: %v | Caller: %s[%d] | Group: %s | Priority: %d | Parent: %s | Perms: %v",
			result,
			ctx.Caller.Name, ctx.Caller.SteamID64,
			groupName,
			priority,
			parentName,
			perms,
		)
	}()

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandDeleteGroup(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		groupName := ctx.Args[0]

		result := permissions.DeleteGroup(c.pluginID, groupName, false)

		var msg string
		switch result {
		case permissions.Status_Success:
			msg = trs.GetTranslation("command.deletegroup.success", ctx.Caller.Language, groupName)
		case permissions.Status_GroupNotFound:
			msg = trs.GetTranslation("command.error.group.not_found", ctx.Caller.Language, groupName)
		default:
			msg = trs.GetTranslation("command.error.unknown", ctx.Caller.Language)
		}

		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, msg)

		c.log.Infof("Command: DeleteGroup | Result: %v | Caller: %s[%d] | Group: %s",
			result,
			ctx.Caller.Name, ctx.Caller.SteamID64,
			groupName,
		)
	}()

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandGroupParentSet(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		childName := ctx.Args[0]
		parentName := ctx.Args[1]

		result := permissions.SetParent(c.pluginID, childName, parentName, false)

		var msg string
		switch result {
		case permissions.Status_Success:
			msg = trs.GetTranslation("command.group.parent.set.success", ctx.Caller.Language, childName, parentName)
		case permissions.Status_ChildGroupNotFound:
			msg = trs.GetTranslation("command.error.group.child_not_found", ctx.Caller.Language, childName)
		case permissions.Status_ParentGroupNotFound:
			msg = trs.GetTranslation("command.error.group.parent_not_found", ctx.Caller.Language, parentName)
		default:
			msg = trs.GetTranslation("command.error.unknown", ctx.Caller.Language)
		}

		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, msg)

		c.log.Infof("Command: SetParent | Result: %v | Caller: %s[%d] | ChildGroup: %s | ParentGroup: %s",
			result,
			ctx.Caller.Name, ctx.Caller.SteamID64,
			childName,
			parentName,
		)
	}()

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandListGroups(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		groups := permissions.GetAllGroups()

		rows := make([][]any, 0, len(groups))
		for i := range groups {
			rows = append(rows, []any{groups[i]})
		}

		utils.PrintConsoleTableToClient(ctx.Caller.Slot, "Available Groups",
			[]string{"name"},
			rows,
		)
	}()

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandGroupPermissions(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		var group = ctx.Args[0]
		var permissionsGroup []string

		result := permissions.DumpPermissionsGroup(group, &permissionsGroup)
		if result == permissions.Status_GroupNotFound {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, trs.GetTranslation("command.error.group.not_found", ctx.Caller.Language, group))
			return
		}

		rows := make([][]any, 0, len(permissionsGroup))
		for i := range permissionsGroup {
			status := "allow"
			if permissionsGroup[i][0] == '-' {
				status = "disallow"
				permissionsGroup[i] = permissionsGroup[i][1:]
			}

			rows = append(rows, []any{permissionsGroup[i], status})
		}

		utils.PrintConsoleTableToClient(ctx.Caller.Slot, group+"`s permissions",
			[]string{"name", "status"},
			rows,
		)
	}()

	return s2sdk.ResultType_Continue
}
