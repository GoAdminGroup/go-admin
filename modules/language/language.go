// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package language

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	"golang.org/x/text/language"
	"html/template"
	"strings"
)

var (
	EN = language.English.String()
	CN = language.Chinese.String()
	JP = language.Japanese.String()
	TC = language.TraditionalChinese.String()
)

// Get return the value of default scope.
func Get(value string) string {
	return GetWithScope(value)
}

// GetWithScope return the value of given scopes.
func GetWithScope(value string, scopes ...string) string {
	if config.Get().Language == "" {
		return value
	}

	if locale, ok := Lang[config.Get().Language][JoinScopes(scopes)+strings.ToLower(value)]; ok {
		return locale
	}

	return value
}

// GetFromHtml return the value of given scopes and template.HTML value.
func GetFromHtml(value template.HTML, scopes ...string) template.HTML {
	if config.Get().Language == "" {
		return value
	}

	if locale, ok := Lang[config.Get().Language][JoinScopes(scopes)+strings.ToLower(string(value))]; ok {
		return template.HTML(locale)
	}

	return value
}

// WithScopes join scopes prefix and the value.
func WithScopes(value string, scopes ...string) string {
	return JoinScopes(scopes) + strings.ToLower(value)
}

// LangMap is the map of language packages.
type LangMap map[string]map[string]string

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
	if config.Get().Language == "" {
		return value
	}

	if locale, ok := lang[config.Get().Language][JoinScopes(scopes)+strings.ToLower(value)]; ok {
		return locale
	}

	return value
}

// Add add a language package to the Lang.
func Add(key string, lang map[string]string) {
	Lang[key] = lang
}

func JoinScopes(scopes []string) string {
	j := ""
	for _, scope := range scopes {
		j += scope + "."
	}
	return j
}

// GetUser return the value of user scope.
func GetUser(value string, uid int64) string {
	return GetUserWithScope(value, uid)
}

// GetUserWithScope return the value of given scopes.
func GetUserWithScope(value string, uid int64, scopes ...string) string {
	if config.GetUserConf(uid).Language == "" {
		return value
	}

	if locale, ok := Lang[config.GetUserConf(uid).Language][JoinScopes(scopes)+strings.ToLower(value)]; ok {
		return locale
	}

	return value
}

// GetUserFromHtml return the value of given scopes and template.HTML value.
func GetUserFromHtml(value template.HTML, uid int64, scopes ...string) template.HTML {
	if config.GetUserConf(uid).Language == "" {
		return value
	}

	if locale, ok := Lang[config.GetUserConf(uid).Language][JoinScopes(scopes)+strings.ToLower(string(value))]; ok {
		return template.HTML(locale)
	}

	return value
}
