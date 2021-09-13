package muon

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Localization struct {
	bundle     *i18n.Bundle
	Localizers map[string]*i18n.Localizer
}

func NewLocalization(defaultLanguage string) *Localization {
	bundle := i18n.NewBundle(language.Make(defaultLanguage))
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	localization := &Localization{
		bundle:     bundle,
		Localizers: make(map[string]*i18n.Localizer),
	}
	return localization
}

func (localization *Localization) Load(logger *Logger, localeDir string) error {
	logger.Info.Println("Loading localizations from", localeDir)
	files, err := filepath.Glob(filepath.Join(localeDir, "*.json"))
	if err != nil {
		return err
	}

	for _, file := range files {
		locale_name := filepath.Base(file)
		locale_name = strings.TrimSuffix(locale_name, ".json")
		_, err := localization.bundle.LoadMessageFile(file)
		if err != nil {
			return err
		}
		localization.Localizers[locale_name] = i18n.NewLocalizer(localization.bundle, locale_name)
	}

	return nil
}

func (localization *Localization) Localizer(locale string) *i18n.Localizer {
	return localization.Localizers[locale]
}

func (localization *Localization) Localize(locale string, messageId string, formatArgs ...interface{}) (string, error) {
	localized, err := localization.Localizer(locale).Localize(&i18n.LocalizeConfig{
		MessageID: messageId,
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(localized, formatArgs...), nil
}
