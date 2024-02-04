package cmd

import (
	"testing"

	"github.com/ddefrancesco/scopectl/configurations"
	"github.com/spf13/cobra"
)

func TestAlignCommandAcknowledge(t *testing.T) {
	err0 := configurations.InitConfig()
	if err0 != nil {
		panic(err0)
	}
	// Test ack flag
	cmd := &cobra.Command{}
	cmd.Flags().Bool("acknowledge", true, "")

	err := alignCmd.RunE(cmd, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

}

func TestAlignCommandAlignInvalid(t *testing.T) {
	err0 := configurations.InitConfig()
	if err0 != nil {
		panic(err0)
	}
	// Test ack flag
	cmd := &cobra.Command{}

	// Test invalid mode

	cmd.Flags().String("mode", "invalid", "")

	err := alignCmd.RunE(cmd, nil)
	if err == nil {
		t.Error("Expected error for invalid mode")
	}

}

func TestAlignCommandAlignValid(t *testing.T) {
	t.Skip()
	err0 := configurations.InitConfig()
	if err0 != nil {
		panic(err0)
	}
	// Test ack flag
	cmd := &cobra.Command{}

	// Test valid mode

	cmd.Flags().String("mode", "land", "")

	err := alignCmd.RunE(cmd, nil)
	if err != nil {
		t.Errorf("Unexpected error for valid mode: %v", err)
	}
}
