package transmission

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createAddCommand(t *testing.T) {

	const (
		expectedCmd = `{"method":"torrent-add","arguments":{"download-dir":"somedir","filename":"bit.ly/1","files-unwanted":[2],"peer-limit":12},"tag":0}`
	)

	params := addClientParams{
		DownloadDir:   "somedir",
		Filename:      "bit.ly/1",
		FilesUnwanted: []int{2},
		PeerLimit:     12,
	}

	cmd, err := createAddCommand(params)
	if err != nil {
		t.Fatal("failed to create command")
	}

	r, _ := json.Marshal(cmd)
	assert.Equal(t, expectedCmd, string(r))

}

func Test_translateResponse(t *testing.T) {
	data := []byte(`{"torrents":[{"id":1,"name":"name","hashString":"hash"}]}`)
	resp, err := translateResponse(data)
	if err != nil {
		t.Fatal("failed to translate response")
	}

	expectedResponse := AddResponse{
		Torrents: []AddedTorrent{
			AddedTorrent{
				ID:         1,
				Name:       "name",
				Hashstring: "hash",
			},
		},
	}

	assert.Equal(t, expectedResponse, resp)
}

func Test_translateResponseNone(t *testing.T) {
	data := []byte(`{"torrents":[]}`)
	resp, err := translateResponse(data)
	if err != nil {
		t.Fatal("failed to translate response")
	}

	expectedResponse := AddResponse{
		Torrents: []AddedTorrent{},
	}

	assert.Equal(t, expectedResponse, resp)
}
