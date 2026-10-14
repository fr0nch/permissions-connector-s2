package internal

import (
	"permissions-connector-s2/internal/pkg/utils"
	"strings"

	permissions "github.com/fr0nch/go-plugify-permissions"

	s2sdk "github.com/fr0nch/go-plugify-s2sdk/v2"
)

func (p *PermsConnectorPlugin) RegisterEvents() {
	s2sdk.OnServerActivate_Register(p.OnServerActivated)
	s2sdk.OnMapEnd_Register(p.OnMapEnd)

	s2sdk.OnClientAuthenticated_Register(p.OnClientAuthenticated)
	s2sdk.OnClientDisconnect_Register(p.OnClientDisconnect)

	permissions.OnUserLoaded_Register(p.OnUserLoaded)

	permissions.OnUserPermissionChange_Register(p.OnUserPermissionChange)
	permissions.OnUserGroupChange_Register(p.OnUserGroupChange)
	permissions.OnUserCookieChange_Register(p.OnUserCookieChange)

	permissions.OnPermissionExpiration_Register(p.OnPermissionExpiration)
	permissions.OnGroupExpiration_Register(p.OnGroupExpiration)
}

func (p *PermsConnectorPlugin) OnServerActivated() {
	go permissions.LoadGroups(p.PluginID, false)
	//p.log.Debugf("OnServerActivated called.")
}

func (p *PermsConnectorPlugin) OnMapEnd() {
	users := permissions.DumpUsersList()
	for _, user := range users {
		if permissions.UserExists(user) == permissions.PlayerState_Offline {
			permissions.DeleteUser(p.PluginID, user)
		}
	}
	//p.log.Debugf("OnMapEnd called.")
}

func (p *PermsConnectorPlugin) OnClientAuthenticated(playerSlot int32, steamID uint64) {
	p.log.Debugf("Client %s authenticated", s2sdk.GetClientName(playerSlot))

	if s2sdk.IsFakeClient(playerSlot) || s2sdk.IsClientSourceTV(playerSlot) {
		return
	}

	go permissions.LoadUser(p.PluginID, steamID, s2sdk.GetClientName(playerSlot), false, false)
}

func (p *PermsConnectorPlugin) OnClientDisconnect(playerSlot int32, reason s2sdk.NetworkDisconnectionReason) {
	if s2sdk.IsFakeClient(playerSlot) || s2sdk.IsClientSourceTV(playerSlot) {
		return
	}

	s2sdk.KillsTimer(p.loadedUserTimers[playerSlot])
	p.loadedUserTimers[playerSlot] = 0

	permissions.DeleteUser(p.PluginID, s2sdk.GetClientSteamID64(playerSlot))
}

func (p *PermsConnectorPlugin) OnUserLoaded(pluginID int64, targetID uint64, playerState permissions.PlayerState) {
	if !p.config.Settings.DisplayActiveGroups {
		return
	}

	s2sdk.QueueTaskForNextFrame(p.userLoadedTask, []any{targetID})
}

func (p *PermsConnectorPlugin) userLoadedTask(userData []any) {
	targetID := userData[0].(uint64)
	playerSlot := utils.GetPlayerSlotBySteamID64(targetID)

	var groups []string
	permissions.GetUserGroups(targetID, &groups)

	if len(groups) == 0 {
		return
	}

	p.log.Debugf("Loaded user '%d' | Groups: '%v'", targetID, groups)

	if !s2sdk.IsClientInGame(playerSlot) || s2sdk.GetClientTeam(playerSlot) == s2sdk.CSTeam_None {
		p.loadedUserTimers[playerSlot] = s2sdk.CreateTimer(1.0, p.timerPostUserLoaded, s2sdk.TimerFlag_Repeat, []any{targetID, groups})
		return
	}

	p.DisplayActiveGroups(playerSlot, groups)
}

func (p *PermsConnectorPlugin) timerPostUserLoaded(timer uint32, userData []any) {
	clientSteamID := userData[0].(uint64)
	playerSlot := utils.GetPlayerSlotBySteamID64(clientSteamID)

	if !s2sdk.IsClientInGame(playerSlot) || s2sdk.GetClientTeam(playerSlot) == s2sdk.CSTeam_None {
		return
	}

	p.DisplayActiveGroups(playerSlot, userData[1].([]string))

	s2sdk.KillsTimer(s2sdk.TimerId(timer))
	p.loadedUserTimers[playerSlot] = 0
}

func (p *PermsConnectorPlugin) OnUserPermissionChange(pluginID int64, action permissions.Action, targetID uint64, perm string, oldState permissions.Status, newState permissions.Status, oldTimestamp int64, newTimestamp int64) {
	playerSlot := utils.GetPlayerSlotBySteamID64(targetID)

	if !utils.IsValidClient(playerSlot) {
		return
	}

	language := s2sdk.GetClientLanguageCode(playerSlot)
	keyPrefix := "user.permission"
	if perm == "*" {
		keyPrefix += ".all"
	}
	translationKey := utils.BuildTranslationPermKey(keyPrefix, action, oldState, newState, oldTimestamp, newTimestamp)

	switch action {
	case permissions.Action_Add:
		p.log.Debugf("Added permission '%s' for user '%d' on %s", perm, targetID, utils.FormatTimestamp(newTimestamp, p.config.Settings.FormatTimestamp))

		if newTimestamp > 0 {
			p.SendTranslatedNotifyPlayer(playerSlot, p.config.Settings.NotifyMode.ToFlag(),
				translationKey, language,
				perm, utils.FormatTimestamp2(newTimestamp, p.config.Settings.FormatTimestamp, language), utils.TranslateTime(language, newTimestamp, true, p.config.Settings.AbbreviateTime),
			)
		} else {
			p.SendTranslatedNotifyPlayer(playerSlot, p.config.Settings.NotifyMode.ToFlag(),
				translationKey, language, perm,
			)
		}

	case permissions.Action_Replace, permissions.Action_ReplaceToWC:
		var diff int64
		var duration string

		if oldTimestamp != 0 && newTimestamp != 0 {
			if oldTimestamp > newTimestamp {
				diff = oldTimestamp - newTimestamp
			} else {
				diff = newTimestamp - oldTimestamp
			}
		}

		duration = utils.TranslateTime(language, diff, false, p.config.Settings.AbbreviateTime)

		if action == permissions.Action_Replace {
			p.log.Debugf("Replace permission '%s'(oldState: %d [%d] | newState: %d [%d]) duration on %s for user '%d'", perm, oldState, oldTimestamp, newState, newTimestamp, utils.FormatTimestamp(newTimestamp, p.config.Settings.FormatTimestamp), targetID)
		} else {
			temp := strings.TrimSuffix(perm, ".*")
			p.log.Debugf("Replace '%s' to wildcard permission '%s' for user '%d'", temp, perm, targetID)
			perm = temp
		}

		if newTimestamp > 0 {
			p.SendTranslatedNotifyPlayer(playerSlot, p.config.Settings.NotifyMode.ToFlag(),
				translationKey, language,
				perm, utils.FormatTimestamp2(oldTimestamp, p.config.Settings.FormatTimestamp, language), utils.FormatTimestamp2(newTimestamp, p.config.Settings.FormatTimestamp, language), duration,
			)
		} else {
			p.SendTranslatedNotifyPlayer(playerSlot, p.config.Settings.NotifyMode.ToFlag(),
				translationKey, language,
				perm, utils.FormatTimestamp2(oldTimestamp, p.config.Settings.FormatTimestamp, language),
			)
		}
	case permissions.Action_Remove:
		p.log.Debugf("Remove permission '%s' for user '%d'", perm, targetID)
		p.SendTranslatedNotifyPlayer(playerSlot, p.config.Settings.NotifyMode.ToFlag(),
			translationKey, language, perm,
		)
	}

	return
}

func (p *PermsConnectorPlugin) OnUserGroupChange(pluginID int64, action permissions.Action, targetID uint64, group string, oldTimestamp int64, newTimestamp int64) {
	playerSlot := utils.GetPlayerSlotBySteamID64(targetID)

	if !utils.IsValidClient(playerSlot) {
		return
	}

	language := s2sdk.GetClientLanguageCode(playerSlot)
	translationKey := utils.BuildTranslationGroupKey("user.group", action, oldTimestamp, newTimestamp)

	switch action {
	case permissions.Action_Add:
		p.log.Debugf("Added group '%s' for user '%d' on %s", group, targetID, utils.FormatTimestamp(newTimestamp, p.config.Settings.FormatTimestamp))
		if newTimestamp > 0 {
			p.SendTranslatedNotifyPlayer(playerSlot, p.config.Settings.NotifyMode.ToFlag(),
				translationKey, language,
				group, utils.FormatTimestamp(newTimestamp, p.config.Settings.FormatTimestamp),
			)
		} else {
			p.SendTranslatedNotifyPlayer(playerSlot, p.config.Settings.NotifyMode.ToFlag(),
				translationKey, language, group,
			)
		}
	case permissions.Action_Replace:
		p.log.Debugf("Set group '%s' for user '%d' on %s", group, targetID, utils.FormatTimestamp(newTimestamp, p.config.Settings.FormatTimestamp))
		if newTimestamp > 0 {
			p.SendTranslatedNotifyPlayer(playerSlot, p.config.Settings.NotifyMode.ToFlag(),
				translationKey, language,
				group, utils.FormatTimestamp(newTimestamp, p.config.Settings.FormatTimestamp),
			)
		} else {
			p.SendTranslatedNotifyPlayer(playerSlot, p.config.Settings.NotifyMode.ToFlag(),
				translationKey, language, group,
			)
		}
	case permissions.Action_Remove:
		p.log.Debugf("Remove group '%s' for user '%d'", group, targetID)
		p.SendTranslatedNotifyPlayer(playerSlot, p.config.Settings.NotifyMode.ToFlag(),
			translationKey, language, group,
		)
	}

	return
}

func (p *PermsConnectorPlugin) OnUserCookieChange(pluginID int64, targetID uint64, name string, cookie any) {
	playerSlot := utils.GetPlayerSlotBySteamID64(targetID)

	if !utils.IsValidClient(playerSlot) {
		return
	}

	// ...
	return
}

func (p *PermsConnectorPlugin) OnPermissionExpiration(targetID uint64, perm string, state permissions.Status) {
	playerSlot := utils.GetPlayerSlotBySteamID64(targetID)

	if !utils.IsValidClient(playerSlot) {
		return
	}

	p.log.Debugf("Permission '%s'[state: %v] expiration from user '%d'", perm, state, targetID)

	language := s2sdk.GetClientLanguageCode(playerSlot)

	var stateKey string

	if state == permissions.Status_Disallow {
		stateKey = ".disallow"
	} else {
		stateKey = ".allow"
	}

	keyPrefix := "user.permission"
	if strings.TrimPrefix(perm, "-") == "*" {
		keyPrefix += ".all"
	}

	p.SendTranslatedNotifyPlayer(playerSlot, p.config.Settings.NotifyMode.ToFlag(), keyPrefix+".expired"+stateKey, language, perm)
}

func (p *PermsConnectorPlugin) OnGroupExpiration(targetID uint64, group string) {
	playerSlot := utils.GetPlayerSlotBySteamID64(targetID)

	if !utils.IsValidClient(playerSlot) {
		return
	}

	p.log.Debugf("Group '%s' expiration from user '%d'", group, targetID)

	language := s2sdk.GetClientLanguageCode(playerSlot)
	p.SendTranslatedNotifyPlayer(playerSlot, p.config.Settings.NotifyMode.ToFlag(),
		"user.group.expired", language, group,
	)
}
