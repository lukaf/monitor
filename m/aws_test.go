package monitor

import (
	"os"
	"testing"
)

const (
	fakeAccessKey = "fakeId"
	fakeSecretKey = "fakeSecret"
	fakeRegion    = "eu-west-1"
)

func TestNewCW(t *testing.T) {
	if err := os.Setenv("AWS_ACCESS_KEY_ID", fakeAccessKey); err != nil {
		t.Errorf("Unable to set AWS_ACCESS_KEY_ID environment variable.")
	}

	if err := os.Setenv("AWS_SECRET_ACCESS_KEY", fakeSecretKey); err != nil {
		t.Errorf("Unable to set AWS_SECRET_ACCESS_KEY environment variable.")
	}

	// Fetch credentials from the environment variables created above.
	var cw *CW
	var err error
	cw, err = NewCW("", "", fakeRegion)
	if err != nil {
		t.Errorf("Failed to create a new CW struct: %s", err.Error())
	}

	t.Logf("Return value of NewCW: %#v", cw)
}
