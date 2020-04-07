package form

import "errors"

const (
	PostTypeKey           = "__go_admin_post_type"
	PostResultKey         = "__go_admin_post_result"
	PostIsSingleUpdateKey = "__go_admin_is_single_update"

	PreviousKey = "__go_admin_previous_"
	TokenKey    = "__go_admin_t_"
	MethodKey   = "__go_admin_method_"

	NoAnimationKey = "__go_admin_no_animation_"
)

// Values maps a string key to a list of values.
// It is typically used for query parameters and form values.
// Unlike in the http.Header map, the keys in a Values map
// are case-sensitive.
type Values map[string][]string

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (f Values) Get(key string) string {
	if len(f[key]) > 0 {
		return f[key][0]
	}
	return ""
}

// Add adds the value to key. It appends to any existing
// values associated with key.
func (f Values) Add(key string, value string) {
	f[key] = []string{value}
}

// IsEmpty check the key is empty or not.
func (f Values) IsEmpty(key ...string) bool {
	for _, k := range key {
		if f.Get(k) == "" {
			return true
		}
	}
	return false
}

// Has check the key exists or not.
func (f Values) Has(key ...string) bool {
	for _, k := range key {
		if f.Get(k) != "" {
			return true
		}
	}
	return false
}

// Delete deletes the values associated with key.
func (f Values) Delete(key string) {
	delete(f, key)
}

func (f Values) ToMap() map[string]string {
	var m = make(map[string]string)
	for key, v := range f {
		if len(v) > 0 {
			m[key] = v[0]
		}
	}
	return m
}

func (f Values) IsUpdatePost() bool {
	return f.Get(PostTypeKey) == "0"
}

func (f Values) IsInsertPost() bool {
	return f.Get(PostTypeKey) == "1"
}

func (f Values) PostError() error {
	msg := f.Get(PostResultKey)
	if msg == "" {
		return nil
	}
	return errors.New(msg)
}

func (f Values) IsSingleUpdatePost() bool {
	return f.Get(PostIsSingleUpdateKey) == "1"
}

func (f Values) RemoveRemark() Values {
	f.Delete(PostTypeKey)
	f.Delete(PostIsSingleUpdateKey)
	return f
}

func (f Values) RemoveSysRemark() Values {
	f.Delete(PostTypeKey)
	f.Delete(PostIsSingleUpdateKey)
	f.Delete(PostResultKey)
	f.Delete(PreviousKey)
	f.Delete(TokenKey)
	f.Delete(MethodKey)
	f.Delete(NoAnimationKey)
	return f
}
