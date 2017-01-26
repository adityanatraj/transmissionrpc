package transmission

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createIDOnlyCommand(t *testing.T) {
	const (
		expectedCmd = `{"method":"torrent-start","arguments":{"ids":[1,2,"asf","fda"]},"tag":0}`
	)
	cmd, err := createIDOnlyCommand("torrent-start", WhichTorrents{
		IDs:  []int{1, 2},
		SHAs: []string{"asf", "fda"},
	})
	if err != nil {
		t.Fatal("failed to create command")
	}

	r, _ := json.Marshal(cmd)
	assert.Equal(t, expectedCmd, string(r))

}

func Test_createIDOnlyCommandRecent(t *testing.T) {
	const (
		expectedCmd = `{"method":"torrent-reannounce","arguments":{"ids":"recently-active"},"tag":0}`
	)

	cmd, err := createIDOnlyCommand("torrent-reannounce", WhichTorrents{
		Recent: true,
	})
	if err != nil {
		t.Fatal("failed to create command")
	}

	r, _ := json.Marshal(cmd)
	assert.Equal(t, expectedCmd, string(r))
}

func Test_createIDOnlyCommandBad(t *testing.T) {

	_, err := createIDOnlyCommand("torrent-start", WhichTorrents{
		Recent: true,
		IDs:    []int{1, 2},
	})

	if err == nil {
		t.Fatal("expected error for mixing recent and ids")
	}
}
