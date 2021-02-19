package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/system"
	"github.com/GoAdminGroup/go-admin/modules/utils"
)

func buildProjectWeb(port string) {

	var startChan = make(chan struct{})

	type InstallationPage struct {
		Version       string
		GoVer         string
		Port          string
		CSRFToken     string
		CurrentLang   string
		GOOS          string
		DefModuleName string
	}

	var tokens = make([]string, 0)

	rootPath, err := os.Getwd()

	if err != nil {
		rootPath = "."
	} else {
		rootPath = filepath.ToSlash(rootPath)
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	defModuleName := filepath.Base(dir)

	go func(sc chan struct{}) {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

			lang := defaultLang
			if r.URL.Query().Get("lang") != "" {
				lang = r.URL.Query().Get("lang")
			}

			w.Header().Add("Content-Type", "text/html; charset=utf-8")
			t, err := template.New("web_installation").Funcs(map[string]interface{}{
				"local": local(lang),
			}).Parse(projectWebTmpl)
			if err != nil {
				fmt.Println("get web installation page template error", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			curLang := "web.simplified chinese"
			if lang == "en" {
				curLang = "web.english"
			}
			tk := utils.Uuid(25)
			tokens = append(tokens, tk)
			err = t.Execute(w, InstallationPage{
				Version:       system.Version(),
				GoVer:         strings.Title(runtime.Version()),
				Port:          port,
				CSRFToken:     tk,
				CurrentLang:   curLang,
				GOOS:          runtime.GOOS,
				DefModuleName: defModuleName,
			})
			if err != nil {
				fmt.Println("get web installation page template error", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		})

		var checkEmpty = func(list []string, r *http.Request) (string, bool) {
			for i := 0; i < len(list); i++ {
				if r.PostFormValue(list[i]) == "" {
					return list[i], false
				}
			}
			return "", true
		}

		http.HandleFunc("/install", func(w http.ResponseWriter, r *http.Request) {

			if r.Method != "POST" {
				w.WriteHeader(http.StatusBadRequest)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"code": 400, "msg": "wrong method"}`))
				return
			}

			lang := defaultLang
			if r.URL.Query().Get("lang") != "" {
				lang = r.URL.Query().Get("lang")
			}

			_ = r.ParseForm()

			fields := []string{"db_type", "web_framework", "module_name", "theme", "language", "http_port", "prefix",
				"web_title", "login_page_logo", "sidebar_logo", "sidebar_min_logo"}

			if field, ok := checkEmpty(fields, r); !ok {
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"code": 400, "msg": "` + local(lang)("web.wrong parameter") + `: ` + field + `"}`))
				return
			}

			if r.PostFormValue("db_type") == "sqlite" {
				if r.PostFormValue("db_path") == "" {
					w.WriteHeader(http.StatusOK)
					w.Header().Add("Content-Type", "application/json")
					_, _ = w.Write([]byte(`{"code": 400, "msg": "` + local(lang)("web.wrong parameter") + `"}`))
					return
				}
			} else if r.PostFormValue("db_type") == "postgresql" {
				if field, ok := checkEmpty([]string{"db_schema", "db_port", "db_host", "db_user", "db_passwd", "db_name"}, r); !ok {
					w.WriteHeader(http.StatusOK)
					w.Header().Add("Content-Type", "application/json")
					_, _ = w.Write([]byte(`{"code": 400, "msg": "` + local(lang)("web.wrong parameter") + `: ` + field + `"}`))
					return
				}
			} else {
				if field, ok := checkEmpty([]string{"db_port", "db_host", "db_user", "db_passwd", "db_name"}, r); !ok {
					w.WriteHeader(http.StatusOK)
					w.Header().Add("Content-Type", "application/json")
					_, _ = w.Write([]byte(`{"code": 400, "msg": "` + local(lang)("web.wrong parameter") + `: ` + field + `"}`))
					return
				}
			}

			var p = Project{
				Port:         r.PostFormValue("http_port"),
				Theme:        r.PostFormValue("theme"),
				Prefix:       r.PostFormValue("prefix"),
				Language:     r.PostFormValue("language"),
				Driver:       r.PostFormValue("db_type"),
				DriverModule: r.PostFormValue("db_type"),
				Framework:    r.PostFormValue("web_framework"),
				Module:       r.PostFormValue("module_name"),
				Orm:          r.PostFormValue("use_gorm"),
			}

			if p.Driver == db.DriverPostgresql {
				p.DriverModule = "postgres"
			}

			var (
				info = &dbInfo{
					DriverName: r.PostFormValue("db_type"),
					Host:       r.PostFormValue("db_host"),
					Port:       r.PostFormValue("db_port"),
					File:       r.PostFormValue("db_path"),
					User:       r.PostFormValue("db_user"),
					Password:   r.PostFormValue("db_passwd"),
					Schema:     r.PostFormValue("db_schema"),
					Database:   r.PostFormValue("db_name"),
				}
				dbList config.DatabaseList
			)

			if info.DriverName != db.DriverSqlite {
				dbList = map[string]config.Database{
					"default": {
						Host:       info.Host,
						Port:       info.Port,
						User:       info.User,
						Pwd:        info.Password,
						Name:       info.Database,
						MaxIdleCon: 5,
						MaxOpenCon: 10,
						Driver:     info.DriverName,
					},
				}
			} else {
				dbList = map[string]config.Database{
					"default": {
						Driver: info.DriverName,
						File:   info.File,
					},
				}
			}

			cfg := config.SetDefault(&config.Config{
				Debug: true,
				Env:   config.EnvLocal,
				Theme: p.Theme,
				Store: config.Store{
					Path:   "./uploads",
					Prefix: "uploads",
				},
				Language:          p.Language,
				UrlPrefix:         p.Prefix,
				IndexUrl:          "/",
				AccessLogPath:     rootPath + "/logs/access.log",
				ErrorLogPath:      rootPath + "/logs/error.log",
				InfoLogPath:       rootPath + "/logs/info.log",
				BootstrapFilePath: rootPath + "/bootstrap.go",
				GoModFilePath:     rootPath + "/go.mod",
				Logo:              template.HTML(r.PostFormValue("sidebar_logo")),
				LoginLogo:         template.HTML(r.PostFormValue("login_page_logo")),
				MiniLogo:          template.HTML(r.PostFormValue("sidebar_min_logo")),
				Title:             r.PostFormValue("web_title"),
				Databases:         dbList,
			})

			installProjectTmpl(p, cfg, "", info)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"code": 0, "msg": "` + local(lang)("web.install success") + `", "data": {"readme": ""}}`))
		})

		l, err := net.Listen("tcp", ":"+port)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}

		fmt.Println("GoAdmin web install program start.")
		sc <- struct{}{}

		if err := http.Serve(l, nil); err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}(startChan)

	<-startChan
	err = open("http://localhost:" + port + "/")
	if err != nil {
		fmt.Println("failed to open browser", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
