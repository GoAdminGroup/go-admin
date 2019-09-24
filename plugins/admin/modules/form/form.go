package form

type Values map[string][]string

func (f Values) Get(key string) string {
	return f[key][0]
}

func (f Values) IsEmpty(key ...string) bool {
	for _, k := range key {
		if f.Get(k) == "" {
			return true
		}
	}
	return false
}
