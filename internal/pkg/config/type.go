package config

type ConfigsType interface {
	Settings
}

type Config struct {
	Settings *Settings `yaml:"settings"`
}

type NotifyFlag uint8

const (
	Console NotifyFlag = 1 << iota
	Chat
	Hint
)

type NotifyMode struct {
	Console bool `yaml:"console" line_comment:""`
	Chat    bool `yaml:"chat" line_comment:""`
	Hint    bool `yaml:"hint" line_comment:""`
}

func (n NotifyMode) ToFlag() NotifyFlag {
	var f NotifyFlag

	if n.Console {
		f |= Console
	}
	if n.Chat {
		f |= Chat
	}
	if n.Hint {
		f |= Hint
	}

	return f
}

type TimeMode uint8

const (
	Seconds TimeMode = iota
	Minutes
	Hours
	Days
)

type Settings struct {
	NotifyMode   NotifyMode `yaml:"notify_mode" head_comment:"Controls how players are notified when permissions or groups are granted or removed.\n- console: Sends a message to the player's console.\n- chat: Sends a message in the in-game chat.\n- hint: Displays a hint notification on the player's screen."`
	HintDuration int32      `yaml:"hint_duration" head_comment:"Duration (in seconds) for which hint notifications are displayed on the player's screen.\n- Only applies if hint notifications are enabled in notify_mode."`

	TimeMode        TimeMode `yaml:"time_mode" head_comment:"Specifies the time format used for displaying durations.\n \n- Possible options:\n  |=> 0 (Seconds)\n  |=> 1 (Minutes)\n  |=> 2 (Hours)\n  |=> 3 (Days)\n \n- Affects how temporary durations are shown to players."`
	AbbreviateTime  bool     `yaml:"abbreviate_time" head_comment:"# Whether to use abbreviated time units (e.g., '5 min.' instead of '5 minutes').\n- If set to true, units will be: d., h., min., sec.\n- If set to false, units will use full localized forms with proper declension."`
	FormatTimestamp string   `yaml:"format_timestamp" head_comment:"The format used for displaying date and time.\n- Uses the Go formatting standard (t.Format): https://pkg.go.dev/time#pkg-constants\n- Example: 2006-01-02 15:04:05"`

	DisplayActiveGroups bool     `yaml:"display_active_groups" head_comment:"Determines whether to display the list of active groups to the player when they join the server."`
	HiddenGroups        []string `yaml:"hidden_groups" head_comment:"A list of groups that will be hidden from the player (e.g., to hide the default group)."`
}
