package types

import (
	"fmt"
	"html"
	"html/template"
	"strings"
)

type FieldDisplay struct {
	Display              FieldFilterFn
	DisplayProcessChains DisplayProcessFnChains
}

func (f FieldDisplay) ToDisplay(value FieldModel) interface{} {
	val := f.Display(value)

	if _, ok := val.(template.HTML); !ok {
		if _, ok2 := val.([]string); !ok2 {
			valStr := fmt.Sprintf("%v", val)
			for _, process := range f.DisplayProcessChains {
				valStr = process(valStr)
			}
			return valStr
		}
	}

	return val
}

func (f FieldDisplay) AddLimit(limit int) DisplayProcessFnChains {
	return f.DisplayProcessChains.Add(func(value string) string {
		if limit > len(value) {
			return value
		} else if limit < 0 {
			return ""
		} else {
			return value[:limit]
		}
	})
}

func (f FieldDisplay) AddTrimSpace() DisplayProcessFnChains {
	return f.DisplayProcessChains.Add(func(value string) string {
		return strings.TrimSpace(value)
	})
}

func (f FieldDisplay) AddSubstr(start int, end int) DisplayProcessFnChains {
	return f.DisplayProcessChains.Add(func(value string) string {
		if start > end || start > len(value) || end < 0 {
			return ""
		}
		if start < 0 {
			start = 0
		}
		if end > len(value) {
			end = len(value)
		}
		return value[start:end]
	})
}

func (f FieldDisplay) AddToTitle() DisplayProcessFnChains {
	return f.DisplayProcessChains.Add(func(value string) string {
		return strings.Title(value)
	})
}

func (f FieldDisplay) AddToUpper() DisplayProcessFnChains {
	return f.DisplayProcessChains.Add(func(value string) string {
		return strings.ToUpper(value)
	})
}

func (f FieldDisplay) AddToLower() DisplayProcessFnChains {
	return f.DisplayProcessChains.Add(func(value string) string {
		return strings.ToLower(value)
	})
}

type DisplayProcessFn func(string) string

type DisplayProcessFnChains []DisplayProcessFn

func (d DisplayProcessFnChains) Valid() bool {
	return len(d) > 0
}

func (d DisplayProcessFnChains) Add(f DisplayProcessFn) DisplayProcessFnChains {
	return append(d, f)
}

func (d DisplayProcessFnChains) Append(f DisplayProcessFnChains) DisplayProcessFnChains {
	return append(d, f...)
}

func (d DisplayProcessFnChains) Copy() DisplayProcessFnChains {
	if len(d) == 0 {
		return make(DisplayProcessFnChains, 0)
	} else {
		var newDisplayProcessFnChains = make(DisplayProcessFnChains, len(d))
		copy(newDisplayProcessFnChains, d)
		return newDisplayProcessFnChains
	}
}

func chooseDisplayProcessChains(internal DisplayProcessFnChains) DisplayProcessFnChains {
	if len(internal) > 0 {
		return internal
	}
	return globalDisplayProcessChains.Copy()
}

var globalDisplayProcessChains = make(DisplayProcessFnChains, 0)

func AddGlobalDisplayProcessFn(f DisplayProcessFn) {
	globalDisplayProcessChains = globalDisplayProcessChains.Add(f)
}

func AddLimit(limit int) DisplayProcessFnChains {
	return addLimit(limit, globalDisplayProcessChains)
}

func AddTrimSpace() DisplayProcessFnChains {
	return addTrimSpace(globalDisplayProcessChains)
}

func AddSubstr(start int, end int) DisplayProcessFnChains {
	return addSubstr(start, end, globalDisplayProcessChains)
}

func AddToTitle() DisplayProcessFnChains {
	return addToTitle(globalDisplayProcessChains)
}

func AddToUpper() DisplayProcessFnChains {
	return addToUpper(globalDisplayProcessChains)
}

func AddToLower() DisplayProcessFnChains {
	return addToLower(globalDisplayProcessChains)
}

func AddXssFilter() DisplayProcessFnChains {
	return addXssFilter(globalDisplayProcessChains)
}

func AddXssJsFilter() DisplayProcessFnChains {
	return addXssJsFilter(globalDisplayProcessChains)
}

func addLimit(limit int, chains DisplayProcessFnChains) DisplayProcessFnChains {
	chains = chains.Add(func(value string) string {
		if limit > len(value) {
			return value
		} else if limit < 0 {
			return ""
		} else {
			return value[:limit]
		}
	})
	return chains
}

func addTrimSpace(chains DisplayProcessFnChains) DisplayProcessFnChains {
	chains = chains.Add(func(value string) string {
		return strings.TrimSpace(value)
	})
	return chains
}

func addSubstr(start int, end int, chains DisplayProcessFnChains) DisplayProcessFnChains {
	chains = chains.Add(func(value string) string {
		if start > end || start > len(value) || end < 0 {
			return ""
		}
		if start < 0 {
			start = 0
		}
		if end > len(value) {
			end = len(value)
		}
		return value[start:end]
	})
	return chains
}

func addToTitle(chains DisplayProcessFnChains) DisplayProcessFnChains {
	chains = chains.Add(func(value string) string {
		return strings.Title(value)
	})
	return chains
}

func addToUpper(chains DisplayProcessFnChains) DisplayProcessFnChains {
	chains = chains.Add(func(value string) string {
		return strings.ToUpper(value)
	})
	return chains
}

func addToLower(chains DisplayProcessFnChains) DisplayProcessFnChains {
	chains = chains.Add(func(value string) string {
		return strings.ToLower(value)
	})
	return chains
}

func addXssFilter(chains DisplayProcessFnChains) DisplayProcessFnChains {
	chains = chains.Add(func(value string) string {
		return html.EscapeString(value)
	})
	return chains
}

func addXssJsFilter(chains DisplayProcessFnChains) DisplayProcessFnChains {
	chains = chains.Add(func(value string) string {
		s := strings.Replace(value, "<script>", "&lt;script&gt;", -1)
		return strings.Replace(s, "</script>", "&lt;/script&gt;", -1)
	})
	return chains
}
