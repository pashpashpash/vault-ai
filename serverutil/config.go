package serverutil

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"
)

var (
	// Ensure files only have to be read once
	CONFIG_READ_MUTEX sync.Mutex
	CONFIG            *ConstantsConfig
)

const (
	WEB_METADATA_PATH = "./config/websites.json"
)

type ConstantsConfig struct {
	SiteMetaData SiteConfig
}

type WebPageTemplate struct {
	PageTitle       string
	PageIcon        string
	MetaTitle       string
	MetaDescription string
	MetaVideo       string
	MetaType        string
	MetaImage       string
	MetaUrl         string
	TwitterUsername string
	MetaKeywords    string
}

type SiteConfig struct {
	SitePath map[string]WebPageTemplate
}

func GetConfig() *ConstantsConfig {
	CONFIG_READ_MUTEX.Lock()
	defer CONFIG_READ_MUTEX.Unlock()

	if CONFIG != nil {
		return CONFIG
	}

	CONFIG = &ConstantsConfig{}

	MetaData := &SiteConfig{}
	unmarshalFile(WEB_METADATA_PATH, &MetaData.SitePath)
	CONFIG.SiteMetaData = *MetaData

	log.Println("[Config] Loaded ./config/ files")

	return CONFIG
}

func unmarshalFile(path string, targetStruct interface{}) {
	path, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("Failed to parse %s: %v", path, err)
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to parse %s: %v", path, err)
	}

	if err := json.Unmarshal(file, targetStruct); err != nil {
		log.Fatalf("Failed to parse %s: %v", path, err)
	}
}
