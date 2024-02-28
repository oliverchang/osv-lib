package output_test

import (
	"os"
	"testing"

	"github.com/oliverchang/osv-lib/internal/testutility"
)

func TestMain(m *testing.M) {
	code := m.Run()

	testutility.CleanSnapshots(m)

	os.Exit(code)
}
