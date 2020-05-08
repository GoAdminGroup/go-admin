package system

const version = "v1.2.10"

var requireThemeVersion = map[string][]string{
	"adminlte": {"v0.0.32"},
	"sword":    {"v0.0.32"},
}

// Version return the version of framework.
func Version() string {
	return version
}

// RequireThemeVersion return the require official version
func RequireThemeVersion() map[string][]string {
	return requireThemeVersion
}
