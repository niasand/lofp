// Package i18n provides simple internationalization for the LoFP game engine.
// It looks up translation keys from a locale-specific map, falling back to
// the key itself (English) when no translation is found.
package i18n

import (
	"os"
	"strings"
	"sync"
)

// currentLocale is the active locale. Defaults to "en".
var (
	currentLocale = "en"
	localeMu      sync.RWMutex
)

func init() {
	if loc := os.Getenv("LOCALE"); loc != "" {
		currentLocale = strings.ToLower(loc)
	}
}

// SetLocale overrides the current locale at runtime.
func SetLocale(loc string) {
	localeMu.Lock()
	defer localeMu.Unlock()
	currentLocale = strings.ToLower(loc)
}

// GetLocale returns the current locale.
func GetLocale() string {
	localeMu.RLock()
	defer localeMu.RUnlock()
	return currentLocale
}

// T translates a key into the current locale. If no translation is found,
// the key itself is returned (assumed to be the English fallback).
func T(key string) string {
	localeMu.RLock()
	loc := currentLocale
	localeMu.RUnlock()
	if loc == "zh" {
		if msg, ok := zhMessages[key]; ok {
			return msg
		}
	}
	return key
}

// Tformat is a convenience wrapper that returns the translated key.
// Format args are NOT interpolated here — callers use fmt.Sprintf on the result.
func Tformat(key string) string {
	return T(key)
}
