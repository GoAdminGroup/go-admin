// Copyright 2019 GoAdmin Core Team.  All rights reserved.
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

func Get(value string) string {
	return GetWithScope(value)
}

func GetWithScope(value string, scopes ...string) string {
	if config.Get().Language == "" {
		return value
	}

	if locale, ok := Lang[config.Get().Language][joinScopes(scopes)+strings.ToLower(value)]; ok {
		return locale
	} else {
		return value
	}
}

func GetFromHtml(value template.HTML, scopes ...string) template.HTML {
	if config.Get().Language == "" {
		return value
	}

	if locale, ok := Lang[config.Get().Language][joinScopes(scopes)+strings.ToLower(string(value))]; ok {
		return template.HTML(locale)
	} else {
		return value
	}
}

func WithScopes(value string, scopes ...string) string {
	return joinScopes(scopes) + strings.ToLower(value)
}

type LangMap map[string]map[string]string

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

func (lang LangMap) Get(value string) string {
	return lang.GetWithScope(value)
}

func (lang LangMap) GetWithScope(value string, scopes ...string) string {
	if config.Get().Language == "" {
		return value
	}

	if locale, ok := lang[config.Get().Language][joinScopes(scopes)+strings.ToLower(value)]; ok {
		return locale
	} else {
		return value
	}
}

func Add(key string, lang map[string]string) {
	Lang[key] = lang
}

func joinScopes(scopes []string) string {
	j := ""
	for _, scope := range scopes {
		j += scope + "."
	}
	return j
}
