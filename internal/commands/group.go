package commands

import (
	"fmt"
	"permissions-connector-s2/internal/pkg/utils"
	"strconv"
	"strings"

	permissions "github.com/fr0nch/go-plugify-permissions"
	t "github.com/fr0nch/go-plugify-translations"

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
	c.addCommand("perms_group_set_option", CommandInfo{
		Handler:     c.commandGroupOptionSet,
		MinArgs:     3,
		Permission:  "permissions.group.option.set",
		Description: "command.description.group.option.set",
		Example:     "command.example.group.option.set",
	})
	c.addCommand("perms_group_option", CommandInfo{
		Handler:     c.commandGroupOption,
		MinArgs:     2,
		Permission:  "permissions.group.option",
		Description: "command.description.group.option",
		Example:     "command.example.group.option",
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

		if strings.TrimSpace(groupName) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				t.TranslateFormat(ctx.Caller.Language, "command.error.group.empty_name", []any{t.Translate(ctx.Caller.Language, ctx.Info.Example)}),
			)
			return
		}

		var priority int32
		result := permissions.GetPriorityGroup(groupName, &priority)
		if result == permissions.Status_GroupNotFound {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, t.TranslateFormat(ctx.Caller.Language, "command.error.group.not_found", []any{groupName}))
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

		utils.PrintToConsole(ctx.Caller.Slot, "[Group Info]\n ")

		utils.PrintToConsole(ctx.Caller.Slot, fmt.Sprintf("Name: %s", groupName))
		utils.PrintToConsole(ctx.Caller.Slot, fmt.Sprintf("Priority: %d", priority))
		utils.PrintToConsole(ctx.Caller.Slot, fmt.Sprintf("Parent: %s\n ", parentName))

		var data [][]any
		for i := range permissionsGroup {
			status := "allow"
			name := permissionsGroup[i]
			if strings.HasPrefix(name, "-") {
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

		if strings.TrimSpace(groupName) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				t.TranslateFormat(ctx.Caller.Language, "command.error.group.empty_name", []any{t.Translate(ctx.Caller.Language, ctx.Info.Example)}),
			)
			return
		}

		priority, err := strconv.ParseInt(ctx.Args[1], 10, 32)
		if err != nil {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				t.TranslateFormat(ctx.Caller.Language, "command.error.invalid_number", []any{t.Translate(ctx.Caller.Language, ctx.Info.Example)}),
			)
			return
		}

		parentName := ctx.Args[2]
		perms := strings.Split(ctx.Args[3], " ")

		result := permissions.CreateGroup(c.pluginID, groupName, perms, int32(priority), parentName, false)

		var msg string
		switch result {
		case permissions.Status_Success:
			msg = t.TranslateFormat(ctx.Caller.Language, "command.creategroup.success", []any{groupName})
		case permissions.Status_GroupAlreadyExist:
			msg = t.TranslateFormat(ctx.Caller.Language, "command.error.creategroup.already_exist", []any{groupName})
		case permissions.Status_ParentGroupNotFound:
			msg = t.TranslateFormat(ctx.Caller.Language, "command.error.group.parent_not_found", []any{parentName})
		default:
			msg = t.Translate(ctx.Caller.Language, "command.error.unknown")
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

		if strings.TrimSpace(groupName) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				t.TranslateFormat(ctx.Caller.Language, "command.error.group.empty_name", []any{t.Translate(ctx.Caller.Language, ctx.Info.Example)}),
			)
			return
		}

		result := permissions.DeleteGroup(c.pluginID, groupName, false)

		var msg string
		switch result {
		case permissions.Status_Success:
			msg = t.TranslateFormat(ctx.Caller.Language, "command.deletegroup.success", []any{groupName})
		case permissions.Status_GroupNotFound:
			msg = t.TranslateFormat(ctx.Caller.Language, "command.error.group.not_found", []any{groupName})
		default:
			msg = t.Translate(ctx.Caller.Language, "command.error.unknown")
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

func (c *Commands) commandGroupOptionSet(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		groupName := ctx.Args[0]
		optionName := ctx.Args[1]
		value := ctx.Args[2]

		if strings.TrimSpace(groupName) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				t.TranslateFormat(ctx.Caller.Language, "command.error.group.empty_name", []any{t.Translate(ctx.Caller.Language, ctx.Info.Example)}),
			)
			return
		}

		result := permissions.SetOptionGroup(c.pluginID, groupName, optionName, value, false)

		var msg string
		switch result {
		case permissions.Status_Success:
			msg = t.TranslateFormat(ctx.Caller.Language, "command.group.option.set.success", []any{groupName, optionName, value})
		case permissions.Status_GroupNotFound:
			msg = t.TranslateFormat(ctx.Caller.Language, "command.error.group.not_found", []any{groupName})
		case permissions.Status_InvalidOptionName:
			msg = t.Translate(ctx.Caller.Language, "command.error.option.invalid")
		default:
			msg = t.Translate(ctx.Caller.Language, "command.error.unknown")
		}

		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, msg)

		c.log.Infof("Command: SetOptionGroup | Result: %v | Caller: %s[%d] | Group: %s | Option: %s | Value: %s",
			result,
			ctx.Caller.Name, ctx.Caller.SteamID64,
			groupName,
			optionName,
			value,
		)
	}()

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandGroupOption(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		groupName := ctx.Args[0]
		optionName := ctx.Args[1]

		if strings.TrimSpace(groupName) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				t.TranslateFormat(ctx.Caller.Language, "command.error.group.empty_name", []any{t.Translate(ctx.Caller.Language, ctx.Info.Example)}),
			)
			return
		}

		var value any
		result := permissions.GetOptionGroup(groupName, optionName, &value)

		var msg string
		switch result {
		case permissions.Status_Success:
			msg = t.TranslateFormat(ctx.Caller.Language, "command.group.option.info", []any{groupName, optionName, fmt.Sprint(value)})
		case permissions.Status_OptionNotFound:
			msg = t.TranslateFormat(ctx.Caller.Language, "command.error.group.option.not_found", []any{groupName, optionName})
		case permissions.Status_GroupNotFound:
			msg = t.TranslateFormat(ctx.Caller.Language, "command.error.group.not_found", []any{groupName})
		default:
			msg = t.Translate(ctx.Caller.Language, "command.error.unknown")
		}

		utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, msg)
	}()

	return s2sdk.ResultType_Continue
}

func (c *Commands) commandGroupParentSet(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		childName := ctx.Args[0]
		parentName := ctx.Args[1]

		if strings.TrimSpace(childName) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				t.TranslateFormat(ctx.Caller.Language, "command.error.group.empty_name", []any{t.Translate(ctx.Caller.Language, ctx.Info.Example)}),
			)
			return
		}

		if strings.TrimSpace(parentName) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				t.TranslateFormat(ctx.Caller.Language, "command.error.group.empty_name", []any{t.Translate(ctx.Caller.Language, ctx.Info.Example)}),
			)
			return
		}

		result := permissions.SetParent(c.pluginID, childName, parentName, false)

		var msg string
		switch result {
		case permissions.Status_Success:
			msg = t.TranslateFormat(ctx.Caller.Language, "command.group.parent.set.success", []any{childName, parentName})
		case permissions.Status_ChildGroupNotFound:
			msg = t.TranslateFormat(ctx.Caller.Language, "command.error.group.child_not_found", []any{childName})
		case permissions.Status_ParentGroupNotFound:
			msg = t.TranslateFormat(ctx.Caller.Language, "command.error.group.parent_not_found", []any{parentName})
		case permissions.Status_GroupHierarchyCycle:
			msg = t.TranslateFormat(ctx.Caller.Language, "command.error.group.hierarchy_cycle", []any{childName, parentName})
		default:
			msg = t.Translate(ctx.Caller.Language, "command.error.unknown")
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

		if strings.TrimSpace(groupName) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				t.TranslateFormat(ctx.Caller.Language, "command.error.group.empty_name", []any{t.Translate(ctx.Caller.Language, ctx.Info.Example)}),
			)
			return
		}

		var parentName string
		result := permissions.GetParent(groupName, &parentName)
		if result == permissions.Status_ChildGroupNotFound {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, t.TranslateFormat(ctx.Caller.Language, "command.error.group.not_found", []any{groupName}))
			return
		}

		var msg string
		if parentName == "" {
			msg = t.TranslateFormat(ctx.Caller.Language, "command.group.parent.info.none", []any{groupName})
		} else {
			msg = t.TranslateFormat(ctx.Caller.Language, "command.group.parent.info", []any{groupName, parentName})
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

		if strings.TrimSpace(childName) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				t.TranslateFormat(ctx.Caller.Language, "command.error.group.empty_name", []any{t.Translate(ctx.Caller.Language, ctx.Info.Example)}),
			)
			return
		}

		if strings.TrimSpace(parentName) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				t.TranslateFormat(ctx.Caller.Language, "command.error.group.empty_name", []any{t.Translate(ctx.Caller.Language, ctx.Info.Example)}),
			)
			return
		}

		result := permissions.HasParentGroup(childName, parentName)

		var msg string
		switch result {
		case permissions.Status_Allow:
			msg = t.TranslateFormat(ctx.Caller.Language, "command.group.parent.check.has", []any{childName, parentName})
		case permissions.Status_Disallow:
			msg = t.TranslateFormat(ctx.Caller.Language, "command.group.parent.check.has_not", []any{childName, parentName})
		case permissions.Status_ChildGroupNotFound:
			msg = t.TranslateFormat(ctx.Caller.Language, "command.error.group.child_not_found", []any{childName})
		case permissions.Status_ParentGroupNotFound:
			msg = t.TranslateFormat(ctx.Caller.Language, "command.error.group.parent_not_found", []any{parentName})
		default:
			msg = t.Translate(ctx.Caller.Language, "command.error.unknown")
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

		if strings.TrimSpace(group) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				t.TranslateFormat(ctx.Caller.Language, "command.error.group.empty_name", []any{t.Translate(ctx.Caller.Language, ctx.Info.Example)}),
			)
			return
		}

		var permissionsGroup []string

		result := permissions.DumpPermissionsGroup(group, &permissionsGroup)
		if result == permissions.Status_GroupNotFound {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext, t.TranslateFormat(ctx.Caller.Language, "command.error.group.not_found", []any{group}))
			return
		}

		rows := make([][]any, 0, len(permissionsGroup))
		for i := range permissionsGroup {
			status := "allow"
			if strings.HasPrefix(permissionsGroup[i], "-") {
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

		if strings.TrimSpace(groupName) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				t.TranslateFormat(ctx.Caller.Language, "command.error.group.empty_name", []any{t.Translate(ctx.Caller.Language, ctx.Info.Example)}),
			)
			return
		}

		if strings.TrimSpace(permName) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				t.TranslateFormat(ctx.Caller.Language, "command.error.permission.empty", []any{t.Translate(ctx.Caller.Language, ctx.Info.Example)}),
			)
			return
		}

		result := permissions.HasPermissionGroup(groupName, permName)

		var msg string
		switch result {
		case permissions.Status_Allow:
			msg = t.TranslateFormat(ctx.Caller.Language, "command.group.permission.allow", []any{groupName, permName})
		case permissions.Status_Disallow:
			msg = t.TranslateFormat(ctx.Caller.Language, "command.group.permission.disallow", []any{groupName, permName})
		case permissions.Status_PermNotFound:
			msg = t.Translate(ctx.Caller.Language, "command.error.permission.not_found")
		case permissions.Status_GroupNotFound:
			msg = t.TranslateFormat(ctx.Caller.Language, "command.error.group.not_found", []any{groupName})
		default:
			msg = t.Translate(ctx.Caller.Language, "command.error.unknown")
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

func groupPermissionActionMsg(lang string, prefix string, result permissions.Status, wildcard bool, disallow bool, group string, perm string) string {
	switch result {
	case permissions.Status_Success:
		if perm == "*" {
			key := prefix + ".all.allow"
			if disallow {
				key = prefix + ".all.disallow"
			}
			return t.TranslateFormat(lang, key, []any{group})
		}

		key := prefix + ".allow"
		switch {
		case wildcard && disallow:
			key = prefix + ".wc.disallow"
		case wildcard:
			key = prefix + ".wc.allow"
		case disallow:
			key = prefix + ".disallow"
		}
		return t.TranslateFormat(lang, key, []any{group, perm})
	case permissions.Status_GroupNotFound:
		return t.TranslateFormat(lang, "command.error.group.not_found", []any{group})
	case permissions.Status_PermAlreadyGranted:
		return t.TranslateFormat(lang, "command.error.group.permission.already_granted", []any{group, perm})
	case permissions.Status_PermConflict:
		key := "command.error.group.permission.conflict.disallow"
		if disallow {
			key = "command.error.group.permission.conflict.allow"
		}
		return t.TranslateFormat(lang, key, []any{group, perm})
	case permissions.Status_InvalidPermission:
		return t.Translate(lang, "command.error.permission.invalid")
	default:
		return t.Translate(lang, "command.error.unknown")
	}
}

func (c *Commands) commandGroupPermissionAdd(ctx *CommandContext) s2sdk.ResultType {
	go func() {
		groupName := ctx.Args[0]
		permName := ctx.Args[1]

		if strings.TrimSpace(groupName) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				t.TranslateFormat(ctx.Caller.Language, "command.error.group.empty_name", []any{t.Translate(ctx.Caller.Language, ctx.Info.Example)}),
			)
			return
		}

		if strings.TrimSpace(permName) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				t.TranslateFormat(ctx.Caller.Language, "command.error.permission.empty", []any{t.Translate(ctx.Caller.Language, ctx.Info.Example)}),
			)
			return
		}

		permCleared := permName
		disallow := strings.HasPrefix(permName, "-")
		if disallow {
			permCleared = permCleared[1:]
		}

		wildcard := strings.HasSuffix(permCleared, ".*")
		if wildcard {
			permCleared = strings.TrimSuffix(permCleared, ".*")
		}

		result := permissions.AddPermissionGroup(c.pluginID, groupName, permName, false)

		msg := groupPermissionActionMsg(ctx.Caller.Language, "command.group.permission.add", result, wildcard, disallow, groupName, permCleared)

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

		if strings.TrimSpace(groupName) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				t.TranslateFormat(ctx.Caller.Language, "command.error.group.empty_name", []any{t.Translate(ctx.Caller.Language, ctx.Info.Example)}),
			)
			return
		}

		if strings.TrimSpace(permName) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				t.TranslateFormat(ctx.Caller.Language, "command.error.permission.empty", []any{t.Translate(ctx.Caller.Language, ctx.Info.Example)}),
			)
			return
		}

		permCleared := permName
		disallow := strings.HasPrefix(permName, "-")
		if disallow {
			permCleared = permCleared[1:]
		}

		wildcard := strings.HasSuffix(permCleared, ".*")
		if wildcard {
			permCleared = strings.TrimSuffix(permCleared, ".*")
		}

		result := permissions.SetPermissionGroup(c.pluginID, groupName, permName, false)

		msg := groupPermissionActionMsg(ctx.Caller.Language, "command.group.permission.set", result, wildcard, disallow, groupName, permCleared)

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

		if strings.TrimSpace(groupName) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				t.TranslateFormat(ctx.Caller.Language, "command.error.group.empty_name", []any{t.Translate(ctx.Caller.Language, ctx.Info.Example)}),
			)
			return
		}

		if strings.TrimSpace(permName) == "" {
			utils.ReplyToCommandCleared(ctx.Caller.Slot, ctx.ConCommandContext,
				t.TranslateFormat(ctx.Caller.Language, "command.error.permission.empty", []any{t.Translate(ctx.Caller.Language, ctx.Info.Example)}),
			)
			return
		}

		permCleared := permName
		if strings.HasPrefix(permName, "-") {
			permCleared = permCleared[1:]
		}

		var oldState permissions.Status
		if permCleared == "*" {
			oldState = permissions.HasPermissionGroupExtended(groupName, permCleared, true)
		}

		result := permissions.RemovePermissionGroup(c.pluginID, groupName, permName, false, false)

		var msg string

		switch result {
		case permissions.Status_Success:
			switch {
			case permCleared == "*" && oldState == permissions.Status_Disallow:
				msg = t.TranslateFormat(ctx.Caller.Language, "command.group.permission.remove.all.disallow", []any{groupName})
			case permCleared == "*":
				msg = t.TranslateFormat(ctx.Caller.Language, "command.group.permission.remove.all.allow", []any{groupName})
			default:
				msg = t.TranslateFormat(ctx.Caller.Language, "command.group.permission.remove.success", []any{groupName, permCleared})
			}
		case permissions.Status_GroupNotFound:
			msg = t.TranslateFormat(ctx.Caller.Language, "command.error.group.not_found", []any{groupName})
		case permissions.Status_PermNotFound:
			msg = t.Translate(ctx.Caller.Language, "command.error.permission.not_found")
		case permissions.Status_InvalidPermission:
			msg = t.Translate(ctx.Caller.Language, "command.error.permission.invalid")
		default:
			msg = t.Translate(ctx.Caller.Language, "command.error.unknown")
		}

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
