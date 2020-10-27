// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package language

import (
	"html/template"
	"strings"

	"github.com/GoAdminGroup/go-admin/modules/config"
	"golang.org/x/text/language"
)

var (
	EN = language.English.String()
	CN = language.Chinese.String()
	JP = language.Japanese.String()
	TC = language.TraditionalChinese.String()
)

func FixedLanguageKey(key string) string {
	if key == "en" {
		return EN
	}
	if key == "cn" {
		return CN
	}
	if key == "jp" {
		return JP
	}
	if key == "tc" {
		return TC
	}
	return key
}

var Langs = [...]string{EN, CN, JP, TC}

// Get return the value of default scope.
func Get(value string) string {
	return GetWithScope(value)
}

// GetWithScope return the value of given scopes.
func GetWithScope(value string, scopes ...string) string {
	if config.GetLanguage() == "" {
		return value
	}

	return GetWithScopeAndLanguageSet(value, config.GetLanguage(), scopes...)
}

// GetWithLang return the value of given language set.
func GetWithLang(value, lang string) string {
	if lang == "" {
		lang = config.GetLanguage()
	}
	return GetWithScopeAndLanguageSet(value, lang)
}

// GetWithScopeAndLanguageSet return the value of given scopes and language set.
func GetWithScopeAndLanguageSet(value, lang string, scopes ...string) string {
	if locale, ok := Lang[lang][JoinScopes(scopes)+strings.ToLower(value)]; ok {
		return locale
	}

	return value
}

// GetFromHtml return the value of given scopes and template.HTML value.
func GetFromHtml(value template.HTML, scopes ...string) template.HTML {
	if config.GetLanguage() == "" {
		return value
	}

	if locale, ok := Lang[config.GetLanguage()][JoinScopes(scopes)+strings.ToLower(string(value))]; ok {
		return template.HTML(locale)
	}

	return value
}

// WithScopes join scopes prefix and the value.
func WithScopes(value string, scopes ...string) string {
	return JoinScopes(scopes) + strings.ToLower(value)
}

type LangSet map[string]string

func (l LangSet) Add(key, value string) {
	l[key] = value
}

func (l LangSet) Combine(set LangSet) LangSet {
	for k, s := range set {
		l[k] = s
	}
	return l
}

// LangMap is the map of language packages.
type LangMap map[string]LangSet

// Lang is the global LangMap.
var Lang = LangMap{
	language.Chinese.String():            cn,
	language.English.String():            en,
	language.Japanese.String():           jp,
	language.TraditionalChinese.String(): tc,

	"cn": cn,
	"en": en,
	"jp": jp,
	"tc": tc,
}

// Get get the value from LangMap.
func (lang LangMap) Get(value string) string {
	return lang.GetWithScope(value)
}

// GetWithScope get the value from LangMap with given scopes.
func (lang LangMap) GetWithScope(value string, scopes ...string) string {
	if config.GetLanguage() == "" {
		return value
	}

	if locale, ok := lang[config.GetLanguage()][JoinScopes(scopes)+strings.ToLower(value)]; ok {
		return locale
	}

	return value
}

// Add add a language package to the Lang.
func Add(key string, lang map[string]string) {
	Lang[key] = lang
}

// AppendTo add more language translations to the given language set.
func AppendTo(lang string, set map[string]string) {
	for key, value := range set {
		Lang[lang][key] = value
	}
}

func JoinScopes(scopes []string) string {
	j := ""
	for _, scope := range scopes {
		if scope != "" {
			j += scope + "."
		}
	}
	return j
}
