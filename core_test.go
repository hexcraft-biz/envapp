package envapp

import "testing"

func TestAppSanitizeBaseURLPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "root path",
			path: "/",
			want: "/operating-reports/id",
		},
		{
			name: "application prefix",
			path: "/taiwandreamer-console",
			want: "/taiwandreamer-console/operating-reports/id",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := App{
				Env:        EnvProduction,
				Host:       "console.taiwandreamer.com.tw",
				Path:       tt.path,
				Timezone:   "UTC",
				Visibility: VisibilityExternal,
			}

			if err := app.Sanitize(); err != nil {
				t.Fatalf("Sanitize() error = %v", err)
			}

			got := app.BaseURL.JoinPath("operating-reports", "id").RequestURI()
			if got != tt.want {
				t.Fatalf("BaseURL.JoinPath().RequestURI() = %q, want %q", got, tt.want)
			}
		})
	}
}
