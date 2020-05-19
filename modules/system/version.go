package system

const version = "v1.2.12"

var requireThemeVersion = map[string][]string{
	"adminlte": {">=v0.0.34"},
	"sword":    {">=v0.0.34"},
}

// Version return the version of framework.
func Version() string {
	return version
}

// RequireThemeVersion return the require official version
func RequireThemeVersion() map[string][]string {
	return requireThemeVersion
}
