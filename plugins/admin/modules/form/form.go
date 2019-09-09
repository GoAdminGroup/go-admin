package form

type FormValue map[string][]string

func (f FormValue) Get(key string) string {
	return f[key][0]
}

func (f FormValue) IsEmpty(key ...string) bool {
	for _, k := range key {
		if f.Get(k) == "" {
			return true
		}
	}
	return false
}
