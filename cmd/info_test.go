package cmd

import (
	"testing"

	"github.com/ddefrancesco/scopectl/configurations"
	"github.com/spf13/cobra"
)

func TestInfoCommand(t *testing.T) {
	t.Skip("Skipping test that requires command execution")
	err0 := configurations.InitConfig()
	if err0 != nil {
		panic(err0)
	}

	cmd := &cobra.Command{}
	cmd.Flags().String("infos", "altitude", "")

	err := infoCmd.RunE(cmd, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestInfoCommandInvalidFlag(t *testing.T) {
	t.Skip("Skipping test that requires command execution")
	err0 := configurations.InitConfig()
	if err0 != nil {
		panic(err0)
	}
	cmd := &cobra.Command{}
	cmd.Flags().String("infos", "invalid", "")

	err := infoCmd.RunE(cmd, nil)
	if err == nil {
		t.Error("Expected error for invalid flag value")
	}
}
