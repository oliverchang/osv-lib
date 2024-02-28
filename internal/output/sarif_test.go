package output_test

import (
	"bytes"
	"testing"

	"github.com/oliverchang/osv-lib/internal/output"
	"github.com/oliverchang/osv-lib/internal/testutility"
	"github.com/oliverchang/osv-lib/pkg/models"
)

func TestGroupFixedVersions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args []models.VulnerabilityFlattened
		want testutility.Snapshot
	}{
		{
			name: "",
			args: testutility.LoadJSONFixtureWithWindowsReplacements[[]models.VulnerabilityFlattened](t,
				"fixtures/flattened_vulns.json",
				map[string]string{
					"/path/to/scorecard-check-osv-e2e/sub-rust-project/Cargo.lock": "D:\\\\path\\\\to\\\\scorecard-check-osv-e2e\\\\sub-rust-project\\\\Cargo.lock",
					"/path/to/scorecard-check-osv-e2e/go.mod":                      "D:\\\\path\\\\to\\\\scorecard-check-osv-e2e\\\\go.mod",
				},
			),
			want: testutility.NewSnapshot().WithWindowsReplacements(
				map[string]string{
					"D:\\\\path\\\\to\\\\scorecard-check-osv-e2e\\\\sub-rust-project\\\\Cargo.lock": "/path/to/scorecard-check-osv-e2e/sub-rust-project/Cargo.lock",
					"D:\\\\path\\\\to\\\\scorecard-check-osv-e2e\\\\go.mod":                         "/path/to/scorecard-check-osv-e2e/go.mod",
				},
			),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := output.GroupFixedVersions(tt.args)
			tt.want.MatchJSON(t, got)
		})
	}
}

func TestPrintSARIFReport(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args models.VulnerabilityResults
		want testutility.Snapshot
	}{
		{
			name: "",
			args: testutility.LoadJSONFixtureWithWindowsReplacements[models.VulnerabilityResults](t,
				"fixtures/test-vuln-results-a.json",
				map[string]string{
					"/path/to/sub-rust-project/Cargo.lock": "D:\\\\path\\\\to\\\\sub-rust-project\\\\Cargo.lock",
					"/path/to/go.mod":                      "D:\\\\path\\\\to\\\\go.mod",
				},
			),
			want: testutility.NewSnapshot().WithWindowsReplacements(
				map[string]string{
					"lockfile:D:\\\\path\\\\to\\\\sub-rust-project\\\\Cargo.lock": "lockfile:/path/to/sub-rust-project/Cargo.lock",
					"lockfile:D:\\\\path\\\\to\\\\go.mod":                         "lockfile:/path/to/go.mod",
					"D:\\\\path\\\\to\\\\sub-rust-project/osv-scanner.toml":       "/path/to/sub-rust-project/osv-scanner.toml",
					"D:\\\\path\\\\to/osv-scanner.toml":                           "/path/to/osv-scanner.toml",
					"file:///D:/path/to":                                          "file:///path/to",
				},
			),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			bufOut := bytes.Buffer{}
			err := output.PrintSARIFReport(&tt.args, &bufOut)
			if err != nil {
				t.Errorf("Error writing SARIF output: %s", err)
			}
			tt.want.MatchText(t, bufOut.String())
		})
	}
}
