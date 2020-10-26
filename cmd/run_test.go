package cmd

import "testing"

func TestIntegr1(t *testing.T) {

}

func TestCommon(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
}
