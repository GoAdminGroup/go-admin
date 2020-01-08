package utils

import (
	"html/template"
	"strings"
)

func CompressedContent(h *template.HTML) {
	st := strings.Split(string(*h), "\n")
	ss := []string{}
	for i := 0;i < len(st);i++ {
		st[i] = strings.TrimSpace(st[i])
		if st[i] != "" {
			ss = append(ss, st[i])
		}
	}
	*h = template.HTML(strings.Join(ss, "\n"))
}
