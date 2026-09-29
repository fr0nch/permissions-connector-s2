package internal

import (
	"permissions-connector-s2/internal/pkg/config"
	"permissions-connector-s2/internal/pkg/utils"

	s2sdk "github.com/fr0nch/go-plugify-s2sdk/v2"
	t "github.com/fr0nch/go-plugify-translations"
)

func (p *PermsConnectorPlugin) SendTranslatedNotifyPlayer(playerSlot int32, flags config.NotifyFlag, translationKey, language string, args ...any) {
	var msg string
	if flags&(config.Console|config.Chat) != 0 {
		msg = t.TranslateFormat(language, translationKey, args)
		p.SendNotifyPlayer(playerSlot, config.Console|config.Chat, msg)
	}
	if flags&config.Hint != 0 {
		msg = t.TranslateFormat(language, translationKey+".hint", args)
		p.SendNotifyPlayer(playerSlot, config.Hint, msg)
	}
}

func (p *PermsConnectorPlugin) SendNotifyPlayer(playerSlot int32, flags config.NotifyFlag, text string) {
	if flags&config.Console != 0 {
		msg := utils.RemoveTags(text)
		// Sending a null string to the console will crash the server (bug S2SDK).
		if len(msg) == 0 {
			return
		}

		s2sdk.PrintToConsole(playerSlot, msg)
	}
	if flags&config.Chat != 0 {
		s2sdk.PrintToChatColored(playerSlot, text)
	}
	if flags&config.Hint != 0 {
		s2sdk.PrintCentreHtml(playerSlot, utils.RemoveTags(text), p.config.Settings.HintDuration)
	}
}

func (p *PermsConnectorPlugin) DisplayActiveGroups(playerSlot int32, groups []string) {
	parsedGroups := utils.ParseGroupsOrNodes(groups, p.config.Settings.HiddenGroups)
	if len(parsedGroups) == 0 {
		return
	}

	notifyMode := p.config.Settings.NotifyMode.ToFlag()
	language := s2sdk.GetClientLanguageCode(playerSlot)

	if notifyMode&(config.Console|config.Chat) != 0 {
		mode := config.Console | config.Chat
		p.SendTranslatedNotifyPlayer(playerSlot, mode, "user.loaded.active_groups.header", language)

		for _, group := range parsedGroups {
			var text string

			if group.Timestamp > 0 {
				text = t.TranslateFormat(language, "user.loaded.active_groups.temporary", []any{group.Name,
					utils.FormatTimestamp(group.Timestamp, p.config.Settings.FormatTimestamp)})
			} else {
				text = t.TranslateFormat(language, "user.loaded.active_groups.permanent", []any{group.Name})
			}

			p.SendNotifyPlayer(playerSlot, mode, text)
		}
	}
}
