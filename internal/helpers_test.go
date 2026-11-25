package internal

import "testing"

func TestTranslateWithFallback(t *testing.T) {
	calls := make([]string, 0, 2)
	stub := func(messages map[string]string) func(string, string, []any) string {
		calls = calls[:0]
		return func(lang, key string, _ []any) string {
			calls = append(calls, lang+":"+key)
			return messages[lang+":"+key]
		}
	}

	got := translateWithFallback("zh-CN", "k.hint", nil, stub(map[string]string{
		"zh-CN:k.hint": "中文提示",
		"en:k.hint":    "English hint",
	}))
	if got != "中文提示" || len(calls) != 1 {
		t.Fatalf("expected direct locale hit without fallback, got %q calls=%v", got, calls)
	}

	got = translateWithFallback("zh-CN", "k.hint", nil, stub(map[string]string{
		"en:k.hint": "English hint",
	}))
	if got != "English hint" || len(calls) != 2 || calls[1] != "en:k.hint" {
		t.Fatalf("expected English fallback, got %q calls=%v", got, calls)
	}

	got = translateWithFallback("en", "k.hint", nil, stub(map[string]string{}))
	if got != "" || len(calls) != 1 {
		t.Fatalf("expected single lookup for missing English key, got %q calls=%v", got, calls)
	}
}

func TestHasVisibleText(t *testing.T) {
	if hasVisibleText("") || hasVisibleText("   \t\n") || hasVisibleText("{red}{default}") {
		t.Fatal("expected empty, whitespace, and tag-only strings not to count as visible text")
	}
	if !hasVisibleText("{green}Granted") {
		t.Fatal("expected tagged non-empty text to be visible")
	}
}
