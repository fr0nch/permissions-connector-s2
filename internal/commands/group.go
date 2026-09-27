package commands

import (
	"fmt"
	trs "permissions-connector-s2/internal/pkg/translations"
	"permissions-connector-s2/internal/pkg/utils"
	"strconv"
	"strings"

	permissions "github.com/fr0nch/go-plugify-permissions"

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
	c.addCommand("perms_group_info", CommandInfo{
		Handler:     c.commandGroupInfo,
		MinArgs:     1,
		Permission:  "permissions.group.info",
		Description: "command.description.group.info",
		Example:     "command.example.group.info",
	})
	c.addCommand("perms_group_permissions", CommandInfo{
		Handler:     c.commandGroupPermissions,
		MinArgs:     1,
		Permission:  "permissions.group.permissions",
		Description: "command.description.group.permissions",
		Example:     "command.example.group.permissions",
	})
	c.addCommand("perms_group_check_permission", CommandInfo{
		Handler:     c.commandGroupPermissionCheck,
		MinArgs:     2,
		Permission:  "permissions.group.permission.check",
		Description: "command.description.group.permission.check",
		Example:     "command.example.group.permission.check",
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
	c.addCommand("perms_group_parent_info", CommandInfo{
		Handler:     c.commandGroupParentInfo,
		MinArgs:     1,
		Permission:  "permissions.group.parent.info",
		Description: "command.description.group.parent.info",
		Example:     "command.example.group.parent.info",
	})
	c.addCommand("perms_group_parent_check", CommandInfo{
		Handler:     c.commandGroupParentCheck,
		MinArgs:     2,
		Permission:  "permissions.group.parent.check",
		Description: "command.description.group.parent.check",
		Example:     "command.example.group.parent.check",
	})
	c.addCommand("perms_group_add_permission", CommandInfo{
		Handler:     c.commandGroupPermissionAdd,
		MinArgs:     2,
		Permission:  "permissions.group.permission.add",
		Description: "command.description.group.permission.add",
		Example:     "command.example.group.permission.add",
	})
	c.addCommand("perms_group_set_permission", CommandInfo{
		Handler:     c.commandGroupPermissionSet,
		MinArgs:     2,
		Permission:  "permissions.group.permission.set",
		Description: "command.description.group.permission.set",
		Example:     "command.example.group.permission.set",
	})
	c.addCommand("perms_group_remove_permission", CommandInfo{
		Handler:     c.commandGroupPermissionRemove,
		MinArgs:     2,
		Permission:  "permissions.group.permission.remove",
		Description: "command.description.group.permission.remove",
		Example:     "command.example.group.permission.remove",
	})
}

func (c *Commands) commandGroupInfo(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		groupName := ctx.Args[0]

		var priority int32
		result := permissions.GetPriorityGroup(groupName, &priority)
		if result == permissions.Status_GroupNotFound {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, trs.GetTranslation("command.error.group.not_found", ctx.Caller.Language, groupName))
			return
		}

		var parentName string
		permissions.GetParent(groupName, &parentName)
		if parentName == "" {
			parentName = "-"
		}

		var permissionsGroup []string
		permissions.DumpPermissionsGroup(groupName, &permissionsGroup)

		var optionNames []string
		var optionValues []any
		permissions.GetAllOptionsGroup(groupName, &optionNames, &optionValues)

		s2sdk.PrintToConsole(ctx.Caller.Slot, "[Group Info]\n ")

		s2sdk.PrintToConsole(ctx.Caller.Slot, fmt.Sprintf("Name: %s", groupName))
		s2sdk.PrintToConsole(ctx.Caller.Slot, fmt.Sprintf("Priority: %d", priority))
		s2sdk.PrintToConsole(ctx.Caller.Slot, fmt.Sprintf("Parent: %s\n ", parentName))

		var data [][]any
		for i := range permissionsGroup {
			status := "allow"
			name := permissionsGroup[i]
			if name[0] == '-' {
				status = "disallow"
				name = name[1:]
			}

			data = append(data, []any{name, status})
		}

		utils.PrintConsoleTableToClient(ctx.Caller.Slot, "Permissions",
			[]string{"name", "state"},
			data,
		)

		clear(data)
		data = data[:0]
		for i := 0; i < len(optionNames); i++ {
			data = append(data, []any{optionNames[i], optionValues[i]})
		}

		utils.PrintConsoleTableToClient(ctx.Caller.Slot, "Options",
			[]string{"name", "value"},
			data,
		)
	}()

	return s2sdk.ResultType_Continue
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

func (c *Commands) commandGroupParentInfo(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		groupName := ctx.Args[0]

		var parentName string
		result := permissions.GetParent(groupName, &parentName)
		if result == permissions.Status_ChildGroupNotFound {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, trs.GetTranslation("command.error.group.not_found", ctx.Caller.Language, groupName))
			return
		}

		var msg string
		if parentName == "" {
			msg = trs.GetTranslation("command.group.parent.info.none", ctx.Caller.Language, groupName)
		} else {
			msg = trs.GetTranslation("command.group.parent.info", ctx.Caller.Language, groupName, parentName)
		}

		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, msg)

		c.log.Infof("Command: GroupParentInfo | Caller: %s[%d] | Group: %s | Parent: %s",
			ctx.Caller.Name, ctx.Caller.SteamID64,
			groupName,
			parentName,
		)
	}()

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandGroupParentCheck(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		childName := ctx.Args[0]
		parentName := ctx.Args[1]

		result := permissions.HasParentGroup(childName, parentName)

		var msg string
		switch result {
		case permissions.Status_Allow:
			msg = trs.GetTranslation("command.group.parent.check.has", ctx.Caller.Language, childName, parentName)
		case permissions.Status_Disallow:
			msg = trs.GetTranslation("command.group.parent.check.has_not", ctx.Caller.Language, childName, parentName)
		case permissions.Status_ChildGroupNotFound:
			msg = trs.GetTranslation("command.error.group.child_not_found", ctx.Caller.Language, childName)
		case permissions.Status_ParentGroupNotFound:
			msg = trs.GetTranslation("command.error.group.parent_not_found", ctx.Caller.Language, parentName)
		default:
			msg = trs.GetTranslation("command.error.unknown", ctx.Caller.Language)
		}

		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, msg)

		c.log.Infof("Command: GroupParentCheck | Result: %v | Caller: %s[%d] | Child: %s | Parent: %s",
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

func (c *Commands) commandGroupPermissionCheck(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		groupName := ctx.Args[0]
		permName := ctx.Args[1]

		result := permissions.HasPermissionGroup(groupName, permName)

		var msg string
		switch result {
		case permissions.Status_Allow:
			msg = trs.GetTranslation("command.group.permission.allow", ctx.Caller.Language, groupName, permName)
		case permissions.Status_Disallow:
			msg = trs.GetTranslation("command.group.permission.disallow", ctx.Caller.Language, groupName, permName)
		case permissions.Status_PermNotFound:
			msg = trs.GetTranslation("command.error.permission.not_found", ctx.Caller.Language)
		case permissions.Status_GroupNotFound:
			msg = trs.GetTranslation("command.error.group.not_found", ctx.Caller.Language, groupName)
		default:
			msg = trs.GetTranslation("command.error.unknown", ctx.Caller.Language)
		}

		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, msg)

		c.log.Infof("Command: CheckPermissionGroup | Result: %v | Caller: %s[%d] | Group: %s | Perm: %s",
			result,
			ctx.Caller.Name, ctx.Caller.SteamID64,
			groupName,
			permName,
		)
	}()

	return s2sdk.ResultType_Continue
}

// buildGroupPermissionActionKey resolves the translation key for the result of
// AddPermissionGroup/SetPermissionGroup, given the "command.group.permission.add"
// or "command.group.permission.set" prefix.
func buildGroupPermissionActionKey(prefix string, result permissions.Status, wildcard bool, disallow bool) string {
	switch result {
	case permissions.Status_Success:
		switch {
		case wildcard && disallow:
			return prefix + ".wc.disallow"
		case wildcard:
			return prefix + ".wc.allow"
		case disallow:
			return prefix + ".disallow"
		default:
			return prefix + ".allow"
		}
	case permissions.Status_GroupNotFound:
		return "command.error.group.not_found"
	case permissions.Status_PermAlreadyGranted:
		return "command.error.group.permission.already_granted"
	default:
		return "command.error.unknown"
	}
}

// buildGroupPermissionRemoveKey resolves the translation key for the result of RemovePermissionGroup.
func buildGroupPermissionRemoveKey(result permissions.Status) string {
	switch result {
	case permissions.Status_Success:
		return "command.group.permission.remove.success"
	case permissions.Status_GroupNotFound:
		return "command.error.group.not_found"
	case permissions.Status_PermNotFound:
		return "command.error.permission.not_found"
	default:
		return "command.error.unknown"
	}
}

func (c *Commands) commandGroupPermissionAdd(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		groupName := ctx.Args[0]
		permName := ctx.Args[1]

		permCleared := permName
		disallow := permName[0] == '-'
		if disallow {
			permCleared = permCleared[1:]
		}

		wildcard := strings.HasSuffix(permCleared, ".*")
		if wildcard {
			permCleared = strings.TrimSuffix(permCleared, ".*")
		}

		result := permissions.AddPermissionGroup(c.pluginID, groupName, permName, false)

		key := buildGroupPermissionActionKey("command.group.permission.add", result, wildcard, disallow)
		msg := trs.GetTranslation(key, ctx.Caller.Language, groupName, permCleared)

		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, msg)

		c.log.Infof("Command: AddPermissionGroup | Result: %v | Caller: %s[%d] | Group: %s | Perm: %s",
			result,
			ctx.Caller.Name, ctx.Caller.SteamID64,
			groupName,
			permName,
		)
	}()

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandGroupPermissionSet(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		groupName := ctx.Args[0]
		permName := ctx.Args[1]

		permCleared := permName
		disallow := permName[0] == '-'
		if disallow {
			permCleared = permCleared[1:]
		}

		wildcard := strings.HasSuffix(permCleared, ".*")
		if wildcard {
			permCleared = strings.TrimSuffix(permCleared, ".*")
		}

		result := permissions.SetPermissionGroup(c.pluginID, groupName, permName, false)

		key := buildGroupPermissionActionKey("command.group.permission.set", result, wildcard, disallow)
		msg := trs.GetTranslation(key, ctx.Caller.Language, groupName, permCleared)

		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, msg)

		c.log.Infof("Command: SetPermissionGroup | Result: %v | Caller: %s[%d] | Group: %s | Perm: %s",
			result,
			ctx.Caller.Name, ctx.Caller.SteamID64,
			groupName,
			permName,
		)
	}()

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandGroupPermissionRemove(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		groupName := ctx.Args[0]
		permName := ctx.Args[1]

		permCleared := permName
		if permName[0] == '-' {
			permCleared = permCleared[1:]
		}

		result := permissions.RemovePermissionGroup(c.pluginID, groupName, permName, false, false)

		key := buildGroupPermissionRemoveKey(result)
		msg := trs.GetTranslation(key, ctx.Caller.Language, groupName, permCleared)

		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, msg)

		c.log.Infof("Command: RemovePermissionGroup | Result: %v | Caller: %s[%d] | Group: %s | Perm: %s",
			result,
			ctx.Caller.Name, ctx.Caller.SteamID64,
			groupName,
			permName,
		)
	}()

	return s2sdk.ResultType_Continue
}
