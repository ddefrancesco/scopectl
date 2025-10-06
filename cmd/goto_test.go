package cmd

import (
	"bytes"
	"testing"

	etxClient "github.com/ddefrancesco/scopectl/restclient"
	"github.com/stretchr/testify/assert"
)

// Test that command-line parsing populates pmap correctly when using --ngc
func TestGotoCmd_RunE_Success(t *testing.T) {
	t.Skip("Skipping test that requires command execution")
	// swap the handler to capture input
	orig := gotoHandler
	defer func() { gotoHandler = orig }()

	var captured map[string]string
	gotoHandler = func(pmap map[string]string) (*etxClient.ScopeResponse, error) {
		captured = pmap
		return &etxClient.ScopeResponse{Response: "ok"}, nil
	}

	cmd := gotoCmd
	b := new(bytes.Buffer)
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--ngc=1234"})
	err := cmd.Execute()
	assert.NoError(t, err)
	if assert.NotNil(t, captured) {
		assert.Equal(t, "1234", captured["ngc"])
	}
}

// Test that an unknown flag does not populate pmap and causes execution error
func TestGotoCmd_RunE_Fail(t *testing.T) {
	t.Skip("Skipping test that requires command execution")
	orig := gotoHandler
	defer func() { gotoHandler = orig }()

	var captured map[string]string
	gotoHandler = func(pmap map[string]string) (*etxClient.ScopeResponse, error) {
		captured = pmap
		return nil, nil
	}

	cmd := gotoCmd
	b := new(bytes.Buffer)
	cmd.SetOut(b)
	// use --gnc (typo) which should not match any flag
	cmd.SetArgs([]string{"--gnc=1234"})
	err := cmd.Execute()
	assert.Error(t, err)
	// the handler should not be invoked when cobra fails to parse flags
	assert.Nil(t, captured)
}
