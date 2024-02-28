package lockfile_test

import (
	"testing"

	"github.com/oliverchang/osv-lib/pkg/lockfile"
)

func TestYarnLockExtractor_ShouldExtract(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "",
			path: "",
			want: false,
		},
		{
			name: "",
			path: "yarn.lock",
			want: true,
		},
		{
			name: "",
			path: "path/to/my/yarn.lock",
			want: true,
		},
		{
			name: "",
			path: "path/to/my/yarn.lock/file",
			want: false,
		},
		{
			name: "",
			path: "path/to/my/yarn.lock.file",
			want: false,
		},
		{
			name: "",
			path: "path.to.my.yarn.lock",
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			e := lockfile.YarnLockExtractor{}
			got := e.ShouldExtract(tt.path)
			if got != tt.want {
				t.Errorf("Extract() got = %v, want %v", got, tt.want)
			}
		})
	}
}
