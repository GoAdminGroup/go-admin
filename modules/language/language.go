package language

import (
	"github.com/chenhg5/go-admin/modules/config"
	"strings"
	"html/template"
)

func Get(value string) string {

	if config.Get().LANGUAGE == "" {
		return value
	}

	if locale, ok := Lang[config.Get().LANGUAGE][strings.ToLower(value)]; ok {
		return locale
	} else {
		return value
	}
}

func GetFromHtml(value template.HTML) template.HTML {

	if config.Get().LANGUAGE == "" {
		return value
	}

	if locale, ok := Lang[config.Get().LANGUAGE][strings.ToLower(string(value))]; ok {
		return template.HTML(locale)
	} else {
		return value
	}
}

type LangMap map[string]map[string]string

var Lang = LangMap{
	"cn": cn,
	"en": en,
	"jp": jp,
}

func (lang LangMap) Get(value string) string {
	if config.Get().LANGUAGE == "" {
		return value
	}

	if locale, ok := lang[config.Get().LANGUAGE][value]; ok {
		return locale
	} else {
		return value
	}
}
