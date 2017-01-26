package transmission

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createRemoveCommand(t *testing.T) {
	const (
		expectedCmd = `{"method":"torrent-remove","arguments":{"ids":[1,"a"],"delete-local-data":true},"tag":0}`
	)

	which := WhichTorrents{IDs: []int{1}, SHAs: []string{"a"}}
	cmd, err := createRemoveCommand(which, true)
	if err != nil {
		t.Fatal("failed to create command")
	}

	r, _ := json.Marshal(cmd)
	assert.Equal(t, expectedCmd, string(r))

}

func Test_createRemoveCommandAll(t *testing.T) {
	const (
		expectedCmd = `{"method":"torrent-remove","arguments":{"ids":[],"delete-local-data":true},"tag":0}`
	)

	which := WhichTorrents{}
	cmd, err := createRemoveCommand(which, true)
	if err != nil {
		t.Fatal("failed to create command")
	}

	r, _ := json.Marshal(cmd)
	assert.Equal(t, expectedCmd, string(r))

}

func Test_createRemoveCommandBad(t *testing.T) {

	which := WhichTorrents{Recent: true}
	if _, err := createRemoveCommand(which, true); err == nil {
		t.Fatal("expected error for using recent in command")
	}

}
