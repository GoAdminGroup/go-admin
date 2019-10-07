package form

type Values map[string][]string

func (f Values) Get(key string) string {
	if len(f[key]) > 0 {
		return f[key][0]
	}
	return ""
}

func (f Values) Add(key string, value string) {
	f[key] = []string{value}
}

func (f Values) IsEmpty(key ...string) bool {
	for _, k := range key {
		if f.Get(k) == "" {
			return true
		}
	}
	return false
}

func (f Values) Delete(key string) {
	delete(f, key)
}
