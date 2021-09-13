package muon

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ContentLink struct {
	Text string `json:"text"`
	Url  string `json:"url"`
}

type ContentConfiguration struct {
	DefaultLanguage string        `json:"defaultLanguage"`
	Languages       []string      `json:"languages"`
	Index           string        `json:"index"`
	NotFound        string        `json:"notFound"`
	Links           []ContentLink `json:"links"`
	ContactUri      string        `json:"contactUri"`
}

type ServiceConfiguration struct {
	BindTo           string
	ArticlesRoot     string
	StaticRoot       string
	TemplateRoot     string
	LocalizationRoot string
	ContentConfig    *ContentConfiguration
}

func LoadContentConfig(logger *Logger, configPath string) (*ContentConfiguration, error) {
	var config ContentConfiguration
	rawContent, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(rawContent, &config)
	if err != nil {
		return nil, err
	}

	logger.Info.Println("Loaded content configuration from ", configPath, ": ", string(rawContent))
	return &config, nil
}

func loadEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.New(fmt.Sprintf("%v environment variable is empty", key))
	} else {
		return value, nil
	}
}

func LoadConfigFromEnv(logger *Logger) (*ServiceConfiguration, error) {
	dataRoot, err := loadEnv("MUON_DATA")
	if err != nil {
		return nil, err
	}
	listenTo, err := loadEnv("MUON_LISTEN")
	if err != nil {
		return nil, err
	}

	config := ServiceConfiguration{
		BindTo:           listenTo,
		ArticlesRoot:     filepath.Join(dataRoot, "articles"),
		StaticRoot:       filepath.Join(dataRoot, "static"),
		TemplateRoot:     filepath.Join(dataRoot, "templates"),
		LocalizationRoot: filepath.Join(dataRoot, "i18n"),
	}
	logger.Info.Println("Instantiated service configuration from environment variables")
	logger.Info.Println("  * Bind to: ", config.BindTo)
	logger.Info.Println("  * Article catalogue: ", config.ArticlesRoot)
	logger.Info.Println("  * Static files: ", config.StaticRoot)
	logger.Info.Println("  * Templates: ", config.TemplateRoot)
	logger.Info.Println("  * Localization: ", config.LocalizationRoot)

	contentConfig, err := LoadContentConfig(logger, filepath.Join(dataRoot, "content.json"))
	if err != nil {
		return nil, err
	}
	config.ContentConfig = contentConfig
	return &config, nil
}
