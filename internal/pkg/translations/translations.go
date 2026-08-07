package translations

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/untrustedmodders/go-plugify"
	"gopkg.in/yaml.v3"
)

var translations map[string]map[string]string

var steamLanguages = map[string]string{
	"english":    "en",
	"german":     "de",
	"french":     "fr",
	"italian":    "it",
	"koreana":    "ko",
	"spanish":    "es",
	"schinese":   "zh-CN",
	"tchinese":   "zh-TW",
	"russian":    "ru",
	"thai":       "th",
	"japanese":   "ja",
	"portuguese": "pt",
	"polish":     "pl",
	"danish":     "da",
	"dutch":      "nl",
	"finnish":    "fi",
	"norwegian":  "no",
	"swedish":    "sv",
	"hungarian":  "hu",
	"czech":      "cs",
	"romanian":   "ro",
	"turkish":    "tr",
	"brazilian":  "pt-BR",
	"bulgarian":  "bg",
	"greek":      "el",
	"ukrainian":  "uk",
	"latam":      "es-419",
	"vietnamese": "vi",
	"indonesian": "id",
}

func LoadTranslation(fileName string) error {
	filePath := filepath.Join(plugify.BaseDir(), "translations", fileName+".yaml")

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to read translation file: %w", err)
	}

	defer file.Close()

	var data map[string]map[string]string

	err = yaml.NewDecoder(file).Decode(&data)
	if err != nil {
		return fmt.Errorf("failed to parse YAML: %w", err)
	}

	translations = make(map[string]map[string]string)

	for key, translation := range data {
		translations[key] = translation
	}

	return nil
}

func getTranslationText(key, lang string) (string, bool) {
	if trans, ok := translations[key]; ok {
		if text, ok := trans[lang]; ok {
			return text, true
		}
		if text, ok := trans["en"]; ok {
			return text, true
		}
	}

	return "", false
}

func processArgs(text string, args ...any) string {
	text = strings.ReplaceAll(text, "\n", "\u2029")
	if len(args) == 0 {
		return text
	}

	return safeSprintf(text, args...)
}

func safeSprintf(format string, args ...any) string {
	result := fmt.Sprintf(format, args...)

	if strings.Contains(result, "%!") {
		log.Printf("FORMAT ERROR:\nformat: %q\nargs: %#v\nresult: %q\n", format, args, result)
	}

	return result
}

func GetTranslation(key string, lang string, args ...any) string {
	if result, ok := steamLanguages[lang]; ok {
		lang = result
	} else {
		lang = "en"
	}

	text, found := getTranslationText(key, lang)
	if !found {
		return key
	}

	return processArgs(text, args...)
}
