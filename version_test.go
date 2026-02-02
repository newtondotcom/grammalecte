package grammalecte

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestVersionFromPythonSource(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		source  string
		want    string
		wantErr error
	}{
		{
			name: "prefers dunder version",
			source: `
__version__ = "1.2.3"
version = "9.9.9"
`,
			want: "1.2.3",
		},
		{
			name: "accepts single quotes",
			source: `
__version__ = '2.0.1'
`,
			want: "2.0.1",
		},
		{
			name: "falls back to version",
			source: `
version = "3.4.5"
`,
			want: "3.4.5",
		},
		{
			name: "ignores commented version",
			source: `
# __version__ = "0.1.0"
version = "4.5.6"
`,
			want: "4.5.6",
		},
		{
			name: "handles inline comments",
			source: `
__version__ = "7.8.9"  # trailing comment
`,
			want: "7.8.9",
		},
		{
			name: "missing version returns error",
			source: `
print("no version here")
`,
			wantErr: ErrVersionNotFound,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := VersionFromPythonSource(tc.source)
			if tc.wantErr != nil {
				if !errors.Is(err, tc.wantErr) {
					t.Fatalf("expected error %v, got %v", tc.wantErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Fatalf("expected version %q, got %q", tc.want, got)
			}
		})
	}
}

func TestVersionFromPythonFile(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "version.py")
	if err := os.WriteFile(path, []byte(`__version__ = "9.9.9"`), 0o644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	got, err := VersionFromPythonFile(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "9.9.9" {
		t.Fatalf("expected version %q, got %q", "9.9.9", got)
	}
}
