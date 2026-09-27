package utils

import (
	"bytes"
	"cmp"
	"fmt"
	"io"
	"permissions-connector-s2/internal/pkg/config"
	"permissions-connector-s2/internal/pkg/translations"
	"slices"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
	"unicode/utf8"

	permissions "github.com/fr0nch/go-plugify-permissions"

	s2sdk "github.com/fr0nch/go-plugify-s2sdk/v2"
)

func FormatTimestamp(ts int64, fTs string) string {
	t := time.Unix(ts, 0)
	return t.Format(fTs)
}

func FormatTimestamp2(ts int64, fTs string, lang string) string {
	if ts == 0 {
		return translations.GetTranslation("time.forever", lang)
	}
	return FormatTimestamp(ts, fTs)
}

func TranslateTime(lang string, ts int64, compute bool, abbr bool) string {
	if ts == 0 {
		return translations.GetTranslation("time.forever", lang)
	}

	if compute {
		ts = ts - time.Now().Unix()
	}

	values := []int64{
		ts / 86400,
		(ts % 86400) / 3600,
		(ts % 3600) / 60,
		ts % 60,
	}

	keys := []string{"time.days", "time.hours", "time.minutes", "time.seconds"}
	decline := make([]string, 3)
	if !abbr {
		decline[0], decline[1], decline[2] = ".1", ".2", ".5"
	}

	var result string

	for i := range values {
		if values[i] > 0 {
			result += fmt.Sprintf("%d %s", values[i], DeclineWord(values[i], lang, keys[i]+decline[0], keys[i]+decline[1], keys[i]+decline[2]))
			if i >= 0 && i < len(values)-1 {
				result += " "
			}
		}
	}

	return result
}

func TimeToSeconds(time uint64, timeMode config.TimeMode) uint64 {
	switch timeMode {
	case config.Minutes:
		return time * 60
	case config.Hours:
		return time * 3600
	case config.Days:
		return time * 86400
	default:
		return time
	}
}

type GroupsOrNodes struct {
	Name      string
	Timestamp int64
}

func ParseGroupsOrNodes(inputs []string, filter []string) []GroupsOrNodes {
	result := make([]GroupsOrNodes, 0, len(inputs))

	for _, input := range inputs {
		input = strings.TrimSpace(input)
		item := GroupsOrNodes{}

		i := strings.LastIndex(input, " ")
		if i == -1 {
			item.Name = input
		} else {
			namePart := input[:i]
			tsPart := input[i+1:]

			if ts, err := strconv.ParseInt(tsPart, 10, 64); err == nil {
				item.Timestamp = ts
				item.Name = strings.TrimSpace(namePart)
			} else {
				item.Name = input
			}
		}

		if slices.Contains(filter, item.Name) {
			continue
		}

		result = append(result, item)
	}

	return result
}

func BuildTranslationGroupKey(key string, action permissions.Action, oldTimestamp int64, newTimestamp int64) string {
	switch action {
	case permissions.Action_Add:
		key += ".add"
	case permissions.Action_Replace:
		key += ".replace"
	case permissions.Action_Remove:
		key += ".remove"
		return key
	}

	if newTimestamp > 0 {
		key += ".temporary"
	} else {
		key += ".permanent"
	}

	if action == permissions.Action_Replace {
		if oldTimestamp != 0 && newTimestamp != 0 {
			if oldTimestamp > newTimestamp {
				key += ".subtracted"
			} else {
				key += ".added"
			}
		}
	}

	return key
}

func BuildTranslationPermKey(key string, action permissions.Action, oldStatus permissions.Status, newStatus permissions.Status, oldTimestamp int64, newTimestamp int64) string {
	switch action {
	case permissions.Action_Add:
		key += ".add"
	case permissions.Action_Replace:
		key += ".replace"
	case permissions.Action_ReplaceToWC:
		key += ".replace.wc"
	case permissions.Action_Remove:
		key += ".remove"

		if oldStatus == permissions.Status_Disallow {
			key += ".disallow"
		} else {
			key += ".allow"
		}

		return key
	}

	if newStatus == permissions.Status_Disallow {
		key += ".disallow"
	} else {
		key += ".allow"
	}

	if (oldStatus == permissions.Status_Disallow && newStatus == permissions.Status_Allow) ||
		(oldStatus == permissions.Status_Allow && newStatus == permissions.Status_Disallow) {
		key += ".change_state"
	}

	if newTimestamp == oldTimestamp && action != permissions.Action_Add {
		return key
	}

	if newTimestamp > 0 {
		key += ".temporary"
	} else {
		key += ".permanent"
	}

	if action == permissions.Action_Replace || action == permissions.Action_ReplaceToWC {
		if oldTimestamp != 0 && newTimestamp != 0 {
			if oldTimestamp > newTimestamp {
				key += ".subtracted"
			} else {
				key += ".added"
			}
		}
	}

	return key
}

func DeclineWord(n int64, lang, one, two, five string) string {
	return getDeclension(n, translations.GetTranslation(one, lang), translations.GetTranslation(two, lang), translations.GetTranslation(five, lang))
}

// Source: https://github.com/rekurt/go-propisyu/blob/80c2c6af752483498506c191b080423cc7f747d5/declension.go#L17
func getDeclension(n int64, one, two, five string) string {
	n %= 100
	if n < 0 {
		n = -n
	}
	if n >= 11 && n <= 19 {
		return five
	}
	switch n % 10 {
	case 1:
		return one
	case 2, 3, 4:
		return two
	default:
		return five
	}
}

func IsValidClient(playerSlot int32) bool {
	return s2sdk.IsClientInGame(playerSlot) && !s2sdk.IsFakeClient(playerSlot) && !s2sdk.IsClientSourceTV(playerSlot)
}

func GetPlayerSlotBySteamID64(steamID64 uint64) int32 {
	for i := int32(0); i <= s2sdk.GetMaxClients(); i++ {
		if s2sdk.GetClientSteamID64(i) == steamID64 {
			return i
		}
	}

	return -1
}

func PrintConsoleTableToClient(playerSlot int32, tableName string, headers []string, rows [][]any) {
	//if !IsValidClient(playerSlot) {
	//	s2sdk.PrintToConsole(playerSlot, "No valid player.\n")
	//	return
	//}

	if len(tableName) == 0 {
		tableName = "Table"
	}

	if len(rows) == 0 {
		s2sdk.PrintToConsole(playerSlot, fmt.Sprintf("[%s]\n -\n ", tableName))
		return
	}

	slices.SortFunc(rows, func(a, b []any) int {
		return cmp.Compare(fmt.Sprint(a[0]), fmt.Sprint(b[0]))
	})

	colWidths := make([]int, len(headers))

	for i := range headers {
		colWidths[i] = utf8.RuneCountInString(headers[i])
	}

	for i := range rows {
		for j := range rows[i] {
			if j >= len(colWidths) {
				break
			}
			cellStr := fmt.Sprintf("%v", rows[i][j])
			cellLen := utf8.RuneCountInString(cellStr)
			if cellLen > colWidths[j] {
				colWidths[j] = cellLen
			}
		}
	}

	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', tabwriter.TabIndent)

	for i := range 2 {
		for j := range headers {
			if i == 0 {
				io.WriteString(w, headers[j])
			} else {
				io.WriteString(w, strings.Repeat("_", colWidths[j]))
			}

			if j < len(headers)-1 {
				w.Write([]byte{'\t'})
			}
		}

		w.Write([]byte{'\n'})
	}

	for i := range rows {
		for j, cell := range rows[i] {
			fmt.Fprint(w, cell)
			if j < len(rows[i])-1 {
				w.Write([]byte{'\t'})
			}
		}

		w.Write([]byte{'\n'})
	}

	w.Flush()

	s2sdk.PrintToConsole(playerSlot, fmt.Sprintf("[%s]\n", tableName))

	lines := strings.Split(buf.String(), "\n")
	for i := range lines {
		if lines[i] == "" {
			if i < len(lines)-1 {
				continue
			}
		}
		s2sdk.PrintToConsole(playerSlot, " "+lines[i])
	}
}

func RemoveTags(text string) string {
	var builder strings.Builder
	builder.Grow(len(text))

	tag := false

	for i := 0; i < len(text); i++ {
		c := text[i]
		switch c {
		case '{':
			tag = true
			continue
		case '}':
			if tag {
				tag = false
				continue
			}
		}
		if !tag {
			builder.WriteByte(c)
		}
	}

	return builder.String()
}

func ReplyToCommandCleared(caller int32, context s2sdk.ConCommandContext, message string) {
	if len(message) == 0 {
		return
	}

	if context == s2sdk.ConCommandContext_Console {
		message = RemoveTags(message)
	}

	s2sdk.ReplyToCommand(context, caller, message)
}
