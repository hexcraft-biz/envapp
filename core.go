package envapp

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"time"
)

// ================================================================
type App struct {
	Env        string         `json:"env"`
	Host       string         `json:"host"`
	Path       string         `json:"path"`
	Port       string         `json:"port"`
	Timezone   string         `json:"timezone"`
	Location   *time.Location `json:"-"`
	TrustProxy string         `json:"trustProxy"`
	Visibility string         `json:"visibility"`
	BaseURL    *url.URL       `json:"-"`
}

const (
	EnvDevelopment = "development"
	EnvDebug       = "debug"
	EnvStage       = "stage"
	EnvProduction  = "production"

	VisibilityInternal = "internal"
	VisibilityExternal = "external"
)

func (r *App) Sanitize() error {
	var err error

	switch r.Env {
	case EnvDevelopment, EnvDebug, EnvStage, EnvProduction:
	default:
		return fmt.Errorf("app: invalid env (%s|%s|%s|%s)", EnvDevelopment, EnvDebug, EnvStage, EnvProduction)
	}

	switch r.Visibility {
	case VisibilityInternal, VisibilityExternal:
	default:
		return fmt.Errorf("app: invalid visibility. (%s|%s)", VisibilityInternal, VisibilityExternal)
	}

	r.Location, err = time.LoadLocation(r.Timezone)
	if err != nil {
		return err
	}

	r.BaseURL, err = url.ParseRequestURI("https://" + path.Join(r.Host, r.Path))
	return err
}

// ================================================================
func New() (*App, error) {
	env := App{
		Env:        os.Getenv("APP_ENV"),
		Host:       os.Getenv("APP_HOST"),
		Path:       path.Join("/", os.Getenv("APP_PATH")),
		Port:       os.Getenv("APP_PORT"),
		Timezone:   os.Getenv("TIMEZONE"),
		TrustProxy: os.Getenv("TRUST_PROXY"),
		Visibility: os.Getenv("VISIBILITY"),
	}

	if err := env.Sanitize(); err != nil {
		return nil, err
	}

	return &env, nil
}
