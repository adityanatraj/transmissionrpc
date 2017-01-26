package transmission

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createMoveCommand(t *testing.T) {
	const (
		expectedCmd = `{"method":"torrent-set-location","arguments":{"ids":[1,"a"],"location":"here","move":true},"tag":0}`
	)

	which := WhichTorrents{IDs: []int{1}, SHAs: []string{"a"}}
	cmd, err := createMoveCommand(which, "here", true)
	if err != nil {
		t.Fatal("failed to create command")
	}

	r, _ := json.Marshal(cmd)
	assert.Equal(t, expectedCmd, string(r))

}

func Test_createMoveCommandAll(t *testing.T) {
	const (
		expectedCmd = `{"method":"torrent-set-location","arguments":{"ids":[],"location":"here","move":true},"tag":0}`
	)

	which := WhichTorrents{}
	cmd, err := createMoveCommand(which, "here", true)
	if err != nil {
		t.Fatal("failed to create command")
	}

	r, _ := json.Marshal(cmd)
	assert.Equal(t, expectedCmd, string(r))

}

func Test_createMoveCommandBad(t *testing.T) {
	if _, err := createMoveCommand(WhichTorrents{Recent: true}, "yes", false); err == nil {
		t.Fatal("expected error for using recent in command")
	}
}
